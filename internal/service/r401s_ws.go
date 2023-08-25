package service

import (
	"context"

	"c3h/api/r401s"
	"c3h/third_party/websocket"
)

func (s *R401SService) SetWebSocketServer(ws *websocket.Server) {
	s.ws = ws
}

func (s *R401SService) OnWebsocketConnect(sessionId websocket.SessionID, register bool) {
	if register {
		s.logger.Infof("%s connected", sessionId)
	} else {
		s.logger.Infof("%s disconnect", sessionId)
	}
}

func (s *R401SService) OnMessage(sessionId websocket.SessionID, msg *r401s.R401SWsMessage) error {
	s.logger.Debugf("on message")
	res, err := s.getWsMessage(nil)
	if err != nil {
		s.logger.Warnf("get collected data failed,err :%s", err.Error())
		return err
	}
	s.ws.Broadcast(websocket.MessageType(r401s.MessageType_R401S_WS), res)
	return nil
}

func (s *R401SService) SendMessage(ctx context.Context) {
	res, err := s.getWsMessage(ctx)
	if err != nil {
		s.logger.Warnf("get collected data failed,err :%s", err.Error())
		return
	}
	s.ws.Broadcast(websocket.MessageType(r401s.MessageType_R401S_WS), res)
}

func (s *R401SService) getWsMessage(ctx context.Context) (*r401s.R401SWsMessage, error) {
	res := &r401s.R401SWsMessage{Code: 20000, Data: &r401s.R401SWsMessage_Data{}}
	apcReply, err := s.GetAPCControl(ctx, nil)
	if err != nil {
		return nil, err
	}
	res.Data.ApcVars = apcReply.Data.List

	opVars, err := s.GetOperationVars(ctx, nil)
	if err != nil {
		return nil, err
	}
	res.Data.OperationVars = opVars.Data.List

	cfVars, err := s.GetConfoundingVars(ctx, nil)
	if err != nil {
		return nil, err
	}
	res.Data.ConfoundingVars = cfVars.Data.List

	stVars, err := s.GetStatusVars(ctx, nil)
	if err != nil {
		return nil, err
	}
	res.Data.StatusVars = stVars.Data.List

	rpVars, err := s.GetReactorPerformance(ctx, nil)
	if err != nil {
		return nil, err
	}
	res.Data.ReactorPerformanceVars = rpVars.Data.List

	return res, nil
}
