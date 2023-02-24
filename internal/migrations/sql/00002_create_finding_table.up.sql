CREATE TABLE `finding` (
      `id` varchar(55) NOT NULL,
      `info_id` varchar(55) NOT NULL,
      `rule_id` varchar(20) NOT NULL,
      `type` varchar(20) NOT NULL,
      `commit` varchar(40) NOT NULL,
      `created_at` timestamp NOT NULL,
      `updated_at` timestamp NOT NULL,
      `deleted_at` timestamp NULL DEFAULT NULL,
      PRIMARY KEY (`id`),
      KEY `idx_info_id` (`info_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
