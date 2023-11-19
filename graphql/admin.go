package graphql

import (
	"github.com/jinzhu/copier"
)

/********** Types ***********/

type Admin struct {
	Id       string `json:"id"`
	Password string `json:"password"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
}

type AdminOutput struct {
	Id       *string `json:"id"`
	Password *string `json:"password"`
	Name     *string `json:"name"`
	LastName *string `json:"lastName"`
	Email    *string `json:"email"`
}

type AdminInput struct {
	Password string `json:"password"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

/**************** GraphQL types *****************/

type getAdminByEmailResponse struct {
	Content AdminOutput `json:"getAdminByEmail"`
}

type createAdminResponse struct {
	Content AdminOutput `json:"createAdmin"`
}

/*************** Implementations *****************/

func GetAdminByEmail(email string) (Admin, error) {
	var admin getAdminByEmailResponse
	var resp Admin
	query := `query getAdminByEmail($email: String!) {
                getAdminByEmail(email: $email) {
                    id,
                    password,
                    name,
                    lastName,
                    email
                }
            }`

	err := Query(query, map[string]interface{}{
		"email": email,
	}, &admin)
	_ = copier.Copy(&resp, &admin.Content)
	return resp, err
}

func CreateAdmin(newAdmin AdminInput) (Admin, error) {
	var admin createAdminResponse
	var resp Admin
	query := `mutation createAdmin($email: String!, $password: String!, $name: String!, $lastName: String!) {
        createAdmin(email:$email, password:$password, name:$name, lastName:$lastName) {
                    id,
                    name,
                    lastName,
                    email,
                    password,
                }
            }`
	err := Query(query, map[string]interface{}{
		"email":    newAdmin.Email,
		"name":     newAdmin.Name,
		"lastName": newAdmin.LastName,
		"password": newAdmin.Password,
	}, &admin)
	_ = copier.Copy(&resp, &admin.Content)
	return resp, err
}
