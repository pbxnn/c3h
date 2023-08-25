package server

import (
	"errors"

	"c3h/api/r401s"
	"c3h/internal/conf"
	"c3h/internal/service"
	"c3h/third_party/websocket"

	"github.com/go-kratos/kratos/v2/log"
)

// NewWebsocketServer create a websocket server.
func NewWebsocketServer(bc *conf.Bootstrap, logger log.Logger, svc *service.R401SService) *websocket.Server {
	srv := websocket.NewServer(
		websocket.WithAddress(bc.Server.Websocket.Addr),
		websocket.WithPath(bc.Server.Websocket.Path),
		websocket.WithConnectHandle(svc.OnWebsocketConnect),
		websocket.WithCodec("json"),
	)

	svc.SetWebSocketServer(srv)

	srv.RegisterMessageHandler(websocket.MessageType(r401s.MessageType_R401S_WS),
		func(sessionId websocket.SessionID, payload websocket.MessagePayload) error {
			switch t := payload.(type) {
			case *r401s.R401SWsMessage:
				return svc.OnMessage(sessionId, t)
			default:
				return errors.New("invalid payload type")
			}
		},
		func() websocket.Any { return &r401s.R401SWsMessage{} },
	)

	return srv
}
