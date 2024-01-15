CREATE TABLE USER_ACCOUNT (
    profile_id INTEGER,
    user_id INTEGER UNIQUE,
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
    -- UNIQUE(user_id, username)  -- Add this line
);

CREATE TABLE EMAIL_VALIDATION_STATUS (
    email_validation_status_id INTEGER PRIMARY KEY,
    status_desc VARCHAR(100)
);

CREATE TABLE USER_AUTH_INFO (
    user_id INTEGER,
    username VARCHAR(32),
    login_name VARCHAR(100),
    password_hash VARCHAR(100),
    password_salt VARCHAR(100),
    email_id VARCHAR(100),
    created_at TIMESTAMP,
    password_updated_at TIMESTAMP,
    confirmation_token VARCHAR(100),
    token_generation_time TIMESTAMP,
    email_validation_status_id INTEGER REFERENCES EMAIL_VALIDATION_STATUS(email_validation_status_id),
    email_validation_time TIMESTAMP,
    pwd_recovery_token VARCHAR(100),
    token_recovery_time TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES USER_ACCOUNT(user_id)
);


CREATE TABLE USER_ROLE (
    role_id SERIAL PRIMARY KEY,
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
    id SERIAL,
    user_id INTEGER,
    account_status VARCHAR(20),
    last_login TIMESTAMP,
    creation_date TIMESTAMP,
    modified_date TIMESTAMP,
    reason VARCHAR(100),
    FOREIGN KEY (user_id) REFERENCES USER_ACCOUNT(user_id)
);


CREATE TABLE USER_ACCOUNT_LOG (
    log_id INTEGER PRIMARY KEY,
    profile_id INTEGER,
    event_type VARCHAR(50),
    timestamp TIMESTAMP,
    ip_address VARCHAR(20),
    description TEXT,
    status VARCHAR(10),
    FOREIGN KEY (profile_id) REFERENCES USER_ACCOUNT(user_id)
);

CREATE TABLE EXTERNAL_AUTH_PROVIDERS (
    external_provider_id SERIAL PRIMARY KEY,
    provider_name VARCHAR(100)
);


CREATE TABLE USER_EXTERNAL_LOGIN (
    user_id INTEGER,
    external_provider_id SERIAL REFERENCES EXTERNAL_AUTH_PROVIDERS(external_provider_id),
    external_provider_token VARCHAR(100),
    FOREIGN KEY (user_id) REFERENCES USER_ACCOUNT(user_id)
);


CREATE TABLE PAYMENT_INFO (
    subscription_info VARCHAR(50),
    payment_info VARCHAR(50),
    user_id INTEGER,
    profile_id INTEGER,
    FOREIGN KEY (user_id, profile_id) REFERENCES USER_ACCOUNT(user_id, profile_id)
);

CREATE TABLE INTEREST (
    interest_id INTEGER PRIMARY KEY,
    interest_name VARCHAR(100)
);

CREATE TABLE USER_INTEREST (
    user_id INTEGER,
    interest_id INTEGER,
    PRIMARY KEY (user_id, interest_id),
    FOREIGN KEY (user_id) REFERENCES USER_ACCOUNT(user_id),
    FOREIGN KEY (interest_id) REFERENCES INTEREST(interest_id)
);



-- Insert predefined roles
INSERT INTO user_role (role_desc) VALUES ('admin');
INSERT INTO user_role (role_desc) VALUES ('editor');
INSERT INTO user_role (role_desc) VALUES ('author');
INSERT INTO user_role (role_desc) VALUES ('subscriber');


-- Insert Predefined auth-providers
INSERT INTO EXTERNAL_AUTH_PROVIDERS (provider_name) VALUES ('google-oauth2')