package database

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/the-monkeys/the_monkeys/config"

	"github.com/the-monkeys/the_monkeys/apis/serviceconn/gateway_user/pb"
	"github.com/the-monkeys/the_monkeys/microservices/the_monkeys_users/internal/models"
)

type UserDb interface {
	CheckIfEmailExist(email string) (*models.TheMonkeysUser, error)
	CheckIfUsernameExist(username string) (*models.TheMonkeysUser, error)
	GetMyProfile(email string) (*pb.UserProfileRes, error)
}

type uDBHandler struct {
	db  *sql.DB
	log *logrus.Logger
}

func NewUserDbHandler(cfg *config.Config, log *logrus.Logger) (UserDb, error) {
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

	return &uDBHandler{db: dbPsql, log: log}, nil
}

// To get User Profile
func (uh *uDBHandler) GetUserProfile(profile_id string) (*models.TheMonkeysUser, error) {
	var tmu models.TheMonkeysUser
	if err := uh.db.QueryRow(`SELECT ua.user_id, ua.profile_id, ua.username, ua.first_name, ua.last_name, 
		uai.email_id, uai.password_hash, evs.ev_status, us.usr_status, uai.email_validation_token,
		uai.email_verification_timeout
		FROM USER_ACCOUNT ua
		LEFT JOIN USER_AUTH_INFO uai ON ua.user_id = uai.user_id
		LEFT JOIN email_validation_status evs ON uai.email_validation_status = evs.id
		LEFT JOIN user_status us ON ua.user_status = us.id
		WHERE uai.profile_id = $1;
	`, profile_id).Scan(&tmu.Id, &tmu.ProfileId, &tmu.Username, &tmu.FirstName, &tmu.LastName, &tmu.Email, &tmu.Password,
		&tmu.EmailVerificationStatus, &tmu.UserStatus, &tmu.EmailVerificationToken, &tmu.EmailVerificationTimeout); err != nil {
		logrus.Errorf("can't find a user existing with this profile id  %s, error: %+v", profile_id, err)
		return nil, err
	}

	return &tmu, nil
}

// TODO: Find all the fields of models.TheMonkeysUser
func (uh *uDBHandler) CheckIfEmailExist(email string) (*models.TheMonkeysUser, error) {
	var tmu models.TheMonkeysUser
	if err := uh.db.QueryRow(`
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

func (uh *uDBHandler) CheckIfUsernameExist(username string) (*models.TheMonkeysUser, error) {
	var tmu models.TheMonkeysUser
	if err := uh.db.QueryRow(`
			SELECT ua.user_id, ua.profile_id, ua.username, ua.first_name, ua.last_name, 
			uai.email_id, uai.password_hash, uai.password_recovery_token, uai.password_recovery_timeout,
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

func (uh *uDBHandler) GetMyProfile(email string) (*pb.UserProfileRes, error) {
	var profile *pb.UserProfileRes
	if err := uh.db.QueryRow(`
			SELECT ua.profile_id, ua.username, ua.first_name, ua.last_name,  ua.date_of_birth, 
			ua.bio, ua.avatar_url, ua.created_at, ua.updated_at, ua.address,
			ua.contact_number, us.usr_status
			FROM USER_ACCOUNT ua
			LEFT JOIN USER_AUTH_INFO uai ON ua.user_id = uai.user_id
			LEFT JOIN email_validation_status evs ON uai.email_validation_status = evs.id
			LEFT JOIN user_status us ON ua.user_status = us.id
			WHERE uai.email_id = $1;
		`, email).
		Scan(&profile.ProfileId, &profile.Username, &profile.FirstName, &profile.LastName, &profile.DateOfBirth, &profile.Bio, &profile.AvatarUrl,
			&profile.CreatedAt, &profile.UpdatedAt, &profile.Address, &profile.ContactNumber, &profile.UserStatus); err != nil {
		logrus.Errorf("can't find a user profile existing with email %s, error: %+v", email, err)
		return nil, err
	}

	return profile, nil
}

// TODO: If the record doesn't exist throw 404 error
// func (uh *uDBHandler) UpdateMyProfile(info *pb.SetMyProfileReq) error {
// 	stmt, err := uh.db.Prepare(`UPDATE the_monkeys_user SET first_name=$1, last_name=$2,
// 	country_code=$3, mobile_no=$4, about=$5, instagram=$6, twitter=$7, update_time=$8 WHERE id=$9`)
// 	if err != nil {
// 		uh.log.Errorf("cannot prepare update profile statement, error: %v", err)
// 		return err
// 	}
// 	defer stmt.Close()
// 	time := time.Now().Format(common.DATE_TIME_FORMAT)
// 	res, err := stmt.Exec(info.FirstName, info.LastName, info.CountryCode, info.MobileNo,
// 		info.About, info.Instagram, info.Twitter, time, info.Id)
// 	if err != nil {
// 		uh.log.Errorf("cannot execute update profile statement, error: %v", err)
// 		return err
// 	}

// 	row, err := res.RowsAffected()
// 	if err != nil {
// 		logrus.Errorf("error while checking rows affected for %s, error: %v", info.Email, err)
// 		return err
// 	}
// 	if row > 1 {
// 		logrus.Errorf("more or less than 1 row is affected for %s, error: %v", info.Email, err)
// 		return errors.New("more or less than 1 row is affected")
// 	}

// 	return nil
// }

// TODO: If the record doesn't exist throw 404 error
func (uh *uDBHandler) UploadProfilePic(pic []byte, id int64) error {
	stmt, err := uh.db.Prepare(`UPDATE the_monkeys_user SET profile_pic=$1 WHERE id=$2`)
	if err != nil {
		uh.log.Errorf("cannot prepare upload profile pic statement, error: %v", err)
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(pic, id)
	if err != nil {
		uh.log.Errorf("cannot execute update profile pic statement, error: %v", err)
		return err
	}

	row, err := res.RowsAffected()
	if err != nil {
		logrus.Errorf("error while checking rows affected for %d, error: %v", id, err)
		return err
	}
	if row != 1 {
		logrus.Errorf("more or less than 1 row is affected for %d, error: %v", id, err)
		return errors.New("more or less than 1 row is affected")
	}

	return nil
}

func (uh *uDBHandler) DeactivateMyAccount(id int64) error {
	stmt, err := uh.db.Prepare(`UPDATE the_monkeys_user SET deactivated=true WHERE id=$1`)
	if err != nil {
		uh.log.Errorf("cannot prepare deactivate profile statement, error: %v", err)
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(id)
	if err != nil {
		uh.log.Errorf("cannot execute deactivate profile statement, error: %v", err)
		return err
	}

	row, err := res.RowsAffected()
	if err != nil {
		logrus.Errorf("error while checking rows affected for %d, error: %v", id, err)
		return err
	}
	if row != 1 {
		logrus.Errorf("more or less than 1 row is affected for %d, error: %v", id, err)
		return errors.New("more or less than 1 row is affected")
	}

	return nil
}
