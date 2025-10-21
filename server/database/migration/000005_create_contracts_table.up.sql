CREATE TABLE contracts (
    id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT 'Unique identifier for the contract',
    user_id BIGINT NOT NULL COMMENT 'Reference to the user',
    company_id BIGINT NOT NULL COMMENT 'Reference to the company',
    position_id BIGINT DEFAULT NULL COMMENT 'Reference to the user position',
    contract_number VARCHAR(50) UNIQUE NOT NULL COMMENT 'Contract reference number',
    contract_type ENUM('Probation', 'Fixed-term', 'Permanent', 'Freelance', 'Internship') DEFAULT 'Fixed-term' COMMENT 'Type of contract',
    start_date DATE NOT NULL COMMENT 'Contract start date',
    end_date DATE COMMENT 'Contract end date (nullable for permanent contracts)',
    salary DECIMAL(15,2) NOT NULL COMMENT 'Monthly salary amount',
    status ENUM('Active', 'Expired', 'Terminated', 'Pending') DEFAULT 'Pending' COMMENT 'Current contract status',
    file_path VARCHAR(500) NULL COMMENT 'Link to store contract in s3',
    notes TEXT COMMENT 'Additional notes about the contract',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'Record creation timestamp',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Record update timestamp',
    
    CONSTRAINT fk_contracts_user FOREIGN KEY (user_id) REFERENCES users(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE,
    CONSTRAINT fk_contracts_company FOREIGN KEY (company_id) REFERENCES companies(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE,
    CONSTRAINT fk_contracts_position FOREIGN KEY (position_id) REFERENCES positions(id)
        ON DELETE SET NULL
        ON UPDATE CASCADE
);

