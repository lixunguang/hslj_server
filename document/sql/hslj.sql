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
-- Table structure for action
-- ----------------------------
DROP TABLE IF EXISTS `action`;
CREATE TABLE `action`
(
    `id`          int(11)                                                       NOT NULL AUTO_INCREMENT COMMENT '主键',
    `title`       varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '动作名称',
    `content`     text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci         NULL COMMENT '动作介绍',
    `created_at`  timestamp(0)                                                  NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
    `updated_at`  timestamp(0)                                                  NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 103
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT = '毽子动作'
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of action
-- ----------------------------
INSERT INTO `action`
VALUES (1, '盘踢', '盘踢，即用左右两脚互换踢毽。髋关节和膝关节放松，踝关节发力带动小腿上摆，膝关节向外摆，大腿自内向外翻转，用脚的内侧向上垂直踢起毽子。一般踢起的高度不超过下颌。
初练习时，可以先不用毽子，模仿踢毽子的动作做“空踢”练习。练习时，一腿站立，另一腿膝关节外张，大腿翻转至内侧向上抬起，脚尖向前，脚跟与直立腿保持一脚远的距离，高度约与直立腿膝关节相同。抬好后，坚持几秒钟再放下，换另一条腿练习。两条腿的动作基本准确定型后，就可以用毽子直接练习了。这种方法可以比较快地掌握入门技术。',
        '2021-01-22 11:06:09', '2021-01-22 17:39:40');
INSERT INTO `action`
VALUES (2, '绷踢', '绷踢，即用脚尖外三趾部分互换踢毽子（其他部位踢出的毽子不稳，并且容易砸脚），髋关节、膝关节、踝关节放松，大腿向前抬起，身体成150°～160°夹角。踢毽时，脚尖外三趾部位与脚跟同时发力，使脚尖外三趾向上发力时带动全脚向上勾起。两脚跟发力带动小腿向前摆出，大腿保持原角度，将毽子踢起，高低均可。',
        '2023-03-14 11:06:09', '2022-05-25 17:39:40');
INSERT INTO `action`
VALUES (3, '磕踢', '磕踢，即用两腿膝盖部分互换踢毽子。髋关节、膝关节、踝关节放松，小腿自然下垂，足尖稍指地，膝关节发力，带动大腿上摆，将毽子撞起，一般不超过下颌。
练习时，双肘放于腰间，掌心向下，前臂前伸，同上臂成90°夹角不动。用“空踢”的方法进行练习，膝关节发力，用膝盖部位撞击双手发出声响，很像用手击打膝盖的声音(但绝不是击打)，如果声音不像，那么就是动作错误，原因不外是膝关节、髋关节没有放松，小腿没有自然下垂，造成膝关节发力受到限制，此时需要及时调整动作。声音相像后，再使用毽子进行练习，就容易多了。',
        '2023-03-14 11:06:09', '2022-05-25 17:39:40');
INSERT INTO `action`
VALUES (4, '拐踢', '拐踢，即用两脚外侧互换踢毽子。髋关节、膝关节放松自然下垂，勾脚尖，踝关节发力带动小腿，向体侧后上方摆动，当踢毽的瞬间，踢毽脚的脚内侧应平行于地面，高度约为30厘米左右，踢起的毽子一般与肩膀平齐。
练习时，先采用“空踢”的方法进行练习，动作基本准确后，再用“一踢一接”的方法练习，即用手将毽子在体侧抛起，高度约与肩部平齐，用拐踢将毽子踢起，用手将毽子接住，再抛、再踢、再接（这种方法适用于各种踢法的练习)。动作基本准确后，改为两踢一接、三踢一接……熟练后不再用手接毽子，进行双腿互换的练习，次数越多越好。',
        '2023-03-14 11:06:09', '2022-05-25 17:39:40');


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
  COLLATE = utf8mb4_general_ci COMMENT = '用户表，用来存储用户信息'
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
    `id`         int(11)                                                       NOT NULL AUTO_INCREMENT COMMENT '主键',
    `area_name`  varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '区域标题，区域可能为一个省，也可能为一个地区',
    `area_code`  int(11)                                                       NULL     DEFAULT NULL COMMENT '根据腾讯地图的分类code',
    `created_at` timestamp(0)                                                  NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
    `updated_at` timestamp(0)                                                  NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 103
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT = '区域,踢毽地点所属区域'
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of area
-- ----------------------------

