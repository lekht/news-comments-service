CREATE SCHEMA IF NOT EXISTS comments;

DROP TABLE IF EXISTS comments.messages;

CREATE TABLE IF NOT EXISTS comments.messages (
    id SERIAL PRIMARY KEY,
    parent_id INT NOT NULL,
    news_id INT NOT NULL,
    msg TEXT NOT NULL,
    pubTime BIGINT NOT NULL CHECK (pubTime > 0)
);