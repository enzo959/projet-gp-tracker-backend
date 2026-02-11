CREATE TABLE concerts (
    id SERIAL PRIMARY KEY,
    artist_id INT NOT NULL REFERENCES artists(id) ON DELETE CASCADE,
    artist_name TEXT,
    date TIMESTAMP NOT NULL,
    location TEXT NOT NULL,
    price_cents INT NOT NULL,
    total_tickets INT NOT NULL,
    detail TEXT,
    image_url TEXT
    
);
