ALTER TABLE employee_positions DROP FOREIGN KEY fk_employee_roles_employee;
ALTER TABLE employee_positions DROP FOREIGN KEY fk_employee_roles_role;
DROP TABLE IF EXISTS employee_positions;