
-- +migrate Up
ALTER TABLE users
    DROP COLUMN password;

ALTER TABLE users
    ADD COLUMN password VARCHAR(1024);

COMMENT ON COLUMN users.password   IS 'password';

-- +migrate Down
ALTER TABLE users
    DROP COLUMN password;

ALTER TABLE users
    ADD COLUMN password VARCHAR(64);

COMMENT ON COLUMN users.password   IS 'password';