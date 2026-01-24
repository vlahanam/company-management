CREATE TABLE roles (
    id INT AUTO_INCREMENT PRIMARY KEY COMMENT 'Unique role ID',
    name VARCHAR(50) NOT NULL UNIQUE COMMENT 'Role name (e.g., admin, editor)',
    description VARCHAR(255) COMMENT 'Short description of the role purpose',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'Timestamp when the role was created'
) COMMENT='Defines roles assigned to users';
