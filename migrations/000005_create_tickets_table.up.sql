CREATE TABLE IF NOT EXISTS tickets (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    concert_id INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),

    CONSTRAINT fk_ticket_user
        FOREIGN KEY (user_id) REFERENCES users(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_ticket_concert
        FOREIGN KEY (concert_id) REFERENCES concerts(id)
        ON DELETE CASCADE

    CONSTRAINT unique_user_concert
        UNIQUE (user_id, concert_id)
);
