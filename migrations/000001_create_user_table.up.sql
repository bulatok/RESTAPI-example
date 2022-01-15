CREATE TABLE IF NOT EXISTS users(
  user_id SERIAL PRIMARY KEY,
  name VARCHAR(50),
  surname VARCHAR(50),
  phone_number VARCHAR(50)
);