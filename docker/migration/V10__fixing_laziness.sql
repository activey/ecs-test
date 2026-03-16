DROP  TABLE ability_scores;

CREATE TABLE ability_scores
(
    id           BIGSERIAL PRIMARY KEY,
    created_at   TIMESTAMPTZ DEFAULT NOW(),
    updated_at   TIMESTAMPTZ DEFAULT NOW(),
    deleted_at   TIMESTAMPTZ DEFAULT NULL,

    character_id BIGINT NOT NULL,
    strength     INT DEFAULT 8 NOT NULL,
    dexterity    INT DEFAULT 8 NOT NULL,
    constitution INT DEFAULT 8 NOT NULL,
    intelligence INT DEFAULT 8 NOT NULL,
    wisdom       INT DEFAULT 8 NOT NULL,
    charisma     INT DEFAULT 8 NOT NULL,

    CONSTRAINT fk_character
        FOREIGN KEY (character_id)
            REFERENCES characters(id)
            ON DELETE CASCADE,

    CONSTRAINT unique_character_id UNIQUE (character_id)
);

-- Insert example record
INSERT INTO ability_scores (character_id, strength, dexterity, constitution, intelligence, wisdom, charisma)
VALUES (1, 8, 8, 8, 8, 8, 8);
