CREATE TABLE companies (
    id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT 'Unique identifier for the company',
    name VARCHAR(150) NOT NULL COMMENT 'Company name',
    parent_id BIGINT DEFAULT NULL COMMENT 'Parent company ID (self-reference)',
    description TEXT COMMENT 'Detailed description of the company',
    founded_date DATE COMMENT 'Date when the company was established',
    address VARCHAR(255) COMMENT 'Company address',
    phone_number VARCHAR(20) COMMENT 'Company contact number',
    email VARCHAR(100) UNIQUE COMMENT 'Company email address',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'Record creation timestamp',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Record update timestamp',
    
    CONSTRAINT fk_companies_parent FOREIGN KEY (parent_id) REFERENCES companies(id)
        ON DELETE SET NULL
        ON UPDATE CASCADE
);
