# Gaming Ad Agent SaaS - 系统架构设计 (v1.0)

## 1. 总体架构图 (High-Level Architecture)

```mermaid
graph TD
    Client[Unity SDK] -->|HTTP/gRPC Log Stream| Ingress[K8s Ingress / API Gateway]
    
    subgraph "Data Ingestion Layer (高吞吐接入)"
        Ingress --> Collector[Log Collector Service (Go)]
        Collector --> Kafka[Kafka / Redpanda]
    end
    
    subgraph "The Brain (决策核心)"
        Kafka --> Consumer[Event Consumer (Go)]
        Consumer --> Cache{L2 Redis Cache}
        
        Cache -- "Miss" --> Reasoner[Reasoning Agent (LangChainGo)]
        Reasoner --> LLM[LLM Service (vLLM/Ollama)]
        Reasoner -- "Write Policy" --> Cache
        
        Cache -- "Hit" --> Output[Strategy Output]
    end
    
    subgraph "Management Plane (管理面)"
        DevPortal[Developer Portal] --> Postgres[(PostgreSQL)]
        DevPortal --> ClickHouse[(ClickHouse - Analytics)]
    end
    
    Consumer -- "Archive Raw" --> ClickHouse
```

## 2. 核心组件设计

### 2.1 接入服务 (Log Collector)
*   **职责**: 仅负责接收 SDK 上报的数据，鉴权，并快速写入 Kafka。不做任何业务逻辑处理。
*   **技术**: Golang, Gin/Fiber, Kafka Producer (Sarama)。
*   **性能目标**: 单 Pod 支持 5k+ QPS。
*   **数据协议**: Protobuf over HTTP/2 (以减少传输体积)。

### 2.2 消息队列 (Event Bus)
*   **选型**: Redpanda (兼容 Kafka，但单二进制文件，运维更简单，适合 K8s)。
*   **Topic 设计**:
    *   `raw-events`: 原始日志。
    *   `analyzed-intents`: 推理后的意图结果（用于离线分析）。

### 2.3 决策消费者 (Decision Consumer)
*   **职责**: 消费 Kafka 消息，执行“缓存查询 -> 推理 -> 响应”的逻辑。
*   **关键逻辑 (L2 Cache 机制)**:
    1.  **Key 生成**: `FormatKey(GameID, EventType, StandardizedContext)`
        *   例如: `game_123:level_fail:level_10:retry_5`
    2.  **查缓存**: Redis GET Key。
    3.  **命中**: 直接返回 TTL 内的策略 (例如 "Wait 5 mins before showing ad")。
    4.  **未命中**: 调用 Reasoning Agent。

### 2.4 推理 Agent (Reasoning Agent)
*   **技术**: LangChainGo。
*   **流程**:
    1.  **Prompt 组装**: 加载该游戏的 `SystemPrompt` (包含游戏术语定义)。
    2.  **LLM 调用**: 发送请求给 vLLM (部署 Qwen2.5-7B-Int4)。
    3.  **结果解析**: 将 LLM 的自然语言回复解析为结构化 JSON (`Action`, `AdTemplateID`)。
    4.  **回写缓存**: 将结果写入 Redis，TTL 设置为 10-30 分钟。

## 3. 数据存储设计

### 3.1 Redis (热数据)
*   **用途**: L2 语义缓存、频控计数器。
*   **结构**:
    *   `policy:{hash}` -> JSON 策略
    *   `user_freq:{uid}:{ad_id}` -> 计数器 (防止同一广告轰炸)

### 3.2 PostgreSQL (元数据)
*   **用途**: 开发者账号、游戏配置、广告模版库。
*   **表结构**:
    *   `games`: 游戏基础信息、AppSecret。
    *   `ad_templates`: 开发者自定义的广告样式模版。

### 3.3 ClickHouse (冷数据/分析)
*   **用途**: 存储海量原始日志，用于 BI 看板和后续模型微调。
*   **表引擎**: MergeTree，按时间分区。

## 4. 接口协议设计 (SDK <-> Server)

### 4.1 上报接口
`POST /v1/events/collect`
```json
{
  "trace_id": "uuid",
  "events": [
    {
      "name": "player_death",
      "timestamp": 1700000000,
      "payload": {
        "level": 3,
        "killer": "boss_1"
      }
    }
  ]
}
```

### 4.2 策略下发 (通过 SSE 或 Polling)
*   为了支持主动推送，SDK 需保持一个 Long Polling 或 SSE 连接。
*   或者简化为 Request-Response 模式：SDK 发送关键事件时，同步等待 Server 返回策略（会有延迟，需前端做“假加载”掩盖）。

## 5. 扩展性考虑 (Scalability)
*   **水平扩展**: Collector 和 Consumer 均无状态，可随意 HPA (Pod 自动扩缩容)。
*   **LLM 瓶颈**: 单独部署 LLM 推理集群（或使用第三方 API 兜底）。
