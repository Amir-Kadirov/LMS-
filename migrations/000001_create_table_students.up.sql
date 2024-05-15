CREATE TABLE IF NOT EXISTS students (
  id uuid PRIMARY KEY,
  first_name varchar(50),
  last_name varchar(50),
  age integer,
  external_id varchar(50),
  phone varchar(50),
  mail varchar(50),
  pasword varchar(50),
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated timestamp,
  active boolean
);

ALTER TABLE IF EXISTS students
ALTER COLUMN pasword TYPE varchar;

ALTER TABLE IF EXISTS students
ADD CONSTRAINT unique_phone UNIQUE (phone);

ALTER TABLE IF EXISTS students
ADD CONSTRAINT unique_mail UNIQUE (mail);