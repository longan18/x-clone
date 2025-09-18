CREATE TABLE roles (
    id INT AUTO_INCREMENT PRIMARY KEY COMMENT 'Primary key - role ID',
    role_name VARCHAR(50) NOT NULL UNIQUE COMMENT 'Role name (e.g. admin, user)',
    description TEXT COMMENT 'Description of the role',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'Role creation timestamp',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Last update timestamp',
    deleted_at TIMESTAMP DEFAULT NULL COMMENT 'Soft delete timestamp'
);
