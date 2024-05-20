package chat

import (
	"context"
	"errors"
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
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
	gqlClient := graphql.CreateClient()

	chat, err := graphql.GetChatById(context.Background(), gqlClient, content.ChatId)
	if err != nil {
		return UpdateChatResponse{Chat: model.Chat{}, Code: 400, Err: errors.New("id does not correspond to a chat")}
	}

	for i, participant := range chat.GetChatById.Participants {
		if participant.Participant_id == patientID {
			chat.GetChatById.Participants[i].Last_seen = int(time.Now().Unix())
			break
		}
	}

	participants := make([]graphql.ChatParticipantsInput, len(chat.GetChatById.Participants))
	for i, p := range chat.GetChatById.Participants {
		participants[i] = graphql.ChatParticipantsInput{
			Participant_id: p.Participant_id,
			Last_seen:      p.Last_seen,
		}
	}

	messages := make([]graphql.ChatMessagesInput, len(chat.GetChatById.Messages)+1)
	for i, m := range chat.GetChatById.Messages {
		messages[i] = graphql.ChatMessagesInput{
			Owner_id:    m.Owner_id,
			Message:     m.Message,
			Sended_time: m.Sended_time,
		}
	}
	messages[len(messages)-1] = graphql.ChatMessagesInput{
		Owner_id:    patientID,
		Message:     content.Message,
		Sended_time: int(time.Now().Unix()),
	}

	updatedChat, err := graphql.UpdateChat(context.Background(), gqlClient, chat.GetChatById.Id, participants, messages)
	if err != nil {
		return UpdateChatResponse{Chat: model.Chat{}, Code: 400, Err: errors.New("update failed: " + err.Error())}
	}

	participantReturn := make([]*model.ChatParticipants, len(updatedChat.UpdateChat.Participants))
	for i, p := range updatedChat.UpdateChat.Participants {
		participantReturn[i] = &model.ChatParticipants{
			ParticipantID: p.Participant_id,
			LastSeen:      p.Last_seen,
		}
	}

	messageReturn := make([]*model.ChatMessages, len(updatedChat.UpdateChat.Messages))
	for i, m := range updatedChat.UpdateChat.Messages {
		messageReturn[i] = &model.ChatMessages{
			OwnerID:    m.Owner_id,
			Message:    m.Message,
			SendedTime: m.Sended_time,
		}
	}

	return UpdateChatResponse{
		Chat: model.Chat{
			ID:           updatedChat.UpdateChat.Id,
			Participants: participantReturn,
			Messages:     messageReturn,
		},
		Code: 200,
		Err:  nil,
	}
}
