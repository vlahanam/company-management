CREATE TABLE role_permissions (
    role_id INT NOT NULL COMMENT 'Reference to roles.id',
    permission_id INT NOT NULL COMMENT 'Reference to permissions.id',
    granted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'Timestamp when the permission was granted to the role',

    CONSTRAINT pk_role_permission PRIMARY KEY (role_id, permission_id),
    CONSTRAINT fk_role_permissions_role FOREIGN KEY (role_id) REFERENCES roles(id) 
        ON DELETE CASCADE 
        ON UPDATE CASCADE,
    CONSTRAINT fk_role_permissions_permission FOREIGN KEY (permission_id) REFERENCES permissions(id) 
        ON DELETE CASCADE 
        ON UPDATE CASCADE
) COMMENT='Many-to-many relationship between roles and permissions';
