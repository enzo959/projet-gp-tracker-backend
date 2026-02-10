CREATE TABLE artists (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    image_url TEXT,
    created_at TIMESTAMP DEFAULT now()
);

