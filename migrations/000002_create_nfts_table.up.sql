CREATE TABLE nfts
(
    id          SERIAL PRIMARY KEY,
    token_id    TEXT,
    token_uri   TEXT,
    name        VARCHAR(255),
    description TEXT,
    price       NUMERIC(38, 18),
    owner_id    BIGINT NOT NULL,
    image_path  TEXT,
    in_sales    bool DEFAULT false,
    proposed    bool DEFAULT false,
    votes_amount BIGINT DEFAULT 0,
    FOREIGN KEY (owner_id) REFERENCES users (id) ON DELETE CASCADE
);

INSERT INTO nfts (token_id, token_uri, name, description, price, owner_id, image_path, in_sales, proposed, votes_amount)
VALUES
    ('1', 'http://localhost:8080/metadata/El Capitan 2.jpg', 'El Capitan 2', 'High quality image of El Capitan 2', 0.000000000000000000, 1, 'El Capitan 2.jpg', false, true, 0),
    ('2', 'http://localhost:8080/metadata/_21x_19gjcxlw2_Kikis-Delivery-Service-Blu-ray.jpg', 'Kikis Delivery Service Blu-ray', 'A beautiful artwork of Kikis Delivery Service Blu-ray', 0.000000000000000000, 1, '_21x_19gjcxlw2_Kikis-Delivery-Service-Blu-ray.jpg', false, false, 0),
    ('3', 'http://localhost:8080/metadata/119038-fictional_character-creative_arts-woman_warrior-woman-art-3840x2160.jpg', 'Fictional Character Warrior Art', 'An artistic depiction of a woman warrior.', 0.000000000000000000, 1, '119038-fictional_character-creative_arts-woman_warrior-woman-art-3840x2160.jpg', false, false, 0),
    ('4', 'http://localhost:8080/metadata/16248.jpg', 'Abstract Artwork', 'Abstract artwork with vibrant colors.', 0.000000000000000000, 1, '16248.jpg', true, false, 0),
    ('5', 'http://localhost:8080/metadata/1MS7cWbJ-wallha.com.jpg', 'Wallha Art', 'A captivating landscape painting.', 0.000000000000000000, 1, '1MS7cWbJ-wallha.com.jpg', true, true, 0),
    ('6', 'http://localhost:8080/metadata/2024-Formula1-Red-Bull-Racing-RB20-005-2160.jpg', 'Formula 1 Red Bull Racing', 'A high-quality image of a Red Bull Racing Formula 1 car.', 0.000000000000000000, 1, '2024-Formula1-Red-Bull-Racing-RB20-005-2160.jpg', true, false, 0),
    ('7', 'http://localhost:8080/metadata/2968674.jpg', 'Nature Landscape', 'Beautiful nature landscape with mountains and lake.', 0.000000000000000000, 1, '2968674.jpg', false, true, 0),
    ('8', 'http://localhost:8080/metadata/3333.jpg', 'Blue Sky', 'A calm blue sky with soft clouds.', 0.000000000000000000, 1, '3333.jpg', true, true, 0),
    ('9', 'http://localhost:8080/metadata/4k Landscape Wallpaper (2).jpg', '4K Landscape', 'A serene 4K landscape wallpaper.', 0.000000000000000000, 1, '4k Landscape Wallpaper (2).jpg', true, false, 0),
    ('10', 'http://localhost:8080/metadata/4k Landscape Wallpaper.jpg', '4K Landscape', 'Another beautiful 4K landscape wallpaper.', 0.000000000000000000, 1, '4k Landscape Wallpaper.jpg', false, true, 0);
