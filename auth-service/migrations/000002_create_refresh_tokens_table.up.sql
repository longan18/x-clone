CREATE TABLE refresh_tokens (
    id INT AUTO_INCREMENT PRIMARY KEY COMMENT 'Primary key',
    user_id INT NOT NULL COMMENT 'Foreign key referencing users(id)',
    token VARCHAR(255) NOT NULL COMMENT 'Refresh token string',
    expiry_date TIMESTAMP NOT NULL COMMENT 'Expiration date of the token',
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);