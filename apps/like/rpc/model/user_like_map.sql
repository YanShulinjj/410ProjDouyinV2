CREATE DATABASE IF NOT EXISTS douyin;
USE douyin;
CREATE TABLE `user_like_map` (
          `like_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '关系ID',
          `user_id` bigint(20) unsigned NOT NULL COMMENT '用户id',
          `video_id` bigint(20) unsigned NOT NULL COMMENT '视频id',
          `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
          `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
          PRIMARY KEY (`like_id`),
          UNIQUE KEY `uniq_like` (`user_id`, `video_id`),
          KEY `ix_update_time` (`update_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='点赞映射表';