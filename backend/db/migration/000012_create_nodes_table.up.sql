CREATE TABLE nodes (
    id INT AUTO_INCREMENT PRIMARY KEY,
    workflows_id INT NOT NULL,
    workflow_node_id VARCHAR(255) NOT NULL,
    type VARCHAR(255) NOT NULL,
    data JSON,
    position_x FLOAT NOT NULL,
    position_y FLOAT NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (workflows_id) REFERENCES workflows (id) ON DELETE CASCADE
);
