CREATE TABLE IF NOT EXISTS administrators (
	`id` INT NOT NULL AUTO_INCREMENT,
	`name` VARCHAR (255) NOT NULL,
	`email` VARCHAR (255) NOT NULL,
	`password` VARCHAR (255) NOT NULL,
	`created_by` INT NULL,
    `updated_by` INT NULL,
	`created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	`updated_at` TIMESTAMP NULL,
	`deleted_at` TIMESTAMP NULL,
	PRIMARY KEY (`id`),
	UNIQUE KEY admin_email_unique (`email`),
    INDEX idx_admin_email (`email`),
    INDEX idx_admin_created_by (`created_by`),
    INDEX idx_admin_updated_by (`updated_by`),
    CONSTRAINT FK_admin_created_by FOREIGN KEY (`created_by`) REFERENCES administrators(`id`) ON DELETE SET NULL,
    CONSTRAINT FK_admin_updated_by FOREIGN KEY (`updated_by`) REFERENCES administrators(`id`) ON DELETE SET NULL
) ENGINE = INNODB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8;