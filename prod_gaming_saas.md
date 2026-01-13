# Gaming Ad Agent SaaS - 产品规格书 (MVP)

## 1. 产品概述
**产品名称**: Gaming Ad Agent (GAA)
**定位**: 面向独立游戏开发者与中小型游戏工作室的 **"原生化广告决策引擎 SaaS"**。
**核心价值**: 
1.  **高 eCPM**: 捕捉玩家情绪最高昂/脆弱的时刻 (Moment of Truth) 进行投放。
2.  **低干扰**: 将广告包装为游戏内福利、任务或 NPC 对话，保护留存率。
3.  **零门槛**: 通过 SDK 自动采集并理解游戏日志，无需开发者手动打复杂的广告点位。

## 2. 核心用户流 (User Flow)
1.  **开发者**: 注册 SaaS -> 下载 Unity/Unreal SDK -> 拖入游戏工程 -> 配置广告样式 (如“宝箱样式的广告位”)。
2.  **玩家**: 玩游戏 -> 触发特定事件 (如连续失败) -> SDK 上报日志 -> Cloud Agent 推理出“挫败情绪” -> 返回“复活道具”广告策略 -> SDK 渲染原生广告。
3.  **广告主**: (初期对接 AdMob/Pangle, 后期自建) 获得高转化的流量。

## 3. 功能模块详解

### 3.1 开发者端 (Developer Portal)

#### A. 游戏接入与 SDK 配置
*   **Game Key 管理**: 生成唯一的 `app_id` 和 `secret_key`。
*   **事件映射 (Event Mapping)**:
    *   虽然 SDK 自动抓取，但允许开发者手动修正。
    *   *界面*: "我们在日志里发现了 `PlayerDeath` 事件，是否将其标记为【负面/挫败】信号？" [是/否]

#### B. 原生广告位设计器 (Native Ad Designer)
不再是 Banner/插屏，而是**"游戏资产 (Game Asset)"**。
*   **样式模板库**:
    1.  **NPC Dialog**: 气泡对话框样式（"嘿，听说你需要把新武器？"）。
    2.  **Loot Box**: 宝箱掉落样式（点击宝箱 -> 播放激励视频 -> 获得奖励）。
    3.  **System Notice**: 顶部公告样式（"限时福利：领取 50 金币"）。
*   **素材绑定**: 开发者上传自己的 UI 素材（背景图、按钮样式），SDK 动态替换文字，确保风格统一。

#### C. 数据看板 (Dashboard)
*   **实时收入**: eCPM, 总收益。
*   **情绪分布 (Emotion Heatmap)**:
    *   "您的玩家在第 3 关产生大量【愤怒】情绪，建议降低难度或增加【复活】广告投放。"
    *   *价值*: 这不仅是广告后台，还是**游戏运营分析工具** (这是吸引开发者的杀手锏)。

### 3.2 客户端 SDK (Unity/Unreal)

#### A. 数据采集 (Data Capture)
*   **接口**: `GAA.Track(string eventName, Dictionary<string, object> params)`
*   **自动采集**:
    *   FPS/Ping 值 (判断卡顿/焦虑)。
    *   点击频率 (判断急躁)。
    *   游戏内资源变化 (金币归零 -> 贫穷 -> 推送“金币大礼包”)。

#### B. 渲染引擎 (Renderer)
*   **预加载 (Pre-load)**: 预测即将发生广告展示，提前缓存素材。
*   **原生渲染**: 使用游戏引擎原生的 UI 系统 (UGUI / UMG) 绘制广告，而不是 WebView，确保无违和感。

### 3.3 云端核心引擎 (The Core)

#### A. 游戏垂直意图模型 (Gaming Intent Model)
基于 LLM 微调，专门理解游戏黑话。
*   **输入**: `{ "event": "rank_drop", "hero": "yasuo", "kda": "0/10/0" }`
*   **推理**: `Intent: "Extreme Frustration" (极度挫败) + "Blame Teammates" (甩锅队友) -> Strategy: "Boost Service" (代练/陪玩) OR "High Damage Skin" (皮肤转运)`

#### B. 意图拍卖行 (Intent Exchange)
*   **缓存层 (L2 Cache)**:
    *   Key: `GameType:MOBA` + `Event:Lose_Streak_3`
    *   Value: `AdStrategy: "Push_Coaching_App"` (推陪玩 App)
*   **决策逻辑**:
    *   IF `User_Value` > Threshold THEN 调用 LLM (精准生成文案)。
    *   ELSE 走缓存策略 (低成本)。

## 4. 数据结构示例 (Schema)

### 4.1 原始日志 (Raw Log)
```json
{
  "game_id": "game_123",
  "user_id": "u_999",
  "timestamp": 1705308000,
  "event_name": "level_fail",
  "context": {
    "level": 15,
    "retry_count": 4,
    "resource_left": 0,
    "death_reason": "boss_skill_A"
  }
}
```

### 4.2 推理结果 (Inferred Intent)
```json
{
  "emotion": "anxious",
  "intensity": 0.8,
  "need": "resource_recovery",
  "recommended_action": "reward_video",
  "copywriting_tone": "encouraging"
}
```

### 4.3 广告策略 (Ad Strategy)
```json
{
  "ad_type": "loot_box",
  "content": {
    "title": "Don't give up!",
    "body": "Watch this to revive with full HP.",
    "reward_id": "revive_potion"
  },
  "dsp_payload": { ... } // 具体的广告素材链接
}
```

## 5. MVP 开发计划
1.  **Week 1-2**: 定义 SDK 接口，开发 Mock Server (模拟云端)。
2.  **Week 3-4**: 训练/微调一个小型的“游戏意图理解模型” (基于 Llama-3-8B 或 Qwen-7B)。
3.  **Week 5-6**: 开发 Unity Demo 游戏，跑通“死亡 -> 推理 -> 弹窗”闭环。

