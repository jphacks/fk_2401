CREATE TABLE edges (
    id INT AUTO_INCREMENT PRIMARY KEY,
    workflows_id INT NOT NULL,
    source_node_id VARCHAR(255) NOT NULL,
    target_node_id VARCHAR(255) NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (workflows_id) REFERENCES workflows (id) ON DELETE CASCADE
);
