CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE tasks (
  task_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  title VARCHAR(255) NOT NULL,
  status BOOLEAN NOT NULL DEFAULT FALSE,
  user_id UUID,
  CONSTRAINT fk_parent FOREIGN KEY (user_id) REFERENCES users(id)
);