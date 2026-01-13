# Gaming Ad Agent SaaS - MVP 开发计划

## 0. 总体目标
在 **6 周** 内完成最小可行性产品 (MVP)，验证核心闭环：
**SDK 采集日志 -> 云端推理意图 -> 返回原生策略 -> 客户端渲染展示**

## Phase 1: 基础设施与数据接入 (Weeks 1-2)
**目标**: 打通数据流，模拟生成游戏日志并成功入库。

### 技术栈 (Tech Stack)
*   **Language**: Golang 1.22+
*   **AI Framework**: LangChainGo
*   **Infrastructure**: Kubernetes (K8s), Docker
*   **Messaging**: Kafka / Redpanda

### 任务列表
1.  **后端基础 (Backend - Go)**
    *   [ ] 初始化 Go Module (`go mod init github.com/EliotXuHKUST/ADAgent`).
    *   [ ] 搭建 HTTP Server (Gin/Echo) 接收日志。
    *   [ ] 集成 LangChainGo 用于后续 LLM 调用。
    *   [ ] 编写 Dockerfile 和 K8s Helm Charts。
2.  **SDK 原型 (SDK Proto)**
    *   [ ] 定义跨平台 SDK 接口规范 (C# for Unity)。
    *   [ ] 编写 Go 脚本 `cmd/mock_client`，模拟发送 1000 QPS 的游戏战斗日志。
3.  **数据清洗 (ETL)**
    *   [ ] 实现 Log Parser：将 JSON 日志扁平化并存入 ClickHouse/Postgres。

**交付物**:
*   一个能跑通的 API Server。
*   一个自动发日志的脚本。
*   数据库里有模拟的战斗数据。

## Phase 2: 核心大脑与模型微调 (Weeks 3-4)
**目标**: 让系统能“看懂”日志，输出合理的广告策略。

### 任务列表
1.  **意图模型 (Intent Model)**
    *   [ ] 准备数据集：手动构造 500 条“游戏事件 -> 玩家情绪 -> 推荐策略”的样本对。
    *   [ ] 微调 (Fine-tune)：基于 Qwen2.5-7B 或 Llama-3-8B，使用 LoRA 进行微调，使其理解游戏术语（如 "gank", "kda", "wipe"）。
    *   [ ] 部署模型：使用 vLLM 部署推理服务，封装为内部 API。
2.  **决策引擎 (Decision Engine)**
    *   [ ] 实现 **L2 语义缓存**：Redis 缓存高频场景的策略（避免每次都调 LLM）。
    *   [ ] 实现简单的规则引擎：`IF emotion == 'frustrated' AND user_level < 5 THEN strategy = 'revive_potion'`。
3.  **广告库 (Ad Inventory)**
    *   [ ] 建立一个 Mock 广告库 (JSON)，包含 10 个虚拟广告（复活药水、皮肤折扣、陪玩服务）。

**交付物**:
*   一个推理 API：输入 Log，输出 JSON 策略。
*   命中率测试报告：缓存命中率 > 50%。

## Phase 3: 客户端集成与 Demo (Weeks 5-6)
**目标**: 做出一个能给人看的 Demo。

### 任务列表
1.  **Unity Demo 游戏**
    *   [ ] 找一个开源的 Unity 小游戏 (如 2D 射击或跑酷)。
    *   [ ] 集成 Phase 1 定义的 C# SDK。
2.  **原生渲染器 (Native Renderer)**
    *   [ ] 在 Unity 中实现 3 种广告模版：
        *   **NPC 对话框** (Dialog Box)
        *   **宝箱掉落** (Loot Crate)
        *   **顶部横幅** (Toast)
3.  **联调与验收**
    *   [ ] 场景测试：在 Demo 游戏中故意连续死 3 次，验证是否弹出“复活药水”广告。
    *   [ ] 延迟测试：从死亡到弹窗的端到端延迟 < 1s。

**交付物**:
*   一个可运行的 `.apk` 或 `.exe` Demo 游戏。
*   演示视频：展示从操作到广告触发的全过程。

## 资源需求
*   **后端开发**: 1人 (Golang/K8s)
*   **AI/算法**: 1人 (Prompt/Fine-tuning)
*   **Unity 开发**: 1人 (或外包，负责 Demo 制作)
*   **算力**: 1 台 A100/A10 (训练) + K8s Cluster (推理部署)

