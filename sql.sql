/*
Navicat MySQL Data Transfer

Source Server         : sql
Source Server Version : 100109
Source Host           : localhost:3306
Source Database       : beeapp

Target Server Type    : MYSQL
Target Server Version : 100109
File Encoding         : 65001

Date: 2016-08-02 00:33:37
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for comment
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment` (
`id`  int(11) NOT NULL AUTO_INCREMENT ,
`post_id`  int(11) NOT NULL ,
`author`  int(11) NOT NULL ,
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
AUTO_INCREMENT=73

;

-- ----------------------------
-- Records of comment
-- ----------------------------
BEGIN;
INSERT INTO `comment` VALUES ('1', '22', '6', '0', '踩踩踩踩踩', '1', '2016-07-03 16:03:31', '2016-07-03 16:03:31', null), ('2', '22', '6', '0', 'vv', '1', '2016-07-03 16:06:20', '2016-07-03 16:06:20', null), ('3', '22', '6', '0', 'vv', '1', '2016-07-03 16:07:02', '2016-07-03 16:07:02', null), ('4', '22', '6', '0', '哈哈哈哈', '1', '2016-07-03 16:11:24', '2016-07-03 16:11:24', null), ('5', '22', '6', '0', '男男女女', '1', '2016-07-03 16:13:03', '2016-07-03 16:13:03', null), ('6', '22', '6', '0', '男男女女', '1', '2016-07-03 16:13:33', '2016-07-03 16:13:33', null), ('7', '22', '6', '0', '哈哈哈', '1', '2016-07-03 16:14:22', '2016-07-03 16:14:22', null), ('8', '22', '6', '0', '哈哈哈', '1', '2016-07-03 16:15:40', '2016-07-03 16:15:40', null), ('9', '22', '6', '0', 'LoadComments', '1', '2016-07-03 16:50:54', '2016-07-03 16:50:54', null), ('10', '22', '6', '0', 'Why you need computed for that? As it is computed property in common understanding.\nUser is providing some date and total value is computed based on that input values.\n\nTrust me, you will find computed really useful with time.', '1', '2016-07-03 16:51:09', '2016-07-03 16:51:09', null), ('11', '22', '6', '0', '@Hee#213 ', '1', '2016-07-03 17:32:42', '2016-07-03 17:32:42', null), ('12', '22', '6', '0', '踩踩踩', '1', '2016-07-03 17:32:54', '2016-07-03 17:32:54', null), ('13', '22', '6', '0', 'hhhhhhhggg', '1', '2016-07-03 17:43:15', '2016-07-03 17:43:15', null), ('14', '22', '6', '0', 'hhhhhhhhhhhhjjjj', '1', '2016-07-03 17:44:27', '2016-07-03 17:44:27', null), ('15', '22', '6', '0', 'f', '1', '2016-07-03 17:45:14', '2016-07-03 17:45:14', null), ('16', '22', '6', '0', 'ggggg', '1', '2016-07-03 17:50:18', '2016-07-03 17:50:18', null), ('17', '22', '6', '0', '@Hee#213 ', '1', '2016-07-03 17:51:57', '2016-07-03 17:51:57', null), ('18', '22', '6', '0', '@Hee#213 ', '1', '2016-07-03 17:53:51', '2016-07-03 17:53:51', null), ('19', '22', '6', '0', '@Hee#213 ', '1', '2016-07-03 18:23:48', '2016-07-03 18:23:48', null), ('20', '21', '6', '0', 'jjjj', '1', '2016-07-03 18:24:16', '2016-07-03 18:24:16', null), ('21', '21', '6', '0', '@Hee#213 ', '1', '2016-07-03 18:24:28', '2016-07-03 18:24:28', null), ('22', '21', '6', '0', 'ttttttttttt', '1', '2016-07-03 18:24:59', '2016-07-03 18:24:59', null), ('23', '21', '6', '0', '@Hee#213 ', '1', '2016-07-03 19:34:39', '2016-07-03 19:34:39', null), ('24', '21', '6', '0', '@Hee#213 n', '1', '2016-07-03 19:35:45', '2016-07-03 19:35:45', null), ('25', '21', '6', '0', '@Hee#213 rrrr', '1', '2016-07-03 19:40:20', '2016-07-03 19:40:20', null), ('26', '21', '6', '0', '@Hee#213 ', '1', '2016-07-03 20:09:35', '2016-07-03 20:09:35', null), ('27', '21', '6', '0', '@Hee#213 ', '1', '2016-07-03 20:10:16', '2016-07-03 20:10:16', null), ('28', '21', '6', '0', 'ggggg', '1', '2016-07-03 20:17:11', '2016-07-03 20:17:11', null), ('29', '18', '6', '0', 'hi', '1', '2016-07-03 20:17:39', '2016-07-03 20:17:39', null), ('30', '18', '6', '0', '@Hee#213 shit', '1', '2016-07-03 20:17:51', '2016-07-03 20:17:51', null), ('31', '17', '6', '0', 'MY erl', '1', '2016-07-04 16:37:49', '2016-07-04 16:37:49', null), ('32', '23', '6', '0', 'HI,I\'m a dog!', '1', '2016-07-06 15:46:05', '2016-07-06 15:46:05', null), ('33', '21', '6', '0', '点点滴滴', '1', '2016-07-06 16:14:04', '2016-07-06 16:14:04', null), ('34', '21', '6', '0', '踩踩踩', '1', '2016-07-06 16:21:08', '2016-07-06 16:21:08', null), ('35', '22', '6', '0', '哈哈哈', '1', '2016-07-06 16:24:05', '2016-07-06 16:24:05', null), ('36', '21', '6', '0', 'Hello', '1', '2016-07-06 17:35:48', '2016-07-06 17:35:48', null), ('37', '21', '6', '0', 'jjj', '1', '2016-07-06 17:37:17', '2016-07-06 17:37:17', null), ('38', '21', '6', '0', '@Hee#213  hi', '1', '2016-07-06 17:39:22', '2016-07-06 17:39:22', null), ('39', '21', '6', '0', '@Hee#213  jjjj', '1', '2016-07-06 17:40:27', '2016-07-06 17:40:27', null), ('40', '21', '6', '0', '@Hee#213  ok', '1', '2016-07-06 17:41:51', '2016-07-06 17:41:51', null), ('41', '21', '6', '0', '@Hee#213 是吗？', '1', '2016-07-06 17:42:22', '2016-07-06 17:42:22', null), ('42', '21', '6', '0', '@Hee#213 嘿嘿嘿', '1', '2016-07-09 11:24:49', '2016-07-09 11:24:49', null), ('43', '21', '6', '0', '@Hee#213  哈哈哈哈哈哈哈哈哈哈哈哈哈 ', '1', '2016-07-09 11:25:25', '2016-07-09 11:25:25', null), ('44', '21', '6', '0', '@Hee#213 ', '1', '2016-07-09 11:27:15', '2016-07-09 11:27:15', null), ('45', '21', '6', '0', '@Hee#213 ', '1', '2016-07-09 11:27:41', '2016-07-09 11:27:41', null), ('46', '16', '6', '0', 'bbbb', '1', '2016-07-12 19:00:46', '2016-07-12 19:00:46', null), ('47', '16', '6', '0', 'qqqq', '1', '2016-07-15 15:21:59', '2016-07-15 15:21:59', null), ('48', '16', '6', '0', '3333', '1', '2016-07-15 15:22:19', '2016-07-15 15:22:19', null), ('49', '16', '6', '0', 'test re', '1', '2016-07-15 15:27:10', '2016-07-15 15:27:10', null), ('50', '16', '6', '0', 'g', '1', '2016-07-15 15:30:34', '2016-07-15 15:30:34', null), ('51', '16', '6', '0', 'dd', '1', '2016-07-15 15:31:57', '2016-07-15 15:31:57', null), ('52', '16', '6', '0', '哈哈哈哈', '1', '2016-07-15 15:34:39', '2016-07-15 15:34:39', null), ('53', '16', '6', '0', '灌灌灌灌灌灌灌灌灌灌', '1', '2016-07-15 15:35:17', '2016-07-15 15:35:17', null), ('54', '17', '6', '0', 'dddd', '1', '2016-07-15 15:37:37', '2016-07-15 15:37:37', null), ('55', '17', '6', '0', 'ddddd', '1', '2016-07-15 15:43:32', '2016-07-15 15:43:32', null), ('56', '17', '6', '0', '@Hee#213  fhhaaa~', '1', '2016-07-15 15:44:07', '2016-07-15 15:44:07', null), ('57', '23', '6', '0', 'https://discuss.flarum.org/assets/avatars/vknletk4grtrlylm.jpg', '1', '2016-07-16 17:08:42', '2016-07-16 17:08:42', null), ('58', '21', '6', '0', 'hhhA', '1', '2016-07-29 16:33:30', '2016-07-29 16:33:30', null), ('59', '21', '6', '0', 'hhtestH哈哈哈', '1', '2016-07-29 16:34:22', '2016-07-29 16:34:22', null), ('60', '21', '6', '0', '@Hee#213 哈哈哈哈', '1', '2016-07-29 16:34:39', '2016-07-29 16:34:39', null), ('61', '21', '6', '0', '@Hee#213 ', '1', '2016-07-29 16:36:30', '2016-07-29 16:36:30', null), ('62', '21', '6', '0', '@Hee#213 bug Here', '1', '2016-07-29 16:40:14', '2016-07-29 16:40:14', null), ('63', '25', '6', '0', 'test', '1', '2016-07-31 11:48:35', '2016-07-31 11:48:35', null), ('64', '19', '6', '0', 'firste re', '1', '2016-07-31 20:00:12', '2016-07-31 20:00:12', null), ('65', '19', '6', '0', '@Hee#213 hi~~~~~~~', '1', '2016-07-31 20:00:55', '2016-07-31 20:00:55', null), ('66', '19', '6', '0', '@Hee#213 ............... ', '1', '2016-07-31 20:03:25', '2016-07-31 20:03:25', null), ('67', '19', '6', '0', '@Hee#213 ', '1', '2016-07-31 20:03:57', '2016-07-31 20:03:57', null), ('68', '19', '6', '0', '@Hee#213 ', '1', '2016-07-31 20:17:48', '2016-07-31 20:17:48', null), ('69', '19', '6', '0', 'eee', '1', '2016-07-31 20:18:42', '2016-07-31 20:18:42', null), ('70', '19', '6', '0', 'l啦啦啦啦啦', '1', '2016-07-31 20:21:28', '2016-07-31 20:21:28', null), ('71', '18', '6', '0', 'hi', '1', '2016-07-31 23:07:02', '2016-07-31 23:07:02', null), ('72', '18', '6', '0', 'ffffff', '1', '2016-07-31 23:09:46', '2016-07-31 23:09:46', null);
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
AUTO_INCREMENT=3

;

-- ----------------------------
-- Records of feedback
-- ----------------------------
BEGIN;
INSERT INTO `feedback` VALUES ('1', '0', '0', 'rrr', 'rrrr', '2016-07-27 23:58:49', '2016-07-27 23:58:49', null), ('2', '0', '0', 'kkk', 'kkkkk', '2016-07-28 00:16:01', '2016-07-28 00:16:01', null);
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
INSERT INTO `follows` VALUES ('5', '6'), ('6', '5'), ('6', '11'), ('6', '22'), ('6', '23'), ('6', '31'), ('11', '6'), ('22', '6'), ('23', '6');
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
AUTO_INCREMENT=27

;

-- ----------------------------
-- Records of notification
-- ----------------------------
BEGIN;
INSERT INTO `notification` VALUES ('2', '23', '6', '2', 0x7B22757365726E616D65223A2248656C6C6F227D, '0', '2016-08-02 00:06:03', '2016-08-02 00:06:03', null);
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
`type`  int(11) NOT NULL ,
`post_id`  int(11) NOT NULL ,
`post_title`  varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL ,
`summary`  varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`quote`  varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '@see: post.summary' ,
`created_at`  datetime NOT NULL DEFAULT '0000-00-00 00:00:00' ,
`updated_at`  datetime NOT NULL DEFAULT '0000-00-00 00:00:00' ,
`deleted_at`  datetime NULL DEFAULT NULL ,
PRIMARY KEY (`id`)
)
ENGINE=InnoDB
DEFAULT CHARACTER SET=utf8 COLLATE=utf8_general_ci
AUTO_INCREMENT=1

;

-- ----------------------------
-- Records of post_message
-- ----------------------------
BEGIN;
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
AUTO_INCREMENT=26

;

-- ----------------------------
-- Records of posts
-- ----------------------------
BEGIN;
INSERT INTO `posts` VALUES ('15', '5', '1', 'hhhhhhhhhhhh', '', '<p>请输入hhhh内容...</p><p><br></p>', '1', '0', '0', '0', '2016-06-29 11:03:24', '1', '2016-06-29 11:03:24', '2016-06-29 11:03:24', null), ('16', '5', '1', '没标题?', '', '<p>请输入内容...</p><p><br></p>', '1', '0', '8', '0', '2016-06-29 11:04:18', '1', '2016-06-29 11:04:18', '2016-06-29 11:04:18', null), ('17', '5', '1', 'yyy', '', '<p>请输入内容...</p><p>yy</p>', '1', '0', '4', '0', '2016-06-29 11:05:45', '1', '2016-06-29 11:05:45', '2016-06-29 11:05:45', null), ('18', '5', '1', '标题', '', '<p>请输入内容...</p><p>hhha</p>', '1', '0', '4', '0', '2016-06-29 12:12:43', '1', '2016-06-29 12:12:43', '2016-06-29 12:12:43', null), ('19', '6', '1', 'jjjj', '', '<p>请输入内容...</p><p><br></p>', '1', '0', '7', '0', '2016-06-29 23:53:31', '1', '2016-06-29 23:53:31', '2016-06-29 23:53:31', null), ('20', '6', '1', 'hhh', '', '<p>请输入内容...</p><p><br></p>', '1', '0', '0', '0', '2016-06-30 11:40:01', '1', '2016-06-30 11:40:01', '2016-06-30 11:40:01', null), ('21', '6', '1', 'title', '', '<p>请输入内容...content</p><p><br></p>', '1', '0', '22', '0', '2016-07-01 00:41:58', '1', '2016-07-01 00:41:58', '2016-07-01 00:41:58', null), ('22', '6', '1', 'H', '', '<p class=\"\">请输入内容...</p><p class=\"\">Hil</p><img src=\"#BASE_URL#FgxdvCRQ6asrgwnXCNagXSW-otjl\" style=\"max-width: 100%;\">', '1', '0', '2', '0', '2016-07-03 09:58:37', '1', '2016-07-03 09:58:37', '2016-07-03 09:58:37', null), ('23', '6', '1', 'Test Title', '', '<p class=\"\"><b>请输入内容...</b></p><p><br></p>', '1', '0', '2', '0', '2016-07-06 15:40:29', '1', '2016-07-06 15:40:29', '2016-07-06 15:40:29', null), ('24', '6', '1', '标题', '', '<p class=\"\">请输入内容...</p><img src=\"#BASE_URL#Fi5AOEz8NtrWNpeT5c6PeTK9c5iy\" style=\"max-width: 100%;\"><p><br></p>', '1', '0', '0', '0', '2016-07-22 18:38:05', '1', '2016-07-22 18:38:05', '2016-07-22 18:38:05', null), ('25', '6', '1', '1111', '请输入内容...Htitle', '<blockquote w_editor_quote_style=\"1\" style=\"display: block; border-left: 5px solid rgb(208, 229, 242); padding: 4px 0px 4px 10px; margin: 4px 0px; background-color: rgb(241, 241, 241);\">请输入内容...</blockquote><p class=\"\"><br></p><p class=\"\"><b>Htitle</b></p><p class=\"\"><br></p><img src=\"#BASE_URL#Fo8kkhlmcrUUdtu2gh1A_fC8NrNm\" style=\"max-width: 100%;\">', '1', '0', '1', '0', '2016-07-31 11:48:04', '1', '2016-07-31 11:48:04', '2016-07-31 11:48:04', null);
COMMIT;

-- ----------------------------
-- Table structure for profile
-- ----------------------------
DROP TABLE IF EXISTS `profile`;
CREATE TABLE `profile` (
`user_refer`  int(11) NOT NULL ,
`avatar`  varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL ,
`coins`  int(11) NOT NULL DEFAULT 0 ,
`bio`  varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`post_count`  int(11) NOT NULL DEFAULT 0 ,
`comment_count`  int(11) NOT NULL DEFAULT 0 ,
`following_count`  int(11) NOT NULL DEFAULT 0 ,
`followed_count`  int(11) NOT NULL DEFAULT 0 ,
PRIMARY KEY (`user_refer`)
)
ENGINE=InnoDB
DEFAULT CHARACTER SET=utf8 COLLATE=utf8_general_ci

;

-- ----------------------------
-- Records of profile
-- ----------------------------
BEGIN;
INSERT INTO `profile` VALUES ('5', 'deaf.png', '9', '这家伙很懒,什么都没有', '0', '3', '0', '0'), ('6', '/static/img/default.png', '9', '这家伙很懒,什么都没有', '0', '8', '0', '0'), ('7', 'deaf.png', '9', '这家伙很懒,什么都没有', '0', '3', '0', '0'), ('11', 'deaf.png', '9', '这家伙很懒,什么都没有', '0', '3', '0', '0'), ('23', '/static/img/default.png', '9', '这家伙很懒,什么都没有', '0', '3', '0', '0'), ('30', '/static/img/default.png', '0', null, '0', '3', '0', '0'), ('31', '/static/img/default.png', '0', null, '0', '3', '0', '0');
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
`visible`  tinyint(4) NOT NULL ,
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
INSERT INTO `swipe` VALUES ('1', 'https://baidu.com', '/static/img/swipe1.jpeg', 'empty', '1', '0000-00-00 00:00:00', '0000-00-00 00:00:00', null), ('2', 'https://baidu.com', '/static/img/swipe2.jpeg', 'empty', '1', '0000-00-00 00:00:00', '0000-00-00 00:00:00', null), ('3', 'https://baidu.com', '/static/img/swipe3.jpeg', 'empty', '1', '0000-00-00 00:00:00', '0000-00-00 00:00:00', null);
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
AUTO_INCREMENT=2

;

-- ----------------------------
-- Records of tag
-- ----------------------------
BEGIN;
INSERT INTO `tag` VALUES ('1', '自拍', null, '#0eb400', '1', '2016-07-07 12:30:30', '2016-07-07 12:30:30', null);
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
INSERT INTO `topic` VALUES ('1', '闲聊', '我们可以在这随便聊什么都行', 'chat', '#795548', '1', '0000-00-00 00:00:00', '0000-00-00 00:00:00', null), ('2', '创意', '好的创意,四溢横飞', 'v', '#3F51B5', '1', '0000-00-00 00:00:00', '0000-00-00 00:00:00', null);
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
AUTO_INCREMENT=32

;

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO `user` VALUES ('5', '', '储根深', '1690512717@qq.com', '273a0c7bd3c679ba9a6f5d99078e36e85d02b952', '', null, '10', '2016-06-24 14:57:23', '2016-06-24 14:57:23', null), ('6', '', 'Hello', 'genshenchu@gmail.com', '7c4a8d09ca3762af61e59520943dc26494f8941b', '', '', '10', '2016-06-24 15:16:00', '2016-06-24 15:16:00', null), ('7', '', 'Hello1', 'genshen220284@163.com', '7c4a8d09ca3762af61e59520943dc26494f8941b', '', '', '0', '2016-06-24 15:52:54', '2016-06-24 15:52:54', null), ('10', '', 'Hello2', '1690512717@qq.com', '7c4a8d09ca3762af61e59520943dc26494f8941b', '', '', '10', '2016-07-06 19:01:14', '2016-07-06 19:01:14', null), ('11', '', 'Hello3', 'genshenchu@gail.com', '7c4a8d09ca3762af61e59520943dc26494f8941b', '', '', '10', '2016-07-09 17:23:29', '2016-07-09 17:23:29', null), ('12', '', 'Hell4', 'genshen220284@163.com', '7c4a8d09ca3762af61e59520943dc26494f8941b', '', '', '10', '2016-07-09 17:39:58', '2016-07-09 17:39:58', null), ('13', '', 'Hello6', 'genshenchu@gmail.co', '7c4a8d09ca3762af61e59520943dc26494f8941b', '', '', '10', '2016-07-09 18:18:16', '2016-07-09 18:18:16', null), ('14', '', 'Hello7', 'ggg@qq.com', '3d4f2bf07dc1be38b20cd6e46949a1071f9d0e3d', '', '', '10', '2016-07-09 21:11:12', '2016-07-09 21:11:12', null), ('15', '', 'Hello56', '11111@qq.com', '3d4f2bf07dc1be38b20cd6e46949a1071f9d0e3d', '', '', '10', '2016-07-09 21:12:19', '2016-07-09 21:12:19', null), ('16', '', 'Hello53', '111@qq.com', '7c4a8d09ca3762af61e59520943dc26494f8941b', '', '', '10', '2016-07-09 21:13:31', '2016-07-09 21:13:31', null), ('19', '', '', '1112@qq.com', '3d4f2bf07dc1be38b20cd6e46949a1071f9d0e3d', '', '', '10', '2016-07-16 11:26:08', '2016-07-16 11:26:08', null), ('20', '', '', '11113@qq.com', '77bce9fb18f977ea576bbcd143b2b521073f0cd6', '', '', '10', '2016-07-16 11:46:12', '2016-07-16 11:46:12', null), ('21', '', '', '1111111@qq.com', '3d4f2bf07dc1be38b20cd6e46949a1071f9d0e3d', '', '', '10', '2016-07-16 12:02:26', '2016-07-16 12:02:26', null), ('22', '', '', 'www@ww.com', 'b33b5e3e04dae7c04d1e4dc759ca5c80e26e576a', '', '', '10', '2016-07-16 12:04:06', '2016-07-16 12:04:06', null), ('23', '', '无名氏', '1111211@qq.com', '3d4f2bf07dc1be38b20cd6e46949a1071f9d0e3d', '', '', '10', '2016-07-16 12:04:36', '2016-07-16 12:04:36', null), ('24', '', '', '1234@111.com', '3d4f2bf07dc1be38b20cd6e46949a1071f9d0e3d', '', '', '10', '2016-07-16 16:10:46', '2016-07-16 16:10:46', null), ('25', '', 'half储根深', 'gens@gmal.com', '3d4f2bf07dc1be38b20cd6e46949a1071f9d0e3d', '', '', '10', '2016-07-16 16:11:41', '2016-07-16 16:11:41', null), ('30', '', '123', 'm18813128117@163.com', '7c4a8d09ca3762af61e59520943dc26494f8941b', '', '', '9', '2016-07-26 01:33:53', '2016-07-26 01:33:53', null), ('31', '', '2222', 'genshen220284@foxmail.com', '273a0c7bd3c679ba9a6f5d99078e36e85d02b952', '', '', '9', '2016-07-27 10:46:06', '2016-07-27 10:46:06', null);
COMMIT;

-- ----------------------------
-- Auto increment value for comment
-- ----------------------------
ALTER TABLE `comment` AUTO_INCREMENT=73;

-- ----------------------------
-- Auto increment value for feedback
-- ----------------------------
ALTER TABLE `feedback` AUTO_INCREMENT=3;

-- ----------------------------
-- Auto increment value for notification
-- ----------------------------
ALTER TABLE `notification` AUTO_INCREMENT=27;

-- ----------------------------
-- Auto increment value for post_message
-- ----------------------------
ALTER TABLE `post_message` AUTO_INCREMENT=1;

-- ----------------------------
-- Auto increment value for posts
-- ----------------------------
ALTER TABLE `posts` AUTO_INCREMENT=26;

-- ----------------------------
-- Auto increment value for swipe
-- ----------------------------
ALTER TABLE `swipe` AUTO_INCREMENT=4;

-- ----------------------------
-- Auto increment value for tag
-- ----------------------------
ALTER TABLE `tag` AUTO_INCREMENT=2;

-- ----------------------------
-- Auto increment value for topic
-- ----------------------------
ALTER TABLE `topic` AUTO_INCREMENT=3;

-- ----------------------------
-- Auto increment value for user
-- ----------------------------
ALTER TABLE `user` AUTO_INCREMENT=32;
