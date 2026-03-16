CREATE TABLE users
(
    id         BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ  DEFAULT NOW(),
    updated_at TIMESTAMPTZ  DEFAULT NOW(),
    deleted_at TIMESTAMPTZ  DEFAULT NULL,
    username   VARCHAR(256) NOT NULL,
    password   VARCHAR(255) DEFAULT NULL
);

INSERT INTO users (username, password) values ('test', '$2a$10$4AKpPI5r8NcuYCjtWCD5TuQgVrjC.c0UC5Nb0Fe/IWEqMN7QuLajy');

ALTER TABLE characters
    DROP COLUMN username;
ALTER TABLE characters
    ADD COLUMN user_id BIGSERIAL;
ALTER TABLE characters
    ADD CONSTRAINT fk_user
        FOREIGN KEY (user_id)
            REFERENCES users (id)
            ON DELETE CASCADE;

