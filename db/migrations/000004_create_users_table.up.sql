CREATE TABLE `test`
(
    `id`         int      NOT NULL COMMENT '主键',
    `name`       varchar(255) NULL COMMENT '名称',
    `password`        varchar(255) NULL DEFAULT 0 COMMENT '密码',
    `created_at` datetime NOT NULL COMMENT '创建时间',
    `updated_at` datetime NOT NULL COMMENT '更新时间',
    `deleted_at` varchar(255) NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    INDEX        `idx_deleted_at`(`deleted_at`)
);