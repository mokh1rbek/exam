
CREATE TABLE categorys (
    category_id UUID PRIMARY KEY NOT NULL,
    name VARCHAR(255) NOT NULL,
    parent_uuid UUID,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);

CREATE TABLE products (
    product_id UUID PRIMARY KEY NOT NULL,
    name VARCHAR(255) NOT NULL,
    price NUMERIC NOT NULL DEFAULT 0,
    category_id UUID NOT NULL REFERENCES categorys(category_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);

CREATE TABLE orders (
    order_id UUID PRIMARY KEY NOT NULL,
    description VARCHAR(255) NOT NULL,
    product_id UUID NOT NULL REFERENCES products(product_id), 
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);
