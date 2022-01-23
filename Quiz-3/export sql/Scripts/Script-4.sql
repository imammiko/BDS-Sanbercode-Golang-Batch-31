ALTER TABLE admin_users
ADD COLUMN salt varchar(50),
add column forgotten_password_code varchar (50),
add column remember_code varchar (50),
add column created_on TIMESTAMP;
