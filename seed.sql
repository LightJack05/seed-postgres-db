CREATE TABLE foo (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO foo (name) VALUES
('Alice'),
('Bob'),
('Charlie');
