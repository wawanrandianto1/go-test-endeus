CREATE TABLE `categories` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `reseps` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `category_id` bigint NOT NULL,
  `judul` text NOT NULL,
  `video_url` text,
  `deskripsi` text NOT NULL,
  `publish` tinyint(1) DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `category_id` (`category_id`),
  CONSTRAINT `reseps_ibfk_1` FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `bahans` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `resep_id` bigint DEFAULT NULL,
  `porsi` int DEFAULT NULL,
  `deskripsi` text,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `resep_id` (`resep_id`),
  CONSTRAINT `bahans_ibfk_1` FOREIGN KEY (`resep_id`) REFERENCES `reseps` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `cara_buats` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `resep_id` bigint DEFAULT NULL,
  `lama_waktu` int DEFAULT NULL,
  `deskripsi` text,
  `tips` text,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `resep_id` (`resep_id`),
  CONSTRAINT `cara_buats_ibfk_1` FOREIGN KEY (`resep_id`) REFERENCES `reseps` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
