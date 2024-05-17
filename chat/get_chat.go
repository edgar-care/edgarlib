package chat

import (
	"context"
	"errors"
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
	"time"
)

type ChatParticipants struct {
	ParticipantId string `json:"participant_id"`
	LastSeen      int    `json:"last_seen"`
}

type ChatMessages struct {
	OwnerId  string `json:"owner_id"`
	Message  string `json:"message"`
	SendTime int    `json:"send_time"`
}

type GetChatByIdResponse struct {
	Chat model.Chat
	Code int
	Err  error
}

type GetChatsResponse struct {
	Chats []model.Chat
	Code  int
	Err   error
}

func UpdateMessageRead(patientID string, id string) GetChatByIdResponse {
	gqlClient := graphql.CreateClient()
	var res model.Chat

	chat, err := graphql.GetChatById(context.Background(), gqlClient, id)
	if err != nil {
		return GetChatByIdResponse{model.Chat{}, 400, errors.New("id does not correspond to a chat")}
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

	var messages []graphql.ChatMessagesInput
	for _, m := range chat.GetChatById.Messages {
		messages = append(messages, graphql.ChatMessagesInput{
			Owner_id:    m.Owner_id,
			Message:     m.Message,
			Sended_time: m.Sended_time,
		})
	}

	updatedChat, err := graphql.UpdateChat(context.Background(), gqlClient, chat.GetChatById.Id, participants, messages)
	if err != nil {
		return GetChatByIdResponse{model.Chat{}, 400, errors.New("update failed: " + err.Error())}
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

	res = model.Chat{
		ID:           updatedChat.UpdateChat.Id,
		Participants: participantReturn,
		Messages:     messageReturn,
	}
	return GetChatByIdResponse{res, 200, nil}
}

func GetChat(patientId string) GetChatsResponse {
	gqlClient := graphql.CreateClient()
	var res []model.Chat

	chats, err := graphql.GetChats(context.Background(), gqlClient, patientId)
	if err != nil {
		return GetChatsResponse{[]model.Chat{}, 400, errors.New("invalid input: " + err.Error())}
	}

	for _, chat := range chats.GetChats {
		var participants []*model.ChatParticipants
		for _, p := range chat.Participants {
			participants = append(participants, &model.ChatParticipants{
				ParticipantID: p.Participant_id,
				LastSeen:      p.Last_seen,
			})
		}
		var messages []*model.ChatMessages
		for _, m := range chat.Messages {
			messages = append(messages, &model.ChatMessages{
				OwnerID:    m.Owner_id,
				Message:    m.Message,
				SendedTime: m.Sended_time,
			})
		}

		res = append(res, model.Chat{
			ID:           chat.Id,
			Participants: participants,
			Messages:     messages,
		})
	}
	return GetChatsResponse{res, 200, nil}
}
