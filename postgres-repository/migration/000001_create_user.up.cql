CREATE TABLE IF NOT EXISTS users (
                                     id UUID PRIMARY KEY,
                                     name TEXT,
                                     password TEXT,
                                     email TEXT,
                                     fullname TEXT,
                                     phone TEXT,
                                     created_at TIMESTAMP,
                                     updated_at TIMESTAMP,
                                     deleted_at TIMESTAMP
);
