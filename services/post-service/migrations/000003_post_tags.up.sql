CREATE TABLE post_tags (
    id INT AUTO_INCREMENT PRIMARY KEY COMMENT 'Primary key',
    post_id INT NOT NULL COMMENT 'Foreign key referencing posts(id)',
    tag_name VARCHAR(50) NOT NULL COMMENT 'Tag name for categorizing posts',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'Tag creation timestamp',
    FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE
);