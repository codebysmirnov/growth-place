-- +migrate Up
CREATE TABLE users(
    id       UUID CONSTRAINT users_pk PRIMARY KEY,
    login    VARCHAR (32) CONSTRAINT users__login_uniq_idx UNIQUE NOT NULL,
    name     VARCHAR (32),
    email    VARCHAR (32) CONSTRAINT users__email_uniq_idx UNIQUE,
    phone    VARCHAR (32) CONSTRAINT users__phone_uniq_idx UNIQUE,
    password VARCHAR (64),

    created_at  TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
    updated_at  TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
    deleted_at  TIMESTAMP WITH TIME ZONE
);

COMMENT ON TABLE  users            IS 'Users';
COMMENT ON COLUMN users.id         IS 'identifier';
COMMENT ON COLUMN users.login      IS 'login';
COMMENT ON COLUMN users.name       IS 'name';
COMMENT ON COLUMN users.email      IS 'email';
COMMENT ON COLUMN users.phone      IS 'phone';
COMMENT ON COLUMN users.password   IS 'password';
COMMENT ON COLUMN users.created_at IS 'record created date';
COMMENT ON COLUMN users.updated_at IS 'record modify date';
COMMENT ON COLUMN users.deleted_at IS 'record delete date';

-- +migrate Down
DROP TABLE users;