CREATE TABLE `test`
(
    `id`         int      NOT NULL COMMENT '主键',
    `title`      varchar(255) NULL COMMENT '名称',
    `content`    text NULL COMMENT '内容',
    `sorce`      int(11) NULL DEFAULT 0 COMMENT '分数',
    `category`   enum('apple','samsang','oppo') NOT NULL DEFAULT 'apple' COMMENT '分类 苹果:apple  三星:samsang 步步高:oppo',
    `created_at` datetime NOT NULL COMMENT '创建时间',
    `updated_at` datetime NOT NULL COMMENT '更新时间',
    `deleted_at` varchar(255) NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    INDEX        `idx_deleted_at`(`deleted_at`)
);