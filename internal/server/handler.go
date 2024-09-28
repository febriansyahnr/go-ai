package server

import (
	"log"
	"log/slog"
	"net/http"
	"strings"

	"github.com/coder/websocket"
	"github.com/sashabaranov/go-openai"
)

func (s *Server) websocketHandler(w http.ResponseWriter, r *http.Request) {
	socket, err := websocket.Accept(w, r, nil)
	tokenBearer := r.Header.Get("Authorization")
	if tokenBearer != "" {
		slog.Info("Auth", slog.String("token", tokenBearer))
	}

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

	messages := make([]openai.ChatCompletionMessage, 0)
	messages = append(messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleSystem,
		Content: "You are a helpful assistant.",
	})

	if s.AI == nil {
		socket.Close(websocket.StatusNormalClosure, "AI Service is nill")
		return
	}

	for !quit {
		select {
		case <-ctx.Done():
			slog.Info("context done")
			return
		default:
			_, msg, err := socket.Read(ctx)
			if err != nil {
				socket.Close(websocket.StatusNormalClosure, "error reading message")
				return
			}

			strMsg := string(msg)
			if strings.EqualFold(strMsg, "quit()") {
				quit = true
				continue
			}

			response, err := s.AI.SendMessage(ctx, strMsg, &messages)
			if err != nil {
				slog.Error("error sending message", "error", err)
				break
			}

			if err := socket.Write(ctx, websocket.MessageText, []byte(response)); err != nil {
				slog.Error("error writing message", "error", err)
				break
			}
		}
	}
	socket.Close(websocket.StatusNormalClosure, "")
}
