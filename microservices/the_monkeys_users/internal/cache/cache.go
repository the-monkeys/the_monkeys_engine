package cache

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/the-monkeys/the_monkeys/microservices/the_monkeys_users/internal/database"
	"github.com/the-monkeys/the_monkeys/microservices/the_monkeys_users/internal/models"
)

func AddUserLog(dbConn database.UserDb, user *models.UserLogs, activity, service, event string, logger *logrus.Logger) {
	err := dbConn.CreateUserLog(user, fmt.Sprintf(activity, user.AccountId))
	if err != nil {
		logger.Errorf("failed to store user registration log: %v. service: %s, method: %s", err, service, event)
	}
}
