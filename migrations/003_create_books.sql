-- +migrate Up
CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    image_url TEXT,
    release_year INTEGER CHECK (release_year BETWEEN 1980 AND 2024),
    price INTEGER,
    total_page INTEGER,
    thickness VARCHAR(50),
    category_id INTEGER REFERENCES categories(id) ON DELETE SET NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(100),
    modified_at TIMESTAMP,
    modified_by VARCHAR(100)
);

-- +migrate Down
DROP TABLE IF EXISTS books;
