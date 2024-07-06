CREATE TABLE todos
(
    id          SERIAL PRIMARY KEY,
    title       VARCHAR,
    description TEXT,
    done        BOOLEAN DEFAULT FALSE
);
