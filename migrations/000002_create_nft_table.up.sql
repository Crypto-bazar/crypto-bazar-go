CREATE TABLE nft (
    id SERIAL PRIMARY KEY,
    token_id CHAR(42) UNIQUE NOT NULL,
    owner_id INT NOT NULL,
		image_path VARCHAR(255),
    FOREIGN KEY (owner_id) REFERENCES users (id) ON DELETE CASCADE
);
