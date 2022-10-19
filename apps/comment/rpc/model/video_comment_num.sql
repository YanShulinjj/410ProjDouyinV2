CREATE DATABASE IF NOT EXISTS douyin;
USE douyin;
CREATE TABLE `video_comment_num` (
          `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
          `video_id`  bigint(20) unsigned NOT NULL COMMENT '视频id',
          `comment_num` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '视频的点赞数',
          `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
          `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
          PRIMARY KEY (`id`),
          UNIQUE KEY `uniq_like` (`video_id`),
          KEY `ix_update_time` (`update_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='视频评论数量表';