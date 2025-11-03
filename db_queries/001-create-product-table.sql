create table products (
    id BIGSERIAL PRIMARY KEY,
    title VARCHAR(250) NOT NULL,
    description TEXT,
    price DOUBLE PRECISION NOT NULL,
    image_url VARCHAR(500),
    created_at TIMESTAMP
    WITH
        TIME ZONE DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP
    WITH
        TIME ZONE DEFAULT CURRENT_TIMESTAMP
);