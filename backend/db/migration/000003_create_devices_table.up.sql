CREATE TABLE devices (
    id INT AUTO_INCREMENT PRIMARY KEY,
    house_id INT NOT NULL,
    climate_data_id INT NOT NULL,
    set_point FLOAT,
    duration INT,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (house_id) REFERENCES houses (id) ON DELETE CASCADE,
    FOREIGN KEY (climate_data_id) REFERENCES climate_datas (id) ON DELETE RESTRICT
);
