package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type TheMonkeysGateway struct {
	HTTPS string `mapstructure:"HTTPS"`
	HTTP  string `mapstructure:"HTTP"`
}

type Microservices struct {
	TheMonkeysAuthz     string `mapstructure:"the_monkeys_authz"`
	TheMonkeysBlog      string `mapstructure:"the_monkeys_blog"`
	TheMonkeysUser      string `mapstructure:"the_monkeys_user"`
	TheMonkeysFileStore string `mapstructure:"the_monkeys_file_storage"`
}

type Database struct {
	DBUsername string `mapstructure:"db_username"`
	DBPassword string `mapstructure:"db_password"`
	DBHost     string `mapstructure:"db_host"`
	DBPort     int    `mapstructure:"db_port"`
	DBName     string `mapstructure:"db_name"`
}

type Postgresql struct {
	PrimaryDB Database `mapstructure:"primary_db"`
	Replica1  Database `mapstructure:"replica_1"`
}

type JWT struct {
	SecretKey string `mapstructure:"secret_key"`
}

type Opensearch struct {
	Address  string `mapstructure:"address"`
	Host     string `mapstructure:"os_host"`
	Username string `mapstructure:"os_username"`
	Password string `mapstructure:"os_password"`
}

type Email struct {
	SMTPAddress  string `mapstructure:"smtp_address"`
	SMTPMail     string `mapstructure:"smtp_mail"`
	SMTPPassword string `mapstructure:"smtp_password"`
	SMTPHost     string `mapstructure:"smtp_host"`
}

type Authentication struct {
	EmailVerificationAddr string `mapstructure:"email_verification_addr"`
}

type RabbitMQ struct {
	Protocol    string   `mapstructure:"protocol"`
	Host        string   `mapstructure:"host"`
	Port        string   `mapstructure:"port"`
	Username    string   `mapstructure:"username"`
	Password    string   `mapstructure:"password"`
	VirtualHost string   `mapstructure:"virtual_host"`
	Exchange    string   `mapstructure:"exchange"`
	Queues      []string `yaml:"queues"`
	RoutingKeys []string `yaml:"routingKeys"`
}

type Config struct {
	TheMonkeysGateway TheMonkeysGateway `mapstructure:"the_monkeys_gateway"`
	Microservices     Microservices     `mapstructure:"microservices"`
	Postgresql        Postgresql        `mapstructure:"postgresql"`
	JWT               JWT               `mapstructure:"jwt"`
	Opensearch        Opensearch        `mapstructure:"opensearch"`
	Email             Email             `mapstructure:"email"`
	Authentication    Authentication    `mapstructure:"authentication"`
	RabbitMQ          RabbitMQ          `mapstructure:"rabbitMQ"`
	GoogleApiKey      GoogleAPI_KEY     `mapstructure:"google"`
}

// TODO: remove the print statement and add logger instead
func GetConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")

	config := &Config{}
	if err := viper.ReadInConfig(); err != nil {
		logrus.Errorf("Error reading config file, %v", err)
		return config, err
	}

	if err := viper.Unmarshal(config); err != nil {
		logrus.Errorf("Unable to decode into struct, %v", err)
		return config, err
	}

	logrus.Infof("Configuration file loaded: %+v", config) // Add this line to print the loaded configuration

	var googleApiKey GoogleApiKey
	err := ReadSecret("google", &googleApiKey)
	if err != nil {
		logrus.Errorf("Error reading secret from googleApiKey: %v", err)
		return config, err
	}

	var primaryDBPassword PrimaryDBPassword
	err = ReadSecret("db_password", &primaryDBPassword)
	if err != nil {
		logrus.Errorf("Error reading secret from primaryDBPassword: %v", err)
		return config, err
	}

	var replicaOneDBPassword ReplicaOneDBPassword
	err = ReadSecret("replica_db_password", &replicaOneDBPassword)
	if err != nil {
		logrus.Errorf("Error reading secret from replicaOneDBPassword: %v", err)
		return config, err
	}

	var jWTSecret JWTSecret
	err = ReadSecret("JWT", &jWTSecret)
	if err != nil {
		logrus.Errorf("Error reading secret from jWTSecret: %v", err)
		return config, err
	}

	var opensearchPassword OpensearchPassword
	err = ReadSecret("os_password", &opensearchPassword)
	if err != nil {
		logrus.Errorf("Error reading secret from opensearchPassword: %v", err)
		return config, err
	}

	var emailPassword EmailPassword
	err = ReadSecret("smtp_password", &emailPassword)
	if err != nil {
		logrus.Errorf("Error reading secret from emailPassword: %v", err)
		return config, err
	}

	var rabbitMQPassword RabbitMQPassword
	err = ReadSecret("rabbitPassword", &rabbitMQPassword)
	if err != nil {
		logrus.Errorf("Error reading secret from rabbitMQPassword: %v", err)
		return config, err
	}

	config.JWT.SecretKey = jWTSecret.Key
	config.Postgresql.PrimaryDB.DBPassword = primaryDBPassword.Key
	config.Postgresql.Replica1.DBPassword = replicaOneDBPassword.Key
	config.Opensearch.Password = opensearchPassword.Key
	config.GoogleApiKey.API_KEY = googleApiKey.Key
	config.RabbitMQ.Password = rabbitMQPassword.Key
	config.Email.SMTPPassword = emailPassword.Key

	logrus.Infof("config vlaue %v", config)

	return config, nil
}
