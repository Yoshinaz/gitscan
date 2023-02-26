CREATE TABLE `info` (
      `id` varchar(55) NOT NULL,
      `name` varchar(150) NOT NULL,
      `url` varchar(150) NOT NULL,
      `rule_set` varchar(40) NOT NULL,
      `status` varchar(20) NOT NULL,
      `commit` varchar(40) NOT NULL,
      `description` varchar(150) NOT NULL,
      `enqueued_at` timestamp NOT NULL,
      `started_at` timestamp NULL DEFAULT NULL,
      `finished_at` timestamp NULL DEFAULT NULL,
      `created_at` timestamp NOT NULL,
      `updated_at` timestamp NOT NULL,
      `deleted_at` timestamp NULL DEFAULT NULL,
      PRIMARY KEY (`id`),
      KEY `idx_url` (`url`),
      KEY `idx_commit` (`commit`),
      KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;