CREATE TABLE nfts (
    id SERIAL PRIMARY KEY,
    token_id TEXT,
    token_uri TEXT,
    name VARCHAR(255),
    description TEXT,
    price NUMERIC(38,18), 
    owner_id BIGINT NOT NULL,
    image_path TEXT,
    FOREIGN KEY (owner_id) REFERENCES users (id) ON DELETE CASCADE
);
