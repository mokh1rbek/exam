
CREATE TABLE category (
    category_id UUID PRIMARY KEY NOT NULL,
    name VARCHAR(255) NOT NULL,
    parent_uuid UUID,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE product (
    product_id UUID PRIMARY KEY NOT NULL,
    name VARCHAR(255) NOT NULL,
    price NUMERIC NOT NULL DEFAULT 0,
    category_id UUID NOT NULL REFERENCES category(category_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE order (
    order_id UUID PRIMARY KEY NOT NULL,
    description VARCHAR(255) NOT NULL,
    product_id UUID NOT NULL REFERENCES product, 
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);
