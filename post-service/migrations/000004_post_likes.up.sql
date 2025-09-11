CREATE TABLE post_likes (
    id INT AUTO_INCREMENT PRIMARY KEY COMMENT 'Primary key',
    post_id INT NOT NULL COMMENT 'Foreign key referencing posts(id)',
    user_id INT NOT NULL COMMENT 'User ID from auth-service',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'Timestamp when like was made',
    FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE
);
