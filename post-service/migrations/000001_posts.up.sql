CREATE TABLE posts (
    id INT AUTO_INCREMENT PRIMARY KEY COMMENT 'Primary key - post ID',
    user_id INT NOT NULL COMMENT 'Foreign key referencing posts(id)',
    title VARCHAR(255) NOT NULL COMMENT 'Title of the post',
    content TEXT COMMENT 'Content body of the post',
    status ENUM('draft', 'published', 'archived') DEFAULT 'draft' COMMENT 'Post status',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'Post creation timestamp',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Post last update timestamp',
    deleted_at TIMESTAMP DEFAULT NULL COMMENT 'Soft delete timestamp'
);