CREATE DATABASE IF NOT EXISTS todo_srv DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

use todo_srv;

DROP TABLE IF EXISTS todo_srv.`user`;
CREATE TABLE todo_srv.`user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `provider` varchar(32) NOT NULL COMMENT '登录第三方(gmail/facebook/github)',
  `provider_user_id` varchar(128) NOT NULL COMMENT '第三方用户ID',
  `name` varchar(64) DEFAULT '' COMMENT '用户名',
  `email` varchar(64) DEFAULT '' COMMENT '邮箱',
  `created_at` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_provider_user` (`provider`,`provider_user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

DROP TABLE IF EXISTS todo_srv.`todo`;
CREATE TABLE todo_srv.`todo` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'TODO ID',
  `user_id` bigint(20) unsigned NOT NULL COMMENT '用户ID',
  `content` varchar(1024) NOT NULL COMMENT 'TODO内容',
  `status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否完成，0否、1是)',
  `created_at` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='todo表';


INSERT INTO todo_srv`user` (`provider`, `provider_user_id`, `name`, `email`, `created_at`, `updated_at`)
VALUES ('google', '1002jsldfj78ksahe21faslj9sfsjakj229dfjkj', 'zhusen', 'zhusen@google.com', 1768559361, 1768559361);
