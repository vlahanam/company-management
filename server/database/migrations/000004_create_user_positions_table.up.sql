CREATE TABLE user_positions (
    id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT 'Unique identifier for the user-position relation',
    user_id BIGINT NOT NULL COMMENT 'Reference to the user',
    position_id BIGINT NOT NULL COMMENT 'Reference to the position',
    start_date DATE NOT NULL COMMENT 'Date when the user started this position',
    end_date DATE DEFAULT NULL COMMENT 'Date when the user left this position (NULL if current)',
    is_primary BOOLEAN DEFAULT FALSE COMMENT 'Indicates if this is the main/primary position',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'Record creation timestamp',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Record update timestamp',
    
    CONSTRAINT fk_user_positions_user FOREIGN KEY (user_id) REFERENCES users(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE,
    CONSTRAINT fk_user_positions_position FOREIGN KEY (position_id) REFERENCES positions(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE,
    CONSTRAINT uq_user_position UNIQUE (user_id, position_id, start_date)
);
