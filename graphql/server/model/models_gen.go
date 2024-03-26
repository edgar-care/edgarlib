// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Address struct {
	Street  string `json:"street" bson:"street"`
	ZipCode string `json:"zip_code" bson:"zip_code"`
	Country string `json:"country" bson:"country"`
	City    string `json:"city" bson:"city"`
}

type AddressInput struct {
	Street  string `json:"street" bson:"street"`
	ZipCode string `json:"zip_code" bson:"zip_code"`
	Country string `json:"country" bson:"country"`
	City    string `json:"city" bson:"city"`
}

type Admin struct {
	ID       string `json:"id" bson:"_id"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
	Name     string `json:"name" bson:"name"`
	LastName string `json:"last_name" bson:"last_name"`
}

type Alert struct {
	ID       string   `json:"id" bson:"_id"`
	Name     string   `json:"name" bson:"name"`
	Sex      *string  `json:"sex,omitempty" bson:"sex"`
	Height   *int     `json:"height,omitempty" bson:"height"`
	Weight   *int     `json:"weight,omitempty" bson:"weight"`
	Symptoms []string `json:"symptoms" bson:"symptoms"`
	Comment  string   `json:"comment" bson:"comment"`
}

type AnteChir struct {
	ID              string   `json:"id" bson:"_id"`
	Name            string   `json:"name" bson:"name"`
	Localisation    string   `json:"localisation" bson:"localisation"`
	InducedSymptoms []string `json:"induced_symptoms,omitempty" bson:"induced_symptoms"`
}

type AnteDisease struct {
	ID            string   `json:"id" bson:"_id"`
	Name          string   `json:"name" bson:"name"`
	Chronicity    float64  `json:"chronicity" bson:"chronicity"`
	SurgeryIds    []string `json:"surgery_ids,omitempty" bson:"surgery_ids"`
	Symptoms      []string `json:"symptoms,omitempty" bson:"symptoms"`
	TreatmentIds  []string `json:"treatment_ids,omitempty" bson:"treatment_ids"`
	StillRelevant bool     `json:"still_relevant" bson:"still_relevant"`
}

type AnteFamily struct {
	ID      string   `json:"id" bson:"_id"`
	Name    string   `json:"name" bson:"name"`
	Disease []string `json:"disease" bson:"disease"`
}

type DemoAccount struct {
	ID       string `json:"id" bson:"_id"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

type Disease struct {
	ID               string           `json:"id" bson:"_id"`
	Code             string           `json:"code" bson:"code"`
	Name             string           `json:"name" bson:"name"`
	Symptoms         []string         `json:"symptoms" bson:"symptoms"`
	SymptomsAcute    []*SymptomWeight `json:"symptoms_acute,omitempty" bson:"symptoms_acute"`
	SymptomsSubacute []*SymptomWeight `json:"symptoms_subacute,omitempty" bson:"symptoms_subacute"`
	SymptomsChronic  []*SymptomWeight `json:"symptoms_chronic,omitempty" bson:"symptoms_chronic"`
	Advice           *string          `json:"advice,omitempty" bson:"advice"`
}

type Doctor struct {
	ID            string    `json:"id" bson:"_id"`
	Email         string    `json:"email" bson:"email"`
	Password      string    `json:"password" bson:"password"`
	Name          string    `json:"name" bson:"name"`
	Firstname     string    `json:"firstname" bson:"firstname"`
	Address       *Address  `json:"address" bson:"address"`
	RendezVousIds []*string `json:"rendez_vous_ids,omitempty" bson:"rendez_vous_ids"`
	PatientIds    []*string `json:"patient_ids,omitempty" bson:"patient_ids"`
}

type Document struct {
	ID           string       `json:"id" bson:"_id"`
	OwnerID      string       `json:"owner_id" bson:"owner_id"`
	Name         string       `json:"name" bson:"name"`
	DocumentType DocumentType `json:"document_type" bson:"document_type"`
	Category     Category     `json:"category" bson:"category"`
	IsFavorite   bool         `json:"is_favorite" bson:"is_favorite"`
	DownloadURL  string       `json:"download_url" bson:"download_url"`
}

type Logs struct {
	Question string `json:"question" bson:"question"`
	Answer   string `json:"answer" bson:"answer"`
}

type LogsInput struct {
	Question string `json:"question" bson:"question"`
	Answer   string `json:"answer" bson:"answer"`
}

type MedicalAntecedents struct {
	ID            string       `json:"id" bson:"_id"`
	Name          string       `json:"name" bson:"name"`
	Medicines     []*Treatment `json:"medicines" bson:"medicines"`
	StillRelevant bool         `json:"still_relevant" bson:"still_relevant"`
}

type MedicalAntecedentsInput struct {
	Name          string            `json:"name" bson:"name"`
	Medicines     []*TreatmentInput `json:"medicines" bson:"medicines"`
	StillRelevant bool              `json:"still_relevant" bson:"still_relevant"`
}

type MedicalInfo struct {
	ID                   string           `json:"id" bson:"_id"`
	Name                 string           `json:"name" bson:"name"`
	Firstname            string           `json:"firstname" bson:"firstname"`
	Birthdate            int              `json:"birthdate" bson:"birthdate"`
	Sex                  Sex              `json:"sex" bson:"sex"`
	Height               int              `json:"height" bson:"height"`
	Weight               int              `json:"weight" bson:"weight"`
	PrimaryDoctorID      string           `json:"primary_doctor_id" bson:"primary_doctor_id"`
	OnboardingStatus     OnboardingStatus `json:"onboarding_status" bson:"onboarding_status"`
	AntecedentDiseaseIds []string         `json:"antecedent_disease_ids" bson:"antecedent_disease_ids"`
}

type Medicament struct {
	ID              string       `json:"id" bson:"_id"`
	Name            string       `json:"name" bson:"name"`
	Unit            MedicineUnit `json:"unit" bson:"unit"`
	TargetDiseases  []string     `json:"target_diseases" bson:"target_diseases"`
	TreatedSymptoms []string     `json:"treated_symptoms" bson:"treated_symptoms"`
	SideEffects     []string     `json:"side_effects" bson:"side_effects"`
}

type MedicamentInput struct {
	Name            string       `json:"name" bson:"name"`
	Unit            MedicineUnit `json:"unit" bson:"unit"`
	TargetDiseases  []string     `json:"target_diseases" bson:"target_diseases"`
	TreatedSymptoms []string     `json:"treated_symptoms" bson:"treated_symptoms"`
	SideEffects     []string     `json:"side_effects" bson:"side_effects"`
}

type Mutation struct {
}

type Notification struct {
	ID      string `json:"id" bson:"_id"`
	Token   string `json:"token" bson:"token"`
	Title   string `json:"title" bson:"title"`
	Message string `json:"message" bson:"message"`
}

type Patient struct {
	ID            string    `json:"id" bson:"_id"`
	Email         string    `json:"email" bson:"email"`
	Password      string    `json:"password" bson:"password"`
	RendezVousIds []*string `json:"rendez_vous_ids,omitempty" bson:"rendez_vous_ids"`
	MedicalInfoID *string   `json:"medical_info_id,omitempty" bson:"medical_info_id"`
	DocumentIds   []*string `json:"document_ids,omitempty" bson:"document_ids"`
}

type Query struct {
}

type Rdv struct {
	ID                string            `json:"id" bson:"_id"`
	DoctorID          string            `json:"doctor_id" bson:"doctor_id"`
	IDPatient         string            `json:"id_patient" bson:"id_patient"`
	StartDate         int               `json:"start_date" bson:"start_date"`
	EndDate           int               `json:"end_date" bson:"end_date"`
	CancelationReason *string           `json:"cancelation_reason,omitempty" bson:"cancelation_reason"`
	AppointmentStatus AppointmentStatus `json:"appointment_status" bson:"appointment_status"`
	SessionID         string            `json:"session_id" bson:"session_id"`
}

type Session struct {
	ID           string             `json:"id" bson:"_id"`
	Diseases     []*SessionDiseases `json:"diseases" bson:"diseases"`
	Symptoms     []*SessionSymptom  `json:"symptoms" bson:"symptoms"`
	Age          int                `json:"age" bson:"age"`
	Height       int                `json:"height" bson:"height"`
	Weight       int                `json:"weight" bson:"weight"`
	Sex          string             `json:"sex" bson:"sex"`
	AnteChirs    []string           `json:"ante_chirs" bson:"ante_chirs"`
	AnteDiseases []string           `json:"ante_diseases" bson:"ante_diseases"`
	Treatments   []string           `json:"treatments" bson:"treatments"`
	LastQuestion string             `json:"last_question" bson:"last_question"`
	Logs         []*Logs            `json:"logs" bson:"logs"`
	Alerts       []string           `json:"alerts" bson:"alerts"`
}

type SessionDiseases struct {
	Name     string  `json:"name" bson:"name"`
	Presence float64 `json:"presence" bson:"presence"`
}

type SessionDiseasesInput struct {
	Name     string  `json:"name" bson:"name"`
	Presence float64 `json:"presence" bson:"presence"`
}

type SessionSymptom struct {
	Name     string `json:"name" bson:"name"`
	Presence *bool  `json:"presence,omitempty" bson:"presence"`
	Duration *int   `json:"duration,omitempty" bson:"duration"`
}

type SessionSymptomInput struct {
	Name     string `json:"name" bson:"name"`
	Presence *bool  `json:"presence,omitempty" bson:"presence"`
	Duration *int   `json:"duration,omitempty" bson:"duration"`
}

type Symptom struct {
	ID       string   `json:"id" bson:"_id"`
	Code     string   `json:"code" bson:"code"`
	Name     string   `json:"name" bson:"name"`
	Location *string  `json:"location,omitempty" bson:"location"`
	Duration *int     `json:"duration,omitempty" bson:"duration"`
	Acute    *int     `json:"acute,omitempty" bson:"acute"`
	Subacute *int     `json:"subacute,omitempty" bson:"subacute"`
	Chronic  *int     `json:"chronic,omitempty" bson:"chronic"`
	Symptom  []string `json:"symptom" bson:"symptom"`
	Advice   *string  `json:"advice,omitempty" bson:"advice"`
	Question string   `json:"question" bson:"question"`
}

type SymptomWeight struct {
	Key   string  `json:"key" bson:"key"`
	Value float64 `json:"value" bson:"value"`
}

type SymptomWeightInput struct {
	Key   string  `json:"key" bson:"key"`
	Value float64 `json:"value" bson:"value"`
}

type TestAccount struct {
	ID       string `json:"id" bson:"_id"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

type Treatment struct {
	ID         string   `json:"id" bson:"_id"`
	Period     []Period `json:"period" bson:"period"`
	Day        []Day    `json:"day" bson:"day"`
	Quantity   int      `json:"quantity" bson:"quantity"`
	MedicineID string   `json:"medicine_id" bson:"medicine_id"`
}

type TreatmentInput struct {
	Period     []*Period `json:"period" bson:"period"`
	Day        []*Day    `json:"day" bson:"day"`
	Quantity   int       `json:"quantity" bson:"quantity"`
	MedicineID string    `json:"medicine_id" bson:"medicine_id"`
}

type AppointmentStatus string

const (
	AppointmentStatusWaitingForReview    AppointmentStatus = "WAITING_FOR_REVIEW"
	AppointmentStatusAcceptedDueToReview AppointmentStatus = "ACCEPTED_DUE_TO_REVIEW"
	AppointmentStatusCanceledDueToReview AppointmentStatus = "CANCELED_DUE_TO_REVIEW"
	AppointmentStatusCanceled            AppointmentStatus = "CANCELED"
	AppointmentStatusOpened              AppointmentStatus = "OPENED"
)

var AllAppointmentStatus = []AppointmentStatus{
	AppointmentStatusWaitingForReview,
	AppointmentStatusAcceptedDueToReview,
	AppointmentStatusCanceledDueToReview,
	AppointmentStatusCanceled,
	AppointmentStatusOpened,
}

func (e AppointmentStatus) IsValid() bool {
	switch e {
	case AppointmentStatusWaitingForReview, AppointmentStatusAcceptedDueToReview, AppointmentStatusCanceledDueToReview, AppointmentStatusCanceled, AppointmentStatusOpened:
		return true
	}
	return false
}

func (e AppointmentStatus) String() string {
	return string(e)
}

func (e *AppointmentStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AppointmentStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AppointmentStatus", str)
	}
	return nil
}

func (e AppointmentStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Category string

const (
	CategoryGeneral Category = "GENERAL"
	CategoryFinance Category = "FINANCE"
)

var AllCategory = []Category{
	CategoryGeneral,
	CategoryFinance,
}

func (e Category) IsValid() bool {
	switch e {
	case CategoryGeneral, CategoryFinance:
		return true
	}
	return false
}

func (e Category) String() string {
	return string(e)
}

func (e *Category) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Category(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Category", str)
	}
	return nil
}

func (e Category) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Day string

const (
	DayMonday    Day = "MONDAY"
	DayTuesday   Day = "TUESDAY"
	DayWednesday Day = "WEDNESDAY"
	DayThursday  Day = "THURSDAY"
	DayFriday    Day = "FRIDAY"
	DaySaturday  Day = "SATURDAY"
	DaySunday    Day = "SUNDAY"
)

var AllDay = []Day{
	DayMonday,
	DayTuesday,
	DayWednesday,
	DayThursday,
	DayFriday,
	DaySaturday,
	DaySunday,
}

func (e Day) IsValid() bool {
	switch e {
	case DayMonday, DayTuesday, DayWednesday, DayThursday, DayFriday, DaySaturday, DaySunday:
		return true
	}
	return false
}

func (e Day) String() string {
	return string(e)
}

func (e *Day) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Day(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Day", str)
	}
	return nil
}

func (e Day) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type DocumentType string

const (
	DocumentTypeXray         DocumentType = "XRAY"
	DocumentTypePrescription DocumentType = "PRESCRIPTION"
	DocumentTypeOther        DocumentType = "OTHER"
	DocumentTypeCertificate  DocumentType = "CERTIFICATE"
)

var AllDocumentType = []DocumentType{
	DocumentTypeXray,
	DocumentTypePrescription,
	DocumentTypeOther,
	DocumentTypeCertificate,
}

func (e DocumentType) IsValid() bool {
	switch e {
	case DocumentTypeXray, DocumentTypePrescription, DocumentTypeOther, DocumentTypeCertificate:
		return true
	}
	return false
}

func (e DocumentType) String() string {
	return string(e)
}

func (e *DocumentType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = DocumentType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid DocumentType", str)
	}
	return nil
}

func (e DocumentType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type MedicineUnit string

const (
	MedicineUnitApplication MedicineUnit = "APPLICATION"
	MedicineUnitTablet      MedicineUnit = "TABLET"
	MedicineUnitTablespoon  MedicineUnit = "TABLESPOON"
	MedicineUnitCoffeespoon MedicineUnit = "COFFEESPOON"
)

var AllMedicineUnit = []MedicineUnit{
	MedicineUnitApplication,
	MedicineUnitTablet,
	MedicineUnitTablespoon,
	MedicineUnitCoffeespoon,
}

func (e MedicineUnit) IsValid() bool {
	switch e {
	case MedicineUnitApplication, MedicineUnitTablet, MedicineUnitTablespoon, MedicineUnitCoffeespoon:
		return true
	}
	return false
}

func (e MedicineUnit) String() string {
	return string(e)
}

func (e *MedicineUnit) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = MedicineUnit(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid MedicineUnit", str)
	}
	return nil
}

func (e MedicineUnit) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type OnboardingStatus string

const (
	OnboardingStatusNotStarted OnboardingStatus = "NOT_STARTED"
	OnboardingStatusInProgress OnboardingStatus = "IN_PROGRESS"
	OnboardingStatusDone       OnboardingStatus = "DONE"
)

var AllOnboardingStatus = []OnboardingStatus{
	OnboardingStatusNotStarted,
	OnboardingStatusInProgress,
	OnboardingStatusDone,
}

func (e OnboardingStatus) IsValid() bool {
	switch e {
	case OnboardingStatusNotStarted, OnboardingStatusInProgress, OnboardingStatusDone:
		return true
	}
	return false
}

func (e OnboardingStatus) String() string {
	return string(e)
}

func (e *OnboardingStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OnboardingStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OnboardingStatus", str)
	}
	return nil
}

func (e OnboardingStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Period string

const (
	PeriodMorning Period = "MORNING"
	PeriodNoon    Period = "NOON"
	PeriodEvening Period = "EVENING"
	PeriodNight   Period = "NIGHT"
)

var AllPeriod = []Period{
	PeriodMorning,
	PeriodNoon,
	PeriodEvening,
	PeriodNight,
}

func (e Period) IsValid() bool {
	switch e {
	case PeriodMorning, PeriodNoon, PeriodEvening, PeriodNight:
		return true
	}
	return false
}

func (e Period) String() string {
	return string(e)
}

func (e *Period) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Period(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Period", str)
	}
	return nil
}

func (e Period) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Sex string

const (
	SexMale   Sex = "MALE"
	SexFemale Sex = "FEMALE"
	SexOther  Sex = "OTHER"
)

var AllSex = []Sex{
	SexMale,
	SexFemale,
	SexOther,
}

func (e Sex) IsValid() bool {
	switch e {
	case SexMale, SexFemale, SexOther:
		return true
	}
	return false
}

func (e Sex) String() string {
	return string(e)
}

func (e *Sex) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Sex(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Sex", str)
	}
	return nil
}

func (e Sex) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
