CREATE TABLE post_attachments (
    id INT AUTO_INCREMENT PRIMARY KEY COMMENT 'Primary key',
    post_id INT NOT NULL COMMENT 'Foreign key referencing posts(id)',
    file_url VARCHAR(255) NOT NULL COMMENT 'URL of the attached file',
    file_name VARCHAR(255) COMMENT 'Name of the attached file',
    uploaded_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'Timestamp when file was uploaded',
    deleted_at TIMESTAMP DEFAULT NULL COMMENT 'Soft delete timestamp',
    FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE
);
