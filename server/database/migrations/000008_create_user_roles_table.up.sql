CREATE TABLE user_roles (
    user_id BIGINT NOT NULL COMMENT 'Reference to users.id',
    role_id INT NOT NULL COMMENT 'Reference to roles.id',
    assigned_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'Timestamp when the role was assigned to the user',

    CONSTRAINT pk_user_roles PRIMARY KEY (user_id, role_id),
    CONSTRAINT fk_user_roles_user FOREIGN KEY (user_id) REFERENCES users(id) 
        ON DELETE CASCADE 
        ON UPDATE CASCADE,
    CONSTRAINT fk_user_roles_role FOREIGN KEY (role_id) REFERENCES roles(id) 
        ON DELETE CASCADE 
        ON UPDATE CASCADE
) COMMENT='Many-to-many relationship between users and roles';
