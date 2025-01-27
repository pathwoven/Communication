/*
 Navicat Premium Data Transfer

 Source Server         : test new sql
 Source Server Type    : MySQL
 Source Server Version : 90001
 Source Host           : localhost:3306
 Source Schema         : communication

 Target Server Type    : MySQL
 Target Server Version : 90001
 File Encoding         : 65001

 Date: 27/01/2025 20:05:14
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for chat
-- ----------------------------
DROP TABLE IF EXISTS `chat`;
CREATE TABLE `chat`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` int UNSIGNED NOT NULL,
  `target_id` int UNSIGNED NOT NULL COMMENT '对方的id',
  `unread_count` tinyint NOT NULL DEFAULT 0,
  `last_message` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `last_time` datetime(3) NULL DEFAULT NULL,
  `is_pinned` tinyint(1) NOT NULL DEFAULT 0,
  `is_mute` tinyint(1) NOT NULL DEFAULT 0,
  `is_blocked` tinyint(1) NOT NULL DEFAULT 0,
  `is_deleted` tinyint(1) NOT NULL DEFAULT 0,
  `tag1_id` int UNSIGNED NULL DEFAULT NULL,
  `tag2_id` int UNSIGNED NULL DEFAULT NULL,
  `tag3_id` int UNSIGNED NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `user_target_id_idx`(`user_id` ASC, `target_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for chat_tag
-- ----------------------------
DROP TABLE IF EXISTS `chat_tag`;
CREATE TABLE `chat_tag`  (
  `user_id` int NOT NULL,
  `name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`user_id`, `name`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for code
-- ----------------------------
DROP TABLE IF EXISTS `code`;
CREATE TABLE `code`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` int UNSIGNED NOT NULL,
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `save_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `details` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `create_time` datetime(3) NOT NULL,
  `update_time` datetime(3) NOT NULL,
  `is_deleted` tinyint(1) NOT NULL DEFAULT 0,
  `tag1_id` int UNSIGNED NULL DEFAULT NULL,
  `tag2_id` int UNSIGNED NULL DEFAULT NULL,
  `tag3_id` int UNSIGNED NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `user_id_name_idx`(`user_id` ASC, `name` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for code_tag
-- ----------------------------
DROP TABLE IF EXISTS `code_tag`;
CREATE TABLE `code_tag`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` int UNSIGNED NOT NULL,
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `user_id_idx`(`user_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for contact
-- ----------------------------
DROP TABLE IF EXISTS `contact`;
CREATE TABLE `contact`  (
  `user_id` int UNSIGNED NOT NULL,
  `target_id` int UNSIGNED NOT NULL,
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `divide_id` int UNSIGNED NOT NULL,
  `is_blacklisted` tinyint(1) NOT NULL DEFAULT 0,
  `background_img` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`user_id`, `target_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for divide_friend
-- ----------------------------
DROP TABLE IF EXISTS `divide_friend`;
CREATE TABLE `divide_friend`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` int UNSIGNED NOT NULL,
  `name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `user_id_idx`(`user_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for divide_group
-- ----------------------------
DROP TABLE IF EXISTS `divide_group`;
CREATE TABLE `divide_group`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `group_id` int UNSIGNED NOT NULL,
  `name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `group_id_idx`(`group_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for entity
-- ----------------------------
DROP TABLE IF EXISTS `entity`;
CREATE TABLE `entity`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `is_group` tinyint(1) NOT NULL,
  `display_id` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '展示给用户的id（账号或群号）',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `online_status` tinyint NOT NULL COMMENT '0为离线，1为在线，2为隐身',
  `status` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `introduction` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `group_owner_id` int UNSIGNED NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for favorite
-- ----------------------------
DROP TABLE IF EXISTS `favorite`;
CREATE TABLE `favorite`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` int UNSIGNED NOT NULL,
  `type` tinyint NOT NULL COMMENT '0:消息 1:多选消息 2:笔记 3:代码',
  `target_id` int UNSIGNED NOT NULL COMMENT '在对应表中的id',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `user_id_idx`(`user_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for favorite_tag
-- ----------------------------
DROP TABLE IF EXISTS `favorite_tag`;
CREATE TABLE `favorite_tag`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` int UNSIGNED NOT NULL,
  `name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `user_id_idx`(`user_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for friend_application
-- ----------------------------
DROP TABLE IF EXISTS `friend_application`;
CREATE TABLE `friend_application`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `sender_id` int UNSIGNED NOT NULL,
  `receiver_id` int UNSIGNED NOT NULL,
  `reason` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `time` datetime(3) NOT NULL,
  `status` tinyint NOT NULL COMMENT '0:未处理 1:已同意 2:已拒绝 3:已忽略',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `sender_id_idx`(`sender_id` ASC) USING BTREE,
  INDEX `receiver_id_idx`(`receiver_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for group_member_info
-- ----------------------------
DROP TABLE IF EXISTS `group_member_info`;
CREATE TABLE `group_member_info`  (
  `group_id` int UNSIGNED NOT NULL,
  `user_id` int UNSIGNED NOT NULL,
  `group_nickname` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `is_banned` tinyint(1) NOT NULL DEFAULT 0,
  `is_manager` tinyint(1) NOT NULL,
  PRIMARY KEY (`group_id`, `user_id`) USING BTREE,
  INDEX `user_id_idx`(`user_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for group_notice
-- ----------------------------
DROP TABLE IF EXISTS `group_notice`;
CREATE TABLE `group_notice`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `group_id` int UNSIGNED NOT NULL,
  `sender_id` int UNSIGNED NOT NULL,
  `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `content_type` tinyint NOT NULL,
  `create_time` datetime(3) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `group_id_idx`(`group_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for group_notification
-- ----------------------------
DROP TABLE IF EXISTS `group_notification`;
CREATE TABLE `group_notification`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `group_id` int UNSIGNED NOT NULL,
  `sender_id` int UNSIGNED NOT NULL,
  `receiver_id` int UNSIGNED NOT NULL,
  `type` tinyint NOT NULL COMMENT '0:申请 1:邀请 2:邀请后的申请 3:通知',
  `reason` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `time` datetime(3) NOT NULL,
  `status` tinyint NOT NULL COMMENT '0:未处理 1:已同意 2:已拒绝 3:已忽略 4:已撤销',
  `pend_status` tinyint NULL DEFAULT NULL COMMENT '（用于邀请后的申请）0:未处理 1:已同意 2:已拒绝 3:已忽略',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `sender_id_idx`(`sender_id` ASC) USING BTREE,
  INDEX `receiver_id_idx`(`receiver_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for group_setting
-- ----------------------------
DROP TABLE IF EXISTS `group_setting`;
CREATE TABLE `group_setting`  (
  `group_id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `allow_found_by_id` tinyint(1) NOT NULL DEFAULT 1 COMMENT '允许被id搜索',
  `allow_found_by_name` tinyint(1) NOT NULL DEFAULT 1 COMMENT '允许被群名搜索',
  `allow_member_invite` tinyint(1) NOT NULL DEFAULT 1 COMMENT '允许成员邀请',
  `join_auth_method` tinyint UNSIGNED NOT NULL DEFAULT 1 COMMENT '0表示无需验证，1表示需要发送验证消息，2表示设置了问题',
  `allow_direct_invite` tinyint(1) NOT NULL DEFAULT 1 COMMENT '成员邀请无需验证',
  `auth_question` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '验证问题',
  `auth_answer` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '答案',
  PRIMARY KEY (`group_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for message
-- ----------------------------
DROP TABLE IF EXISTS `message`;
CREATE TABLE `message`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `sender_id` int UNSIGNED NOT NULL,
  `receiver_id` int UNSIGNED NOT NULL,
  `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `content_type` tinyint NOT NULL COMMENT '0表示文本，1表示图片，2表示文件，3表示群投票，4表示多选消息',
  `create_time` datetime(3) NOT NULL,
  `is_recall` tinyint(1) NOT NULL DEFAULT 0,
  `operation` tinyint NULL DEFAULT NULL,
  `op_message_id` int UNSIGNED NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `sender_id_idx`(`sender_id` ASC) USING BTREE,
  INDEX `receiver_id_idx`(`receiver_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for message_deleted
-- ----------------------------
DROP TABLE IF EXISTS `message_deleted`;
CREATE TABLE `message_deleted`  (
  `user_id` int UNSIGNED NOT NULL,
  `start_message_id` int UNSIGNED NOT NULL,
  `end_message_id` int UNSIGNED NOT NULL,
  PRIMARY KEY (`user_id`, `start_message_id`, `end_message_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for multiple_message
-- ----------------------------
DROP TABLE IF EXISTS `multiple_message`;
CREATE TABLE `multiple_message`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `message_id` int UNSIGNED NOT NULL,
  PRIMARY KEY (`id`, `message_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for note
-- ----------------------------
DROP TABLE IF EXISTS `note`;
CREATE TABLE `note`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` int UNSIGNED NOT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `save_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `details` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `create_time` datetime(3) NOT NULL,
  `update_time` datetime(3) NOT NULL,
  `is_deleted` tinyint(1) NOT NULL DEFAULT 0,
  `tag1_id` int UNSIGNED NULL DEFAULT NULL,
  `tag2_id` int UNSIGNED NULL DEFAULT NULL,
  `tag3_id` int UNSIGNED NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `user_id_idx`(`user_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for note_tag
-- ----------------------------
DROP TABLE IF EXISTS `note_tag`;
CREATE TABLE `note_tag`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` int UNSIGNED NOT NULL,
  `name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `user_id_idx`(`user_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for recycle_bin
-- ----------------------------
DROP TABLE IF EXISTS `recycle_bin`;
CREATE TABLE `recycle_bin`  (
  `user_id` int UNSIGNED NOT NULL,
  `type` tinyint NOT NULL COMMENT '0：笔记，1：代码',
  `target_id` int UNSIGNED NOT NULL COMMENT '在对应表中的id',
  PRIMARY KEY (`user_id`, `type`, `target_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for schedule
-- ----------------------------
DROP TABLE IF EXISTS `schedule`;
CREATE TABLE `schedule`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` int UNSIGNED NOT NULL,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `details` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `start_time` datetime(3) NOT NULL,
  `end_time` datetime(3) NULL DEFAULT NULL,
  `need_remind` tinyint(1) NOT NULL DEFAULT 0,
  `color` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `user_id_idx`(`user_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for user_setting
-- ----------------------------
DROP TABLE IF EXISTS `user_setting`;
CREATE TABLE `user_setting`  (
  `user_id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `birthday` date NULL DEFAULT NULL,
  `sex` tinyint NOT NULL DEFAULT 2 COMMENT '0表示男，1表示女，2表示未知',
  `font_size` tinyint NOT NULL DEFAULT 16,
  `font_style` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `theme` tinyint NOT NULL DEFAULT 0,
  `background` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '背景图片的存储地址',
  `receive_notice` tinyint(1) NOT NULL DEFAULT 1 COMMENT '0表示没有新消息提醒，1表示有',
  `notice_sound` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '提示音的存储地址',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `email` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `create_time` datetime(3) NOT NULL,
  `last_login_time` datetime(3) NULL DEFAULT NULL,
  `is_deleted` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否已注销',
  `found_by_account_id` tinyint(1) NOT NULL DEFAULT 1 COMMENT '可以通过id被搜索',
  `found_by_nickname` tinyint(1) NOT NULL DEFAULT 1 COMMENT '可以通过昵称被搜索',
  `found_by_group` tinyint(1) NOT NULL DEFAULT 1 COMMENT '是否可通过群聊添加',
  `found_by_friend` tinyint(1) NOT NULL DEFAULT 1 COMMENT '是否可被好友推荐',
  `allow_stranger_message` tinyint(1) NOT NULL DEFAULT 1 COMMENT '允许陌生人发送消息',
  `is_invisible` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否隐私登录',
  PRIMARY KEY (`user_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
