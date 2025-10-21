ALTER TABLE user_positions DROP FOREIGN KEY fk_user_roles_user;
ALTER TABLE user_positions DROP FOREIGN KEY fk_user_roles_role;
DROP TABLE IF EXISTS user_positions;