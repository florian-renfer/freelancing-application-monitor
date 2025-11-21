CREATE TABLE users (
    id TEXT PRIMARY KEY,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    firstname TEXT NOT NULL,
    lastname TEXT NOT NULL,
    date_of_birth DATE NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL
);
