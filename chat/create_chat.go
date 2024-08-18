package chat

import (
	"errors"
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
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
	time := int(time.Now().Unix())

	var chatParticipants []*model.ChatParticipantsInput
	recipientIDs := content.RecipientIds
	for i := 0; i < len(recipientIDs); i++ {
		recipientID := recipientIDs[i]
		lastSeen := 0
		if recipientID == patientID {
			lastSeen = time
		}
		chatParticipants = append(chatParticipants, &model.ChatParticipantsInput{ParticipantID: recipientID, LastSeen: lastSeen})
	}

	chatMessages := []*model.ChatMessagesInput{
		{OwnerID: patientID, Message: content.Message, SendedTime: time},
	}

	newChat, err := graphql.CreateChat(model.CreateChatInput{
		Participants: chatParticipants,
		Messages:     chatMessages,
	})
	if err != nil {
		return CreateChatResponse{Chat: model.Chat{}, Code: 400, Err: errors.New("Creation failed" + err.Error())}
	}

	patient, err := graphql.GetPatientById(patientID)
	if err != nil {
		return CreateChatResponse{Chat: model.Chat{}, Code: 400, Err: errors.New("Id does not correspond to a patient")}
	}

	_, err = graphql.UpdatePatient(patientID, model.UpdatePatientInput{
		ChatIds: append(patient.ChatIds, &newChat.ID),
	})
	if err != nil {
		return CreateChatResponse{Chat: model.Chat{}, Code: 500, Err: errors.New("Unable to update patient")}
	}

	for _, chatParticipant := range chatParticipants {
		if chatParticipant.ParticipantID == patientID {
			continue
		}
		doctor, err := graphql.GetDoctorById(chatParticipant.ParticipantID)
		if err != nil {
			return CreateChatResponse{Chat: model.Chat{}, Code: 500, Err: errors.New("Unable to update doctor")}
		}
		_, err = graphql.UpdateDoctor(doctor.ID, model.UpdateDoctorInput{
			ChatIds: append(doctor.ChatIds, &newChat.ID),
		})
		if err != nil {
			return CreateChatResponse{Chat: model.Chat{}, Code: 500, Err: errors.New("Unable to update doctor")}
		}

		if !containsPatientID(doctor.PatientIds, &patientID) {
			_, err = graphql.UpdateDoctor(doctor.ID, model.UpdateDoctorInput{
				PatientIds: append(doctor.PatientIds, &patientID),
				ChatIds:    append(doctor.ChatIds, &newChat.ID),
			})
			if err != nil {
				return CreateChatResponse{Chat: model.Chat{}, Code: 500, Err: errors.New("Unable to update doctor")}
			}
		}

	}

	return CreateChatResponse{
		Chat: newChat,
		Code: 201,
		Err:  nil,
	}
}

func containsPatientID(patientIDs []*string, patientID *string) bool {
	for _, id := range patientIDs {
		if id == patientID {
			return true
		}
	}
	return false
}
