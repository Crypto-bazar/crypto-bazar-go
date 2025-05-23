CREATE TABLE
    nfts (
        id SERIAL PRIMARY KEY,
        proposal_id INT DEFAULT 0,
        token_id INT DEFAULT 0,
        token_uri TEXT,
        name VARCHAR(255),
        description TEXT,
        price NUMERIC(38, 18),
        owner_id BIGINT NOT NULL,
        image_path TEXT,
        in_sales bool DEFAULT false,
        proposed bool DEFAULT false,
        votes_amount NUMERIC(78, 0) DEFAULT 0,
        FOREIGN KEY (owner_id) REFERENCES users (id) ON DELETE CASCADE
    );

INSERT INTO
    nfts (
        proposal_id,
        token_id,
        token_uri,
        name,
        description,
        price,
        owner_id,
        image_path,
        in_sales,
        proposed,
        votes_amount
    )
VALUES
    (
        0,
        0,
        'http://localhost:8080/metadata/8e0c57be-bb84-47c1-ae7f-e269c0cb67e4',
        'Эль-Капитан 2',
        'Высококачественное изображение Эль-Капитана 2',
        0.000000000000000000,
        1,
        'El Capitan 2.jpg',
        false,
        false,
        0
    ),
    (
        0,
        0,
        'http://localhost:8080/metadata/3b10a7f2-4ddf-4e3a-b4ee-90e1dfe98cc4',
        'Ведьмина служба доставки (Blu-ray)',
        'Красочное арт-изображение Blu-ray-диска "Ведьмина служба доставки"',
        0.000000000000000000,
        2,
        '_21x_19gjcxlw2_Kikis-Delivery-Service-Blu-ray.jpg',
        false,
        false,
        0
    ),
    (
        0,
        0,
        'http://localhost:8080/metadata/427a5ac3-8e82-4dbe-aac8-e21913475da1',
        'Воин — вымышленный персонаж',
        'Художественное изображение женщины-воина',
        0.000000000000000000,
        1,
        '119038-fictional_character-creative_arts-woman_warrior-woman-art-3840x2160.jpg',
        false,
        false,
        0
    ),
    (
        0,
        0,
        'http://localhost:8080/metadata/60cdd273-14a6-4f2f-83db-932fbe32c590',
        'Абстрактное искусство',
        'Абстрактная работа с яркими цветами',
        0.000000000000000000,
        2,
        '16248.jpg',
        false,
        false,
        0
    ),
    (
        0,
        0,
        'http://localhost:8080/metadata/fef47b65-1de4-4f87-9bfc-157e998f62f1',
        'Wallha Арт',
        'Захватывающая пейзажная живопись',
        0.000000000000000000,
        2,
        '1MS7cWbJ-wallha.com.jpg',
        false,
        false,
        0
    ),
    (
        0,
        0,
        'http://localhost:8080/metadata/f2c2292d-6a3a-4d07-90ed-c7b08c8bce6e',
        'Формула 1 — Red Bull Racing',
        'Изображение болида Red Bull Racing Формулы 1 в высоком качестве',
        0.000000000000000000,
        1,
        '2024-Formula1-Red-Bull-Racing-RB20-005-2160.jpg',
        false,
        false,
        0
    ),
    (
        0,
        0,
        'http://localhost:8080/metadata/a190a206-f964-4d1c-bc85-4fe1b83e688a',
        'Пейзаж природы',
        'Красивый природный пейзаж с горами и озером',
        0.000000000000000000,
        1,
        '2968674.jpg',
        false,
        false,
        0
    ),
    (
        0,
        0,
        'http://localhost:8080/metadata/3e4c6b62-7e92-4c6b-8f90-f3b275b7648d',
        'Голубое небо',
        'Спокойное голубое небо с мягкими облаками',
        0.000000000000000000,
        1,
        '3333.jpg',
        false,
        false,
        0
    ),
    (
        0,
        0,
        'http://localhost:8080/metadata/46f3d650-116f-4a63-902a-403edb3ff38a',
        'Пейзаж в 4K (2)',
        'Спокойный пейзаж в разрешении 4K',
        0.000000000000000000,
        1,
        '4k Landscape Wallpaper (2).jpg',
        false,
        false,
        0
    ),
    (
        0,
        0,
        'http://localhost:8080/metadata/bc7a7cd0-9940-4cfd-9b46-645fc2a3c17e',
        'Пейзаж в 4K',
        'Ещё один красивый пейзаж в разрешении 4K',
        0.000000000000000000,
        1,
        '4k Landscape Wallpaper.jpg',
        false,
        false,
        0
    );