-- ----------------------------
-- Table structure for blog_audios
-- ----------------------------
DROP TABLE IF EXISTS `blog_audios`;
CREATE TABLE `blog_audios`
(
    id            bigint auto_increment
        primary key,
    name          varchar(100)                       null comment '名称',
    cover         varchar(200)                       null comment '封面',
    author        varchar(200)                       null comment '作者信息',
    audio_url     varchar(200)                       null comment '媒体路径',
    audio_lrc_url varchar(200)                       null comment 'lrc路径',
    audio_type    int      default 0                 null comment '媒体类型
0：mp3',
    has_del       int      default 0                 null comment '逻辑删标志位',
    create_time   datetime                           null comment '创建时间',
    update_time   datetime default CURRENT_TIMESTAMP null comment '更新时间'
)
    comment '博客媒体表';