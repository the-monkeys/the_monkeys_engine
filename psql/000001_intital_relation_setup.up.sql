-- User Service
CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);

CREATE TABLE user_profiles (
    profile_id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(user_id),
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    bio TEXT,
    avatar_url TEXT,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);

CREATE TABLE password_resets (
    reset_id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(user_id),
    reset_token VARCHAR(255) NOT NULL,
    expires_at TIMESTAMP NOT NULL
);

CREATE TABLE user_preferences (
    preference_id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(user_id),
    preference_key VARCHAR(50) NOT NULL,
    preference_value VARCHAR(255) NOT NULL
);

-- Auth Service
CREATE TABLE auth_tokens (
    token_id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(user_id),
    access_token VARCHAR(255) NOT NULL,
    refresh_token VARCHAR(255) NOT NULL,
    expires_at TIMESTAMP NOT NULL
);

CREATE TABLE user_roles (
    role_id SERIAL PRIMARY KEY,
    role_name VARCHAR(50) UNIQUE NOT NULL
);

-- Insert predefined roles
INSERT INTO user_roles (role_name) VALUES ('Admin');
INSERT INTO user_roles (role_name) VALUES ('Editor');
INSERT INTO user_roles (role_name) VALUES ('Author');
INSERT INTO user_roles (role_name) VALUES ('Subscriber');

CREATE TABLE user_role_mappings (
    mapping_id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(user_id),
    role_id INTEGER REFERENCES user_roles(role_id)
);

CREATE TABLE role_permissions (
    permission_id SERIAL PRIMARY KEY,
    role_id INTEGER REFERENCES user_roles(role_id),
    permission_name VARCHAR(50) NOT NULL
);