CREATE TABLE positions (
    id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT 'Unique identifier for the position',
    company_id BIGINT NOT NULL COMMENT 'Reference to the company',
    name VARCHAR(100) NOT NULL COMMENT 'Position title or role name',
    description TEXT COMMENT 'Detailed description of the position',
    level INT COMMENT 'Job level (e.g., Junior = 1, Senior = 2, Manager = 3)',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'Record creation timestamp',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Record update timestamp',
    
    CONSTRAINT fk_positions_company FOREIGN KEY (company_id) REFERENCES companies(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);
