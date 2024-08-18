package chat

import (
	"errors"
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"time"
)

type ContentMessage struct {
	Message string `json:"message"`
	ChatId  string `json:"chat_id"`
}

type UpdateChatResponse struct {
	Chat model.Chat
	Code int
	Err  error
}

func AddMessageChat(patientID string, content ContentMessage) UpdateChatResponse {

	chat, err := graphql.GetChatById(content.ChatId)
	if err != nil {
		return UpdateChatResponse{Chat: model.Chat{}, Code: 400, Err: errors.New("id does not correspond to a chat")}
	}

	for i, participant := range chat.Participants {
		if participant.ParticipantID == patientID {
			chat.Participants[i].LastSeen = int(time.Now().Unix())
			break
		}
	}

	participants := make([]*model.ChatParticipantsInput, len(chat.Participants))
	for i, p := range chat.Participants {
		participants[i] = &model.ChatParticipantsInput{
			ParticipantID: p.ParticipantID,
			LastSeen:      p.LastSeen,
		}
	}

	messages := make([]*model.ChatMessagesInput, len(chat.Messages)+1)
	for i, m := range chat.Messages {
		messages[i] = &model.ChatMessagesInput{
			OwnerID:    m.OwnerID,
			Message:    m.Message,
			SendedTime: m.SendedTime,
		}
	}
	messages[len(messages)-1] = &model.ChatMessagesInput{
		OwnerID:    patientID,
		Message:    content.Message,
		SendedTime: int(time.Now().Unix()),
	}

	updatedChat, err := graphql.UpdateChat(chat.ID, model.UpdateChatInput{
		Participants: participants,
		Messages:     messages,
	})
	if err != nil {
		return UpdateChatResponse{Chat: model.Chat{}, Code: 400, Err: errors.New("update failed: " + err.Error())}
	}

	return UpdateChatResponse{
		Chat: updatedChat,
		Code: 200,
		Err:  nil,
	}
}
