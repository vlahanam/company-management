ALTER TABLE employee_positions DROP FOREIGN KEY fk_employee_positions_employee;
ALTER TABLE employee_positions DROP FOREIGN KEY fk_employee_positions_position;

DROP TABLE IF EXISTS employee_positions;