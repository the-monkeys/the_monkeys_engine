package db

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/the-monkeys/the_monkeys/common"
	"github.com/the-monkeys/the_monkeys/config"
	"github.com/the-monkeys/the_monkeys/microservices/service_types"
	"github.com/the-monkeys/the_monkeys/microservices/the_monkeys_authz/internal/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthDBHandler interface {
	// Get Operations
	CheckIfEmailExist(email string) (*models.TheMonkeysUser, error)
	CheckIfUsernameExist(username string) (*models.TheMonkeysUser, error)

	// Create Operations
	RegisterUser(user *models.TheMonkeysUser) (int64, error)

	// Update Operations
	UpdatePasswordRecoveryToken(hash string, req *models.TheMonkeysUser) error
	UpdatePassword(password string, user *models.TheMonkeysUser) error

	// UpdateEmailVerToken(user models.TheMonkeysUser) error
	// GetNamesEmailFromEmail(req *pb.ForgotPasswordReq) (*models.TheMonkeysUser, error)

}
type authDBHandler struct {
	db *sql.DB
}

func NewAuthDBHandler(cfg *config.Config) (AuthDBHandler, error) {
	url := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.Postgresql.PrimaryDB.DBUsername,
		cfg.Postgresql.PrimaryDB.DBPassword,
		cfg.Postgresql.PrimaryDB.DBHost,
		cfg.Postgresql.PrimaryDB.DBPort,
		cfg.Postgresql.PrimaryDB.DBName,
	)

	dbPsql, err := sql.Open("postgres", url)
	if err != nil {
		logrus.Fatalf("cannot connect psql using sql driver, error:, %+v", err)
		return nil, err
	}

	if err = dbPsql.Ping(); err != nil {
		logrus.Errorf("ping test failed to psql using sql driver, error: %+v", err)
		return nil, err

	}

	return &authDBHandler{db: dbPsql}, nil
}

// TODO: Find all the fields of models.TheMonkeysUser
func (adh *authDBHandler) CheckIfEmailExist(email string) (*models.TheMonkeysUser, error) {
	var tmu models.TheMonkeysUser
	if err := adh.db.QueryRow(`
			SELECT ua.user_id, ua.profile_id, ua.username, ua.first_name, ua.last_name, 
			uai.email_id, uai.password_hash, evs.ev_status, us.usr_status, uai.email_validation_token,
			uai.email_verification_timeout
			FROM USER_ACCOUNT ua
			LEFT JOIN USER_AUTH_INFO uai ON ua.user_id = uai.user_id
			LEFT JOIN email_validation_status evs ON uai.email_validation_status = evs.id
			LEFT JOIN user_status us ON ua.user_status = us.id
			WHERE uai.email_id = $1;
		`, email).
		Scan(&tmu.Id, &tmu.ProfileId, &tmu.Username, &tmu.FirstName, &tmu.LastName, &tmu.Email, &tmu.Password,
			&tmu.EmailVerificationStatus, &tmu.UserStatus, &tmu.EmailVerificationToken, &tmu.EmailVerificationTimeout); err != nil {
		logrus.Errorf("can't find a user existing with email %s, error: %+v", email, err)
		return nil, err
	}

	return &tmu, nil
}

func (adh *authDBHandler) RegisterUser(user *models.TheMonkeysUser) (int64, error) {
	tx, err := adh.db.Begin()
	if err != nil {
		return 0, err
	}

	userId, err := adh.insertIntoUserAccount(tx, user)
	if err != nil {
		return 0, err
	}

	authId, err := adh.insertIntoUserAuthInfo(tx, user, userId)
	if err != nil {
		return 0, err
	}

	// USER_ACCOUNT_STATUS

	err = adh.InsertIntoUserLog(tx, &models.TheMonkeysUser{Id: userId}, service_types.EventRegister,
		service_types.ServiceAuth, "", fmt.Sprintf("User containing %v email has called registered API", user.Email))
	if err != nil {
		return 0, err
	}

	// EXTERNAL_AUTH_PROVIDERS

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	if userId != authId {
		logrus.Warnf("we are detecting some data inconsistency for user %s", user.Email)
	}

	return userId, nil
}

func (adh *authDBHandler) insertIntoUserAccount(tx *sql.Tx, user *models.TheMonkeysUser) (int64, error) {
	stmt, err := tx.Prepare(`INSERT INTO USER_ACCOUNT (
		profile_id, username, first_name, last_name, 
		role_id, user_status) VALUES ($1, $2, $3, $4, $5, $6) RETURNING user_id;`)
	if err != nil {
		logrus.Errorf("cannot prepare statement to add user into the USER_ACCOUNT: %v", err)
		return 0, err
	}
	defer stmt.Close()

	var profileId int64
	err = stmt.QueryRow(user.ProfileId, user.Username, user.FirstName, user.LastName, 2, 1).Scan(&profileId)
	if err != nil {
		logrus.Errorf("cannot execute query to add user to the USER_ACCOUNT: %v", err)
		return 0, err
	}

	return profileId, nil
}

