CREATE TABLE
    favourite_nfts (
        id BIGSERIAL PRIMARY KEY,
        user_id INTEGER NOT NULL,
        nft_id INTEGER NOT NULL,
        UNIQUE (user_id, nft_id),
        FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
        FOREIGN KEY (nft_id) REFERENCES nfts (id) ON DELETE CASCADE
    )