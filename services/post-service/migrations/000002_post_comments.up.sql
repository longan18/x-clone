CREATE TABLE post_comments (
    id INT AUTO_INCREMENT PRIMARY KEY COMMENT 'Primary key - comment ID',
    post_id INT NOT NULL COMMENT 'Foreign key referencing posts(id)',
    user_id INT NOT NULL COMMENT 'User ID from auth-service',
    comment_text TEXT NOT NULL COMMENT 'Text content of the comment',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'Comment creation timestamp',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Last update timestamp',
    deleted_at TIMESTAMP DEFAULT NULL COMMENT 'Soft delete timestamp',
    FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE
);