CREATE TABLE role_permissions (
    id INT AUTO_INCREMENT PRIMARY KEY COMMENT 'Primary key',
    role_id INT NOT NULL COMMENT 'Foreign key referencing roles(id)',
    permission_id INT NOT NULL COMMENT 'Foreign key referencing permissions(id)',
    FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE,
    FOREIGN KEY (permission_id) REFERENCES permissions(id) ON DELETE CASCADE
);