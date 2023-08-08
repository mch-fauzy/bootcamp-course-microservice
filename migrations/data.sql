CREATE TABLE bootcamp_courses (
  id CHAR(36) PRIMARY KEY NOT NULL,
  user_id CHAR(36) UNIQUE NOT NULL,
  title VARCHAR(255) NOT NULL,
  content VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  deleted_at TIMESTAMP
);