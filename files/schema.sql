CREATE TABLE gophers(
   id CHAR(32),
   name VARCHAR(100) NOT NULL,
   email VARCHAR(90) NOT NULL,
   company VARCHAR(90),
   timestamp TIMESTAMP default current_timestamp,
   PRIMARY KEY( id )
);