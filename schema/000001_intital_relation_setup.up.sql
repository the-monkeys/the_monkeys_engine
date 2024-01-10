CREATE TABLE USER_ACCOUNT (
    profile_id INTEGER,
    user_id INTEGER,
    username VARCHAR(32),
    first_name VARCHAR(100),
    last_name VARCHAR(100),
    date_of_birth DATE,
    role_id INTEGER,
    bio TEXT,
    avatar_url TEXT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    address VARCHAR(255),
    contact_number INTEGER,
    PRIMARY KEY (profile_id, user_id)
    );


CREATE TABLE USER_AUTH_INFO (
    user_id INTEGER REFERENCES USER_ACCOUNT(user_id),
    username VARCHAR(32) REFERENCES USER_ACCOUNT(username),
    login_name VARCHAR(100),
    password_hash VARCHAR(100),
    password_salt VARCHAR(100),
    email_id VARCHAR(100),
    created_at TIMESTAMP,
    password_updated_at TIMESTAMP,
    confirmation_token VARCHAR(100),
    token_generation_time TIMESTAMP,
    email_valid_status_id INTEGER REFERENCES EMAIL_VALIDATION(email_valid_status_id),
    pwd_recovery_token VARCHAR(100),
    token_recovery_time TIMESTAMP
    );

CREATE TABLE USER_ROLE (
    role_id INTEGER PRIMARY KEY,
    role_desc VARCHAR(50)
);

CREATE TABLE PERMISSIONS (
    permission_id INTEGER PRIMARY KEY,
    permission_desc VARCHAR(50)
);

CREATE TABLE PERMISSIONS_GRANTED (
    role_id INTEGER REFERENCES USER_ROLE(role_id),
    permission_id INTEGER REFERENCES PERMISSIONS(permission_id)
);


CREATE TABLE USER_ACCOUNT_STATUS (
    user_id INTEGER REFERENCES USER_ACCOUNT(user_id),
    account_status VARCHAR(20),
    last_login TIMESTAMP,
    creation_date TIMESTAMP,
    modified_date TIMESTAMP,
    reason VARCHAR(100)
);

-- TODO: Check if more fields are required for email validation
CREATE TABLE EMAIL_VALIDATION (
    email_valid_status_id INTEGER PRIMARY KEY,
    status_desc VARCHAR(100)
);

CREATE TABLE USER_ACCOUNT_LOG (
    log_id INTEGER PRIMARY KEY,
    profile_id INTEGER REFERENCES USER_ACCOUNT(profile_id),
    event_type VARCHAR(50),
    timestamp TIMESTAMP,
    ip_address VARCHAR(20),
    description TEXT,
    status VARCHAR(10)
);

CREATE TABLE USER_LOGIN_EXTERNAL (
    user_id INTEGER REFERENCES USER_ACCOUNT(user_id),
    external_provider_id INTEGER REFERENCES EXTERNAL_PROVIDERS(external_provider_id),
    external_provider_token VARCHAR(100)
    );

-- TODO: Modify the table name for external authentication (eg: google , facebook)
CREATE TABLE EXTERNAL_PROVIDERS (
    external_provider_id INTEGER PRIMARY KEY,
    provider_name VARCHAR(100)
);

CREATE TABLE PAYMENT_INFO (
    subscription_info VARCHAR(50),
    payment_info VARCHAR(50),
    user_id INTEGER REFERENCES USER_ACCOUNT(user_id),
    profile_id INTEGER REFERENCES USER_ACCOUNT(profile_id)
);

CREATE TABLE INTEREST (
    interest_id INTEGER PRIMARY KEY,
    interest_name VARCHAR(100)
);

CREATE TABLE USER_INTEREST (
    user_id INTEGER REFERENCES USER_ACCOUNT(user_id),
    interest_id INTEGER REFERENCES INTEREST(interest_id),
    PRIMARY KEY (user_id, interest_id)
);



-- Insert predefined roles
INSERT INTO user_role (role_desc) VALUES ('Admin');
INSERT INTO user_role (role_desc) VALUES ('Editor');
INSERT INTO user_role (role_desc) VALUES ('Author');
INSERT INTO user_role (role_desc) VALUES ('Subscriber');

