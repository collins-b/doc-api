CREATE DATABASE docapi;
USE docapi;
CREATE TABLE user (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    firstName VARCHAR(50) NOT NULL,
    lastName VARCHAR(50) NOT NULL,
    CreatedOn timestamp NOT NULL DEFAULT current_timestamp
);

CREATE TABLE document (
    id INT AUTO_INCREMENT PRIMARY KEY,
    owner VARCHAR(50) NOT NULL,
    title VARCHAR(50) NOT NULL,
    content VARCHAR(250) NOT NULL,
    CreatedOn timestamp NOT NULL DEFAULT current_timestamp
);
