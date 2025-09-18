-- Clean seed data
DELETE FROM role_permissions;
DELETE FROM permissions;
DELETE FROM roles;

-- Reset auto-increment IDs
ALTER TABLE roles AUTO_INCREMENT = 1;
ALTER TABLE permissions AUTO_INCREMENT = 1;

-- add data to the roles table
INSERT INTO roles (role_name, description)
VALUES
    ('admin', 'Administrator with full access'),
    ('user', 'Standard user with basic access'),
    ('guest', 'Limited access for unauthenticated users');

-- add data to the permissions table
INSERT INTO permissions (permission_name, description)
VALUES
    ('view_users', 'Permission to view user list'),
    ('edit_users', 'Permission to edit user profiles'),
    ('delete_users', 'Permission to delete users');

INSERT INTO permissions (permission_name, description)
VALUES
	('view_posts', 'Permission to create new posts'),
    ('create_posts', 'Permission to create new posts'),
    ('edit_posts', 'Permission to edit posts'),
    ('delete_posts', 'Permission to delete posts');

-- add data to the role_permissions table
-- assign permissions to admin role (id = 1)
INSERT INTO role_permissions (role_id, permission_id)
VALUES
    (1, 1),  -- admin -> view_users
    (1, 2),  -- admin -> edit_users
    (1, 3),  -- admin -> delete_users
	(1, 4),  -- admin -> view_posts
	(1, 5),  -- admin -> create_posts
    (1, 6),  -- admin -> edit_posts
    (1, 7);  -- admin -> delete_posts

-- assign permissions to user role (id = 2)
INSERT INTO role_permissions (role_id, permission_id)
VALUES
	(2, 4),  -- user -> view_posts
	(2, 5),  -- user -> create_posts
    (2, 6),  -- user -> edit_posts
    (2, 7);  -- user -> delete_posts

-- assign permissions to guest role (id = 3)
INSERT INTO role_permissions (role_id, permission_id)
VALUES
	(3, 4);  -- guest -> view_posts