CREATE TABLE characters
(
    id         BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ  DEFAULT NOW(),
    updated_at TIMESTAMPTZ  DEFAULT NOW(),
    deleted_at TIMESTAMPTZ  DEFAULT NULL,
    username   VARCHAR(256) NOT NULL,
    name       VARCHAR(255) DEFAULT NULL
);
