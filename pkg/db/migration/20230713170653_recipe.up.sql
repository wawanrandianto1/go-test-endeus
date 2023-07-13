CREATE TABLE `categories` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `name` longtext,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `reseps` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `category_id` bigint unsigned DEFAULT NULL,
  `judul` longtext NOT NULL,
  `deskripsi` longtext,
  PRIMARY KEY (`id`),
  KEY `fk_reseps_category` (`category_id`),
  CONSTRAINT `fk_reseps_category` FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `bahans` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `resep_id` bigint unsigned NOT NULL,
  `deskripsi` longtext,
  PRIMARY KEY (`id`),
  KEY `fk_bahans_resep` (`resep_id`),
  CONSTRAINT `fk_bahans_resep` FOREIGN KEY (`resep_id`) REFERENCES `reseps` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `cara_masaks` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `resep_id` bigint unsigned NOT NULL,
  `video_url` longtext,
  `lama_masak` bigint unsigned DEFAULT NULL,
  `deskripsi` longtext,
  PRIMARY KEY (`id`),
  KEY `fk_cara_masaks_resep` (`resep_id`),
  CONSTRAINT `fk_cara_masaks_resep` FOREIGN KEY (`resep_id`) REFERENCES `reseps` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
