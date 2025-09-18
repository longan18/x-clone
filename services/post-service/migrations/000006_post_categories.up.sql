CREATE TABLE post_categories (
    id INT AUTO_INCREMENT PRIMARY KEY COMMENT 'Primary key - category ID',
    category_name VARCHAR(100) NOT NULL UNIQUE COMMENT 'Category name',
    description TEXT COMMENT 'Description of the category',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'Category creation timestamp',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Last update timestamp',
    deleted_at TIMESTAMP DEFAULT NULL COMMENT 'Soft delete timestamp'
);
