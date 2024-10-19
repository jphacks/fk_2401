CREATE TABLE uecs_devices (
    id INT AUTO_INCREMENT PRIMARY KEY,
    ccmtype VARCHAR(255),
    room INT,
    region INT,
    `order` INT,
    `priority` INT,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);