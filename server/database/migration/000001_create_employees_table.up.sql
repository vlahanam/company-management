CREATE TABLE employees (
    id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT 'Unique identifier for the employee',
    full_name VARCHAR(100) NOT NULL COMMENT 'Full name of the employee',
    date_of_birth DATE COMMENT 'Employee date of birth',
    gender VARCHAR(10) COMMENT 'Employee gender (e.g., Male, Female, Other)',
    id_card_number VARCHAR(20) UNIQUE NOT NULL COMMENT 'ID Card or CCCD number',
    email VARCHAR(100) UNIQUE NOT NULL COMMENT 'Employee email address',
    phone_number VARCHAR(15) UNIQUE NOT NULL COMMENT 'Employee phone number',
    avatar VARCHAR(255) COMMENT 'URL or path to employee avatar',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'Record creation timestamp',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Record update timestamp'
);
