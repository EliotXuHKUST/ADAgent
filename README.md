# ADAgent - 智能广告投放 Agent

ADAgent 是一个基于大语言模型（LLM）构建的智能广告投放助手，旨在为营销人员提供全流程的自动化广告投放解决方案。从策略制定、文案生成到自动投放和效果监控，ADAgent 能够显著提升广告投放的效率与 ROI。

## 🚀 项目简介

在数字营销领域，广告投放通常涉及繁琐的操作和复杂的数据分析。ADAgent 利用 AI Agent 技术，通过多智能体协作（Multi-Agent Collaboration），实现了广告投放流程的自动化与智能化。

## ✨ 主要功能

- **🧠 智能策略规划**: 根据产品特性、目标受众和预算，自动生成多渠道投放策略。
- **✍️ 创意文案生成**: 集成 AIGC 能力，针对不同平台（如 Google, Facebook, TikTok）自动生成高转化率的广告标题和描述。
- **🎨 素材建议与生成**: 提供广告图片/视频的创意方向建议，并可对接图像生成模型。
- **🤖 自动化投放执行**: 对接主流广告平台 API，自动创建广告系列、广告组和广告创意。
- **📊 实时监控与优化**: 24/7 监控 CTR, CPC, ROAS 等关键指标，基于数据反馈自动调整出价和预算。
- **📈 深度数据报告**: 生成可视化的投放效果分析报告，提供可操作的优化建议。

## 🛠️ 系统架构

ADAgent 采用模块化设计，主要包含以下核心 Agent：

1.  **Strategy Agent (策略官)**: 负责整体投放策略的制定和预算分配。
2.  **Creative Agent (创意官)**: 负责广告文案撰写和素材创意生成。
3.  **Execution Agent (执行官)**: 负责调用广告平台 API 进行实际操作。
4.  **Data Agent (分析师)**: 负责数据收集、效果评估和反馈闭环。

## 🏁 快速开始

### 环境要求

- Python 3.9+
- OpenAI API Key (或兼容的 LLM API Key)
- 目标广告平台的 API Access Token (Google Ads API, Meta Graph API 等)

### 安装

```bash
git clone https://github.com/EliotXuHKUST/ADAgent.git
cd ADAgent
pip install -r requirements.txt
```

### 配置

复制 `.env.example` 文件为 `.env`，并填入必要的配置信息：

```ini
# LLM Configuration
OPENAI_API_KEY=sk-xxxxxxxxxxxxxxxxxxxxxxxx

# Advertising Platforms
GOOGLE_ADS_DEVELOPER_TOKEN=xxxxxxxx
FACEBOOK_APP_ID=xxxxxxxx
FACEBOOK_APP_SECRET=xxxxxxxx
```

### 运行

启动 ADAgent 服务：

```bash
python main.py
```

或者使用 CLI 模式运行特定任务：

```bash
python main.py --task "为一款新的运动鞋在 Instagram 上制定投放计划，预算 $500"
```

## 🗺️ 路线图 (Roadmap)

- [ ] 支持更多广告平台 (Twitter/X, LinkedIn)
- [ ] 引入多模态模型以支持视频素材分析
- [ ] 强化基于强化学习 (RL) 的自动出价策略
- [ ] 开发 Web 可视化管理界面

## 🤝 贡献指南

欢迎提交 Pull Request 或 Issue！在提交代码前，请确保通过了所有的单元测试。

## 📄 许可证

本项目采用 MIT 许可证。详情请参阅 [LICENSE](LICENSE) 文件。
