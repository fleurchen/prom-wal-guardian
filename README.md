# Prometheus WAL Guardian (Go Version)

✨ A safe and pluggable WAL pre-check & cleaning utility for Prometheus written in Go. Prevent OOM crashes caused by oversized WAL segments.

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
4. Logs result and returns exit code for scripting

---

## 🌐 Quick Start

```bash
go build -o prom-wal-guardian ./cmd
docker run -v /prometheus/data:/data prom-wal-guardian --path /data/wal --max-size 5GB
```

---

## 🧰 Example Systemd Hook

```ini
[Unit]
Before=prometheus.service

[Service]
ExecStartPre=/usr/local/bin/prom-wal-guardian --path /data/wal --max-size 5GB
```

---

## 🎯 Roadmap
- Segment age-based cleanup policy
- JSON config file support
- Metrics exporter via HTTP
- WAL corruption detection

---

## 🚀 License
Apache-2.0