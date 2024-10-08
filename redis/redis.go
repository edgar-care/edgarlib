package redis

import (
	"fmt"
	"os"
	"strings"

	"github.com/edgar-care/edgarlib/v2"
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
	resp = strings.TrimSuffix(resp, "\n")

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

func DeleteKey(key string) (string, error) {
	client, err := CreateClient()
	edgarlib.CheckError(err)

	resp, err := ExecuteCommand(client, fmt.Sprintf("redis-cli DEL %s", key))
	edgarlib.CheckError(err)

	return resp, nil
}

func AddTokenToList(token string) (string, error) {
	client, err := CreateClient()
	if err != nil {
		return "", err
	}

	listKey := "user-tokens"
	resp, err := ExecuteCommand(client, fmt.Sprintf("redis-cli RPUSH %s %s", listKey, token))
	if err != nil {
		return "", err
	}

	return resp, nil
}

func RemoveTokenFromList(token string) (string, error) {
	client, err := CreateClient()
	if err != nil {
		return "", err
	}

	listKey := "user-tokens"
	resp, err := ExecuteCommand(client, fmt.Sprintf("redis-cli LREM %s 0 %s", listKey, token))
	if err != nil {
		return "", err
	}

	return resp, nil
}

func StoreUserInfoHash(key, os string, navigation string, location string, expire *int) (string, error) {
	client, err := CreateClient()
	if err != nil {
		return "", err
	}

	command := fmt.Sprintf("redis-cli HSET %s os %s navigation %s location %s", key, os, navigation, location)
	resp, err := ExecuteCommand(client, command)
	if err != nil {
		return "", err
	}

	if expire != nil {
		expireCommand := fmt.Sprintf("redis-cli EXPIRE %s %d", key, *expire)
		_, err = ExecuteCommand(client, expireCommand)
		if err != nil {
			return "", err
		}
	}

	return resp, nil
}

func GetUserInfoHash(key string) (map[string]string, error) {
	client, err := CreateClient()
	if err != nil {
		return nil, err
	}

	resp, err := ExecuteCommand(client, fmt.Sprintf("redis-cli HGETALL %s", key))
	if err != nil {
		return nil, err
	}

	if resp == "" {
		return nil, fmt.Errorf("key does not exist")
	}

	userInfo := make(map[string]string)
	lines := strings.Split(strings.TrimSpace(resp), "\n")
	if len(lines) < 2 {
		return nil, fmt.Errorf("invalid response format")
	}
	for i := 0; i < len(lines); i += 2 {
		userInfo[lines[i]] = lines[i+1]
	}

	return userInfo, nil
}
