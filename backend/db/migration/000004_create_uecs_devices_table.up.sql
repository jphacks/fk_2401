CREATE TABLE uecs_devices (
    id INT AUTO_INCREMENT PRIMARY KEY,
    ccmtype VARCHAR(255) NOT NULL,
    room INT NOT NULL,
    region INT NOT NULL,
    `order` INT NOT NULL,
    `priority` INT NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);