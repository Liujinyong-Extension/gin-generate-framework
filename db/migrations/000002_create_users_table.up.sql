CREATE TABLE `gin_user` (
                            `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
                            `username` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '账号',
                            `password` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '密码',
                            `created_at` datetime NOT NULL COMMENT '创建时间',
                            `updated_at` datetime NOT NULL COMMENT '更新时间',
                            `deleted_at` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '删除时间',
                            PRIMARY KEY (`id`),
                            KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;