package consumer

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/sirupsen/logrus"
	"github.com/the-monkeys/the_monkeys/config"
	"github.com/the-monkeys/the_monkeys/constants"
	"github.com/the-monkeys/the_monkeys/microservices/rabbitmq"
	"github.com/the-monkeys/the_monkeys/microservices/the_monkeys_file_storage/constant"
	"github.com/the-monkeys/the_monkeys/microservices/the_monkeys_file_storage/internal/models"
)

func ConsumeFromQueue(conn rabbitmq.Conn, conf config.RabbitMQ, log *logrus.Logger) {

	// conn, err := queue.GetConn(conf)
	// if err != nil {
	// 	log.Fatalf("Error establishing RabbitMQ connection: %v", err)
	// }

	// defer conn.Channel.Close()

	// Set up signal handling for graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		logrus.Infoln("Received termination signal. Closing connection and exiting gracefully.")
		conn.Channel.Close()
		os.Exit(0)
	}()

	msgs, err := conn.Channel.Consume(
		conf.Queues[0], // queue
		"",             // consumer
		true,           // auto-ack
		false,          // exclusive
		false,          // no-local
		false,          // no-wait
		nil,            // args
	)
	if err != nil {
		logrus.Errorf("Failed to register a consumer: %v", err)
		return
	}

	for d := range msgs {
		user := models.TheMonkeysMessage{}
		if err = json.Unmarshal(d.Body, &user); err != nil {
			logrus.Errorf("Failed to unmarshal user from rabbitMQ: %v", err)
			return
		}

		switch user.Action {
		case constants.USER_PROFILE_DIRECTORY_CREATE:
			if err := CreateUserFolder(user.Username); err != nil {
				logrus.Errorf("Failed to create user folder: %v", err)
			}
		case constants.USER_PROFILE_DIRECTORY_UPDATE:
			if err := UpdateUserFolder(user.Username, user.NewUsername); err != nil {
				logrus.Errorf("Failed to update user folder: %v", err)
			}

		case constants.USER_PROFILE_DIRECTORY_DELETE:
			if err := DeleteUserFolder(user.Username); err != nil {
				logrus.Errorf("Failed to delete user folder: %v", err)
			}
		default:
			logrus.Errorf("Unknown action: %s", user.Action)
		}

	}
}

func CreateUserFolder(userName string) error {
	dirPath, filePath := ConstructPath(constant.ProfileDir, userName, "profile.png")

	// Create directory if it doesn't exist
	err := os.MkdirAll(dirPath, 0755)
	if err != nil {
		logrus.Errorf("Cannot create directory structure for user: %s, error: %v", userName, err)
		return err
	}

	imageByte, err := readImageFromURL(constant.DefaultProfilePhoto)
	if err != nil {
		logrus.Errorf("Error fetching image for user: %s, error: %v", userName, err)
		return fmt.Errorf("error fetching image: %v", err)
	}

	// Write image data to file
	err = os.WriteFile(filePath, imageByte, 0644)
	if err != nil {
		logrus.Errorf("Cannot write profile image file for user: %s, error: %v", userName, err)
		return err
	}

	logrus.Infof("Done uploading profile pic: %s", filePath)
	return nil
}

func readImageFromURL(url string) ([]byte, error) {
	// Create a new HTTP client
	client := http.Client{}

	// Send a GET request to the URL
	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %v", err)
	}
	defer resp.Body.Close()

	// Check for successful response status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP request failed with status code %d", resp.StatusCode)
	}

	// Read the response body into a byte array
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading response body: %v", err)
	}

	return data, nil
}

func ConstructPath(baseDir, userName, fileName string) (string, string) {
	dirPath := filepath.Join(baseDir, userName)
	filePath := filepath.Join(dirPath, fileName)
	return dirPath, filePath
}

// UpdateUserFolder renames a folder to a new name
func UpdateUserFolder(currentName, newName string) error {
	currentPath := filepath.Join(constant.ProfileDir, currentName)
	newPath := filepath.Join(constant.ProfileDir, newName)

	log.Printf("updating user folder %s to %s", currentName, newName)

	from, err := os.Stat(currentPath)
	if err != nil {
		return errors.New("could not stat current directory: " + err.Error())
	}

	if !from.IsDir() {
		return errors.New(currentPath + " is not a directory")
	}

	to := currentPath + "_temp" // Create temporary name

	// Rename the directory
	err = os.Rename(currentPath, to)
	if err != nil {
		return errors.New("failed to rename directory: " + err.Error())
	}

	// Wait for a bit to make sure the directory is released
	// time.Sleep(1 * time.Second)

	// Rename back to the desired name
	err = os.Rename(to, newPath)
	if err != nil {
		return errors.New("failed to rename directory to new name: " + err.Error())
	}

	return nil
}

func DeleteUserFolder(userName string) error {
	dirPath := filepath.Join(constant.ProfileDir, userName)

	// Remove the directory
	err := os.RemoveAll(dirPath)
	if err != nil {
		return errors.New("failed to remove directory: " + err.Error())
	}

	return nil
}
