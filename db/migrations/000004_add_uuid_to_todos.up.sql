ALTER TABLE todos ADD COLUMN uuid UUID;

UPDATE todos SET uuid = gen_random_uuid();

ALTER TABLE todos ALTER COLUMN uuid SET NOT NULL;

CREATE UNIQUE INDEX idx_todos_uuid ON todos(uuid);
