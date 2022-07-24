CREATE DATABASE IF NOT EXISTS shop  DEFAULT CHARSET utf8 COLLATE utf8_general_ci;
use shop;
CREATE TABLE  IF NOT EXISTS `auth` (
                              `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
                              `business_key` varchar(32) NOT NULL DEFAULT '' COMMENT '调用方key',
                              `business_secret` varchar(60) NOT NULL DEFAULT '' COMMENT '调用方secret',
                              `business_developer` varchar(60) NOT NULL DEFAULT '' COMMENT '调用方对接人',
                              `remark` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
                              `is_used` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否启用 1:是  -1:否',
                              `is_deleted` tinyint(1) NOT NULL DEFAULT '-1' COMMENT '是否删除 1:是  -1:否',
                              `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                              `created_user` varchar(60) NOT NULL DEFAULT '' COMMENT '创建人',
                              `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                              `updated_user` varchar(60) NOT NULL DEFAULT '' COMMENT '更新人',
                              PRIMARY KEY (`id`),
                              UNIQUE KEY `unique_business_key` (`business_key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='开发者接口授权表';

INSERT INTO shop.auth (id, business_key, business_secret, business_developer, remark, is_used, is_deleted, created_at, created_user, updated_at, updated_user) VALUES (1, 'AK100523687952', 'W1WTYvJpfeH1YpUjTpeFbEx^DnpQ&35L', '', '', 1, -1, '2022-03-24 15:09:01', 'admin', '2022-07-24 15:10:25', '');