INSERT INTO `area`
VALUES (1, '北京市', 110000, '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `area`
VALUES (2, '天津市', 120000, '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `area`
VALUES (3, '河北省', 130000, '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `area`
VALUES (4, '山西省', 140000, '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `area`
VALUES (5, '内蒙古自治区', 150000, '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `area`
VALUES (6, '辽宁省', 210000, '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `area`
VALUES (7, '吉林省', 220000, '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `area`
VALUES (8, '黑龙江省', 230000, '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `area`
VALUES (9, '上海市', 310000, '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `area`
VALUES (10, '江苏省', 320000, '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `area`
VALUES (11, '浙江省', 330000, '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `area`
VALUES (12, '安徽省', 340000, '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `area`
VALUES (13, '福建省', 350000, '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `area`
VALUES (14, '江西省', 360000, '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `area`
VALUES (15, '山东省', 370000, '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `area`
VALUES (16, '河南省', 410000, '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `area`
VALUES (17, '湖北省', 420000, '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `area`
VALUES (18, '湖南省', 430000, '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `area`
VALUES (19, '广东省', 440000, '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `area`
VALUES (20, '广西壮族自治区', 450000, '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `area`
VALUES (21, '海南省', 460000, '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `area`
VALUES (22, '重庆市', 500000, '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `area`
VALUES (23, '四川省', 510000, '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `area`
VALUES (24, '贵州省', 520000, '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `area`
VALUES (25, '云南省', 530000, '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `area`
VALUES (26, '西藏自治区', 540000, '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `area`
VALUES (27, '陕西省', 610000, '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `area`
VALUES (28, '甘肃省', 620000, '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `area`
VALUES (29, '青海省', 630000, '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `area`
VALUES (30, '宁夏回族自治区', 640000, '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `area`
VALUES (31, '新疆维吾尔自治区', 650000, '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `area`
VALUES (32, '台湾省', 710000, '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `area`
VALUES (33, '香港特别行政区', 810000, '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `area`
VALUES (34, '澳门特别行政区', 820000, '2023-03-09 16:53:08', '2023-04-17 10:29:32');

-- ----------------------------
-- Table structure for active_location
-- ----------------------------
DROP TABLE IF EXISTS `active_location`;
CREATE TABLE `active_location`
(
    `id`           int(11)      NOT NULL AUTO_INCREMENT COMMENT '主键',
    `location_id`  int(11)      NOT NULL DEFAULT 0 COMMENT '地点id',
    `minute_start` int(11)      NOT NULL DEFAULT 0 COMMENT '距离开始需要的时间，单位分钟',
    `hot`        int(11)         NOT NULL DEFAULT 0 COMMENT '热度，统计多少人来这里踢过',
    `area_code`    int(11)      NOT NULL DEFAULT 0 COMMENT '省或者地区一级的编码，如北京110000 朝阳区110105',
    `created_at`   timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
    `updated_at`   timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 22
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT = '活跃地点表，动态生成主要关注地点的动态活跃信息（正在进行或者将要进行，用来在首页提示用户，需要有个后台进程维护这个表），从location表而来，设计的目的是为了减轻location表的负担。'
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of active_location
-- ----------------------------
INSERT INTO `active_location`
VALUES (1, 1, 60, 120, 140000,  '2023-03-09 16:53:08',   '2023-04-17 10:29:32');
INSERT INTO `active_location`
VALUES (2, 3, 180, 120, 140000,  '2023-03-09 16:53:08',   '2023-04-17 10:29:32');
INSERT INTO `active_location`
VALUES (3, 7, 120, 120, 140000,  '2023-03-09 16:53:08',   '2023-04-17 10:29:32');
INSERT INTO `active_location`
VALUES (4, 8, 30, 120, 140000,  '2023-03-09 16:53:08',   '2023-04-17 10:29:32');

-- ----------------------------
-- Table structure for location
-- ----------------------------
DROP TABLE IF EXISTS `location`;
CREATE TABLE `location`
(
    `id`         int(11)                                                       NOT NULL AUTO_INCREMENT COMMENT '主键',
    `name`       varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '地点名字',
    `desc`       text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci         NULL COMMENT '地点描述，整体描述',
    `location_desc`       text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci         NULL COMMENT '地点描述2,位置相关',
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
  COLLATE = utf8mb4_general_ci COMMENT = '地点表，会记录所有的地点及详细信息'
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of location
-- ----------------------------

INSERT INTO `location`
VALUES (1, '北京颐和园北宫门', '颐和园是一个皇家园林。', '', ST_GeomFromText('POINT(121.474103 31.232862)'), 1, "08:00~9:00;3:00~5:00", 2, '0',
        13, 110000,
        '0', '2023-03-09 16:53:08',
        '2023-04-17 10:29:32');
INSERT INTO `location`
VALUES (2, '北京天坛', '天坛是明清两代皇帝“祭天”“祈谷”的场所，总面积273公顷。', '', ST_GeomFromText('POINT(121.474103 31.232862)'), 1,
        "08:00~9:00;3:00~5:00", 2, '0', 20, 110000, '0',
        '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `location`
VALUES (3, '北京玉渊潭',
        '玉渊潭公园是北京三环以内最大的公园。在辽代，玉渊潭即为蓟城的“城外花园”；在金代，玉渊潭已成为金中都西北郊的游览胜地。金章宗完颜璟在此地修建钓鱼台，玉渊潭因而有了“钓鱼台”之名(无怪乎其东湖边上的国宾馆叫做“钓鱼台国宾馆”)。',
        '' , ST_GeomFromText('POINT(121.474103 31.232862)'), 1,
        "08:00~9:00;3:00~5:00", 2,
        '0', 14, 110000, '0', '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `location`
VALUES (4, '北京望京北小河', '望京北小河。', '',  ST_GeomFromText('POINT(121.474103 31.232862)'), 1, "08:00~9:00;3:00~5:00", 2, '0', 3,
        110000, '0',
        '2023-03-09 16:53:08',
        '2023-04-17 10:29:32');
INSERT INTO `location`
VALUES (5, '北京地坛', '原名方泽坛，为明清两代皇帝祭祀皇地祗的场所。', '', ST_GeomFromText('POINT(121.474103 31.232862)'), 1, "08:00~9:00;3:00~5:00",
        2, '0', 15, 110000, '0',
        '2023-03-09 16:53:08', '2023-04-17 10:29:32');

INSERT INTO `location`
VALUES (6, '仰山公园', '一个有很多雕塑的公园。', '', ST_GeomFromText('POINT(121.474103 31.232862)'), 1, "08:00~9:00;3:00~5:00",
        2, '0', 12, 110000, '0',
        '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `location`
VALUES (7, '陶然亭', ' 陶然亭是清代的名亭，也是中国四大名亭之一。', '', ST_GeomFromText('POINT(121.474103 31.232862)'), 1, "08:00~9:00;3:00~5:00",
        2, '0',
        13, 110000,
        '0', '2023-03-09 16:53:08',
        '2023-04-17 10:29:32');

INSERT INTO `location`
VALUES (8, '山西1', ' 陶然亭是清代的名亭，也是中国四大名亭之一。', '', ST_GeomFromText('POINT(121.474103 31.232862)'), 1, "08:00~9:00;3:00~5:00",
        2, '0',
        13, 140000,
        '0', '2023-03-09 16:53:08',
        '2023-04-17 10:29:32');

INSERT INTO `location`
VALUES (9, '山西2', ' 陶然亭是清代的名亭，也是中国四大名亭之一。', '', ST_GeomFromText('POINT(121.474103 31.232862)'), 1, "08:00~9:00;3:00~5:00",
        2, '0',
        13, 140000,
        '0', '2023-03-09 16:53:08',
        '2023-04-17 10:29:32');

INSERT INTO `location`
VALUES (10, '山西3', ' 陶然亭是清代的名亭，也是中国四大名亭之一。', '', ST_GeomFromText('POINT(121.474103 31.232862)'), 1, "08:00~9:00;3:00~5:00",
        2, '0',
        13, 140000,
        '0', '2023-03-09 16:53:08',
        '2023-04-17 10:29:32');


INSERT INTO `location`
VALUES (11, '山东1', ' 陶然亭是清代的名亭，也是中国四大名亭之一。', '', ST_GeomFromText('POINT(121.474103 31.232862)'), 1, "08:00~9:00;3:00~5:00",
        2, '0',
        13, 370000,
        '0', '2023-03-09 16:53:08',
        '2023-04-17 10:29:32');

INSERT INTO `location`
VALUES (12, '山东12', ' 陶然亭是清代的名亭，也是中国四大名亭之一。', '', ST_GeomFromText('POINT(121.474103 31.232862)'), 1, "08:00~9:00;3:00~5:00",
        2, '0',
        13, 370000,
        '0', '2023-03-09 16:53:08',
        '2023-04-17 10:29:32');

INSERT INTO `location`
VALUES (13, '山东13', ' 陶然亭是清代的名亭，也是中国四大名亭之一。', '', ST_GeomFromText('POINT(121.474103 31.232862)'), 1, "08:00~9:00;3:00~5:00",
        2, '0',
        13, 370000,
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
  COLLATE = utf8mb4_general_ci COMMENT = '最新信息表，用来展示毽子的相关新闻'
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of news
-- ----------------------------
INSERT INTO `news`
VALUES (1, '大马路小学', '大马路小学冬季锻炼系列赛之踢毽子比赛，经过为期一周的比赛，近期落下了帷幕。本次比赛考虑到学生的实际水平，采取了趣味比赛的形式，设置了两个集体项目，要求全员参与，分别是踢毽接龙和踢毽比远。',
        'https://2023.ibe.cn/wp-content/uploads/2023/08/2023082201392746-1024x682.png', 'admin', 1,
        '2021-01-22 11:06:09', '2021-01-22 17:39:40');
INSERT INTO `news`
VALUES (2, '彩毽飞扬，凝聚力量', '为营造和谐、快乐的工作环境，丰富活跃职工文化生活，提高职工身体素质与团队凝聚力。5月21日上午，沃尔德北方工厂举办了一场团体踢毽子比赛，来自不同部门的90余名职工参加了比赛。',
        'https://2023.ibe.cn/wp-content/uploads/2023/05/2023051214530675-e1683875105592.png', 'admin', 1,
        '2023-03-14 11:06:09', '2022-05-25 17:39:40');
INSERT INTO `news`
VALUES (3, '喜迎二十大，健康向未来', '为大力营造喜迎党的二十大浓厚氛围，力学与土木工程学院工会将举办首届“砼心圆——踢毽子”趣味比赛。
一、活动主题
喜迎二十大，健康向未来
二、活动时间
10月14日下午4:30
三、活动地点
学院西门外空地
四、活动规则
本次比赛只设置个人单项赛，一分钟计时。比赛开始后，参加者可以用脚、下肢等任意部位接触毽子。根据一分钟内踢的次数进行名次排序，次数相同者中途落地次数少者排名靠前；如仍难以排定次序，另加赛10个，用时短者胜出。
比赛设一、二、三等奖若干名，其中第1名为一等奖、第2名和第3名为二等奖、其余完成比赛的为三等奖。
本次比赛为学院首届踢毽子比赛，让我们走出办公室，活动筋骨、舒缓疲惫、放松身心，以更加积极向上的精气神迎接党的二十大胜利召开。',
        'https://2023.ibe.cn/wp-content/uploads/2023/05/2023051214552348-e1683874847942.png', 'admin', 1,
        '2023-03-14 11:06:09', '2023-03-16 17:39:40');


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
  COLLATE = utf8mb4_general_ci COMMENT = '记录表，用来记录用户的打卡记录'
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
    `id`          int(11)                                                      NOT NULL AUTO_INCREMENT COMMENT '主键',
    `lunar_desc`  tinytext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci    NULL COMMENT '农历节日描述',
    `txt`         text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci        NULL COMMENT '一段美好的文字',
    `picture_url` tinytext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci    NULL COMMENT 'url',
    `date`        varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '日期，如20230902',
    `created_at`  timestamp(0)                                                 NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
    `updated_at`  timestamp(0)                                                 NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 103
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT = '今天，非核心功能，用来显示当日的一些信息'
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of today
-- ----------------------------

INSERT INTO `today`
VALUES (11, '白露:露从今夜白，月是故乡明。',
        '世事早已擦肩而过，我们又何必反复追忆，反复提起。是时候和昨天告别了，忘记一切，也原谅一切。是真的忘记，做到心平气和，在安稳的现世里，循规蹈矩的过日子。不再追求虚浮的奢华，不再喜好俏丽的颜色，不再渴望热烈的爱情。只愿在简约的四季里，穿粗布素衣，和某个平淡的人，一同老去，相约白头。',
        'https://desktop.github.com/images/desktop-icon.svg', '20230902', '2023-07-05 15:50:07', '2023-07-05 19:52:59');
INSERT INTO `today`
VALUES (22, 'aaaaaa', '20230913', 'https://desktop.github.com/images/desktop-icon.svg', '20230903',
        '2023-07-05 15:50:07', '2023-07-05 19:55:59');
INSERT INTO `today`
VALUES (33, 'bbbbbbbb', '20230914', 'https://desktop.github.com/images/desktop-icon.svg', '20230904',
        '2023-07-05 15:50:07', '2023-07-05 19:54:59');
INSERT INTO `today`
VALUES (44, 'bbbbbbbb', '20230915', 'https://desktop.github.com/images/desktop-icon.svg', '20230905',
        '2023-07-05 15:50:07', '2023-07-05 19:54:59');
INSERT INTO `today`
VALUES (55, 'bbbbbbbb', '20230916', 'https://desktop.github.com/images/desktop-icon.svg', '20230906',
        '2023-07-05 15:50:07', '2023-07-05 19:54:59');
INSERT INTO `today`
VALUES (66, 'bbbbbbbb', '20230917', 'https://desktop.github.com/images/desktop-icon.svg', '20230907',
        '2023-07-05 15:50:07', '2023-07-05 19:54:59');
INSERT INTO `today`
VALUES (77, 'bbbbbbbb', '20230908', 'https://desktop.github.com/images/desktop-icon.svg', '20230908',
        '2023-07-05 15:50:07', '2023-07-05 19:54:59');
INSERT INTO `today`
VALUES (90, 'bbbbbbbb', '20230909', 'https://desktop.github.com/images/desktop-icon.svg', '20230909',
        '2023-07-05 15:50:07', '2023-07-05 19:54:59');
INSERT INTO `today`
VALUES (91, 'bbbbbbbb', '20230910', 'https://desktop.github.com/images/desktop-icon.svg', '20230910',
        '2023-07-05 15:50:07', '2023-07-05 19:54:59');
INSERT INTO `today`
VALUES (92, 'bbbbbbbb', '20230911', 'https://desktop.github.com/images/desktop-icon.svg', '20230911',
        '2023-07-05 15:50:07', '2023-07-05 19:54:59');

SET FOREIGN_KEY_CHECKS = 1;


