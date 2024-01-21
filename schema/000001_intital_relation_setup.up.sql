CREATE TABLE IF NOT EXISTS USER_ACCOUNT (
    user_id BIGSERIAL NOT NULL,
    profile_id VARCHAR(32) NOT NULL, -- Will take a UUID
    username VARCHAR(32) NOT NULL,
    first_name VARCHAR(32),
    last_name VARCHAR(32),
    date_of_birth DATE,
    role_id INTEGER,
    bio TEXT,
    avatar_url TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    address VARCHAR(255),
    contact_number INTEGER,
    PRIMARY KEY (profile_id, user_id),
    UNIQUE(user_id, username)  -- Add this line
);
ALTER TABLE USER_ACCOUNT
ADD CONSTRAINT user_id_unique UNIQUE (user_id);

CREATE TABLE IF NOT EXISTS EMAIL_VALIDATION_STATUS (
    email_validation_status_id SERIAL NOT NULL PRIMARY KEY,
    status_desc VARCHAR(100)
);

CREATE TABLE IF NOT EXISTS USER_AUTH_INFO (
    user_auth_info_id BIGSERIAL NOT NULL,
    user_id BIGINT NOT NULL,
    username VARCHAR(32) NOT NULL,
    email_id VARCHAR(100) NOT NULL,
    password_hash VARCHAR(100) NOT NULL,
    -- password_salt VARCHAR(100) NOT NULL,
    -- email_id VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    password_updated_at TIMESTAMP,
    email_validation_token VARCHAR(100),
    token_generation_time TIMESTAMP,
    email_validation_status_id INTEGER NOT NULL,
    email_validation_time TIMESTAMP,
    pwd_recovery_token VARCHAR(100),
    token_recovery_time TIMESTAMP,
    PRIMARY KEY (user_auth_info_id),
    FOREIGN KEY (user_id) REFERENCES USER_ACCOUNT(user_id) ON DELETE CASCADE ON UPDATE NO ACTION,
    FOREIGN KEY (email_validation_status_id) REFERENCES EMAIL_VALIDATION_STATUS(email_validation_status_id) ON DELETE NO ACTION ON UPDATE NO ACTION
);


CREATE TABLE IF NOT EXISTS USER_ROLE (
    role_id SERIAL NOT NULL PRIMARY KEY,
    role_desc VARCHAR(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS PERMISSIONS (
    permission_id SERIAL NOT NULL PRIMARY KEY,
    permission_desc VARCHAR(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS PERMISSIONS_GRANTED (
    role_id BIGINT NOT NULL,
    permission_id INTEGER NOT NULL,
    PRIMARY KEY (role_id, permission_id),
    FOREIGN KEY (role_id) REFERENCES USER_ROLE(role_id) ON DELETE CASCADE ON UPDATE NO ACTION,
    FOREIGN KEY (permission_id) REFERENCES PERMISSIONS(permission_id) ON DELETE CASCADE ON UPDATE NO ACTION
);


CREATE TABLE IF NOT EXISTS USER_ACCOUNT_STATUS (
    id SERIAL NOT NULL,
    user_id BIGSERIAL NOT NULL,
    account_status VARCHAR(20) NOT NULL,
    last_login TIMESTAMP NOT NULL,
    creation_date TIMESTAMP NOT NULL,
    modified_date TIMESTAMP NOT NULL,
    reason VARCHAR(100),
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES USER_ACCOUNT(user_id) ON DELETE CASCADE ON UPDATE NO ACTION
);


CREATE TABLE IF NOT EXISTS USER_ACCOUNT_LOG (
    log_id BIGSERIAL NOT NULL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    event_type VARCHAR(50) NOT NULL,
    timestamp TIMESTAMP NOT NULL,
    ip_address VARCHAR(20) NOT NULL,
    description TEXT,
    status VARCHAR(10) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES USER_ACCOUNT(user_id) ON DELETE CASCADE ON UPDATE NO ACTION
);

CREATE TABLE IF NOT EXISTS EXTERNAL_AUTH_PROVIDERS (
    external_provider_id SERIAL NOT NULL PRIMARY KEY,
    provider_name VARCHAR(100) NOT NULL
);


CREATE TABLE IF NOT EXISTS USER_EXTERNAL_LOGIN (
    user_id BIGINT NOT NULL,
    external_provider_id INTEGER NOT NULL,
    external_provider_token VARCHAR(100) NOT NULL,
    PRIMARY KEY (user_id, external_provider_id),
    FOREIGN KEY (user_id) REFERENCES USER_ACCOUNT(user_id) ON DELETE CASCADE ON UPDATE NO ACTION,
    FOREIGN KEY (external_provider_id) REFERENCES EXTERNAL_AUTH_PROVIDERS(external_provider_id) ON DELETE CASCADE ON UPDATE NO ACTION
);


CREATE TABLE IF NOT EXISTS PAYMENT_INFO (
    subscription_info VARCHAR(50) NOT NULL,
    payment_info VARCHAR(50) NOT NULL,
    user_id BIGINT NOT NULL,
    profile_id VARCHAR(100) NOT NULL,
    PRIMARY KEY (user_id, profile_id),
    FOREIGN KEY (user_id, profile_id) REFERENCES USER_ACCOUNT(user_id, profile_id) ON DELETE CASCADE ON UPDATE NO ACTION
);

CREATE TABLE IF NOT EXISTS INTEREST (
    interest_id INTEGER NOT NULL PRIMARY KEY,
    interest_name VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS USER_INTEREST (
    user_id BIGINT NOT NULL,
    interest_id INTEGER NOT NULL,
    PRIMARY KEY (user_id, interest_id),
    FOREIGN KEY (user_id) REFERENCES USER_ACCOUNT(user_id) ON DELETE CASCADE ON UPDATE NO ACTION,
    FOREIGN KEY (interest_id) REFERENCES INTEREST(interest_id) ON DELETE CASCADE ON UPDATE NO ACTION
);



-- Insert predefined roles
INSERT INTO user_role (role_desc) VALUES ('admin');
INSERT INTO user_role (role_desc) VALUES ('general');
INSERT INTO user_role (role_desc) VALUES ('author');
INSERT INTO user_role (role_desc) VALUES ('subscriber');

INSERT INTO EMAIL_VALIDATION_STATUS (status_desc) VALUES ('validation_link_sent');
INSERT INTO EMAIL_VALIDATION_STATUS (status_desc) VALUES ('verified');
INSERT INTO EMAIL_VALIDATION_STATUS (status_desc) VALUES ('unverified');

-- Insert Predefined auth-providers
INSERT INTO EXTERNAL_AUTH_PROVIDERS (provider_name) VALUES ('google-oauth2');
