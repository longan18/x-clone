CREATE TABLE permissions (
    id INT AUTO_INCREMENT PRIMARY KEY COMMENT 'Primary key - permission ID',
    permission_name VARCHAR(100) NOT NULL UNIQUE COMMENT 'Permission name',
    description TEXT COMMENT 'Description of the permission',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'Permission creation timestamp',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Last update timestamp',
    deleted_at TIMESTAMP DEFAULT NULL COMMENT 'Soft delete timestamp'
);