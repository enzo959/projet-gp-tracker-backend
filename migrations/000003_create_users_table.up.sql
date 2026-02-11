CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    role TEXT NOT NULL DEFAULT 'user',

    first_name TEXT DEFAULT '',
    last_name TEXT DEFAULT '',
    surname TEXT DEFAULT '',
    bio TEXT DEFAULT '',
    image TEXT DEFAULT '',

    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);