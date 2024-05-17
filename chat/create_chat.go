package chat

import (
	"context"
	"errors"
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
	"time"
)

type ParticipantsInput struct {
	ParticipantId string `json:"participant_id"`
	LastSeen      int    `json:"last_seen"`
}

type MessagesInput struct {
	OwnerId  string `json:"owner_id"`
	Message  string `json:"message"`
	SendTime int    `json:"send_time"`
}

type ContentInput struct {
	Message      string   `json:"message"`
	RecipientIds []string `json:"recipient_ids"`
}

type CreateChatResponse struct {
	Chat model.Chat
	Code int
	Err  error
}

func CreateChat(patientID string, content ContentInput) CreateChatResponse {
	gqlClient := graphql.CreateClient()

	var chatParticipants []graphql.ChatParticipantsInput
	recipientIDs := content.RecipientIds
	for i := 0; i < len(recipientIDs); i++ {
		recipientID := recipientIDs[i]
		lastSeen := 0
		if recipientID == patientID {
			lastSeen = int(time.Now().Unix())
		}
		chatParticipants = append(chatParticipants, graphql.ChatParticipantsInput{Participant_id: recipientID, Last_seen: lastSeen})
	}

	// Créer le message pour le patientID
	chatMessages := []graphql.ChatMessagesInput{
		{Owner_id: patientID, Message: content.Message, Sended_time: int(time.Now().Unix())},
	}

	newChat, err := graphql.CreateChat(context.Background(), gqlClient, chatParticipants, chatMessages)
	if err != nil {
		return CreateChatResponse{Chat: model.Chat{}, Code: 400, Err: errors.New("update failed" + err.Error())}
	}

	patient, err := graphql.GetPatientById(context.Background(), gqlClient, patientID)
	if err != nil {
		return CreateChatResponse{Chat: model.Chat{}, Code: 400, Err: errors.New("id does not correspond to a patient")}
	}

	_, err = graphql.UpdatePatient(context.Background(), gqlClient, patientID, patient.GetPatientById.Email, patient.GetPatientById.Password, patient.GetPatientById.Medical_info_id, patient.GetPatientById.Rendez_vous_ids, patient.GetPatientById.Document_ids, patient.GetPatientById.Treatment_follow_up_ids, append(patient.GetPatientById.Chat_ids, newChat.CreateChat.Id))
	if err != nil {
		return CreateChatResponse{Chat: model.Chat{}, Code: 500, Err: errors.New("unable to update patient")}
	}

	participantReturn := make([]*model.ChatParticipants, len(newChat.CreateChat.Participants))
	for i, p := range newChat.CreateChat.Participants {
		participantReturn[i] = &model.ChatParticipants{
			ParticipantID: p.Participant_id,
			LastSeen:      p.Last_seen,
		}
	}

	messageReturn := make([]*model.ChatMessages, len(newChat.CreateChat.Messages))
	for i, m := range newChat.CreateChat.Messages {
		messageReturn[i] = &model.ChatMessages{
			OwnerID:    m.Owner_id,
			Message:    m.Message,
			SendedTime: m.Sended_time,
		}
	}

	return CreateChatResponse{
		Chat: model.Chat{
			ID:           newChat.CreateChat.Id,
			Participants: participantReturn,
			Messages:     messageReturn,
		},
		Code: 201,
		Err:  nil,
	}
}
