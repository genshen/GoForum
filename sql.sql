/*
Navicat MySQL Data Transfer

Source Server         : MySQL
Source Server Version : 50505
Source Host           : 127.0.0.2:3306
Source Database       : goforum

Target Server Type    : MYSQL
Target Server Version : 50505
File Encoding         : 65001

Date: 2016-11-18 11:03:29
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
AUTO_INCREMENT=113

;

-- ----------------------------
-- Records of comment
-- ----------------------------
BEGIN;
INSERT INTO `comment` VALUES ('80', '26', '34', '0', 'wwwww', '1', '2016-08-07 01:06:24', '2016-08-07 01:06:24', null), ('81', '26', '34', '0', 'wwwwwwwwwwwwwwwwwwwwwwwwwwwww', '1', '2016-08-07 01:07:41', '2016-08-07 01:07:41', null), ('89', '26', '34', '0', '@Hee#213 ddd?', '0', '2016-08-07 01:13:30', '2016-08-07 01:13:30', null), ('90', '26', '34', '0', '@Hee#213 ', '0', '2016-08-07 01:13:40', '2016-08-07 01:13:40', null), ('91', '26', '34', '0', '@Hee#213 ', '0', '2016-08-07 01:13:51', '2016-08-07 01:13:51', null), ('92', '26', '34', '0', '@Hee#213 fffff', '0', '2016-08-07 01:15:12', '2016-08-07 01:15:12', null), ('93', '26', '34', '0', '@Hee#213 why?', '0', '2016-08-07 01:16:02', '2016-08-07 01:16:02', null), ('99', '26', '34', '0', '@Hee#213 ', '0', '2016-08-07 01:25:34', '2016-08-07 01:25:34', null), ('100', '26', '34', '0', '@Hee#213 really?', '0', '2016-08-07 01:26:17', '2016-08-07 01:26:17', null), ('101', '26', '34', '0', '@Hee#213  big moutain!', '0', '2016-08-07 01:27:38', '2016-08-07 01:27:38', null), ('102', '26', '34', '0', '哈？', '0', '2016-08-07 01:30:15', '2016-08-07 01:30:15', null), ('103', '27', '34', '0', 'I\'m the first one~', '0', '2016-08-07 01:31:08', '2016-08-07 01:31:08', null), ('104', '27', '34', '0', '@Hee#213 k', '0', '2016-08-07 01:31:55', '2016-08-07 01:31:55', null), ('105', '27', '34', '0', '@Hee#213 ', '0', '2016-08-07 03:07:22', '2016-08-07 03:07:22', null), ('106', '29', '34', '0', ' success', '0', '2016-09-09 01:07:09', '2016-09-09 01:07:09', null), ('107', '26', '34', '0', 'hhhh', '1', '2016-09-10 01:28:31', '2016-09-10 01:28:31', null), ('108', '26', '34', '0', '@Hee#213 ', '1', '2016-09-10 02:09:42', '2016-09-10 02:09:42', null), ('109', '30', '34', '0', 'ok', '1', '2016-09-10 02:15:28', '2016-09-10 02:15:28', null), ('110', '33', '33', '0', '大家快来分享吧~', '1', '2016-11-18 10:43:20', '2016-11-18 10:43:20', null), ('111', '32', '38', '0', '好的,我一定会遵守相关规则的~', '1', '2016-11-18 10:52:51', '2016-11-18 10:52:51', null), ('112', '34', '38', '0', '深度好文,顶顶顶~', '1', '2016-11-18 10:58:39', '2016-11-18 10:58:39', null);
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
AUTO_INCREMENT=6

;

-- ----------------------------
-- Records of feedback
-- ----------------------------
BEGIN;
INSERT INTO `feedback` VALUES ('4', '34', '0', 'wwwww', 'wwwww', '2016-08-07 00:23:10', '2016-08-07 00:23:10', null), ('5', '0', '0', 'wenti 1', '', '2016-09-22 00:37:23', '2016-09-22 00:37:23', null);
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
INSERT INTO `follows` VALUES ('33', '34'), ('33', '38'), ('34', '33');
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
AUTO_INCREMENT=3

;

-- ----------------------------
-- Records of notification
-- ----------------------------
BEGIN;
INSERT INTO `notification` VALUES ('1', '33', '34', '0', 0x7B22757365726E616D65223A2267656E7368227D, '0', '2016-09-22 01:35:21', '2016-09-22 01:35:21', null), ('2', '38', '33', '0', 0x7B22757365726E616D65223A22E5B08FE5BA93227D, '0', '2016-11-18 11:01:28', '2016-11-18 11:01:28', null);
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
AUTO_INCREMENT=12

;

-- ----------------------------
-- Records of post_message
-- ----------------------------
BEGIN;
INSERT INTO `post_message` VALUES ('8', '34', '34', 'gensh', '0', '30', 'Hi ,xss test', 'ok', '请输入内容...M&lt;script&gt;alert(1)&lt;/script&gt;', '0', '2016-09-10 02:15:28', '2016-09-10 02:15:28', null), ('9', '34', '33', '小库', '0', '33', 'Clothes+服装创意论坛情感故事区明细beta1.0', '大家快来分享吧~', '一套情侣装、亲子装、闺蜜装是彼此感情的最好见证.用文字记录下属于你们之间的感情,附上专属你们的服装照片,与大家共同分享属于你们的温暖.除此以外,你还可以上传你亲自设计用来致敬偶像的服装照片,表达对偶像的倾慕之情,展现你的情怀..我们会不定时的选择有温度的故事,免费把你们的情怀转化为衣服,送到你们的手中.', '0', '2016-11-18 10:43:20', '2016-11-18 10:43:20', null), ('10', '34', '38', 'Gesn', '0', '32', 'Clothes+服装创意论坛灌水闲聊区规则beta1.0', '好的,我一定会遵守相关规则的~', '本区为大家提供闲聊的场所,大家可以在此畅所欲言.在这里,大家可以晒出经过自己精心设计的成品,可以分享一些设计的心得,可以吐糟生活中的琐事等等.附注:为维护网上公共秩序和社会稳定，请您自觉遵守以下条款：1. 互相尊重，遵守互联网络道德，遵守中华人民共和国的各项有关法律法规。2. 不得发表损害国家荣誉和利益的，破坏国家宗教政策，宣扬邪教和封建迷信的言论。3. 不得发表煽动民族仇恨、民族歧视，破坏民族团结的言论。4. 不得发表捏造或者歪曲事实，散布谣言，扰乱社会秩序，破坏社会稳定的言论。5. 不得发表侮辱他人或', '0', '2016-11-18 10:52:51', '2016-11-18 10:52:51', null), ('11', '33', '38', 'Gesn', '0', '34', '服装设计专业︱服装设计时需要考虑的五个因素', '深度好文,顶顶顶~', '服装设计，是目前最好就业，也最有发展潜力的一门新型行业。作为企业的骨干力量，主心骨。我们设计时应该注意哪些呢？服装设计时需要考虑的五个因素：1、什么人穿：每个人不仅年龄、性别、体型、相貌不同，而且文化素质、艺术修养、兴趣爱好、社会地位、职业范围、经济能力以及生活的态度也都不一样.因此,在设计服装时必须考虑是为谁设计的,他或她的具体情况是什么.2、什么时候穿:什么时候穿要考虑两个因素:第一,是在一年四季,一般按照春、夏、秋、冬四个季节设计不同的服装,这一点我想大家都知道;第二按照每天不同时间来分,分为日装(', '0', '2016-11-18 10:58:40', '2016-11-18 10:58:40', null);
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
AUTO_INCREMENT=38

;

-- ----------------------------
-- Records of posts
-- ----------------------------
BEGIN;
INSERT INTO `posts` VALUES ('30', '38', '1', 'Hi ,xss test', '请输入内容...M&lt;script&gt;alert(1)&lt;/script&gt;', '<p>请输入内容...</p><h3 class=\"\">M</h3><div class=\"\"><br></div><img src=\"#BASE_URL#FgxdvCRQ6asrgwnXCNagXSW-otjl\" style=\"max-width: 100%;\"><div class=\"\"><br></div><div class=\"\"><script>alert(1)</script></div>', '1', '0', '1', '0', '2016-09-09 00:52:49', '1', '2016-09-09 00:52:49', '2016-09-09 00:52:49', null), ('31', '34', '1', 'Clothes+服装创意论坛创意发布规则beta1.0', '本区专门提供一个创意分享及转化的平台.分为以下几个步骤:⑴分享创意把你的想法创意用文字或图片的形式发布到社区中.⑵创意得到支持.当创意分享到社区之后,其他人可以选择支持该创意.当支持的人数达到 &nbsp;人时,即可进入下一环节(获取免费设计).⑶获取免费设计.随后我们有专业的设计师会和创意提供者联系,把创意转化为实实在在的设计图.⑷众筹转化随后把产品成果图发布在众筹页面,当众筹人数达到 &nbsp;人时,即可进入下一环节(获取私人订制及享受收益)⑸获取私人订制及享受收益众筹结束后,可以免费获得该创意转化', '<p>本区专门提供一个创意分享及转化的平台.分为以下几个步骤:</p><p>⑴分享创意</p><p>把你的想法创意用文字或图片的形式发布到社区中.</p><p>⑵创意得到支持.</p><p>当创意分享到社区之后,其他人可以选择支持该创意.当支持的人数达到 &nbsp;人时,即可进入下一环节(获取免费设计).</p><p>⑶获取免费设计.</p><p>随后我们有专业的设计师会和创意提供者联系,把创意转化为实实在在的设计图.</p><p>⑷众筹转化</p><p>随后把产品成果图发布在众筹页面,当众筹人数达到 &nbsp;人时,即可进入下一环节(获取私人订制及享受收益)</p><p>⑸获取私人订制及享受收益</p><p>众筹结束后,可以免费获得该创意转化的成品一件.并给予众筹件数乘以 &nbsp;元的收益.</p><p><br></p><p>期盼着大家在这里能够享受创造,玩的开心!</p>', '1', '1', '0', '125', '2016-11-18 10:24:18', '1', '2016-11-18 10:24:18', '2016-11-18 10:24:18', null), ('32', '34', '1', 'Clothes+服装创意论坛灌水闲聊区规则beta1.0', '本区为大家提供闲聊的场所,大家可以在此畅所欲言.在这里,大家可以晒出经过自己精心设计的成品,可以分享一些设计的心得,可以吐糟生活中的琐事等等.附注:为维护网上公共秩序和社会稳定，请您自觉遵守以下条款：1. 互相尊重，遵守互联网络道德，遵守中华人民共和国的各项有关法律法规。2. 不得发表损害国家荣誉和利益的，破坏国家宗教政策，宣扬邪教和封建迷信的言论。3. 不得发表煽动民族仇恨、民族歧视，破坏民族团结的言论。4. 不得发表捏造或者歪曲事实，散布谣言，扰乱社会秩序，破坏社会稳定的言论。5. 不得发表侮辱他人或', '<p>本区为大家提供闲聊的场所,大家可以在此畅所欲言.在这里,大家可以晒出经过自己精心设计的成品,可以分享一些设计的心得,可以吐糟生活中的琐事等等.</p><p><br></p><p class=\"\"><b>附注:为维护网上公共秩序和社会稳定，请您自觉遵守以下条款：</b></p><p><br></p><p>1. 互相尊重，遵守互联网络道德，遵守中华人民共和国的各项有关法律法规。</p><p>2. 不得发表损害国家荣誉和利益的，破坏国家宗教政策，宣扬邪教和封建迷信的言论。</p><p>3. 不得发表煽动民族仇恨、民族歧视，破坏民族团结的言论。</p><p>4. 不得发表捏造或者歪曲事实，散布谣言，扰乱社会秩序，破坏社会稳定的言论。</p><p>5. 不得发表侮辱他人或者捏造事实诽谤他人的，侵害他人合法权益的言论。</p><p>6. 不得发表任何对论坛氛围有负面影响的文章。</p><div><br></div>', '1', '1', '1', '18', '2016-11-18 10:30:24', '1', '2016-11-18 10:30:24', '2016-11-18 10:30:24', null), ('33', '34', '1', 'Clothes+服装创意论坛情感故事区明细beta1.0', '一套情侣装、亲子装、闺蜜装是彼此感情的最好见证.用文字记录下属于你们之间的感情,附上专属你们的服装照片,与大家共同分享属于你们的温暖.除此以外,你还可以上传你亲自设计用来致敬偶像的服装照片,表达对偶像的倾慕之情,展现你的情怀..我们会不定时的选择有温度的故事,免费把你们的情怀转化为衣服,送到你们的手中.', '<p class=\"\">一套情侣装、亲子装、闺蜜装是彼此感情的最好见证.用文字记录下属于你们之间的感情,附上专属你们的服装照片,与大家共同分享属于你们的温暖.</p><p class=\"\">除此以外,你还可以上传你亲自设计用来致敬偶像的服装照片,表达对偶像的倾慕之情,展现你的情怀..</p><p class=\"\">我们会不定时的选择有温度的故事,免费把你们的情怀转化为衣服,送到你们的手中.</p><p><br></p>', '1', '1', '1', '50', '2016-11-18 10:31:15', '1', '2016-11-18 10:31:15', '2016-11-18 10:31:15', null), ('34', '33', '1', '服装设计专业︱服装设计时需要考虑的五个因素', '服装设计，是目前最好就业，也最有发展潜力的一门新型行业。作为企业的骨干力量，主心骨。我们设计时应该注意哪些呢？服装设计时需要考虑的五个因素：1、什么人穿：每个人不仅年龄、性别、体型、相貌不同，而且文化素质、艺术修养、兴趣爱好、社会地位、职业范围、经济能力以及生活的态度也都不一样.因此,在设计服装时必须考虑是为谁设计的,他或她的具体情况是什么.2、什么时候穿:什么时候穿要考虑两个因素:第一,是在一年四季,一般按照春、夏、秋、冬四个季节设计不同的服装,这一点我想大家都知道;第二按照每天不同时间来分,分为日装(', '<p class=\"\">服装设计，是目前最好就业，也最有发展潜力的一门新型行业。作为企业的骨干力量，主心骨。</p><p class=\"\"><br></p><p class=\"\">我们设计时应该注意哪些呢？</p><p class=\"\"><br></p><p class=\"\">服装设计时需要考虑的五个因素：</p><p class=\"\"><br></p><p class=\"\">1、什么人穿：每个人不仅年龄、性别、体型、相貌不同，而且文化素质、艺术修养、兴趣爱好、社会地</p><p class=\"\"><br></p><p class=\"\">位、职业范围、经济能力以及生活的态度也都不一样.因此,在设计服装时必须考虑是为谁设计的,他或她的具</p><p class=\"\"><br></p><p class=\"\">体情况是什么.</p><p class=\"\"><br></p><p class=\"\">2、什么时候穿:什么时候穿要考虑两个因素:第一,是在一年四季,一般按照春、夏、秋、冬四个季节设计不同</p><p class=\"\"><br></p><p class=\"\">的服装,这一点我想大家都知道;第二按照每天不同时间来分,分为日装(白天穿)和晚装(晚上穿),大家看过时装</p><p class=\"\"><br></p><p class=\"\">发布会就会知道,几乎所有的作品都分为日装和晚装两大类,还有在结婚时穿的结婚礼服,也分日装和晚装.</p><p class=\"\"><br></p><p class=\"\">3、什么地方穿:从大的方面讲,是考虑地理环境和人文环境因素.例如:让你为海南人和东北人设计服装,你首</p><p class=\"\"><br></p><p class=\"\">先应该想到的是这两个地方地理环境的相差很大.从小的方面讲,要考虑的场合.例如:让你设计一套上班时穿</p><p class=\"\"><br></p><p class=\"\">的服装和出去旅游时穿的服装,这两套服装肯定是不一样的.</p><p class=\"\"><br></p><p class=\"\">4、指在具体的场合设计什么样的服装最合适,假如是参加宴会的,或是在朋友的婚礼上做伴娘,在这些场合,你</p><p class=\"\"><br></p><p class=\"\">生计的服装是由穿着者所扮演的角色所决定.如果是为伴娘设计服装,那么你设计的服装色彩、款式和风格不</p><p class=\"\"><br></p><p class=\"\">能太突出,否则会引起误会.</p><p class=\"\"><br></p><p class=\"\">5、为了什么穿:穿的目的不一样,有些人穿衣服是为了炫耀、展示自己,有写些人是为了工作需要,有些人是为</p><p class=\"\"><br></p><p class=\"\">了突出自己的个性,显得与众不同.你要根据不同的穿衣目的去设计不同的服装</p><p class=\"\"><br></p>', '1', '0', '1', '10', '2016-11-18 10:46:43', '1', '2016-11-18 10:46:43', '2016-11-18 10:46:43', null), ('35', '38', '1', '服装与心情', '我有一幅满意的作品，我的心情会很好；我买了一件乘心的衣服，感觉同样不错；偶然有一次，马路上遇到一个与我前任女友穿着同一款式的女孩在我的面前轻拂而去，那感觉真的...', '<p class=\"\">我有一幅满意的作品，我的心情会很好；</p><p class=\"\">我买了一件乘心的衣服，感觉同样不错；</p><p class=\"\">偶然有一次，马路上遇到一个与我前任女友穿着同一款式的女孩在我的面前轻拂而去，那感觉真的...</p><p><br></p>', '1', '0', '0', '0', '2016-11-18 10:55:50', '1', '2016-11-18 10:55:50', '2016-11-18 10:55:50', null), ('36', '38', '1', '美女们必学的牛仔裤气质搭配', '穿够了职业套装，又不能过于随便地出现在办公室，怎么才能把时尚、多变的Jeans 穿出职场感呢？本月请出艾图服饰的时髦搭配指导Linda 携手著名造型师RUSSELL，用简单的时装搭配技巧，巧妙地将潮流牛仔裤打造成气质职场装.1白色缎带装饰衬衫, 牛仔工装短裤,象牙色针织开衫, 高筒皮靴,水晶装饰宽边手镯，大方得体，时尚气质呼之欲出.', '<p class=\"\">穿够了职业套装，又不能过于随便地出现在办公室，怎么才能把时尚、多变的Jeans 穿出职场感呢？本月请出艾图服饰的时髦搭配指导Linda 携手著名造型师RUSSELL，用简单的时装搭配技巧，巧妙地将潮流牛仔裤打造成气质职场装.</p><p class=\"\"><img src=\"#BASE_URL#Fp-DBhe8oisB1BsIU8p2k2nvjM3o\" style=\"max-width: 100%;\"><br></p><p class=\"\"><br></p><p class=\"\">1白色缎带装饰衬衫, 牛仔工装短裤,象牙色针织开衫, 高筒皮靴,水晶装饰宽边手镯，大方得体，时尚气质呼之欲出.<br></p>', '1', '0', '0', '0', '2016-11-18 10:57:41', '1', '2016-11-18 10:57:41', '2016-11-18 10:57:41', null), ('37', '33', '1', '服装设计师必需的', '在服饰这个行业中，如果想在众多的竞争对手中占有一席之地，那么前提条件就是要有及时的专业资讯。企业只要第一时间有了一手的资讯，那么才能有计划得进行设计与生产加工。&nbsp; &nbsp; &nbsp;对于设计师来说资讯是重要而又必须的，否则是无法设计出来的。看来及时的资讯这么重要，那设计师们一定想知道如何才能得到？我现在告诉大家：POP（全球）时尚网络机构是全球第一家中文服饰设计资讯网站，专业为服装、箱包、鞋(国际鞋子设计在线)类等行业企业、贸易公司以及设计师提供最前端的款式方面的流行资讯，内容囊括全球最', '<p class=\"\">在服饰这个行业中，如果想在众多的竞争对手中占有一席之地，那么前提条件就是要有及时的专业资讯。企业只要第一时间有了一手的资讯，那么才能有计划得进行设计与生产加工。</p><p class=\"\">&nbsp; &nbsp; &nbsp;对于设计师来说资讯是重要而又必须的，否则是无法设计出来的。看来及时的资讯这么重要，那设计师们一定想知道如何才能得到？我现在告诉大家：POP（全球）时尚网络机构是全球第一家中文服饰设计资讯网站，专业为服装、箱包、鞋(国际鞋子设计在线)类等行业企业、贸易公司以及设计师提供最前端的款式方面的流行资讯，内容囊括全球最新的流行趋势分析、款式研究、商场实拍、设计手稿、设计图案、店面设计等十几个方面，新颖时尚，即时实用，是目前国内信息量最大，内容最前端，功能最强大，服务最完善的时尚资讯搜索引擎和网络平台。</p><p><br></p>', '1', '0', '0', '0', '2016-11-18 11:00:11', '1', '2016-11-18 11:00:11', '2016-11-18 11:00:11', null);
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
INSERT INTO `profile` VALUES ('33', '33', '小库', '/static/img/avatars/avatar-gray.jpg', '0', '这家伙很懒,什么都没有', '0', '1', '0', '0'), ('34', '34', '社区管理员', '/static/img/avatars/avatar-gray.jpg', '0', '这家伙很懒,什么都没有', '0', '18', '0', '0'), ('35', '35', 'cgs', '/static/img/default.png', '0', '', '0', '0', '0', '0'), ('36', '36', 'c储根深', '/static/img/default.png', '0', '', '0', '0', '0', '0'), ('37', '37', '储根深', '/static/img/default.png', '0', '', '0', '0', '0', '0'), ('38', '38', 'Gesn', '/static/img/avatars/avatar-green.jpg', '0', '', '0', '2', '0', '0');
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
INSERT INTO `swipe` VALUES ('1', 'https://baidu.com', 'static/img/swipe2.jpg', 'no', '1', '0000-00-00 00:00:00', '0000-00-00 00:00:00', null), ('2', 'https://baidu.com', 'static/img/swipe1.jpg', 'no', '1', '0000-00-00 00:00:00', '0000-00-00 00:00:00', null), ('3', 'https://baidu.com', 'static/img/swipe3.jpg', 'no', '1', '0000-00-00 00:00:00', '0000-00-00 00:00:00', null);
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
AUTO_INCREMENT=4

;

-- ----------------------------
-- Records of topic
-- ----------------------------
BEGIN;
INSERT INTO `topic` VALUES ('1', '服装创意区', '把握创意,做自己的导演', 'chat', '/static/img/topics/topic1.png', '#f44336', '1', '0000-00-00 00:00:00', '0000-00-00 00:00:00', null), ('2', '灌水闲聊区', '一个嬉戏打闹、畅所欲言的港湾', 'discus', '/static/img/topics/topic2.png', '#673ab7', '1', '0000-00-00 00:00:00', '0000-00-00 00:00:00', null), ('3', '情感故事区', '用有温度的文字记录温暖的感情', 'story', '/static/img/topics/topic3.png', '#673ab7', '1', '0000-00-00 00:00:00', '0000-00-00 00:00:00', null);
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
INSERT INTO `user` VALUES ('33', '', '小库', 'm18813128117@163.com', '7c4a8d09ca3762af61e59520943dc26494f8941b', '', '', '2', '2016-08-06 23:24:11', '2016-08-06 23:24:11', null), ('34', '', '社区管理员', 'genshenchu@gmail.com', '7c4a8d09ca3762af61e59520943dc26494f8941b', '', '', '2', '2016-08-07 00:18:38', '2016-08-07 00:18:38', null), ('38', '', 'Gesn', '1690512717@qq.com', '7c4a8d09ca3762af61e59520943dc26494f8941b', '', '', '1', '2016-09-08 23:08:50', '2016-09-08 23:08:50', null);
COMMIT;

-- ----------------------------
-- Auto increment value for comment
-- ----------------------------
ALTER TABLE `comment` AUTO_INCREMENT=113;

-- ----------------------------
-- Auto increment value for feedback
-- ----------------------------
ALTER TABLE `feedback` AUTO_INCREMENT=6;

-- ----------------------------
-- Auto increment value for notification
-- ----------------------------
ALTER TABLE `notification` AUTO_INCREMENT=3;

-- ----------------------------
-- Auto increment value for post_message
-- ----------------------------
ALTER TABLE `post_message` AUTO_INCREMENT=12;

-- ----------------------------
-- Auto increment value for posts
-- ----------------------------
ALTER TABLE `posts` AUTO_INCREMENT=38;

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
ALTER TABLE `topic` AUTO_INCREMENT=4;

-- ----------------------------
-- Auto increment value for user
-- ----------------------------
ALTER TABLE `user` AUTO_INCREMENT=39;
