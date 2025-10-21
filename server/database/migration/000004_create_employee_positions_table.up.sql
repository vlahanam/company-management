CREATE TABLE employee_positions (
    id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT 'Unique identifier for the employee-position relation',
    employee_id BIGINT NOT NULL COMMENT 'Reference to the employee',
    position_id BIGINT NOT NULL COMMENT 'Reference to the position',
    start_date DATE NOT NULL COMMENT 'Date when the employee started this position',
    end_date DATE DEFAULT NULL COMMENT 'Date when the employee left this position (NULL if current)',
    is_primary BOOLEAN DEFAULT FALSE COMMENT 'Indicates if this is the main/primary position',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'Record creation timestamp',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Record update timestamp',
    
    CONSTRAINT fk_employee_positions_employee FOREIGN KEY (employee_id) REFERENCES employees(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE,
    CONSTRAINT fk_employee_positions_position FOREIGN KEY (position_id) REFERENCES positions(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE,
    CONSTRAINT uq_employee_position UNIQUE (employee_id, position_id, start_date)
);
