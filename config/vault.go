package config

import (
	"github.com/sirupsen/logrus"

	"context"
	"os"
	"time"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
	"github.com/mitchellh/mapstructure"
)

type GoogleAPI_KEY struct {
	API_KEY string `mapstructure:"API_KEY"`
}

type GoogleApiKey struct {
	Key string `mapstructure:"API_KEY"`
}

type PrimaryDBPassword struct {
	Key string `mapstructure:"db_password"`
}

type ReplicaOneDBPassword struct {
	Key string `mapstructure:"replica_db_password"`
}

type EmailPassword struct {
	Key string `mapstructure:"smtp_password"`
}

type JWTSecret struct {
	Key string `mapstructure:"JWT"`
}

type OpensearchPassword struct {
	Key string `mapstructure:"os_password"`
}

type RabbitMQPassword struct {
	Key string `mapstructure:"rabbitPassword"`
}

func ReadSecret[T any](secretName string, result *T) error {
	ctx := context.Background()

	client, err := vault.New(
		vault.WithAddress("http://vault:8200"),
		vault.WithRequestTimeout(30*time.Second),
	)
	if err != nil {
		logrus.Errorf("Error reading config file, %v", err)
		return err
	}

	// authenticate with a root approle
	resp, err := client.Auth.AppRoleLogin(
		ctx,
		schema.AppRoleLoginRequest{
			RoleId:   os.Getenv("APPROLE_ROLE_ID"),
			SecretId: os.Getenv("APPROLE_SECRET_ID"),
		},
		vault.WithMountPath("approle"),
	)

	if err != nil {
		logrus.Errorf("error geting os env to Vault: %v", err)
		return err
	}

	if err := client.SetToken(resp.Auth.ClientToken); err != nil {
		logrus.Errorf("error setting Vault token: %v", err)
		return err
	}

	secret, err := client.Secrets.KvV2Read(ctx, secretName, vault.WithMountPath("secret"))
	if err != nil {
		logrus.Errorf("error reading secret from Vault: %w", err)
		return err
	}

	if err := mapstructure.Decode(secret.Data.Data, result); err != nil {
		logrus.Errorf("error decoding secret data: %v", err)
		return err
	}

	return nil
}
