CREATE TABLE IF NOT EXISTS  subjects (
    id uuid PRIMARY KEY,
    name varchar NOT NULL,
    type varchar NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp
);