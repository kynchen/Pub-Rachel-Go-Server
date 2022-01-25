CREATE TABLE `system_users`
(
    `id`          bigint NOT NULL AUTO_INCREMENT,
    `name`        varchar(200) DEFAULT NULL COMMENT '名称',
    `sex`         int          DEFAULT '2' COMMENT '性别：1 男，2 女，0未知',
    `sign`        varchar(200) DEFAULT NULL COMMENT '签名',
    `avatar`      varchar(200) DEFAULT NULL COMMENT '头像',
    `nation`      varchar(100) DEFAULT NULL COMMENT '国家',
    `province`    varchar(200) DEFAULT NULL COMMENT '省份',
    `city`        varchar(200) DEFAULT NULL COMMENT '城市',
    `region`      varchar(100) DEFAULT NULL COMMENT '区',
    `has_del`     int          DEFAULT '0' COMMENT '逻辑删标志位\n0：正常\n1：删除',
    `create_time` datetime     DEFAULT NULL COMMENT '创建时间',
    `update_time` datetime     DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3

