ALTER TABLE role_permissions DROP FOREIGN KEY fk_role_permissions_role;
ALTER TABLE role_permissions DROP FOREIGN KEY fk_role_permissions_permission;
DROP TABLE IF EXISTS role_permissions;