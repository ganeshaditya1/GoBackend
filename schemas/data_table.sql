CREATE TABLE catalog (
    catalog_id SERIAL PRIMARY KEY,
    catalog_name VARCHAR(100) NOT NULL,
    catalog_creator_id INT NOT NULL,
    CHECK (LEFT(catalog_name, 1) IN ('A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M'))
);


CREATE TABLE items (
    item_id SERIAL PRIMARY KEY,
    item_name VARCHAR(100) NOT NULL,
    item_description TEXT,
    catalog_id INT,
    creator_id INT NOT NULL,
    FOREIGN KEY (catalog_id) REFERENCES catalog(catalog_id)
);
