CREATE TABLE users (
  id INTEGER PRIMARY KEY AUTOINCREMENT ,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  username text,
  email text,
  password text
);
