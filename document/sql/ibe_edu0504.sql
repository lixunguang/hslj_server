/*
 Navicat Premium Data Transfer

 Source Server         : centos
 Source Server Type    : MySQL
 Source Server Version : 50742
 Source Host           : 172.30.100.23:3306
 Source Schema         : ibe_edu

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
-- Table structure for course
-- ----------------------------
DROP TABLE IF EXISTS `course`;
CREATE TABLE `course`
(
    `id`         int(11)                                                       NOT NULL AUTO_INCREMENT COMMENT '主键',
    `title`      varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '课程标题',
    `desc`       text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci         NULL COMMENT '课程描述',
    `deleted_at` timestamp(0)                                                  NULL COMMENT '软删',
    `created_at` timestamp(0)                                                  NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
    `updated_at` timestamp(0)                                                  NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE

) ENGINE = InnoDB
  AUTO_INCREMENT = 101
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT = '课程表'
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of course
-- ----------------------------
INSERT INTO `course`
VALUES (1, '发动机传动轴实训课',
        '发动机传动轴在汽车传动过程中最重要、载荷最大、价格最 昂贵的零件之一，在工作过程中，会受到活塞杆的载荷和工作转速引起的离心载荷。要保证传动轴能满足设计要求，必须进行仿真计算，该部分输入发动机传动轴的静强度计算。',
        DEFAULT, '2023-03-21 14:38:04', '2023-04-19 14:02:10');
INSERT INTO `course`
VALUES (2, '厂房除湿、除尘',
        '厂房除湿的必要性：1防腐蚀；2.防止产品结露、仓库结霜；3.防止产品受潮变质；4.为建筑除湿；5.防止发霉；6.为产品干燥冷却；7.防止电子产品性能变化；8.遏制细菌；9.减少工业臭味；10.防止潮湿天气可以减小对人体健康影响。\r\n厂房除尘的必要性：1.符合环保要求；2.有利于安全生产；3.可以减小对车间工人的身体伤害。',
        DEFAULT, '2023-03-21 14:39:38', '2023-04-19 14:02:07');



-- ----------------------------
-- Table structure for learning
-- ----------------------------
DROP TABLE IF EXISTS `learning`;
CREATE TABLE `learning`
(
    `id`               int(11)                                                       NOT NULL AUTO_INCREMENT COMMENT '主键',
    `title`            varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '学习资源标题',
    `desc`             text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci         NULL COMMENT '学习资源描述',
    `author`           varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '作者',
    `cover_picture_id` int(11)                                                       NOT NULL DEFAULT 0 COMMENT '展示图片',
    `category_id`      int(11)                                                       NOT NULL DEFAULT 0 COMMENT '学习资源的分类',
    `created_at`       timestamp(0)                                                  NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
    `updated_at`       timestamp(0)                                                  NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 16
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT = '学习资源表'
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of learning
-- ----------------------------
INSERT INTO `learning`
VALUES (1, 'CAE行业发展及自主仿真之路', '本课程由云道智造资深领域专家从宏观领域分别介绍了CAE概述、CAE发展历程、CAE技术演进、CAE应用、CAE发展趋势以及自主仿真软件不同的发展模式和发展方向。', '云道教育',
        400, 1, '2023-04-26 17:22:19', '2023-04-26 19:27:01');
INSERT INTO `learning`
VALUES (2, '云仿真平台SimCapsule平台简介',
        '本课程为有限元分析课程中的新技术应用部分。主要介绍了自主设计仿真平台Simcapsule的技术先进性和应用特点，阐述了将CAE仿真技术和云技术融合后，如何便捷的支持各行业垂直领域的快速仿真和成果展示。本课程适合设计仿真初学者学习。\r\nSimcapsule 是北京云道智造科技有限公司自主研发的基于云的工业软件开发与仿真平台，致力于通过云原生技术，为细分行业专用软件开发者提供SaaS的快速开发、在线运行和应用推广的一体化平台，最终为中小企业提供高效便捷的仿真服务。\r\n平台提供云桌面、文件管理器、应用商店、材料数据库等系统模块，支持模型导入、前处理设置、后处理等仿真功能，拥有自主高速web渲染引擎和自研界面库等技术优势，方便开发者快速搭建针对细分行业的web应用。\r\n平台采用微服务架构，接口开放，可灵活快速集成各种求解器开发在线应用，云道智造团队基于平台已开发完成CFD内流应用、CAD查看器、CAE后处理、网格剖分等通用在线应用。\r\n平台支持私有云和公有云一键部署，用户无需安装仿真软件，即可通过任何支持HTML5的浏览器跨终端随时随地访问平台、在线使用仿真应用，完成仿真全流程工作。',
        '云道教育', 401, 1, '2023-04-12 17:22:51', '2023-04-19 19:00:43');
INSERT INTO `learning`
VALUES (3, '电磁铁仿真案例', '本视频教程通过一个电磁铁受力仿真APP的开发实例演示Simdroid电磁分析建模、加载、设置材料属性、求解设置和仿真APP开发的基本操作。', '云道教育', 402, 1,
        '2023-04-12 17:31:31', '2023-04-19 19:00:46');



-- ----------------------------
-- Table structure for lesson
-- ----------------------------
DROP TABLE IF EXISTS `lesson`;
CREATE TABLE `lesson`
(
    `id`         int(11)                                                       NOT NULL AUTO_INCREMENT COMMENT '主键',
    `title`      varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '课堂标题',
    `desc`       text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci         NULL COMMENT '课堂描述',
    `deleted_at` timestamp(0)                                                  NULL COMMENT '软删',
    `created_at` timestamp(0)                                                  NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
    `updated_at` timestamp(0)                                                  NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 109
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT = '课程-教师关系表'
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of lesson
-- ----------------------------
INSERT INTO `lesson`
VALUES (1, '课程1-第1节课', '课程1-第1节课课程', DEFAULT, '2023-03-23 16:35:47', '2023-04-03 12:28:30');
INSERT INTO `lesson`
VALUES (2, '课程1-第2节课', '课程1-第2节课课程', DEFAULT, '2023-03-23 16:36:18', '2023-04-03 12:28:31');
INSERT INTO `lesson`
VALUES (3, '课程1-第3节课', '课程1-第3节课课程', DEFAULT, '2023-03-23 16:36:24', '2023-04-03 12:28:34');
INSERT INTO `lesson`
VALUES (4, '课程2-第1节课', '课程2-第1节课课程1-', DEFAULT, '2023-03-23 16:36:33', '2023-03-23 16:36:33');



-- ----------------------------
-- Table structure for news
-- ----------------------------
DROP TABLE IF EXISTS `news`;
CREATE TABLE `news`
(
    `id`         int(11)                                                       NOT NULL AUTO_INCREMENT COMMENT '主键',
    `title`      varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '新闻标题',
    `content`    text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci         NULL COMMENT '新闻内容',
    `picture_id` int(11)                                                       NULL     DEFAULT NULL,
    `publisher`  varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '发布人',
    `is_audit`   tinyint(4)                                                    NOT NULL DEFAULT 1 COMMENT '是否审核 1-否 2-是',
    `created_at` timestamp(0)                                                  NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
    `updated_at` timestamp(0)                                                  NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 103
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT = '新闻动态表'
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of news
-- ----------------------------
INSERT INTO `news`
VALUES (1, '新闻通知1', '111', 0, 'admin', 1, '2023-03-14 11:06:09', '2023-03-15 17:39:40');
INSERT INTO `news`
VALUES (2, '新闻通知2', 'ppppppppppppppppppppppppdiv', 0, 'admin', 1, '2023-03-14 11:06:53', '2023-03-15 17:39:41');
INSERT INTO `news`
VALUES (3, '新闻通知3', '<p style=\"text-align: center;\"><strong>第一条新闻</strong></p>', 8, 'admin', 1, '2023-03-14 11:12:18',
        '2023-03-30 15:35:45');
INSERT INTO `news`
VALUES (4, '新闻通知4',
        '<p style=\"text-align: center;\"><strong>第一条新闻</strong></p><p style=\"text-align: center;\"><span style=\"color: #e03e2d;\"><em><strong>09:34:03 2023-03-14</strong></em></span></p><p>&nbsp; &nbsp;新闻内容段落1</p><p style=\"padding-left: 40px;\"><span style=\"font-family: \'arial black\', sans-serif;\">情况1</span></p><p style=\"padding-left: 40px;\"><span style=\"font-family: \'arial black\', sans-serif;\">情况2</span></p><p style=\"padding-left: 40px;\"><span style=\"font-family: \'arial black\', sans-serif;\">情况3</span></p><p>&nbsp;<span style=\"text-decoration: line-through;\"> &nbsp;新闻内容段落2</span></p><table style=\"border-collapse: collapse; width: 100%;\" border=\"1\"><tbody><tr><td style=\"width: 23.2395%;\">1</td><td style=\"width: 23.2395%;\">2</td><td style=\"width: 23.2395%;\">&nbsp;</td><td style=\"width: 23.2395%;\">&nbsp;</td></tr><tr><td style=\"width: 23.2395%;\">&nbsp;</td><td style=\"width: 23.2395%;\">3</td><td style=\"width: 23.2395%;\">4</td><td style=\"width: 23.2395%;\">4</td></tr><tr><td style=\"width: 23.2395%;\">&nbsp;</td><td style=\"width: 23.2395%;\">5</td><td style=\"width: 23.2395%;\">5</td><td style=\"width: 23.2395%;\">5</td></tr></tbody></table><p><span style=\"text-decoration: underline;\">&nbsp; &nbsp;新闻内容段落3</span></p><p><span style=\"text-decoration: underline;\"><img src=\"http://172.30.100.23:8060/resource/image/1.png\" alt=\"存储在nginx里面的图片\" width=\"300\" height=\"200\" /></span></p><p>&nbsp;</p><p><span style=\"text-decoration: underline;\">图片2</span></p><p><a href=\"http://www.baidu.com\"><strong><u>百度网站</u></strong></a></p>',
        0, 'admin', 1, '2023-03-15 09:08:28', '2023-03-15 17:39:44');
INSERT INTO `news`
VALUES (5, '新闻通知4',
        '<p style=\"text-align: center;\"><strong>第2条新闻</strong></p><p style=\"text-align: center;\"><span style=\"color: #e03e2d;\"><em><strong>09:34:03 2023-03-14</strong></em></span></p><p>&nbsp; &nbsp;新闻内容段落1</p><p style=\"padding-left: 40px;\"><span style=\"font-family: \'arial black\', sans-serif;\">情况1</span></p><p style=\"padding-left: 40px;\"><span style=\"font-family: \'arial black\', sans-serif;\">情况2</span></p><p style=\"padding-left: 40px;\"><span style=\"font-family: \'arial black\', sans-serif;\">情况3</span></p><p>&nbsp;<span style=\"text-decoration: line-through;\"> &nbsp;新闻内容段落2</span></p><table style=\"border-collapse: collapse; width: 100%;\" border=\"1\"><tbody><tr><td style=\"width: 23.2395%;\">1</td><td style=\"width: 23.2395%;\">2</td><td style=\"width: 23.2395%;\">&nbsp;</td><td style=\"width: 23.2395%;\">&nbsp;</td></tr><tr><td style=\"width: 23.2395%;\">&nbsp;</td><td style=\"width: 23.2395%;\">3</td><td style=\"width: 23.2395%;\">4</td><td style=\"width: 23.2395%;\">4</td></tr><tr><td style=\"width: 23.2395%;\">&nbsp;</td><td style=\"width: 23.2395%;\">5</td><td style=\"width: 23.2395%;\">5</td><td style=\"width: 23.2395%;\">5</td></tr></tbody></table><p><span style=\"text-decoration: underline;\">&nbsp; &nbsp;新闻内容段落3</span></p><p><span style=\"text-decoration: underline;\"><img src=\"http://172.30.100.23:8060/resource/image/1.png\" alt=\"存储在nginx里面的图片\" width=\"300\" height=\"200\" /></span></p><p>&nbsp;</p><p><span style=\"text-decoration: underline;\">图片2</span></p><p><a href=\"http://www.baidu.com\"><strong><u>百度网站</u></strong></a></p>',
        0, 'admin', 1, '2023-03-15 09:09:34', '2023-03-15 17:39:45');



-- ----------------------------
-- Table structure for resource
-- ----------------------------
DROP TABLE IF EXISTS `resource`;
CREATE TABLE `resource`
(
    `id`         int(11)                                                       NOT NULL AUTO_INCREMENT COMMENT '主键',
    `title`      varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '资源标题',
    `desc`       text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci         NULL COMMENT '资源描述',
    `type`       tinyint(4)                                                    NOT NULL DEFAULT 0 COMMENT '资源类型，1图片，2 word文档，3 pdf文档，4 apps，5 ibe，6留言7 富文本',
    `content`    tinytext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci     NULL COMMENT '资源内容，存储路径、url及少量的文字',
    `deleted_at` timestamp(0)                                                  NULL COMMENT '软删',
    `created_at` timestamp(0)                                                  NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
    `updated_at` timestamp(0)                                                  NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 522
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT = '存储系统中用到的基础资源'
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of resource
-- ----------------------------
INSERT INTO `resource`
VALUES (1, '一个图片', '图片平铺', 1, 'resource/image/1.webp', DEFAULT, '2023-03-16 16:00:29', '2023-04-19 19:17:21');
INSERT INTO `resource`
VALUES (2, '一个文件', '文件', 1, 'resource/files/1.pdf', DEFAULT, '2023-03-16 16:02:38', '2023-04-19 19:17:18');
INSERT INTO `resource`
VALUES (3, '课堂作业', 'word文件', 2, 'resource/files/1.docx', DEFAULT, '2023-03-16 16:03:35', '2023-04-19 19:17:16');
INSERT INTO `resource`
VALUES (4, '一个apps', 'apps', 4, 'resource/ibe/1.app', DEFAULT, '2023-03-16 16:04:29', '2023-04-19 19:17:14');
INSERT INTO `resource`
VALUES (5, 'ibe文件', 'ibe', 5, 'resource/ibe/1.ibe', DEFAULT, '2023-03-16 16:19:17', '2023-04-19 19:17:12');
INSERT INTO `resource`
VALUES (6, '第二个图片', '图片', 1, 'resource/image/2.webp', DEFAULT, '2023-03-17 09:26:35', '2023-04-19 19:17:09');
INSERT INTO `resource`
VALUES (7, '第san个图片', '图片', 1, 'resource/image/3.webp', DEFAULT, '2023-03-17 11:25:51', '2023-04-19 19:17:07');
INSERT INTO `resource`
VALUES (8, '第si个图片', '图片', 1, 'resource/image/4.webp', DEFAULT, '2023-03-17 11:26:46', '2023-04-19 19:17:05');

-- ----------------------------
-- Table structure for session
-- ----------------------------
DROP TABLE IF EXISTS `session`;
CREATE TABLE `session`
(
    `id`         int(11)                                                      NOT NULL AUTO_INCREMENT COMMENT '主键',
    `user_id`    varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户标识id',
    `token`      text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci        NULL COMMENT 'token',
    `expires_at` timestamp(0)                                                 NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '过期时间，未用',
    `created_at` timestamp(0)                                                 NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
    `updated_at` timestamp(0)                                                 NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE

) ENGINE = InnoDB
  AUTO_INCREMENT = 101
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT = '登录用户session表'
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of session
-- ----------------------------




-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`
(
    `id`              int(11)                                                      NOT NULL AUTO_INCREMENT COMMENT '主键',
    `login_id`        varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '登录id，唯一',
    `name`            varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '姓名',
    `organization_id` int(11)                                                      NOT NULL DEFAULT 0 COMMENT '组织id，包括学校院系班级',
    `password`        varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '密码',
    `phone_number`    varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '电话号码',
    `creator_id`      int(11)                                                      NOT NULL DEFAULT 0 COMMENT '创建人id',
    `is_deleted`      tinyint(4)                                                   NOT NULL DEFAULT 1 COMMENT '是否删除 1-否 2-是',
    `created_at`      timestamp(0)                                                 NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
    `updated_at`      timestamp(0)                                                 NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
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
VALUES (1, '202311111111', 'user1', 1, '11111111', '0', 0, 1, '2023-03-09 16:53:08', '2023-04-17 10:29:32');
INSERT INTO `user`
VALUES (2, '202311111112', 'lixunguang', 1, 'uua', '0', 0, 1, '2023-03-09 16:54:12', '2023-03-09 16:54:12');
INSERT INTO `user`
VALUES (3, 'a1', 'admin', 1, 'admin', '0', 0, 1, '2023-03-09 16:55:09', '2023-03-09 16:55:09');
INSERT INTO `user`
VALUES (4, '202301002', 'admin', 1, 'admin', '0', 0, 1, '2023-03-21 15:46:25', '2023-04-17 10:29:24');
INSERT INTO `user`
VALUES (12, 't1', 'liteacher', 1, 'adminxx', '0', 0, 1, '2023-03-15 09:16:32', '2023-03-22 18:30:07');
INSERT INTO `user`
VALUES (18, '2023123456', 'user2', 1, '123456', '0', 0, 1, '2023-03-21 15:46:25', '2023-04-17 10:29:24');





SET FOREIGN_KEY_CHECKS = 1;
