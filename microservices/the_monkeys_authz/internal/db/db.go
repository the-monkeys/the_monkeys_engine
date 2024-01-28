package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/the-monkeys/the_monkeys/config"
	"github.com/the-monkeys/the_monkeys/microservices/the_monkeys_authz/internal/models"
)

type AuthDBHandler interface {
	RegisterUser(user *models.TheMonkeysUser) (int64, error)
	// UpdateEmailVerToken(user models.TheMonkeysUser) error
	// GetNamesEmailFromEmail(req *pb.ForgotPasswordReq) (*models.TheMonkeysUser, error)
	// UpdatePasswordRecoveryToken(hash string, req *models.TheMonkeysUser) error
	// UpdatePassword(password, email string) error
	CheckIfEmailExist(email string) (*models.TheMonkeysUser, error)
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
		SELECT ua.user_id, ua.username, uai.email_id, uai.password_hash, evs.ev_status
		FROM USER_ACCOUNT ua
		LEFT JOIN USER_AUTH_INFO uai ON ua.user_id = uai.user_id
		LEFT JOIN email_validation_status evs ON uai.email_validation_status = evs.id
		WHERE uai.email_id = $1;
		`, email).
		Scan(&tmu.Id, &tmu.Username, &tmu.Email, &tmu.Password, &tmu.EmailVerified); err != nil {
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

	profileId, err := adh.insertIntoUserAccount(tx, user)
	if err != nil {
		return 0, err
	}

	authId, err := adh.insertIntoUserAuthInfo(tx, user, profileId)
	if err != nil {
		return 0, err
	}

	// USER_ACCOUNT_STATUS

	// USER_ACCOUNT_LOG

	// EXTERNAL_AUTH_PROVIDERS

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	if profileId != authId {
		logrus.Warnf("we are detecting some data inconsistency for user %s", user.Email)
	}

	return profileId, nil
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
	VALUES ($1, $2, $3, $4, $5, (SELECT id FROM email_validation_status where ev_status='%s'), $6, (SELECT id FROM auth_provider where provider_name='%s')) 
	RETURNING id;
	`, "unverified", user.LoginMethod))
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
