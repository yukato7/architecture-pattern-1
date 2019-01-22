CREATE TABLE `users` (
  `id` VARCHAR(28) NOT NULL,
  `name` VARCHAR(50) NOT NULL,
  `icon_url` VARCHAR,
  `created_at` BIGINT(20) NULL,
  `updated_at` BIGINT(20) NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `charge_logs` (
  `id` INT UNSIGNED NOT NULL,
  `user_id` VARCHAR(28) NOT NULL,
  `amount` INT NOT NULL,
  `currency` VARCHAR(3) NOT NULL,
  PRIMARY KEY (`id`, `user_id`)
  CONSTRAINT `fk_charge_logs_users`
    FOREIGN KEY (`user_id`)
    REFERENCES `users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;