CREATE TABLE devices (
    id INT AUTO_INCREMENT PRIMARY KEY,
    house_id INT NOT NULL,
    climate_data_id INT NOT NULL,
    m304_id INT NOT NULL,
    sensor_id INT NOT NULL,
    device_name VARCHAR(255),
    rly INT,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (house_id) REFERENCES houses (id) ON DELETE CASCADE,
    FOREIGN KEY (climate_data_id) REFERENCES climate_datas (id) ON DELETE RESTRICT,
    FOREIGN KEY (m304_id) REFERENCES m304 (id) ON DELETE RESTRICT,
    FOREIGN KEY (uecs_device_id) REFERENCES uecs_devices (id) ON DELETE RESTRICT
);
