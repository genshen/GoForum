/*
Navicat MySQL Data Transfer

Source Server         : sql
Source Server Version : 100109
Source Host           : localhost:3306
Source Database       : goforum

Target Server Type    : MYSQL
Target Server Version : 100109
File Encoding         : 65001

Date: 2016-09-18 15:05:24
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for comment
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment` (
`id`  int(11) NOT NULL AUTO_INCREMENT ,
`post_id`  int(11) NOT NULL ,
`author_id`  int(11) NOT NULL ,
`parent`  int(11) NULL DEFAULT 0 ,
`content`  tinytext CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL ,
`visible`  tinyint(4) NOT NULL DEFAULT 1 ,
`created_at`  datetime NOT NULL ,
`updated_at`  datetime NOT NULL ,
`deleted_at`  datetime NULL DEFAULT NULL ,
PRIMARY KEY (`id`)
)
ENGINE=InnoDB
DEFAULT CHARACTER SET=utf8 COLLATE=utf8_general_ci
AUTO_INCREMENT=110

;

-- ----------------------------
-- Records of comment
-- ----------------------------
BEGIN;
INSERT INTO `comment` VALUES ('80', '26', '34', '0', 'wwwww', '1', '2016-08-07 01:06:24', '2016-08-07 01:06:24', null), ('81', '26', '34', '0', 'wwwwwwwwwwwwwwwwwwwwwwwwwwwww', '1', '2016-08-07 01:07:41', '2016-08-07 01:07:41', null), ('89', '26', '34', '0', '@Hee#213 ddd?', '0', '2016-08-07 01:13:30', '2016-08-07 01:13:30', null), ('90', '26', '34', '0', '@Hee#213 ', '0', '2016-08-07 01:13:40', '2016-08-07 01:13:40', null), ('91', '26', '34', '0', '@Hee#213 ', '0', '2016-08-07 01:13:51', '2016-08-07 01:13:51', null), ('92', '26', '34', '0', '@Hee#213 fffff', '0', '2016-08-07 01:15:12', '2016-08-07 01:15:12', null), ('93', '26', '34', '0', '@Hee#213 why?', '0', '2016-08-07 01:16:02', '2016-08-07 01:16:02', null), ('99', '26', '34', '0', '@Hee#213 ', '0', '2016-08-07 01:25:34', '2016-08-07 01:25:34', null), ('100', '26', '34', '0', '@Hee#213 really?', '0', '2016-08-07 01:26:17', '2016-08-07 01:26:17', null), ('101', '26', '34', '0', '@Hee#213  big moutain!', '0', '2016-08-07 01:27:38', '2016-08-07 01:27:38', null), ('102', '26', '34', '0', '哈？', '0', '2016-08-07 01:30:15', '2016-08-07 01:30:15', null), ('103', '27', '34', '0', 'I\'m the first one~', '0', '2016-08-07 01:31:08', '2016-08-07 01:31:08', null), ('104', '27', '34', '0', '@Hee#213 k', '0', '2016-08-07 01:31:55', '2016-08-07 01:31:55', null), ('105', '27', '34', '0', '@Hee#213 ', '0', '2016-08-07 03:07:22', '2016-08-07 03:07:22', null), ('106', '29', '34', '0', ' success', '0', '2016-09-09 01:07:09', '2016-09-09 01:07:09', null), ('107', '26', '34', '0', 'hhhh', '1', '2016-09-10 01:28:31', '2016-09-10 01:28:31', null), ('108', '26', '34', '0', '@Hee#213 ', '1', '2016-09-10 02:09:42', '2016-09-10 02:09:42', null), ('109', '30', '34', '0', 'ok', '1', '2016-09-10 02:15:28', '2016-09-10 02:15:28', null);
COMMIT;

-- ----------------------------
-- Table structure for feedback
-- ----------------------------
DROP TABLE IF EXISTS `feedback`;
CREATE TABLE `feedback` (
`id`  int(11) NOT NULL AUTO_INCREMENT ,
`user_id`  int(11) NULL DEFAULT NULL ,
`type`  tinyint(4) NOT NULL ,
`content`  varchar(512) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL ,
`contact`  varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL ,
`created_at`  datetime NOT NULL ,
`updated_at`  datetime NOT NULL ,
`deleted_at`  datetime NULL DEFAULT NULL ,
PRIMARY KEY (`id`)
)
ENGINE=InnoDB
DEFAULT CHARACTER SET=utf8 COLLATE=utf8_general_ci
AUTO_INCREMENT=5

;

-- ----------------------------
-- Records of feedback
-- ----------------------------
BEGIN;
INSERT INTO `feedback` VALUES ('4', '34', '0', 'wwwww', 'wwwww', '2016-08-07 00:23:10', '2016-08-07 00:23:10', null);
COMMIT;

-- ----------------------------
-- Table structure for follows
-- ----------------------------
DROP TABLE IF EXISTS `follows`;
CREATE TABLE `follows` (
`follower_id`  int(11) NOT NULL ,
`following_id`  int(11) NOT NULL ,
PRIMARY KEY (`follower_id`, `following_id`)
)
ENGINE=InnoDB
DEFAULT CHARACTER SET=utf8 COLLATE=utf8_general_ci

;

-- ----------------------------
-- Records of follows
-- ----------------------------
BEGIN;
INSERT INTO `follows` VALUES ('33', '34'), ('34', '33');
COMMIT;

-- ----------------------------
-- Table structure for notification
-- ----------------------------
DROP TABLE IF EXISTS `notification`;
CREATE TABLE `notification` (
`id`  int(11) NOT NULL AUTO_INCREMENT ,
`user_id`  int(11) NOT NULL DEFAULT 0 ,
`related_id`  int(11) NOT NULL DEFAULT 0 ,
`subject_type`  int(11) NOT NULL DEFAULT 0 ,
`data`  blob NOT NULL ,
`is_read`  tinyint(4) NOT NULL DEFAULT 0 ,
`created_at`  datetime NOT NULL ,
`updated_at`  datetime NOT NULL DEFAULT '0000-00-00 00:00:00' ,
`deleted_at`  datetime NULL DEFAULT NULL ,
PRIMARY KEY (`id`)
)
ENGINE=InnoDB
DEFAULT CHARACTER SET=utf8 COLLATE=utf8_general_ci
AUTO_INCREMENT=1

;

-- ----------------------------
-- Records of notification
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for post_message
-- ----------------------------
DROP TABLE IF EXISTS `post_message`;
CREATE TABLE `post_message` (
`id`  int(11) NOT NULL AUTO_INCREMENT ,
`user_id`  int(11) NOT NULL ,
`related_id`  int(11) NOT NULL ,
`related_username`  varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL ,
`subject_type`  int(11) NOT NULL ,
`post_id`  int(11) NOT NULL ,
`post_title`  varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL ,
`summary`  varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`quote`  varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '@see: post.summary' ,
`is_read`  tinyint(4) NOT NULL DEFAULT 0 ,
`created_at`  datetime NOT NULL DEFAULT '0000-00-00 00:00:00' ,
`updated_at`  datetime NOT NULL DEFAULT '0000-00-00 00:00:00' ,
`deleted_at`  datetime NULL DEFAULT NULL ,
PRIMARY KEY (`id`)
)
ENGINE=InnoDB
DEFAULT CHARACTER SET=utf8 COLLATE=utf8_general_ci
AUTO_INCREMENT=9

;

-- ----------------------------
-- Records of post_message
-- ----------------------------
BEGIN;
INSERT INTO `post_message` VALUES ('2', '34', '34', 'gensh', '0', '27', 'e', 'I\'m the first one~', '请输入内容...e哈哈哈哈', '0', '2016-08-07 01:31:08', '2016-08-07 01:31:08', null), ('3', '34', '34', 'gensh', '0', '27', 'e', '@Hee#213 k', '请输入内容...e哈哈哈哈', '0', '2016-08-07 01:31:55', '2016-08-07 01:31:55', null), ('4', '34', '34', 'gensh', '0', '27', 'e', '@Hee#213 ', '请输入内容...e哈哈哈哈', '0', '2016-08-07 03:07:22', '2016-08-07 03:07:22', null), ('5', '34', '34', 'gensh', '0', '29', 'Hi ,xss test', ' success', '请输入内容...M&lt;script&gt;alert(1)&lt;/script&gt;', '0', '2016-09-09 01:07:09', '2016-09-09 01:07:09', null), ('6', '33', '34', 'gensh', '0', '26', '222', 'hhhh', '请输222入内容...', '0', '2016-09-10 01:28:31', '2016-09-10 01:28:31', null), ('7', '33', '34', 'gensh', '0', '26', '222', '@Hee#213 ', '请输222入内容...', '0', '2016-09-10 02:09:42', '2016-09-10 02:09:42', null), ('8', '34', '34', 'gensh', '0', '30', 'Hi ,xss test', 'ok', '请输入内容...M&lt;script&gt;alert(1)&lt;/script&gt;', '0', '2016-09-10 02:15:28', '2016-09-10 02:15:28', null);
COMMIT;

-- ----------------------------
-- Table structure for post_tag
-- ----------------------------
DROP TABLE IF EXISTS `post_tag`;
CREATE TABLE `post_tag` (
`post_id`  int(11) NOT NULL ,
`tag_id`  int(11) NOT NULL ,
PRIMARY KEY (`post_id`, `tag_id`)
)
ENGINE=InnoDB
DEFAULT CHARACTER SET=utf8 COLLATE=utf8_general_ci

;

-- ----------------------------
-- Records of post_tag
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for posts
-- ----------------------------
DROP TABLE IF EXISTS `posts`;
CREATE TABLE `posts` (
`id`  int(11) NOT NULL AUTO_INCREMENT ,
`author_id`  int(11) NOT NULL ,
`topic_id`  int(11) NOT NULL DEFAULT 1 ,
`title`  varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL ,
`summary`  varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL ,
`content`  text CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL ,
`is_mobile`  tinyint(4) NULL DEFAULT NULL ,
`sticky`  tinyint(4) NOT NULL DEFAULT 0 ,
`comment_count`  int(11) NOT NULL DEFAULT 0 ,
`view_count`  int(11) NOT NULL DEFAULT 0 ,
`last_replay_at`  datetime NULL DEFAULT NULL ,
`visible`  tinyint(4) NOT NULL DEFAULT 1 ,
`created_at`  datetime NOT NULL DEFAULT '0000-00-00 00:00:00' ,
`updated_at`  datetime NOT NULL DEFAULT '0000-00-00 00:00:00' ,
`deleted_at`  datetime NULL DEFAULT NULL ,
PRIMARY KEY (`id`),
FOREIGN KEY (`author_id`) REFERENCES `user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
INDEX `author` (`author_id`) USING BTREE 
)
ENGINE=InnoDB
DEFAULT CHARACTER SET=utf8 COLLATE=utf8_general_ci
AUTO_INCREMENT=31

;

-- ----------------------------
-- Records of posts
-- ----------------------------
BEGIN;
INSERT INTO `posts` VALUES ('26', '33', '1', '222', '请输222入内容...', '<p>请输222入内容...</p><p><br></p>', '1', '0', '13', '0', '2016-08-07 00:17:41', '1', '2016-08-07 00:17:41', '2016-08-07 00:17:41', null), ('27', '34', '1', 'e', '请输入内容...e哈哈哈哈', '<p>请输入内容...</p><p>e哈哈哈哈</p>', '1', '0', '3', '0', '2016-08-07 00:19:55', '1', '2016-08-07 00:19:55', '2016-08-07 00:19:55', null), ('28', '34', '1', 'e', '请输入内容...e哈哈哈哈', '<p>请输入内容...</p><p>e哈哈哈哈</p>', '1', '0', '0', '0', '2016-08-07 00:21:34', '1', '2016-08-07 00:21:34', '2016-08-07 00:21:34', null), ('29', '34', '1', 'Hi ,xss test', '请输入内容...M&lt;script&gt;alert(1)&lt;/script&gt;', '<p>请输入内容...</p><h3 class=\"\">M</h3><div class=\"\"><br></div><img src=\"#BASE_URL#FgxdvCRQ6asrgwnXCNagXSW-otjl\" style=\"max-width: 100%;\"><div class=\"\"><br></div><div class=\"\"><script>alert(1)</script></div>', '1', '0', '1', '0', '2016-09-09 00:50:03', '1', '2016-09-09 00:50:03', '2016-09-09 00:50:03', null), ('30', '34', '1', 'Hi ,xss test', '请输入内容...M&lt;script&gt;alert(1)&lt;/script&gt;', '<p>请输入内容...</p><h3 class=\"\">M</h3><div class=\"\"><br></div><img src=\"#BASE_URL#FgxdvCRQ6asrgwnXCNagXSW-otjl\" style=\"max-width: 100%;\"><div class=\"\"><br></div><div class=\"\"><script>alert(1)</script></div>', '1', '0', '1', '0', '2016-09-09 00:52:49', '1', '2016-09-09 00:52:49', '2016-09-09 00:52:49', null);
COMMIT;

-- ----------------------------
-- Table structure for profile
-- ----------------------------
DROP TABLE IF EXISTS `profile`;
CREATE TABLE `profile` (
`id`  int(11) NOT NULL ,
`user_refer_id`  int(11) NOT NULL ,
`name`  varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL ,
`avatar`  varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL ,
`coins`  int(11) NOT NULL DEFAULT 0 ,
`bio`  varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`post_count`  int(11) NOT NULL DEFAULT 0 ,
`comment_count`  int(11) NOT NULL DEFAULT 0 ,
`following_count`  int(11) NOT NULL DEFAULT 0 ,
`followed_count`  int(11) NOT NULL DEFAULT 0 ,
PRIMARY KEY (`id`)
)
ENGINE=InnoDB
DEFAULT CHARACTER SET=utf8 COLLATE=utf8_general_ci

;

-- ----------------------------
-- Records of profile
-- ----------------------------
BEGIN;
INSERT INTO `profile` VALUES ('33', '33', '111111', '/static/img/default.png', '0', '这家伙很懒,什么都没有', '0', '0', '0', '0'), ('34', '34', 'gensh', '/static/img/default.png', '0', '这家伙很懒,什么都没有', '0', '18', '0', '0'), ('35', '35', 'cgs', '/static/img/default.png', '0', '', '0', '0', '0', '0'), ('36', '36', 'c储根深', '/static/img/default.png', '0', '', '0', '0', '0', '0'), ('37', '37', '储根深', '/static/img/default.png', '0', '', '0', '0', '0', '0'), ('38', '38', '储根深', '/static/img/default.png', '0', '', '0', '0', '0', '0');
COMMIT;

-- ----------------------------
-- Table structure for swipe
-- ----------------------------
DROP TABLE IF EXISTS `swipe`;
CREATE TABLE `swipe` (
`id`  int(11) NOT NULL AUTO_INCREMENT ,
`url`  varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL ,
`img`  varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL ,
`content`  varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`visible`  tinyint(1) NOT NULL ,
`created_at`  datetime NOT NULL ,
`updated_at`  datetime NOT NULL ,
`deleted_at`  datetime NULL DEFAULT NULL ,
PRIMARY KEY (`id`)
)
ENGINE=InnoDB
DEFAULT CHARACTER SET=utf8 COLLATE=utf8_general_ci
AUTO_INCREMENT=4

;

-- ----------------------------
-- Records of swipe
-- ----------------------------
BEGIN;
INSERT INTO `swipe` VALUES ('1', 'https://baidu.com', 'static/img/swipe1.jpeg', 'no', '1', '0000-00-00 00:00:00', '0000-00-00 00:00:00', null), ('2', 'https://baidu.com', 'static/img/swipe2.jpeg', 'no', '1', '0000-00-00 00:00:00', '0000-00-00 00:00:00', null), ('3', 'https://baidu.com', 'static/img/swipe3.jpeg', 'no', '1', '0000-00-00 00:00:00', '0000-00-00 00:00:00', null);
COMMIT;

-- ----------------------------
-- Table structure for tag
-- ----------------------------
DROP TABLE IF EXISTS `tag`;
CREATE TABLE `tag` (
`id`  int(11) NOT NULL AUTO_INCREMENT ,
`name`  varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL ,
`describe`  varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`color`  varchar(16) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL ,
`visible`  tinyint(4) NOT NULL DEFAULT 1 ,
`created_at`  datetime NOT NULL ,
`updated_at`  datetime NOT NULL ,
`deleted_at`  datetime NULL DEFAULT NULL ,
PRIMARY KEY (`id`)
)
ENGINE=InnoDB
DEFAULT CHARACTER SET=utf8 COLLATE=utf8_general_ci
AUTO_INCREMENT=1

;

-- ----------------------------
-- Records of tag
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for topic
-- ----------------------------
DROP TABLE IF EXISTS `topic`;
CREATE TABLE `topic` (
`id`  int(11) NOT NULL AUTO_INCREMENT ,
`name`  varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL ,
`describe`  varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`slug`  varchar(16) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL ,
`icon`  varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`color`  varchar(16) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`visible`  tinyint(4) NOT NULL DEFAULT 1 ,
`created_at`  datetime NOT NULL ,
`updated_at`  datetime NOT NULL ,
`deleted_at`  datetime NULL DEFAULT NULL ,
PRIMARY KEY (`id`)
)
ENGINE=InnoDB
DEFAULT CHARACTER SET=utf8 COLLATE=utf8_general_ci
AUTO_INCREMENT=3

;

-- ----------------------------
-- Records of topic
-- ----------------------------
BEGIN;
INSERT INTO `topic` VALUES ('1', '聊天', 'chat', 'chat', null, '#f44336', '1', '0000-00-00 00:00:00', '0000-00-00 00:00:00', null), ('2', '讨论', 'slug', 'discus', null, '#673ab7', '1', '0000-00-00 00:00:00', '0000-00-00 00:00:00', null);
COMMIT;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
`id`  int(11) NOT NULL AUTO_INCREMENT ,
`tel`  varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`name`  varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL ,
`email`  varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL ,
`password_hash`  varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL ,
`auth_key`  varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`password_reset_token`  varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`status`  tinyint(4) NOT NULL DEFAULT 10 ,
`created_at`  datetime NOT NULL ,
`updated_at`  datetime NOT NULL ,
`deleted_at`  datetime NULL DEFAULT NULL ,
PRIMARY KEY (`id`)
)
ENGINE=InnoDB
DEFAULT CHARACTER SET=utf8 COLLATE=utf8_general_ci
AUTO_INCREMENT=39

;

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO `user` VALUES ('33', '', '111111', 'm18813128117@163.com', '3d4f2bf07dc1be38b20cd6e46949a1071f9d0e3d', '', '', '2', '2016-08-06 23:24:11', '2016-08-06 23:24:11', null), ('34', '', 'gensh', 'genshenchu@gmail.com', '7c4a8d09ca3762af61e59520943dc26494f8941b', '', '', '2', '2016-08-07 00:18:38', '2016-08-07 00:18:38', null), ('38', '', '储根深', '1690512717@qq.com', '7c4a8d09ca3762af61e59520943dc26494f8941b', '', '', '1', '2016-09-08 23:08:50', '2016-09-08 23:08:50', null);
COMMIT;

-- ----------------------------
-- Auto increment value for comment
-- ----------------------------
ALTER TABLE `comment` AUTO_INCREMENT=110;

-- ----------------------------
-- Auto increment value for feedback
-- ----------------------------
ALTER TABLE `feedback` AUTO_INCREMENT=5;

-- ----------------------------
-- Auto increment value for notification
-- ----------------------------
ALTER TABLE `notification` AUTO_INCREMENT=1;

-- ----------------------------
-- Auto increment value for post_message
-- ----------------------------
ALTER TABLE `post_message` AUTO_INCREMENT=9;

-- ----------------------------
-- Auto increment value for posts
-- ----------------------------
ALTER TABLE `posts` AUTO_INCREMENT=31;

-- ----------------------------
-- Auto increment value for swipe
-- ----------------------------
ALTER TABLE `swipe` AUTO_INCREMENT=4;

-- ----------------------------
-- Auto increment value for tag
-- ----------------------------
ALTER TABLE `tag` AUTO_INCREMENT=1;

-- ----------------------------
-- Auto increment value for topic
-- ----------------------------
ALTER TABLE `topic` AUTO_INCREMENT=3;

-- ----------------------------
-- Auto increment value for user
-- ----------------------------
ALTER TABLE `user` AUTO_INCREMENT=39;
