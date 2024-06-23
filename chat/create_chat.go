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

	time := int(time.Now().Unix())

	var chatParticipants []graphql.ChatParticipantsInput
	recipientIDs := content.RecipientIds
	for i := 0; i < len(recipientIDs); i++ {
		recipientID := recipientIDs[i]
		lastSeen := 0
		if recipientID == patientID {
			lastSeen = time
		}
		chatParticipants = append(chatParticipants, graphql.ChatParticipantsInput{Participant_id: recipientID, Last_seen: lastSeen})
	}

	chatMessages := []graphql.ChatMessagesInput{
		{Owner_id: patientID, Message: content.Message, Sended_time: time},
	}

	newChat, err := graphql.CreateChat(context.Background(), gqlClient, chatParticipants, chatMessages)
	if err != nil {
		return CreateChatResponse{Chat: model.Chat{}, Code: 400, Err: errors.New("Creation failed" + err.Error())}
	}

	patient, err := graphql.GetPatientById(context.Background(), gqlClient, patientID)
	if err != nil {
		return CreateChatResponse{Chat: model.Chat{}, Code: 400, Err: errors.New("Id does not correspond to a patient")}
	}

	_, err = graphql.UpdatePatient(context.Background(), gqlClient, patientID, patient.GetPatientById.Email, patient.GetPatientById.Password, patient.GetPatientById.Medical_info_id, patient.GetPatientById.Rendez_vous_ids, patient.GetPatientById.Document_ids, patient.GetPatientById.Treatment_follow_up_ids, append(patient.GetPatientById.Chat_ids, newChat.CreateChat.Id), patient.GetPatientById.Device_connect, patient.GetPatientById.Double_auth_methods_id)
	if err != nil {
		return CreateChatResponse{Chat: model.Chat{}, Code: 500, Err: errors.New("Unable to update patient")}
	}

	for _, chatParticipant := range chatParticipants {
		if chatParticipant.Participant_id == patientID {
			continue
		}
		doctor, err := graphql.GetDoctorById(context.Background(), gqlClient, chatParticipant.Participant_id)
		if err != nil {
			return CreateChatResponse{Chat: model.Chat{}, Code: 500, Err: errors.New("Unable to update doctor")}
		}
		_, err = graphql.UpdateDoctor(context.Background(), gqlClient, doctor.GetDoctorById.Id, doctor.GetDoctorById.Email, doctor.GetDoctorById.Password, doctor.GetDoctorById.Name, doctor.GetDoctorById.Firstname, doctor.GetDoctorById.Rendez_vous_ids, doctor.GetDoctorById.Patient_ids, graphql.AddressInput{Street: doctor.GetDoctorById.Address.Street, Zip_code: doctor.GetDoctorById.Address.Zip_code, Country: doctor.GetDoctorById.Address.Country}, append(doctor.GetDoctorById.Chat_ids, newChat.CreateChat.Id))
		if err != nil {
			return CreateChatResponse{Chat: model.Chat{}, Code: 500, Err: errors.New("Unable to update doctor")}
		}

		if !containsPatientID(doctor.GetDoctorById.Patient_ids, patientID) {
			_, err = graphql.UpdateDoctor(context.Background(), gqlClient, doctor.GetDoctorById.Id, doctor.GetDoctorById.Email, doctor.GetDoctorById.Password, doctor.GetDoctorById.Name, doctor.GetDoctorById.Firstname, doctor.GetDoctorById.Rendez_vous_ids, append(doctor.GetDoctorById.Patient_ids, patientID), graphql.AddressInput{Street: doctor.GetDoctorById.Address.Street, Zip_code: doctor.GetDoctorById.Address.Zip_code, Country: doctor.GetDoctorById.Address.Country}, append(doctor.GetDoctorById.Chat_ids, newChat.CreateChat.Id))
			if err != nil {
				return CreateChatResponse{Chat: model.Chat{}, Code: 500, Err: errors.New("Unable to update doctor")}
			}
		}

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

func containsPatientID(patientIDs []string, patientID string) bool {
	for _, id := range patientIDs {
		if id == patientID {
			return true
		}
	}
	return false
}
