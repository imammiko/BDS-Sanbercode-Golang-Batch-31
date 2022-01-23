CREATE TABLE admin_users (
  	id serial PRIMARY KEY,
	username VARCHAR ( 50 ) UNIQUE NOT NULL,
	password VARCHAR ( 50 ) NOT NULL,
	email VARCHAR ( 255 ) UNIQUE NOT NULL,
	activation_code varchar  (20) ,
    last_login TIMESTAMP ,
    active boolean,
    first_name VARCHAR (50),
    last_name VARCHAR (50)
);