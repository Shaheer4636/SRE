# SRE Health Watcher – Real-Time Service Monitoring Tool

**SRE Health Watcher** is a lightweight, high-performance Golang-based monitoring tool designed to ensure critical internal services are alive, fast, and reliable.

Built for SRE and Platform teams, this tool actively monitors the health of services, checks for latency and HTTP failures, and alerts on any anomalies via Slack or Email.

---

## Why Use It?

| Feature             | Benefit                                                             |
|---------------------|---------------------------------------------------------------------|
| Lightweight       | No complex setup. Just Go, YAML, and cron logic.                    |
| Continuous Polling| Detect outages and latency spikes within seconds                    |
| Slack & Email     | Get alerts where your team lives                                    |
| Easy Config       | Add/remove services in a YAML file                                  |
| Extendable        | Integrate Prometheus, retries, dashboards, etc.                     |

---

##  Configuration: `config.yaml`

```yaml
interval_seconds: 60

endpoints:
  - name: Auth Service
    url: https://auth.internal.local/health
    latency_threshold_ms: 300
    alert_on: [500, 503]

  - name: Payment Gateway
    url: https://payments.internal/api/status
    latency_threshold_ms: 500
    alert_on: [502, 504]

  - name: ML Inference API
    url: https://ml-infer.internal.local/ping
    latency_threshold_ms: 1000
    alert_on: [400, 500]

alert:
  slack_webhook: "https://hooks.slack.com/services/DUMMY/TOKEN/IGNORE"
  email:
    smtp_server: "smtp.mail.local"
    port: 587
    from: "alerts@internal.local"
    to: "oncall@internal.local"
    username: "alerts@internal.local"
    password: "dummy-password"
