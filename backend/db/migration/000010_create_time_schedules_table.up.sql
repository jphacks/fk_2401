CREATE TABLE time_schedules (
    id INT AUTO_INCREMENT PRIMARY KEY,
    device_condition_id INT NOT NULL,
    start_time VARCHAR(255) NOT NULL,
    end_time VARCHAR(255) NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (device_condition_id) REFERENCES device_conditions (id) ON DELETE CASCADE
);