CREATE TABLE permissions (
    id INT AUTO_INCREMENT PRIMARY KEY COMMENT 'Unique permission ID',
    name VARCHAR(100) NOT NULL UNIQUE COMMENT 'Permission name (e.g., user.create, post.delete)',
    description VARCHAR(255) COMMENT 'Description of what the permission allows',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'Timestamp when the permission was created'
) COMMENT='Defines system permissions or actions that can be granted to roles';
