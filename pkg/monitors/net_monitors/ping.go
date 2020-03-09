package net_monitors

import (
	"fmt"
	"github.com/sparrc/go-ping"
	"time"
)

type PingResponse struct {
	Reachable bool
	ResponseTime time.Duration
}

func Ping(address interface{}) PingResponse {
	pingResp := PingResponse{ Reachable: false }
	pinger, err := ping.NewPinger(address.(string))
	if err != nil {
		panic(err)
	}
	pinger.SetPrivileged(true)
	pinger.Count = 1
	pinger.Run() // blocks until finished
	stats := pinger.Statistics() // get send/receive/rtt stats
	if stats.PacketLoss == 0 {
		pingResp.Reachable = true
	}
	pingResp.ResponseTime = stats.AvgRtt
	fmt.Println(stats.Rtts)
	return pingResp
}