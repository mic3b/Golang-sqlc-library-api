-- migrate:up
CREATE TABLE IF NOT EXISTS book (
  id INTEGER PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  author VARCHAR(255) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  author_id INTEGER, 
  FOREIGN KEY (author_id) REFERENCES author(id)
);

-- migrate:down
DROP TABLE book;

