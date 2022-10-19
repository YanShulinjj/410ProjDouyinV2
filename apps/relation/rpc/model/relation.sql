CREATE DATABASE IF NOT EXISTS douyin;
USE douyin;
CREATE TABLE `relation` (
     `relation_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '关系ID',
     `user_id`  bigint(20) unsigned NOT NULL COMMENT '作者id',
     `to_user_ids` varchar(128) NOT NULL DEFAULT '' COMMENT '关系用户的id，形如 1,2,3',
     `type` bool NOT NULL DEFAULT 0 COMMENT '0-follow, 1-follower',
     `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
     `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
     PRIMARY KEY (`relation_id`),
     UNIQUE KEY `uniq_relation` (`user_id`, `type`),
     KEY `ix_update_time` (`update_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='关系表';