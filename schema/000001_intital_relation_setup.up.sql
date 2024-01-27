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
    id SERIAL NOT NULL PRIMARY KEY,
    status VARCHAR(100)
);

CREATE TABLE IF NOT EXISTS USER_AUTH_INFO (
    id BIGSERIAL NOT NULL,
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
    email_validation_status INTEGER NOT NULL,
    email_validation_time TIMESTAMP,
    pwd_recovery_token VARCHAR(100),
    token_recovery_time TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES USER_ACCOUNT(user_id) ON DELETE CASCADE ON UPDATE NO ACTION,
    FOREIGN KEY (email_validation_status) REFERENCES EMAIL_VALIDATION_STATUS(id) ON DELETE NO ACTION ON UPDATE NO ACTION
);


CREATE TABLE IF NOT EXISTS user_role (
    id SERIAL NOT NULL PRIMARY KEY,
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
    FOREIGN KEY (role_id) REFERENCES USER_ROLE(id) ON DELETE CASCADE ON UPDATE NO ACTION,
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

CREATE TABLE IF NOT EXISTS clients (
    id SERIAL NOT NULL,
    c_name VARCHAR(32),
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS logged_in_devices (
    id          SERIAL NOT NULL,
    device_name VARCHAR(32),
    ip_address  VARCHAR(64),
    operating_sys VARCHAR(32),
    login_time  TIMESTAMP,
    user_id     BIGINT NOT NULL,
    client_id   INTEGER NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES user_account(user_id) ON DELETE CASCADE ON UPDATE NO ACTION,
    FOREIGN KEY (client_id) REFERENCES clients(id) ON DELETE CASCADE ON UPDATE NO ACTION
);

-- CREATE TABLE IF NOT EXISTS test {
--     id          SERIAL NOT NULL,
--     device_name VARCHAR(32)
-- }

-- Insert predefined roles
INSERT INTO user_role (role_desc) VALUES ('Admin'), ('Editor'), ('Author'), ('Subscriber');

-- Insert predefined permissions
INSERT INTO PERMISSIONS (permission_desc) VALUES ('Read'), ('Write'), ('Edit'), ('Delete');

-- Insert predefined clients
INSERT INTO clients (c_name) VALUES ('chrome'), ('firefox'), ('safari'), ('edge'), ('opera'), ('android_os'), ('ios'), ('brave'), ('others');

INSERT INTO EMAIL_VALIDATION_STATUS (status) VALUES ('unverified');
INSERT INTO EMAIL_VALIDATION_STATUS (status) VALUES ('validation_link_sent');
INSERT INTO EMAIL_VALIDATION_STATUS (status) VALUES ('verified');


-- Insert Predefined auth-providers
INSERT INTO EXTERNAL_AUTH_PROVIDERS (provider_name) VALUES ('google-oauth2'), ('instagram-oauth2');



-- Inserting data into PERMISSIONS_GRANTED for all roles
INSERT INTO PERMISSIONS_GRANTED (role_id, permission_id)
SELECT r.id, p.permission_id
FROM user_role r, PERMISSIONS p
WHERE r.role_desc IN ('Admin', 'Editor', 'Author', 'Subscriber')
AND p.permission_desc IN ('Read', 'Write', 'Edit', 'Delete');
