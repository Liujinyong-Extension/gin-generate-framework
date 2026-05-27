CREATE TABLE `test` (
                        `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
                        `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '名称',
                        `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '内容',
                        `sorce` int(11) DEFAULT '0' COMMENT '分数',
                        `category` enum('apple','samsang','oppo') CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'apple' COMMENT '分类 苹果:apple  三星:samsang 步步高:oppo',
                        `created_at` datetime NOT NULL COMMENT '创建时间',
                        `updated_at` datetime NOT NULL COMMENT '更新时间',
                        `deleted_at` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '删除时间',
                        PRIMARY KEY (`id`),
                        KEY `idx_deleted_at` (`deleted_at`),
                        KEY `idx_title` (`title`),
                        KEY `idx_category` (`category`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;