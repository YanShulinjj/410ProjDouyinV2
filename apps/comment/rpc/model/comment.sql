USE douyin;
CREATE TABLE `comment` (
     `comment_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '评论ID',
     `user_id`    bigint(20) unsigned NOT NULL COMMENT '评论作者id',
     `video_id`   bigint(20) unsigned NOT NULL COMMENT '视频id',
     `content`    varchar(128) NOT NULL DEFAULT '' COMMENT '播放外链',
     `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
     `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
     PRIMARY KEY (`comment_id`),
     KEY `ix_update_time` (`update_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='评论表';
