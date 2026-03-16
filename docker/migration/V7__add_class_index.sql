ALTER TABLE characters
    ADD COLUMN class_index VARCHAR(255);

UPDATE characters
SET class_index = 'barbarian';

ALTER TABLE characters
    ALTER COLUMN class_index SET NOT NULL;