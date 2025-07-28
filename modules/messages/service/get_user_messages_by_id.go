package service

import (
	"automatic-doodle/types"
	"context"

	"github.com/google/uuid"
)

func (srv *Service) GetMessagesByUserId(id uuid.UUID, ctx context.Context) ([]types.ReturnMessage, error) {

	messages, err := srv.messageRepository.GetMessagesByUserId(ctx, id)
	if err != nil {
		return nil, err
	}

	var privateMessages []types.ReturnMessage

	for _, msg := range messages {

		privateMessages = append(privateMessages, *types.EntToReturnMessage(msg))
	}

	return privateMessages, nil
}
