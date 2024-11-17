CREATE TABLE device_conditions (
    id INT AUTO_INCREMENT PRIMARY KEY,
    device_id INT NOT NULL,
    operation_id INT NOT NULL,
    valid BOOLEAN NOT NULL DEFAULT FALSE,
    set_point FLOAT,
    duration INT,
    operator INT,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP  ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (device_id) REFERENCES devices (id) ON DELETE CASCADE,
    FOREIGN KEY (operation_id) REFERENCES operations (id) ON DELETE CASCADE
);