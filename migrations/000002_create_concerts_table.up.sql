CREATE TABLE concerts (
    id SERIAL PRIMARY KEY,
    artist_id INT NOT NULL REFERENCES artists(id) ON DELETE CASCADE,
    date TIMESTAMP NOT NULL,
    location TEXT NOT NULL,
    price_cents INT NOT NULL,
    total_tickets INT NOT NULL
    image_url TEXT
);
