CREATE TABLE artists (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,

    bio TEXT DEFAULT '',
    musique_url TEXT DEFAULT '',
    image_url TEXT DEFAULT '',

    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);