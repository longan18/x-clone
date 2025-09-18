CREATE TABLE user_profiles (
    id INT AUTO_INCREMENT PRIMARY KEY COMMENT 'Primary key - user profile ID',
    user_id INT NOT NULL UNIQUE COMMENT 'User ID from auth-service',
    full_name VARCHAR(100) COMMENT 'Full name of the user',
    avatar_url VARCHAR(255) COMMENT 'URL link to user avatar image',
    phone VARCHAR(20) COMMENT 'User phone number',
    address TEXT COMMENT 'User address',
    dob DATE COMMENT 'Date of birth',
    gender ENUM('male', 'female', 'other') COMMENT 'Gender of the user',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'Profile creation timestamp',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Profile last update timestamp',
    deleted_at TIMESTAMP DEFAULT NULL COMMENT 'Soft delete timestamp'
);