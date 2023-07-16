SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for bahans
-- ----------------------------
DROP TABLE IF EXISTS `bahans`;
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
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of bahans
-- ----------------------------
BEGIN;
INSERT INTO `bahans` VALUES (1, 1, 2, 'Aperiam dolor repellat quia animi distinctio nihil.', '2023-07-14 16:19:24', '2023-07-14 16:19:24');
INSERT INTO `bahans` VALUES (2, 2, 4, 'Vitae perspiciatis amet facilis in sit voluptatem provident a optio.', '2023-07-14 16:19:24', '2023-07-14 16:19:24');
INSERT INTO `bahans` VALUES (3, 3, 7, 'Nemo iste nulla explicabo.', '2023-07-14 16:19:24', '2023-07-14 16:19:24');
INSERT INTO `bahans` VALUES (4, 4, 7, 'Occaecati sit sit velit iure illo.', '2023-07-14 16:19:24', '2023-07-14 16:19:24');
INSERT INTO `bahans` VALUES (5, 5, 5, 'Velit a totam.', '2023-07-14 16:19:24', '2023-07-14 16:19:24');
INSERT INTO `bahans` VALUES (6, 6, 6, 'Quam commodi enim tenetur.', '2023-07-14 16:19:24', '2023-07-14 16:19:24');
COMMIT;

-- ----------------------------
-- Table structure for cara_buats
-- ----------------------------
DROP TABLE IF EXISTS `cara_buats`;
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
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of cara_buats
-- ----------------------------
BEGIN;
INSERT INTO `cara_buats` VALUES (1, 1, 51, 'Aliquid beatae non.', 'praesentium alias autem', '2023-07-14 16:19:24', '2023-07-14 16:19:24');
INSERT INTO `cara_buats` VALUES (2, 2, 18, 'Aliquam iusto at.', 'ex atque soluta', '2023-07-14 16:19:24', '2023-07-14 16:19:24');
INSERT INTO `cara_buats` VALUES (3, 3, 37, 'Velit quo dolor dolor quae optio odio maiores tenetur atque.', 'necessitatibus ea odio', '2023-07-14 16:19:24', '2023-07-14 16:19:24');
INSERT INTO `cara_buats` VALUES (4, 4, 56, 'Aliquid asperiores voluptatibus delectus reiciendis possimus rem harum quaerat.', 'quaerat occaecati odit', '2023-07-14 16:19:24', '2023-07-14 16:19:24');
INSERT INTO `cara_buats` VALUES (5, 5, 55, 'Maiores expedita ducimus doloremque nulla porro natus.', 'error nostrum magnam', '2023-07-14 16:19:24', '2023-07-14 16:19:24');
INSERT INTO `cara_buats` VALUES (6, 6, 47, 'Perferendis dicta molestiae molestiae ex ut quasi molestias natus unde.', 'esse facilis voluptatem', '2023-07-14 16:19:24', '2023-07-14 16:19:24');
COMMIT;

-- ----------------------------
-- Table structure for categories
-- ----------------------------
DROP TABLE IF EXISTS `categories`;
CREATE TABLE `categories` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of categories
-- ----------------------------
BEGIN;
INSERT INTO `categories` VALUES (1, 'Serba Ayam', '2023-07-14 16:19:24', '2023-07-14 16:19:24');
INSERT INTO `categories` VALUES (2, 'Serba Ikan', '2023-07-14 16:19:24', '2023-07-14 16:19:24');
INSERT INTO `categories` VALUES (3, 'Serba Daging', '2023-07-14 16:19:24', '2023-07-14 16:19:24');
COMMIT;

-- ----------------------------
-- Table structure for reseps
-- ----------------------------
DROP TABLE IF EXISTS `reseps`;
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
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb;

-- ----------------------------
-- Records of reseps
-- ----------------------------
BEGIN;
INSERT INTO `reseps` VALUES (1, 1, 'Serba Ayam Panggang', 'https://loremflickr.com/640/480?lock=7697108477935616', 'fuga eaque commodi commodi dolor quibusdam explicabo beatae commodi saepe', 1, '2023-07-14 16:19:24', '2023-07-14 16:19:24');
INSERT INTO `reseps` VALUES (2, 1, 'Serba Ayam Bakar', 'https://loremflickr.com/640/480?lock=1513686585835520', 'itaque aspernatur vitae laboriosam quas dignissimos aliquid necessitatibus nemo praesentium', 1, '2023-07-14 16:19:24', '2023-07-14 16:19:24');
INSERT INTO `reseps` VALUES (3, 1, 'Serba Ayam Bumbu Kecap', 'https://loremflickr.com/640/480?lock=3833002389929984', 'dolores illo quasi minus magnam beatae ex tenetur hic animi', 1, '2023-07-14 16:19:24', '2023-07-14 16:19:24');
INSERT INTO `reseps` VALUES (4, 2, 'Serba Ikan Panggang', 'https://picsum.photos/seed/IwCUqsY/640/480', 'voluptate vel minus numquam adipisci facilis hic ullam aliquid quam', 1, '2023-07-14 16:19:24', '2023-07-14 16:19:24');
INSERT INTO `reseps` VALUES (5, 2, 'Serba Ikan Bakar', 'https://loremflickr.com/640/480?lock=2860152623464448', 'atque itaque natus animi blanditiis dolore adipisci soluta unde aut', 1, '2023-07-14 16:19:24', '2023-07-14 16:19:24');
INSERT INTO `reseps` VALUES (6, 3, 'Serba Daging Panggang', 'https://loremflickr.com/640/480?lock=7292917993439232', 'assumenda consequuntur repellendus odit numquam eum sequi vitae ad repudiandae', 1, '2023-07-14 16:19:24', '2023-07-14 16:19:24');
COMMIT;

-- ----------------------------
-- Table structure for schema_migrations
-- ----------------------------
DROP TABLE IF EXISTS `schema_migrations`;
CREATE TABLE `schema_migrations` (
  `version` bigint NOT NULL,
  `dirty` tinyint(1) NOT NULL,
  PRIMARY KEY (`version`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of schema_migrations
-- ----------------------------
BEGIN;
INSERT INTO `schema_migrations` VALUES (20230714162044, 0);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
