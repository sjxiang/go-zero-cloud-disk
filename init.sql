CREATE DATABASE IF NOT EXISTS `cloud-disk` /*!40100 DEFAULT CHARACTER SET utf8 COLLATE utf8_unicode_ci */;

USE `cloud-disk`;

DROP TABLE IF EXISTS `user_basic`;

-- 用户表结构设计
CREATE TABLE `user_basic` (
    `id`          bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `identity`    varchar(36) DEFAULT NULL,
    `name`        varchar(32) DEFAULT NULL,
    `password`    varchar(32) DEFAULT NULL,
    `email`       varchar(32) DEFAULT NULL,
    `created_at`  datetime(3) DEFAULT NULL,
    `updated_at`  datetime(3) DEFAULT NULL,
    `deleted_at`  datetime(3) DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


describe `user_basic`;
show create table `user_basic`; 



-- 存储池表结构设计
CREATE TABLE `repository_pool` (
    `id`          bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `identity`    varchar(36)  DEFAULT NULL,
    `hash`        varchar(32)  DEFAULT NULL COMMENT '文件的唯一标识',
    `name`        varchar(255) DEFAULT NULL,
    `ext`         varchar(30)  DEFAULT NULL COMMENT '文件扩展名',
    `size`        int          DEFAULT NULL COMMENT '文件大小',
    `path`        varchar(255) DEFAULT NULL COMMENT '文件路径',
    `created_at`  datetime(3)  DEFAULT NULL,
    `updated_at`  datetime(3)  DEFAULT NULL,
    `deleted_at`  datetime(3)  DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;



-- 个人存储池表结构设计
CREATE TABLE `user_repository_pool` (
    `id`                  bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `identity`            varchar(36)  DEFAULT NULL,
    `user_identity`       varchar(36)  DEFAULT NULL,
    `parent_id`           int(11)      DEFAULT NULL, 
    `repository_identity` varchar(36)  DEFAULT NULL,
    `ext`                 varchar(255) DEFAULT NULL COMMENT '文件或文件夹',
    `name`                varchar(255) DEFAULT NULL,
    `created_at`  datetime(3) DEFAULT NULL,
    `updated_at`  datetime(3) DEFAULT NULL,
    `deleted_at`  datetime(3) DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


-- 分享表结构设计
CREATE TABLE `share_basic` (
    `id`                  bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `identity`            varchar(36)  DEFAULT NULL,
    `user_identity`       varchar(36)  DEFAULT NULL,
    `parent_id`           int(11)      DEFAULT NULL, 
    `repository_identity` varchar(36)  DEFAULT NULL,
    `expired_time`        int(11)      DEFAULT NULL COMMENT '失效时间；单位时间：秒；0，为永不失效',
    `created_at`  datetime(3) DEFAULT NULL,
    `updated_at`  datetime(3) DEFAULT NULL,
    `deleted_at`  datetime(3) DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


