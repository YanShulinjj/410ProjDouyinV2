USE douyin;
CREATE TABLE `video` (
        `video_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '视频ID',
        `user_id`  bigint(20) unsigned NOT NULL COMMENT '视频作者id',
        `play_url` varchar(128) NOT NULL DEFAULT '' COMMENT '播放外链',
        `cover_url` varchar(128) NOT NULL DEFAULT '' COMMENT '封面外链',
        `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
        `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
        PRIMARY KEY (`video_id`),
        KEY `ix_update_time` (`update_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='视频表';

