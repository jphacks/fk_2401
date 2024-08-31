CREATE TABLE devices (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    house_id BIGINT NOT NULL,
    climate_data_id BIGINT NOT NULL,
    duration INT,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (house_id) REFERENCES houses (id) ON DELETE CASCADE,
    FOREIGN KEY (climate_data_id) REFERENCES climate_datas (id) ON DELETE RESTRICT
);
