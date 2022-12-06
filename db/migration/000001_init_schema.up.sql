CREATE TABLE users (
    id bigint PRIMARY KEY AUTO_INCREMENT,
    email varchar(255) NOT NULL,
    name varchar(255) NOT NULL,
    `password` varchar(255) NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp NULL,
    is_deleted boolean NOT NULL DEFAULT false
);

CREATE INDEX users_email_index ON users (email);
s