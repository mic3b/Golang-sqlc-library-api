-- migrate:up
CREATE TABLE IF NOT EXISTS new_book (
  id INTEGER PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  author VARCHAR(255) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


-- migrate:down
DROP TABLE new_book;

