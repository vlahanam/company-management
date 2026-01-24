ALTER TABLE user_positions DROP FOREIGN KEY fk_user_positions_user;
ALTER TABLE user_positions DROP FOREIGN KEY fk_user_positions_position;

DROP TABLE IF EXISTS user_positions;