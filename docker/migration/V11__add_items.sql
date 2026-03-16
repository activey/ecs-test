CREATE TABLE items
(
    id           BIGSERIAL PRIMARY KEY,
    created_at   TIMESTAMPTZ DEFAULT NOW(),
    updated_at   TIMESTAMPTZ DEFAULT NOW(),
    deleted_at   TIMESTAMPTZ DEFAULT NULL,

    character_id BIGINT NOT NULL,
    item_index   VARCHAR(255) NOT NULL ,

    CONSTRAINT fk_equipment_id
        FOREIGN KEY (character_id)
            REFERENCES characters (id)
            ON DELETE CASCADE
);
