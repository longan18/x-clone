-- Clean seed data
DELETE FROM role_permissions;
DELETE FROM permissions;
DELETE FROM roles;

-- Reset auto-increment IDs
ALTER TABLE roles AUTO_INCREMENT = 1;
ALTER TABLE permissions AUTO_INCREMENT = 1;