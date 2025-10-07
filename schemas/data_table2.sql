CREATE TABLE books (
    book_id SERIAL PRIMARY KEY,
    book_name VARCHAR(100) NOT NULL,
    book_description VARCHAR(100) NOT NULL,
    CHECK (LEFT(book_name, 1) IN ('N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'))
);
