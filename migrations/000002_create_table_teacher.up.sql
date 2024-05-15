CREATE TABLE IF NOT EXISTS teacher (
  id uuid PRIMARY KEY,
  first_name varchar(50),
  last_name varchar(50),
  subject_id uuid,
  start_work timestamp,
  phone varchar(50),
  mail varchar(50),
  created_at timestamp NOT NULL,
  updated timestamp
);

ALTER TABLE IF EXISTS teacher
ALTER COLUMN start_work TYPE varchar(50);

ALTER TABLE IF EXISTS teacher
ADD COLUMN password varchar;

ALTER TABLE IF EXISTS teacher
ADD CONSTRAINT unique_t_phone UNIQUE (phone);

ALTER TABLE IF EXISTS teacher
ADD CONSTRAINT unique_t_mail UNIQUE (mail);