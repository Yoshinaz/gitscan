CREATE TABLE `location` (
       `id` varchar(55) NOT NULL,
       `finding_id` varchar(55) NOT NULL,
       `path` varchar(150) NOT NULL,
       `lines` varchar(150) NOT NULL,
       `status` varchar(20) NOT NULL,
       `created_at` timestamp NOT NULL,
       `updated_at` timestamp NOT NULL,
       `deleted_at` timestamp NULL DEFAULT NULL,
       PRIMARY KEY (`id`),
       KEY `idx_finding_id` (`finding_id`),
       KEY `idx_path` (`path`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
