-- Clean seed data
DELETE FROM user_roles;
DELETE FROM users;

-- Reset auto-increment IDs
ALTER TABLE users AUTO_INCREMENT = 1;

-- add data to the users table
INSERT INTO users (username, email, password_hash)
VALUES
    ('admin', 'admin@yopmail.com', SHA2('admin-salt', 256)),
    ('user-1', 'user-1@yopmail.com', SHA2('user-1-salt', 256)),
    ('guest-1', 'guest-1@yopmail.com', SHA2('guest-1-salt', 256));

-- add data to the user_roles table
INSERT INTO user_roles (user_id, role_id)
VALUES
    (1, 1),  -- admin -> admin
    (2, 2),  -- user-1 -> user
    (3, 3);  -- guest-1 -> guest