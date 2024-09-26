package server

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/coder/websocket"
)

func (s *Server) websocketHandler(w http.ResponseWriter, r *http.Request) {
	socket, err := websocket.Accept(w, r, nil)

	if err != nil {
		log.Printf("could not open websocket: %v", err)
		_, _ = w.Write([]byte("could not open websocket"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer socket.Close(websocket.StatusGoingAway, "server closing websocket")

	ctx := r.Context()
	// socketCtx := socket.CloseRead(ctx)
	quit := false

	for !quit {
		select {
		case <-ctx.Done():
			slog.Info("context done")
			return
		default:
			msgType, msg, err := socket.Read(ctx)
			if err != nil {
				slog.Error("error reading message", "error", err)
				return
			}
			slog.Info("message received", "msgType", msgType, "msg", string(msg))
			response := fmt.Sprintf("server timestamp: %d", time.Now().UnixNano())
			err = socket.Write(ctx, websocket.MessageText, []byte(response))
			if err != nil {
				slog.Error("error writing message", "error", err)
				return
			}
			strMsg := string(msg)
			if strings.EqualFold(strMsg, "quit") {
				quit = true
			}
		}
	}

	// for {
	// 	payload := fmt.Sprintf("server timestamp: %d", time.Now().UnixNano())
	// 	err := socket.Write(socketCtx, websocket.MessageText, []byte(payload))
	// 	if err != nil {
	// 		break
	// 	}
	// 	time.Sleep(time.Second * 2)
	// }
}
