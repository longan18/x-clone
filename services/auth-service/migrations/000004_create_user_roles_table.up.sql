CREATE TABLE user_roles (
    id INT AUTO_INCREMENT PRIMARY KEY COMMENT 'Primary key',
    user_id INT NOT NULL COMMENT 'Foreign key referencing users(id)',
    role_id INT NOT NULL COMMENT 'Foreign key referencing roles(id)',
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE
);
