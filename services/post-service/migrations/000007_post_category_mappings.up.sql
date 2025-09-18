CREATE TABLE post_category_mappings (
    id INT AUTO_INCREMENT PRIMARY KEY COMMENT 'Primary key',
    post_id INT NOT NULL COMMENT 'Foreign key referencing posts(id)',
    category_id INT NOT NULL COMMENT 'Foreign key referencing post_categories(id)',
    FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
    FOREIGN KEY (category_id) REFERENCES post_categories(id) ON DELETE CASCADE
);