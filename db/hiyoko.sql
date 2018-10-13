CREATE DATABASE IF NOT EXISTS Hiyoko_core;

DROP TABLE IF EXISTS Users;
CREATE TABLE Users (
  `user_id` VARCHAR(50) NOT NULL, -- same with line id
  `created_at` datetime NOT NULL,
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS Vocabularies;
CREATE TABLE Vocabularies (
  `voca_id` BIGINT(20) unsigned AUTO_INCREMENT NOT NULL,
  `name` VARCHAR(100) NOT NULL,
  PRIMARY KEY (`voca_id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS Vocabulary_lists;
CREATE TABLE Vocabulary_lists (
  `user_id` VARCHAR(50) NOT NULL,
  `voca_id` BIGINT(20) NOT NULL,
  `meaning` VARCHAR(200) DEFAULT NULL,
  `context_sentence` VARCHAR(500) DEFAULT NULL,
  `context_picture_URL` VARCHAR(500) DEFAULT NULL,
  `created_at` datetime NOT NULL,
  PRIMARY KEY (`user_id`, `voca_id`),
  KEY `user_id_created_at` (`user_id`, `created_at`),
  KEY `user_id` (`user_id`),
  KEY `voca_id` (`voca_id`),
  KEY `created_at` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
