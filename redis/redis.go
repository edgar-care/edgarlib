package redis

import (
	"fmt"
	"os"

	"github.com/edgar-care/edgarlib"
	"golang.org/x/crypto/ssh"
)

func ExecuteCommand(client *ssh.Client, command string) (string, error) {
	session, err := client.NewSession()
	edgarlib.CheckError(err)
	defer session.Close()

	output, err := session.CombinedOutput(command)
	edgarlib.CheckError(err)
	return string(output), nil
}

func CreateClient() (*ssh.Client, error) {
	config := &ssh.ClientConfig{
		User: os.Getenv("REDIS_VM_USERNAME"),
		Auth: []ssh.AuthMethod{
			ssh.Password(os.Getenv("REDIS_VM_PASSWORD")),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%s", os.Getenv("REDIS_VM_HOSTNAME"), os.Getenv("REDIS_VM_PORT")), config)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the VM: %v", err)
	}

	return client, nil
}

func GetKey(key string) (string, error) {
	client, err := CreateClient()
	edgarlib.CheckError(err)

	resp, err := ExecuteCommand(client, fmt.Sprintf("redis-cli GET %s", key))
	edgarlib.CheckError(err)

	return resp, nil
}

func SetKey(key string, value string, expire *int) (string, error) {
	client, err := CreateClient()
	edgarlib.CheckError(err)

	command := fmt.Sprintf("redis-cli SET %s %s", key, value)
	if expire != nil {
		command += fmt.Sprintf(" EX %d", *expire)
	}
	resp, err := ExecuteCommand(client, command)
	edgarlib.CheckError(err)

	return resp, nil

}
