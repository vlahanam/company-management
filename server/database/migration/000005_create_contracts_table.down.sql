
ALTER TABLE contracts DROP FOREIGN KEY fk_contracts_employee;
ALTER TABLE contracts DROP FOREIGN KEY fk_contracts_company;
ALTER TABLE contracts DROP FOREIGN KEY fk_contracts_position;

DROP TABLE IF EXISTS contracts;