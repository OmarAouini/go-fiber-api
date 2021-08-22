CREATE TABLE `user` (
  `id` varchar(128) NOT NULL,
  `name` varchar(128) NOT NULL,
  `surname` varchar(128) NOT NULL,
  `username` varchar(128) NOT NULL,
  `password` varchar(128) NOT NULL,
  `created_at` date NOT NULL,
  `updated_at` date NOT NULL,
  `last_login` date NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uuid index` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci	
