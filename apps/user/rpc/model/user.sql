CREATE DATABASE IF NOT EXISTS douyin;
USE douyin;
CREATE TABLE `user` (
        `user_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID',
        `username` varchar(50) NOT NULL DEFAULT '' COMMENT '用户名',
        `password` varchar(50) NOT NULL DEFAULT '' COMMENT '用户密码，MD5加密',
        `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
        `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
        PRIMARY KEY (`user_id`),
        UNIQUE KEY `uniq_username` (`username`),
        KEY `ix_update_time` (`update_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表';