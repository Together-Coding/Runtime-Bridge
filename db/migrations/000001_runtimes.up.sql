CREATE TABLE IF NOT EXISTS runtime_allocations
(
    `id`               BIGINT AUTO_INCREMENT PRIMARY KEY,
    `user_id`          BIGINT            NOT NULL, -- fk
    `user_ip`          VARCHAR(45),

    `runtime_image_id` BIGINT            NOT NULL, -- fk, ix
    `cont_launched_at` DATETIME(6),
    `cont_ip`          VARCHAR(45),
    `cont_port`        INT,
    `cont_user`        VARCHAR(20),
    `cont_auth_type`   VARCHAR(20),
    `cont_auth`        TEXT,
    `cont_api_key`     VARCHAR(128)      NOT NULL, -- ix
    `health`           TINYINT DEFAULT 0 NOT NULL,

    `created_at`       DATETIME(6)       NOT NULL,
    `updated_at`       DATETIME(6)       NOT NULL
);

CREATE TABLE IF NOT EXISTS runtime_images
(
    `id`            BIGINT AUTO_INCREMENT PRIMARY KEY,
    `name`          VARCHAR(255) NOT NULL,
    `language_name` VARCHAR(20)  NOT NULL, -- fk
    `url`           VARCHAR(255) NOT NULL,
    `tag`           VARCHAR(255) NOT NULL,
    `available`     bool         NOT NULL DEFAULT FALSE,
    `created_at`    DATETIME(6)  NOT NULL
);

CREATE TABLE IF NOT EXISTS supported_languages
(
    `name`  VARCHAR(20) NOT NULL PRIMARY KEY,
    `order` INT         NOT NULL DEFAULT 0
);