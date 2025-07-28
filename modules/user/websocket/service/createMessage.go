package service

import (
	"automatic-doodle/types"
	"context"
	"log"
)

func (srv *WebsocketService) CreateMessage(pm *types.PrivateMessage, ctx context.Context) error {

	messageItem := srv.messageFactory.Create(pm.From.Id, pm.To, pm.Message)

	_, err := srv.messageRepository.Create(messageItem, ctx)

	if err != nil {
		return err
	}

	log.Println("Message Has Been send")
	return nil
}