func (adh *authDBHandler) insertIntoUserAuthInfo(tx *sql.Tx, user *models.TheMonkeysUser, profileId int64) (int64, error) {
	stmt, err := tx.Prepare(fmt.Sprintf(`
	INSERT INTO USER_AUTH_INFO (
	user_id, username, email_id, password_hash, 
	email_validation_token, email_validation_status, email_verification_timeout, auth_provider_id) 
	VALUES ($1, $2, $3, $4, $5, (SELECT id FROM email_validation_status where ev_status='unverified' LIMIT 1), $6, (SELECT id FROM auth_provider where provider_name='%s' LIMIT 1)) 
	RETURNING id;
	`, user.LoginMethod))
	if err != nil {
		logrus.Errorf("cannot prepare statement to add user into the USER_AUTH_INFO: %v", err)
		return 0, err
	}
	defer stmt.Close()

	var authId int64
	err = stmt.QueryRow(profileId, user.Username, user.Email, user.Password,
		user.EmailVerificationToken, user.EmailVerificationTimeout).Scan(&authId)
	if err != nil {
		logrus.Errorf("cannot execute query to add user to the USER_AUTH_INFO: %v", err)
		return 0, err
	}

	return authId, nil
}

func (adh *authDBHandler) UpdatePasswordRecoveryToken(hash string, req *models.TheMonkeysUser) error {
	// TODO: start a database transaction from here till all the process are complete
	tx, err := adh.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(`UPDATE user_auth_info SET pwd_recovery_token=$1,
	pwd_recovery_timeout=$2, pwd_recovery_time=$3 WHERE email_id=$4;`)
	if err != nil {
		logrus.Errorf("cannot prepare the reset link for %s, error: %v", req.Email, err)
		return status.Errorf(codes.Internal, "internal server error, error: %v", err)
	}

	defer stmt.Close()
	result := stmt.QueryRow(hash, time.Now().Add(time.Minute*5), time.Now().Format(common.DATE_TIME_FORMAT), req.Email)
	if result.Err() != nil {
		logrus.Errorf("cannot sent the reset link for %s, error: %v", req.Email, err)
		return status.Errorf(codes.Internal, "internal server error, error: %v", err)
	}
     
	err = tx.Commit()
	if err != nil {
		logrus.Errorf("cannot commit the password recovery token for %s, error: %v", req.Email, err)
		return err
	}
	return nil
}

func (adh *authDBHandler) CheckIfUsernameExist(username string) (*models.TheMonkeysUser, error) {
	var tmu models.TheMonkeysUser
	if err := adh.db.QueryRow(`
			SELECT ua.user_id, ua.profile_id, ua.username, ua.first_name, ua.last_name, 
			uai.email_id, uai.password_hash, uai.pwd_recovery_token, uai.pwd_recovery_timeout,
			evs.ev_status, us.usr_status, uai.email_validation_token, uai.email_verification_timeout
			FROM USER_ACCOUNT ua
			LEFT JOIN USER_AUTH_INFO uai ON ua.user_id = uai.user_id
			LEFT JOIN email_validation_status evs ON uai.email_validation_status = evs.id
			LEFT JOIN user_status us ON ua.user_status = us.id
			WHERE ua.username = $1;
		`, username).
		Scan(&tmu.Id, &tmu.ProfileId, &tmu.Username, &tmu.FirstName, &tmu.LastName, &tmu.Email,
			&tmu.Password, &tmu.PasswordVerificationToken, &tmu.PasswordVerificationTimeout,
			&tmu.EmailVerificationStatus, &tmu.UserStatus, &tmu.EmailVerificationToken,
			&tmu.EmailVerificationTimeout); err != nil {
		logrus.Errorf("can't find a user existing with username %s, error: %+v", username, err)
		return nil, err
	}

	return &tmu, nil
}

func (adh *authDBHandler) UpdatePassword(password string, user *models.TheMonkeysUser) error {
	tx, err := adh.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(`UPDATE user_auth_info SET
	password_hash=$1 WHERE email_id=$2 AND username = $3 RETURNING user_id;`)
	if err != nil {
		logrus.Errorf("cannot prepare statement to update password for %s error: %+v", user.Email, err)
		return err
	}
	defer stmt.Close()

	var userId int64
	err = stmt.QueryRow(password, user.Email, user.Username).Scan(&userId)
	if err != nil {
		logrus.Errorf("cannot update the password for %s, error: %v", user.Email, err)
		return err
	}

	fmt.Printf("userId: %v\n", userId)
	// TODO: Add a record into the log table using the userId

	err = tx.Commit()
	if err != nil {
		logrus.Errorf("cannot commit the password update for %s, error: %v", user.Email, err)
		return err
	}
	return nil
}

func (adh *authDBHandler) InsertIntoUserLog(tx *sql.Tx, user *models.TheMonkeysUser, eventType, serviceType, ipAddress, description string) error {
	stmt, err := tx.Prepare(`INSERT INTO USER_ACCOUNT_LOG (user_id, event_type, service_type, ip_address, description) VALUES ($1, $2, $3, $4, $5);`)
	if err != nil {
		logrus.Errorf("cannot prepare statement to add user into the USER_ACCOUNT_LOG: %v", err)
		return err
	}
	defer stmt.Close()

	row, err := stmt.Exec(user.Id, eventType, serviceType, ipAddress, description)
	if err != nil {
		logrus.Errorf("cannot execute query to add user to the USER_ACCOUNT_LOG: %v", err)
		return err
	}

	affectedRow, err := row.RowsAffected()
	if err != nil {
		logrus.Errorf("error finding affected rows for USER_ACCOUNT_LOG: %v", err)
		return err
	}

	if affectedRow == 0 {
		logrus.Errorf("cannot create a record in the log table for USER_ACCOUNT_LOG: %v", err)
		return errors.New("cannot create a record in the log table")
	}

	return nil
}
