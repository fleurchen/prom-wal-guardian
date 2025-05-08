# Project: prom-wal-guardian

## 📂 Repository Structure

```
prom-wal-guardian/
├── README.md
├── wal_guardian/
│   ├── __init__.py
│   ├── config.py            # 配置加载（阈值、策略）
│   ├── wal_checker.py       # 主逻辑：检测 WAL 目录体积、segment 分析
│   ├── wal_truncator.py     # 安全截断/清理 WAL segment 的逻辑
│   └── utils.py             # 通用函数（如目录体积计算）
├── scripts/
│   └── prestart_hook.sh     # 示例启动前清理脚本（集成到 systemd/k8s init）
├── tests/
│   └── test_core.py         # 单元测试样例
├── .gitignore
└── pyproject.toml           # Poetry 项目管理
```

## 📖 README.md 内容

```markdown
# Prometheus WAL Guardian

✨ A safe and pluggable WAL pre-check & cleaning utility for Prometheus. Prevent OOM crashes caused by oversized WAL segments.

## ✨ Why This Project

Prometheus can crash with OOM if its WAL grows too large (e.g., due to failed remote-write or TSDB compression issues). This tool provides a proactive solution:

- ✅ Pre-check WAL directory size before Prometheus starts
- ✅ Auto-truncate old segments if threshold is exceeded
- ✅ Provide configurable safe-guards to avoid data loss
- ✅ Can be used as a systemd or Kubernetes init hook

---

## ⚙️ How It Works

1. Scans the Prometheus `data/wal/` directory
2. Checks the total size or age of WAL segments
3. Optionally truncates oldest segments beyond threshold
4. Returns exit code for CI/CD or hook script integration

---

## 🌐 Quick Start

```bash
pip install prom-wal-guardian

# or clone and run manually
python3 -m wal_guardian.checker --path=/prometheus/data/wal --max-size=5GB
```

---

## 🧳 Example Use Case: Systemd

```ini
[Unit]
Before=prometheus.service

[Service]
ExecStartPre=/usr/local/bin/prom-wal-guardian --path /data/wal --max-size 5GB
```

---

## 🎓 Advanced Features (Planned)
- Segment-based age policy
- Metrics export via Prometheus exporter
- Slack/Email alert if exceeded
- WAL corruption auto-detector

---

## 🚀 License
Apache-2.0
