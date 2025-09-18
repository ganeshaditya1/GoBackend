CREATE TABLE users (
    userid SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    email_address VARCHAR(100) NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    salt VARCHAR(255) NOT NULL,
    age INT CHECK (age >= 0)
);

CREATE TABLE admin (
    userid INT REFERENCES users(userid)
);