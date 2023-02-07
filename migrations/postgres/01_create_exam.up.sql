
CREATE TABLE IF NOT EXISTS categories (
    id UUID NOT NULL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    parent_id UUID REFERENCES categories(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS products (
    id UUID PRIMARY KEY NOT NULL,
    name VARCHAR(255) NOT NULL,
    price NUMERIC NOT NULL DEFAULT 0,
    category_id UUID NOT NULL REFERENCES categories(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS orders (
    id UUID PRIMARY KEY NOT NULL,
    description VARCHAR(255) NOT NULL,
    product_id UUID NOT NULL REFERENCES products(id), 
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);
