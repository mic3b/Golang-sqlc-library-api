-- migrate:up
ALTER TABLE new_book ADD COLUMN author_id INTEGER;

-- migrate:down
ALTER TABLE new_book DROP COLUMN author_id;
