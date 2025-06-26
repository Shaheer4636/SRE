package internal

import (
	"net/http"
	"time"
	"sre-health-watcher/utils"
)

func CheckEndpoint(ep Endpoint, alert AlertConfig) {
	start := time.Now()
	resp, err := http.Get(ep.URL)
	duration := time.Since(start).Milliseconds()

	if err != nil {
		utils.LogErrorf("🔴 %s - connection failed: %v", ep.Name, err)
		SendAlert(ep.Name, ep.URL, "Connection failed", alert)
		return
	}
	defer resp.Body.Close()

	if duration > int64(ep.LatencyThresholdMS) {
		utils.LogWarnf("⚠️ %s - High latency: %dms", ep.Name, duration)
		SendAlert(ep.Name, ep.URL, "High latency", alert)
	}

	for _, code := range ep.AlertOn {
		if resp.StatusCode == code {
			utils.LogErrorf("🔴 %s - Unexpected status code: %d", ep.Name, resp.StatusCode)
			SendAlert(ep.Name, ep.URL, "HTTP "+http.StatusText(code), alert)
			return
		}
	}

	utils.LogInfof("✅ %s - Healthy (%dms, %d)", ep.Name, duration, resp.StatusCode)
}
