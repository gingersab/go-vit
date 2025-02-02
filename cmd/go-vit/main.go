package main

import (
	"context"
	"go-vit/internal/mon/core"
	"go-vit/internal/mon/ws"
	"net/http"
	"time"
)

func main() {
	var sra core.SystemResourceAcquirer = core.InitResourceAcquirer()
	var rm core.Monitor = core.InitResourceMonitor()
	var wsServer ws.WebSocketServer = ws.InitWebSocketServer()

	ctx, _ := context.WithCancel(context.Background())
	dur := 1000 * time.Millisecond

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		client := ws.InitClient(w, r)
		if client != nil {
			wsServer.AddClient(client)
			go func() {
				client.WaitForClose()
				wsServer.RemoveClient(client)
			}()
		}
	})

	go func() {
		for {
			stats := rm.Start(ctx, sra, dur)
			wsServer.Broadcast(stats)
			time.Sleep(dur)
		}
	}()

	http.ListenAndServe(":8080", nil)
}
