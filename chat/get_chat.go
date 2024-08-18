package chat

import (
	"errors"
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
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
	chat, err := graphql.GetChatById(id)
	if err != nil {
		return GetChatByIdResponse{model.Chat{}, 400, errors.New("id does not correspond to a chat")}
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

	var messages []*model.ChatMessagesInput
	for _, m := range chat.Messages {
		messages = append(messages, &model.ChatMessagesInput{
			OwnerID:    m.OwnerID,
			Message:    m.Message,
			SendedTime: m.SendedTime,
		})
	}

	updatedChat, err := graphql.UpdateChat(chat.ID, model.UpdateChatInput{
		Participants: participants,
		Messages:     messages,
	})
	if err != nil {
		return GetChatByIdResponse{model.Chat{}, 400, errors.New("update failed: " + err.Error())}
	}
	return GetChatByIdResponse{updatedChat, 200, nil}
}

func GetChat(patientId string) GetChatsResponse {
	chats, err := graphql.GetChats(patientId, nil)
	if err != nil {
		return GetChatsResponse{[]model.Chat{}, 400, errors.New("invalid input: " + err.Error())}
	}
	return GetChatsResponse{chats, 200, nil}
}
