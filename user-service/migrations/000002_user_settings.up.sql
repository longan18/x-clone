CREATE TABLE user_settings (
    id INT AUTO_INCREMENT PRIMARY KEY COMMENT 'Primary key',
    user_id INT NOT NULL COMMENT 'User ID from auth-service',
    setting_key VARCHAR(100) NOT NULL COMMENT 'Key of the user setting',
    setting_value TEXT COMMENT 'Value of the user setting',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'Setting creation timestamp',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Last update timestamp',
    deleted_at TIMESTAMP DEFAULT NULL COMMENT 'Soft delete timestamp'
);