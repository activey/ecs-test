ALTER TABLE characters
    ADD COLUMN race_index VARCHAR(255);

UPDATE characters
SET race_index = 'human';

ALTER TABLE characters
    ALTER COLUMN race_index SET NOT NULL;