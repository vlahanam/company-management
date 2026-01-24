CREATE TABLE users (
    id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT 'Unique identifier for the user',
    full_name VARCHAR(100) NOT NULL COMMENT 'Full name of the user',
    hash_password VARCHAR(255) NOT NULL COMMENT 'Password',
    date_of_birth DATE COMMENT 'User date of birth',
    gender VARCHAR(10) COMMENT 'User gender (e.g., Male, Female, Other)',
    id_card_number VARCHAR(20) UNIQUE COMMENT 'ID Card or CCCD number',
    email VARCHAR(100) UNIQUE NOT NULL COMMENT 'User email address',
    phone_number VARCHAR(15) UNIQUE COMMENT 'User phone number',
    avatar VARCHAR(255) COMMENT 'URL or path to user avatar',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'Record creation timestamp',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Record update timestamp'
);
