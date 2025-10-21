CREATE TABLE employee_roles (
    employee_id BIGINT NOT NULL COMMENT 'Reference to employees.id',
    role_id INT NOT NULL COMMENT 'Reference to roles.id',
    assigned_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'Timestamp when the role was assigned to the employee',

    CONSTRAINT pk_employee_roles PRIMARY KEY (employee_id, role_id),
    CONSTRAINT fk_employee_roles_employee FOREIGN KEY (employee_id) REFERENCES employees(id) 
        ON DELETE CASCADE 
        ON UPDATE CASCADE,
    CONSTRAINT fk_employee_roles_role FOREIGN KEY (role_id) REFERENCES roles(id) 
        ON DELETE CASCADE 
        ON UPDATE CASCADE
) COMMENT='Many-to-many relationship between employees and roles';
