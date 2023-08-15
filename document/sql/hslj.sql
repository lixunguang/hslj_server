/*
 Navicat Premium Data Transfer

 Source Server         : centos
 Source Server Type    : MySQL
 Source Server Version : 50742
 Source Host           : 172.30.100.23:3306
 Source Schema         : HSLJ_DB

 Target Server Type    : MySQL
 Target Server Version : 50742
 File Encoding         : 65001

 Date: 04/05/2023 15:16:46
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for admin
-- ----------------------------

DROP TABLE IF EXISTS `admin`;
CREATE TABLE `admin`
(
    `id`           int(11)                                                      NOT NULL AUTO_INCREMENT COMMENT '主键',
    `login_id`     varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '登录id，唯一',
    `name`         varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '姓名',
    `password`     varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '密码',
    `role`         tinyint(4)                                                   NOT NULL DEFAULT 1 COMMENT '管理员类型：1  超级管理员 ',
    `phone_number` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '电话号码',
    `creator_id`   int(11)                                                      NOT NULL DEFAULT 0 COMMENT '创建人id',
    `is_deleted`   tinyint(4)                                                   NOT NULL DEFAULT 1 COMMENT '是否删除 1-否 2-是',
    `created_at`   timestamp(0)                                                 NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
    `updated_at`   timestamp(0)                                                 NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT = '管理员，包括超级管理员，管理员，可以登录后台进行编辑和管理'
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of admin
-- ----------------------------
INSERT INTO `admin`
VALUES (1, '0', 'admin2', 'admin2', 0, '', 0, 0, '2023-06-07 09:49:22', '2023-06-07 09:49:22');
INSERT INTO `admin`
VALUES (2, '0', 'admin1', 'admin1', 0, '', 0, 0, '2023-06-07 09:49:27', '2023-06-07 09:49:27');



-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`
(
    `id`           int(11)                                                      NOT NULL AUTO_INCREMENT COMMENT '主键',
    `login_id`     varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '登录id，唯一',
    `name`         varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '姓名',
    `password`     varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '密码',
    `phone_number` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '电话号码',
    `active`       int(11)                                                      NOT NULL DEFAULT 0 COMMENT '活跃度，参与打卡的次数',
    `rank`         int(11)                                                      NOT NULL DEFAULT 0 COMMENT '用户评级',
    `created_at`   timestamp(0)                                                 NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
    `updated_at`   timestamp(0)                                                 NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 22
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT = '用户表，student表'
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user`
VALUES (1, '202311111111', 'user1', '11111111', '0', 0, 1, '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `user`
VALUES (2, '202311111112', 'lixunguang', 'uua', '0', 0, 1, '2023-03-09 16:54:12', '2023-03-09 16:54:12');
INSERT INTO `user`
VALUES (3, '12134', 'admin', 'admin', '0', 0, 1, '2023-03-09 16:55:09', '2023-03-09 16:55:09');
INSERT INTO `user`
VALUES (4, '111222', 'admin', 'admin', '0', 0, 1, '2023-03-21 15:46:25', '2023-04-17 10:29:24');
INSERT INTO `user`
VALUES (12, '122222', 'liteacher', 'adminxx', '0', 0, 1, '2023-03-15 09:16:32', '2023-03-22 18:30:07');

-- ----------------------------
-- Table structure for area
-- ----------------------------
DROP TABLE IF EXISTS `area`;
CREATE TABLE `area`
(
    `id`          int(11)                                                       NOT NULL AUTO_INCREMENT COMMENT '主键',
    `area_name`     varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '区域标题，区域可能为一个省，也可能为一个地区',
    `area_code`     int(11)                                                       NULL     DEFAULT NULL COMMENT '根据腾讯地图的分类code',
    `created_at`  timestamp(0)                                                  NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
    `updated_at`  timestamp(0)                                                  NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 103
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT = '区域'
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of area
-- ----------------------------

INSERT INTO `area`
VALUES (1, '北京市', 110000 , '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `area`
VALUES (2, '天津市', 120000 , '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `area`
VALUES (3, '河北省', 130000 , '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `area`
VALUES (4, '山西省', 140000 , '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `area`
VALUES (5, '内蒙古自治区', 150000 , '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `area`
VALUES (6, '辽宁省', 210000 , '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `area`
VALUES (7, '吉林省', 220000 , '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `area`
VALUES (8, '黑龙江省', 230000 , '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `area`
VALUES (9, '上海市', 310000 , '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `area`
VALUES (10, '江苏省', 320000 , '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `area`
VALUES (11, '浙江省', 330000 , '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `area`
VALUES (12, '安徽省', 340000 , '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `area`
VALUES (13, '福建省', 350000 , '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `area`
VALUES (14, '江西省', 360000 , '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `area`
VALUES (15, '山东省', 370000 , '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `area`
VALUES (16, '河南省', 410000 , '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `area`
VALUES (17, '湖北省', 420000 , '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `area`
VALUES (18, '湖南省', 430000 , '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `area`
VALUES (19, '广东省', 440000 , '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `area`
VALUES (20, '广西壮族自治区', 450000 , '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `area`
VALUES (21, '海南省', 460000 , '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `area`
VALUES (22, '重庆市', 500000 , '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `area`
VALUES (23, '四川省', 510000 , '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `area`
VALUES (24, '贵州省', 520000 , '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `area`
VALUES (25, '云南省', 530000 , '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `area`
VALUES (26, '西藏自治区', 540000 , '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `area`
VALUES (27, '陕西省', 610000 , '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `area`
VALUES (28, '甘肃省', 620000 , '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `area`
VALUES (29, '青海省', 630000 , '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `area`
VALUES (30, '宁夏回族自治区', 640000 , '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `area`
VALUES (31, '新疆维吾尔自治区', 650000 , '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `area`
VALUES (32, '台湾省', 710000 , '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `area`
VALUES (33, '香港特别行政区', 810000 , '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `area`
VALUES (34, '澳门特别行政区', 820000 , '2023-03-09 16:53:08', '2023-04-17 10:29:32');
-- ----------------------------
-- Table structure for location
-- ----------------------------
DROP TABLE IF EXISTS `location`;
CREATE TABLE `location`
(
    `id`         int(11)                                                       NOT NULL AUTO_INCREMENT COMMENT '主键',
    `name`       varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '地点名字',
    `desc`       text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci         NULL COMMENT '地点描述',
    `location`   geometry                                                      NULL COMMENT '位置',
    `frequency`  int(11)                                                       NOT NULL DEFAULT 0 COMMENT '频率',
    `time_str`   text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci         NULL COMMENT '时间信息',
    `people_num` int(11)                                                       NOT NULL DEFAULT 0 COMMENT '参与人数',
    `rating`     int(11)                                                       NOT NULL DEFAULT 0 COMMENT '评分',
    `hot`        int(11)                                                       NOT NULL DEFAULT 0 COMMENT '热度，统计多少人来这里踢过',
    `area_code`  int(11)                                                       NOT NULL DEFAULT 0 COMMENT '省或者地区一级的编码，如北京110000 朝阳区110105',
    `is_auth`    tinyint                                                       NOT NULL DEFAULT 0 COMMENT '是否已经验证0 未验证，1 已经验证',
    `created_at` timestamp(0)                                                  NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
    `updated_at` timestamp(0)                                                  NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 22
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT = '地点表'
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of location
-- ----------------------------

INSERT INTO `location`
VALUES (1, '北京颐和园北宫门', '颐和园是一个皇家园林。', ST_GeomFromText('POINT(121.474103 31.232862)'), 1, "08:00~9:00;3:00~5:00", 2, '0',
        13,110000,
        '0', '2023-03-09 16:53:08',
        '2023-04-17 10:29:32');
INSERT INTO `location`
VALUES (2, '北京天坛', '天坛是明清两代皇帝“祭天”“祈谷”的场所，总面积273公顷。', ST_GeomFromText('POINT(121.474103 31.232862)'), 1,
        "08:00~9:00;3:00~5:00", 2, '0', 20, 110000,'0',
        '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `location`
VALUES (3, '北京玉渊潭', '玉渊潭公园是北京三环以内最大的公园。在辽代，玉渊潭即为蓟城的“城外花园”；在金代，玉渊潭已成为金中都西北郊的游览胜地。金章宗完颜璟在此地修建钓鱼台，玉渊潭因而有了“钓鱼台”之名(无怪乎其东湖边上的国宾馆叫做“钓鱼台国宾馆”)。', ST_GeomFromText('POINT(121.474103 31.232862)'), 1,
        "08:00~9:00;3:00~5:00", 2,
        '0', 14, 110000,'0', '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `location`
VALUES (4, '北京望京北小河', '望京北小河。', ST_GeomFromText('POINT(121.474103 31.232862)'), 1, "08:00~9:00;3:00~5:00", 2, '0', 3,
        110000,'0',
        '2023-03-09 16:53:08',
        '2023-04-17 10:29:32');
INSERT INTO `location`
VALUES (5, '北京地坛', '原名方泽坛，为明清两代皇帝祭祀皇地祗的场所。', ST_GeomFromText('POINT(121.474103 31.232862)'), 1, "08:00~9:00;3:00~5:00",
        2, '0', 15, 110000,'0',
        '2023-03-09 16:53:08', '2023-04-17 10:29:32');

INSERT INTO `location`
VALUES (6, '仰山公园', '一个有很多雕塑的公园。', ST_GeomFromText('POINT(121.474103 31.232862)'), 1, "08:00~9:00;3:00~5:00",
        2, '0', 12,110000, '0',
        '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `location`
VALUES (7, '陶然亭', ' 陶然亭是清代的名亭，也是中国四大名亭之一。', ST_GeomFromText('POINT(121.474103 31.232862)'), 1, "08:00~9:00;3:00~5:00", 2, '0',
        13,110000,
        '0', '2023-03-09 16:53:08',
        '2023-04-17 10:29:32');

INSERT INTO `location`
VALUES (8, '山西1', ' 陶然亭是清代的名亭，也是中国四大名亭之一。', ST_GeomFromText('POINT(121.474103 31.232862)'), 1, "08:00~9:00;3:00~5:00", 2, '0',
        13,140000,
        '0', '2023-03-09 16:53:08',
        '2023-04-17 10:29:32');

INSERT INTO `location`
VALUES (9, '山西2', ' 陶然亭是清代的名亭，也是中国四大名亭之一。', ST_GeomFromText('POINT(121.474103 31.232862)'), 1, "08:00~9:00;3:00~5:00", 2, '0',
        13,140000,
        '0', '2023-03-09 16:53:08',
        '2023-04-17 10:29:32');

INSERT INTO `location`
VALUES (10, '山西3', ' 陶然亭是清代的名亭，也是中国四大名亭之一。', ST_GeomFromText('POINT(121.474103 31.232862)'), 1, "08:00~9:00;3:00~5:00", 2, '0',
        13,140000,
        '0', '2023-03-09 16:53:08',
        '2023-04-17 10:29:32');


INSERT INTO `location`
VALUES (11, '山东1', ' 陶然亭是清代的名亭，也是中国四大名亭之一。', ST_GeomFromText('POINT(121.474103 31.232862)'), 1, "08:00~9:00;3:00~5:00", 2, '0',
        13,370000,
        '0', '2023-03-09 16:53:08',
        '2023-04-17 10:29:32');

INSERT INTO `location`
VALUES (12, '山东12', ' 陶然亭是清代的名亭，也是中国四大名亭之一。', ST_GeomFromText('POINT(121.474103 31.232862)'), 1, "08:00~9:00;3:00~5:00", 2, '0',
        13,370000,
        '0', '2023-03-09 16:53:08',
        '2023-04-17 10:29:32');

INSERT INTO `location`
VALUES (12, '山东13', ' 陶然亭是清代的名亭，也是中国四大名亭之一。', ST_GeomFromText('POINT(121.474103 31.232862)'), 1, "08:00~9:00;3:00~5:00", 2, '0',
        13,370000,
        '0', '2023-03-09 16:53:08',
        '2023-04-17 10:29:32');

-- ----------------------------
-- Table structure for news
-- ----------------------------
DROP TABLE IF EXISTS `news`;
CREATE TABLE `news`
(
    `id`          int(11)                                                       NOT NULL AUTO_INCREMENT COMMENT '主键',
    `title`       varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '新闻标题',
    `content`     text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci         NULL COMMENT '新闻内容',
    `picture_url` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '图片url',
    `author`      varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '发布人',
    `view_count`  int(11)                                                       NULL     DEFAULT NULL COMMENT '浏览次数',
    `created_at`  timestamp(0)                                                  NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
    `updated_at`  timestamp(0)                                                  NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 103
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT = '新闻表'
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of news
-- ----------------------------
INSERT INTO `news`
VALUES (1, '新闻通知1', '111', '', 'admin', 1, '2023-03-14 11:06:09', '2023-03-16 17:39:40');
INSERT INTO `news`
VALUES (2, '新闻通知2', 'abcd', '', 'admin', 1, '2023-03-14 11:06:09', '2023-03-16 17:39:40');
INSERT INTO `news`
VALUES (3, '新闻通知3', 'abcde', '', 'admin', 1, '2023-03-14 11:06:09', '2023-03-16 17:39:40');


-- ----------------------------
-- Table structure for record
-- ----------------------------
DROP TABLE IF EXISTS `record`;
CREATE TABLE `record`
(
    `id`            int(11)                                                      NOT NULL AUTO_INCREMENT COMMENT '主键',
    `user_login_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '登录id，唯一',
    `location_id`   int(11)                                                      NULL     DEFAULT NULL COMMENT '地点id',
    `created_at`    timestamp(0)                                                 NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
    `updated_at`    timestamp(0)                                                 NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 103
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT = '记录表'
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of record
-- ----------------------------
INSERT INTO `record`
VALUES (11, 12134, 3, '2023-07-05 15:50:07', '2023-07-05 19:55:59');
INSERT INTO `record`
VALUES (12, 111222, 2, '2023-07-05 15:49:46', '2023-07-05 19:56:03');
INSERT INTO `record`
VALUES (103, 12134, 3, '2023-07-05 15:43:34', '2023-07-05 19:56:00');



-- ----------------------------
-- Table structure for today
-- ----------------------------
DROP TABLE IF EXISTS `today`;
CREATE TABLE `today`
(
    `id`          int(11)                                                   NOT NULL AUTO_INCREMENT COMMENT '主键',
    `lunar_desc`  tinytext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '农历节日描述',
    `txt`         text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci     NULL COMMENT '一段美好的文字',
    `picture_url` tinytext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT 'url',
    `created_at`  timestamp(0)                                              NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
    `updated_at`  timestamp(0)                                              NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 103
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT = '今天'
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of today
-- ----------------------------
INSERT INTO `today`
VALUES (11, ‘qqqqq’, ‘ccc’, ‘fsad’, '2023-07-05 15:50:07', '2023-07-05 19:52:59');
INSERT INTO `today`
VALUES (22, ‘aaaaaa’, ‘dd’, ‘fsdf’, '2023-07-05 15:50:07', '2023-07-05 19:55:59');
INSERT INTO `today`
VALUES (33, ‘bbbbbbbb’, ‘dd’, ‘da’, '2023-07-05 15:50:07', '2023-07-05 19:54:59');

SET FOREIGN_KEY_CHECKS = 1;


