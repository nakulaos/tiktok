DROP TABLE IF EXISTS `categories`;
CREATE TABLE `categories`  (
                               `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
                               `vid` int UNSIGNED NOT NULL,
                               `label` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                               `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                               `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                               `delete_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                               `del_state` tinyint NOT NULL DEFAULT '0',
                               `version` bigint NOT NULL DEFAULT '0' COMMENT '版本号',
                               PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB  CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for comments
-- ----------------------------
DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments`  (
                             `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
                             `uid` int UNSIGNED NOT NULL COMMENT '用户id',
                             `vid` int UNSIGNED NOT NULL COMMENT '视频Id',
                             `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '评论内容',
                             `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                             `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                             `delete_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                             `del_state` tinyint NOT NULL DEFAULT '0',
                             `version` bigint NOT NULL DEFAULT '0' COMMENT '版本号',
                             PRIMARY KEY (`id`) USING BTREE,
                             INDEX `uid`(`uid` ASC) USING BTREE,
                             INDEX `vid`(`vid` ASC) USING BTREE
) ENGINE = InnoDB  CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for danmu
-- ----------------------------
DROP TABLE IF EXISTS `danmu`;
CREATE TABLE `danmu`  (
                          `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
                          `uid` int UNSIGNED NOT NULL COMMENT '用户id',
                          `vid` int UNSIGNED NOT NULL COMMENT '视频Id',
                          `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '评论内容',
                          `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                          `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                          `delete_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                          `del_state` tinyint NOT NULL DEFAULT '0',
                          `version` bigint NOT NULL DEFAULT '0' COMMENT '版本号',
                          `send_time` float NOT NULL COMMENT '在视频的哪个点发送的弹幕',
                          PRIMARY KEY (`id`) USING BTREE,
                          INDEX `uid`(`uid` ASC) USING BTREE,
                          INDEX `vid`(`vid` ASC) USING BTREE
) ENGINE = InnoDB  CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for favorites
-- ----------------------------
DROP TABLE IF EXISTS `favorites`;
CREATE TABLE `favorites`  (
                              `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
                              `uid` int UNSIGNED NOT NULL,
                              `vid` int NOT NULL,
                              `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                              `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                              `delete_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                              `del_state` tinyint NOT NULL DEFAULT '0',
                              `version` bigint NOT NULL DEFAULT '0' COMMENT '版本号',
                              PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for history
-- ----------------------------
DROP TABLE IF EXISTS `history`;
CREATE TABLE `history`  (
                            `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
                            `uid` int UNSIGNED NOT NULL,
                            `vid` int NOT NULL,
                            `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                            `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                            `delete_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                            `del_state` tinyint NOT NULL DEFAULT '0',
                            `version` bigint NOT NULL DEFAULT '0' COMMENT '版本号',
                            PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB  CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for relations
-- ----------------------------
DROP TABLE IF EXISTS `relations`;
CREATE TABLE `relations`  (
                              `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
                              `follower_id` int UNSIGNED NOT NULL COMMENT '关注者id',
                              `following_id` int UNSIGNED NOT NULL COMMENT '被关注者id',
                              `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                              `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                              `delete_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                              `del_state` tinyint NOT NULL DEFAULT '0',
                              `version` bigint NOT NULL DEFAULT '0' COMMENT '版本号',
                              PRIMARY KEY (`id`) USING BTREE,
                              INDEX `follower_id`(`follower_id` ASC) USING BTREE,
                              INDEX `following_id`(`following_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for stars
-- ----------------------------
DROP TABLE IF EXISTS `stars`;
CREATE TABLE `stars`  (
                          `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
                          `uid` int UNSIGNED NOT NULL,
                          `vid` int NOT NULL,
                          `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                          `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                          `delete_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                          `del_state` tinyint NOT NULL DEFAULT '0',
                          `version` bigint NOT NULL DEFAULT '0' COMMENT '版本号',
                          PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
                         `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
                         `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '用户账号',
                         `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '用户邮件',
                         `nickname` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '用户昵称',
                         `gender` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户性别',
                         `role` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT 'normalUser' COMMENT  '角色',
                         `status` tinyint UNSIGNED NOT NULL DEFAULT 1 COMMENT '用户状态',
                         `mobile` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '用户电话',
                         `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '用户密码',
                         `dec` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '个性签名',
                         `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '头像',
                         `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                         `delete_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         `del_state` tinyint NOT NULL DEFAULT '0',
                         `version` bigint NOT NULL DEFAULT '0' COMMENT '版本号',
                         `background_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '用户主页背景',
                         PRIMARY KEY (`id`) USING BTREE,
                         UNIQUE INDEX `idx_mobile_unique`(`mobile` ASC) USING BTREE,
                        UNIQUE INDEX `idx_username_unique`(`username` ASC) USING BTREE,
                         UNIQUE INDEX `idx_email_unique`(`email` ASC) USING BTREE
) ENGINE = InnoDB  CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for videos
-- ----------------------------
DROP TABLE IF EXISTS `videos`;
CREATE TABLE `videos`  (
                           `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
                           `author_id` int UNSIGNED NOT NULL COMMENT '上传用户Id',
                           `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '视频标题',
                           `cover_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '封面url',
                           `play_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '视频播放url',
                           `favorite_count` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '点赞数',
                           `star_count` int NOT NULL COMMENT '收藏数',
                           `comment_count` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '评论数目',
                           `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                           `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                           `delete_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                           `del_state` tinyint NOT NULL DEFAULT '0',
                           `category` int NOT NULL COMMENT '视频分类',
                           `version` bigint NOT NULL DEFAULT '0' COMMENT '版本号',
                           `duration` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '视频时长',
                           PRIMARY KEY (`id`) USING BTREE,
                           INDEX `author_id`(`author_id` ASC) USING BTREE
) ENGINE = InnoDB  CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;




