CREATE TABLE m304_record (
    id INT AUTO_INCREMENT PRIMARY KEY,
    m304_id INT NOT NULL,
    device_condition_id INT NOT NULL,
    `block` VARCHAR(1) NOT NULL,
    valid BOOLEAN NOT NULL,
    position INT NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (m304_id) REFERENCES m304 (id) ON DELETE RESTRICT,
    FOREIGN KEY (device_condition_id) REFERENCES device_conditions (id) ON DELETE RESTRICT
);