package service

import (
	"automatic-doodle/types"
	"context"
	"log"

	"github.com/google/uuid"
)

func (srv *Service) FetchConversationsByUserId(
	id uuid.UUID,
	recipentID uuid.UUID,
	ctx context.Context,
) ([]types.ReturnMessage, error) {

	messages, err := srv.messageRepository.FetchConversationsByUserId(id, recipentID, ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var privateMessages []types.ReturnMessage

	for _, msg := range messages {

		privateMessages = append(privateMessages, *types.EntToReturnMessage(msg))
	}

	return privateMessages, nil
}
