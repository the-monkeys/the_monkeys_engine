package database

import (
	"errors"

	"github.com/sirupsen/logrus"
	"github.com/the-monkeys/the_monkeys/microservices/the_monkeys_users/internal/models"
)

func (userDB *uDBHandler) CreateUserLog(user *models.UserLogs, description string) error {
	var userId int64
	var clientId int8
	var err error

	//From username find user id
	if err = userDB.db.QueryRow(`
			SELECT id FROM user_account WHERE account_id = $1;`, user.AccountId).Scan(&userId); err != nil {
		logrus.Errorf("can't get id by using account_id %s, error: %v", user.AccountId, err)
		return err
	}

	//From client name find client id
	if err := userDB.db.QueryRow(`
			SELECT id FROM clients WHERE c_name = $1;`, user.Client).Scan(&clientId); err != nil {
		logrus.Errorf("can't get id by using client name %s, error: %+v", user.Client, err)
		return err
	}

	stmt, err := userDB.db.Prepare(`INSERT INTO user_account_log (user_id, ip_address, description, client_id) VALUES ($1, $2, $3, $4);`)
	if err != nil {
		logrus.Errorf("cannot prepare statement to add user activity into the user_account_log: %v", err)
		return err
	}
	defer stmt.Close()

	row, err := stmt.Exec(userId, user.IpAddress, description, clientId)
	if err != nil {
		logrus.Errorf("cannot execute query to add user to the user_account_log: %v", err)
		return err
	}

	affectedRow, err := row.RowsAffected()
	if err != nil {
		logrus.Errorf("error finding affected rows for user_account_log: %v", err)
		return err
	}

	if affectedRow == 0 {
		logrus.Errorf("cannot create a record in the log table for user_account_log: %v", err)
		return errors.New("cannot create a record in the log table")
	}

	return nil
}

func (uh *uDBHandler) UpdateUserName(currentUsername, newUsername string) error {
	stmt, err := uh.db.Prepare(`UPDATE user_account SET username = $1 WHERE username = $2`)
	if err != nil {
		uh.log.Errorf("cannot prepare statement to update username, error: %v", err)
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(newUsername, currentUsername)
	if err != nil {
		uh.log.Errorf("cannot execute update username query, error: %v", err)
		return err
	}

	row, err := res.RowsAffected()
	if err != nil {
		logrus.Errorf("error while checking rows affected for update username query, error: %v", err)
		return err
	}
	if row != 1 {
		logrus.Errorf("more or less than 1 row is affected for update username query, error: %v", err)
		return errors.New("more or less than 1 row is affected")
	}

	return nil
}
