CREATE TABLE comments
(
    id         SERIAL PRIMARY KEY,
    nft_id     BIGINT      NOT NULL,
    owner_id   BIGINT      NOT NULL,
    content    TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    FOREIGN KEY (owner_id) REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (nft_id) REFERENCES nfts (id) ON DELETE CASCADE
)