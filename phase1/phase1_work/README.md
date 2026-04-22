# Phase1 Work
启动类为 phase1/main.go  
需要在本地启动Mysql作为数据库  
使用PostMan进行接口测试  
启动命令为在phase1目录下执行`go run .`

## 包结构
```text
phase1_work/
├── README.md
├── gin_routing.go # 路由配置
├── handler        # 接口逻辑相关
└── mysql          # 数据库相关
    ├── mysql_init.go # DB初始化
    ├── mysql_logic   # CRUD语句
    └── mysql_model   # 数据模型     
```

## 数据库表
```sql
-- =============================================
-- 用户表 users
-- =============================================
CREATE TABLE `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `username` varchar(64) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(128) NOT NULL DEFAULT '' COMMENT '密码（加密存储）',
  `email` varchar(128) NOT NULL DEFAULT '' COMMENT '邮箱',
  `status` tinyint NOT NULL DEFAULT 1 COMMENT '状态 1-正常 2-禁用',
  `created_at` bigint NOT NULL DEFAULT 0 COMMENT '创建时间',
  `updated_at` bigint NOT NULL DEFAULT 0 COMMENT '更新时间',
  `deleted_at` bigint NOT NULL DEFAULT 0 COMMENT '删除时间（软删除）',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_email_deleted` (`email`,`deleted_at`),
  UNIQUE KEY `uk_username_deleted` (`username`,`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

-- =============================================
-- 文章表 posts
-- =============================================
CREATE TABLE `posts` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '文章ID',
  `title` varchar(255) NOT NULL DEFAULT '' COMMENT '文章标题',
  `content` text NOT NULL COMMENT '文章内容',
  `user_id` bigint unsigned NOT NULL DEFAULT 0 COMMENT '作者ID',
  `created_at` bigint NOT NULL DEFAULT 0 COMMENT '创建时间',
  `updated_at` bigint NOT NULL DEFAULT 0 COMMENT '更新时间',
  `deleted_at` bigint NOT NULL DEFAULT 0 COMMENT '删除时间（软删除）',
  PRIMARY KEY (`id`),
  KEY `idx_userid_deleted` (`user_id`,`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='博客文章表';

-- =============================================
-- 评论表 comments
-- =============================================
CREATE TABLE `comments` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '评论ID',
  `content` varchar(1024) NOT NULL DEFAULT '' COMMENT '评论内容',
  `user_id` bigint unsigned NOT NULL DEFAULT 0 COMMENT '评论用户ID',
  `post_id` bigint unsigned NOT NULL DEFAULT 0 COMMENT '所属文章ID',
  `created_at` bigint NOT NULL DEFAULT 0 COMMENT '创建时间',
  `updated_at` bigint NOT NULL DEFAULT 0 COMMENT '更新时间',
  `deleted_at` bigint NOT NULL DEFAULT 0 COMMENT '删除时间（软删除）',
  PRIMARY KEY (`id`),
  KEY `idx_postid_deleted` (`post_id`,`deleted_at`),
  KEY `idx_userid_deleted` (`user_id`,`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文章评论表';
```

## 各个包的作用


## 接口作用
