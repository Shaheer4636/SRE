package main

import (
	"sre-health-watcher/internal"
	"sre-health-watcher/utils"
	"time"
)

func main() {
	utils.InitLogger()
	cfg := internal.LoadConfig()

	utils.LogInfo("🟢 SRE Health Watcher started")

	for {
		for _, endpoint := range cfg.Endpoints {
			go internal.CheckEndpoint(endpoint, cfg.Alert)
		}
		time.Sleep(time.Duration(cfg.IntervalSeconds) * time.Second)
	}
}
