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
-- Table structure for class
-- ----------------------------
DROP TABLE IF EXISTS `class`;
CREATE TABLE `class`
(
    `id`         int(11)                                                      NOT NULL AUTO_INCREMENT COMMENT '主键',
    `name`       varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '班级名字',
    `created_at` timestamp(0)                                                 NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
    `updated_at` timestamp(0)                                                 NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT = '班级表'
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of class
-- ----------------------------

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
INSERT INTO `course`
VALUES (3, '左心室辅助装置',
        '基于血流动力学仿真结果，可帮助医生在术前即可预测、分析、优化心血管手术辅助装置配置方案，较以往的经验判断更具科学性和可靠性，因而可大大提高手术的成功率，减轻病人负担。术前通过虚拟仿真规划手术方案，是目前临床和仿真领域关注的一大焦点，医学和数字仿真的结合对心血管疾病治疗具有重要意义',
        DEFAULT, '2023-03-21 14:43:24', '2023-04-19 14:01:47');
INSERT INTO `course`
VALUES (4, '绞缆机卷筒',
        '绞缆机作为船舶配套设备的主要设备之一，对船舶的安全营运有着重要的作用。其关键零部件绞缆机卷筒的理论承载能力与实际承载负荷之间存在较大的差异。对绞缆机卷筒缺乏精确计算，使得绞缆机卷筒的理论承载能力远大于实际所受的负载，造成卷筒自身重量过大，使得在生产过程中存在材料的浪费，增加了生产成本，造成了一定的经济损失。为了在制造成本、周期和市场快速响应等方面满足市场需求、提高产品的竞争力，需要对绞缆机进行深入系统的理论分析。\r\n绞缆机作为船舶配套设备的主要设备之一，对船舶的安全营运有着重要的作用。其关键零部件绞缆机卷筒的理论承载能力与实际承载负荷之间存在较大的差异。对绞缆机卷筒缺乏精确计算，使得绞缆机卷筒的理论承载能力远大于实际所受的负载，造成卷筒自身重量过大，使得在生产过程中存在材料的浪费，增加了生产成本，造成了一定的经济损失。为了在制造成本、周期和市场快速响应等方面满足市场需求、提高产品的竞争力，需要对绞缆机进行深入系统的理论分析。\r\n绞缆机作为船舶配套设备的主要设备之一，对船舶的安全营运有着重要的作用。其关键零部件绞缆机卷筒的理论承载能力与实际承载负荷之间存在较大的差异。对绞缆机卷筒缺乏精确计算，使得绞缆机卷筒的理论承载能力远大于实际所受的负载，造成卷筒自身重量过大，使得在生产过程中存在材料的浪费，增加了生产成本，造成了一定的经济损失。为了在制造成本、周期和市场快速响应等方面满足市场需求、提高产品的竞争力，需要对绞缆机进行深入系统的理论分析。',
        DEFAULT, '2023-03-21 15:17:45', '2023-04-19 14:01:42');
INSERT INTO `course`
VALUES (5, '电子机柜三维建模与散热仿真分析',
        '在电力系统和通讯系统中，户外电子机柜设备应用广泛，而且由于处在户外的原因，通常为非全开放式机柜。然而在户外无遮挡情况下，阳光辐射以及设备本身耗散的热量作用，使得密闭机柜内部温度可能超出设备允许的范围，装置长时间在超负荷高温下运行，会引起元件性能的降低，进而导致装置故障，影响整个系统的稳定性。因此，在密闭的户外机柜中如何控制内部温度，成为户外机柜设计的关键。',
        DEFAULT, '2023-03-29 14:06:17', '2023-03-29 14:06:17');
INSERT INTO `course`
VALUES (6, 'Simetherm芯片级 IGBT模块仿真计算方案',
        'IGBT(Insulated Gate Bipolar Transistor)，绝缘栅双极型晶体管，是由BJT(双极结型晶体三极管)和MOS(绝缘栅型场效应管)组成的复合全控型-电压驱动式-功率半导体器件，作为能源转换与传输的核心器件，是电力电子装置的“CPU”，广泛应用于小体积、低噪声、高性能的电子设备中，如逆变器、不间断电源、逆变电焊机以及电机调速的变频器中。近年来，随着电子设备及系统的正向小型化和多功能化发展，IGBT模块也正在向小尺寸、大功率的方向发展，随之而来的是模块内所产生的高热流密度带来的散热问题，因此散热系统的优化设计则成为保障IGBT模块可靠运行的关键环节。',
        DEFAULT, '2023-03-29 14:07:16', '2023-03-29 14:07:16');
INSERT INTO `course`
VALUES (7, 'Simetherm电子设备机箱散热分析方案',
        '随着工业生产需求的高标准化，电子设备都朝着小体积大功率多功能方向发展，随之而来的是高热流密度带来的散热问题，为了有效避免电子设备机箱内温度过高，影响电子器件正常工作，需要在结构设计时对散热做必要的考虑。传统方法是根据指标要求和工程经验设计出样品，加工完成后用环境试验检验，然后根据反馈的问题进行改进，从而得到最优化设计，其研制周期和成本都普遍较高。',
        DEFAULT, '2023-03-29 14:09:07', '2023-03-29 14:09:07');
INSERT INTO `course`
VALUES (100, '塔身吊装实训课',
        '在电力建设施工过程中，抱杆是不可或缺的起重装置。它的主要功能是依托在建铁塔，将塔材提升至需要的位置进行施工安装，因此广泛应用于电力建设施工尤其是吊装施工并且是主要的受力单元。在施工过程中抱杆受力是否均衡合理，是否有超载使用，直接影响着施工的质量和安全。在传统的抱杆受力计算中，一般采用简化的方式，利用三角函数进行计算。然而这种简化方式只能粗略计算简单受力工况，当采用反向拉线等复杂工况或现场地形复杂时，不易分析各个结构的受力结果。对于现场施工人员，抱杆的力学计算所需的专业知识要求较高，繁琐并且无法保证可靠性，因此在施工中主要依靠工作经验来控制抱杆受力。',
        DEFAULT, '2023-04-17 14:13:04', '2023-04-25 14:13:04');

-- ----------------------------
-- Table structure for course_resource
-- ----------------------------
DROP TABLE IF EXISTS `course_resource`;
CREATE TABLE `course_resource`
(
    `id`          int(11)      NOT NULL AUTO_INCREMENT COMMENT '主键',
    `course_id`   int(11)      NOT NULL DEFAULT 0 COMMENT '课程id',
    `resource_id` int(11)      NOT NULL DEFAULT 0 COMMENT '资源id',
    `type`        tinyint(4)   NOT NULL DEFAULT 0 COMMENT '资源用途，1展示图片，2 引导视频，3 课堂id',
    `index`       tinyint(4)   NULL     DEFAULT NULL,
    `deleted_at`  timestamp(0) NULL COMMENT '软删',
    `created_at`  timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
    `updated_at`  timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 163
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT = '课程-资源关系表'
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of course_resource
-- ----------------------------
INSERT INTO `course_resource`
VALUES (30, 5, 11, 2, 0, DEFAULT, '2023-03-22 11:06:23', '2023-03-22 11:06:23');
INSERT INTO `course_resource`
VALUES (31, 5, 1, 3, 1, DEFAULT, '2023-03-22 11:12:49', '2023-03-23 16:01:40');
INSERT INTO `course_resource`
VALUES (32, 5, 2, 3, 2, DEFAULT, '2023-03-22 11:13:16', '2023-03-23 16:01:41');
INSERT INTO `course_resource`
VALUES (33, 5, 3, 3, 3, DEFAULT, '2023-03-22 11:13:22', '2023-03-23 16:01:43');
INSERT INTO `course_resource`
VALUES (34, 5, 10, 2, 0, DEFAULT, '2023-03-23 17:18:48', '2023-03-23 17:18:48');
INSERT INTO `course_resource`
VALUES (35, 5, 11, 2, 0, DEFAULT, '2023-03-30 17:45:23', '2023-04-06 16:56:50');
INSERT INTO `course_resource`
VALUES (37, 6, 11, 2, 0, DEFAULT, '2023-03-22 11:06:23', '2023-03-22 11:06:23');
INSERT INTO `course_resource`
VALUES (38, 6, 1, 3, 1, DEFAULT, '2023-03-22 11:12:49', '2023-03-23 16:01:40');
INSERT INTO `course_resource`
VALUES (39, 6, 2, 3, 2, DEFAULT, '2023-03-22 11:13:16', '2023-03-23 16:01:41');
INSERT INTO `course_resource`
VALUES (40, 6, 3, 3, 3, DEFAULT, '2023-03-22 11:13:22', '2023-03-23 16:01:43');
INSERT INTO `course_resource`
VALUES (41, 6, 10, 2, 0, DEFAULT, '2023-03-23 17:18:48', '2023-03-23 17:18:48');
INSERT INTO `course_resource`
VALUES (42, 6, 11, 2, 0, DEFAULT, '2023-03-30 17:45:23', '2023-04-06 16:56:50');
INSERT INTO `course_resource`
VALUES (44, 7, 11, 2, 0, DEFAULT, '2023-03-22 11:06:23', '2023-03-22 11:06:23');
INSERT INTO `course_resource`
VALUES (45, 7, 1, 3, 1, DEFAULT, '2023-03-22 11:12:49', '2023-03-23 16:01:40');
INSERT INTO `course_resource`
VALUES (46, 7, 2, 3, 2, DEFAULT, '2023-03-22 11:13:16', '2023-03-23 16:01:41');
INSERT INTO `course_resource`
VALUES (47, 7, 3, 3, 3, DEFAULT, '2023-03-22 11:13:22', '2023-03-23 16:01:43');
INSERT INTO `course_resource`
VALUES (48, 7, 10, 2, 0, DEFAULT, '2023-03-23 17:18:48', '2023-03-23 17:18:48');
INSERT INTO `course_resource`
VALUES (49, 7, 11, 2, 0, DEFAULT, '2023-03-30 17:45:23', '2023-04-06 16:56:50');
INSERT INTO `course_resource`
VALUES (51, 8, 9, 1, 0, DEFAULT, '2023-03-22 10:56:59', '2023-03-22 11:16:47');
INSERT INTO `course_resource`
VALUES (52, 8, 11, 2, 0, DEFAULT, '2023-03-22 11:06:23', '2023-03-22 11:06:23');
INSERT INTO `course_resource`
VALUES (53, 8, 1, 3, 1, DEFAULT, '2023-03-22 11:12:49', '2023-03-23 16:01:40');
INSERT INTO `course_resource`
VALUES (54, 8, 2, 3, 2, DEFAULT, '2023-03-22 11:13:16', '2023-03-23 16:01:41');
INSERT INTO `course_resource`
VALUES (55, 8, 3, 3, 3, DEFAULT, '2023-03-22 11:13:22', '2023-03-23 16:01:43');
INSERT INTO `course_resource`
VALUES (56, 8, 10, 2, 0, DEFAULT, '2023-03-23 17:18:48', '2023-03-23 17:18:48');
INSERT INTO `course_resource`
VALUES (57, 8, 11, 2, 0, DEFAULT, '2023-03-30 17:45:23', '2023-04-06 16:56:50');
INSERT INTO `course_resource`
VALUES (61, 9, 9, 1, 0, DEFAULT, '2023-03-22 10:56:59', '2023-03-22 11:16:47');
INSERT INTO `course_resource`
VALUES (62, 9, 11, 2, 0, DEFAULT, '2023-03-22 11:06:23', '2023-03-22 11:06:23');
INSERT INTO `course_resource`
VALUES (63, 9, 1, 3, 1, DEFAULT, '2023-03-22 11:12:49', '2023-03-23 16:01:40');
INSERT INTO `course_resource`
VALUES (64, 9, 2, 3, 2, DEFAULT, '2023-03-22 11:13:16', '2023-03-23 16:01:41');
INSERT INTO `course_resource`
VALUES (65, 9, 3, 3, 3, DEFAULT, '2023-03-22 11:13:22', '2023-03-23 16:01:43');
INSERT INTO `course_resource`
VALUES (66, 9, 10, 2, 0, DEFAULT, '2023-03-23 17:18:48', '2023-03-23 17:18:48');
INSERT INTO `course_resource`
VALUES (67, 9, 11, 2, 0, DEFAULT, '2023-03-30 17:45:23', '2023-04-06 16:56:50');
INSERT INTO `course_resource`
VALUES (71, 10, 9, 1, 0, DEFAULT, '2023-03-22 10:56:59', '2023-03-22 11:16:47');
INSERT INTO `course_resource`
VALUES (72, 10, 11, 2, 0, DEFAULT, '2023-03-22 11:06:23', '2023-03-22 11:06:23');
INSERT INTO `course_resource`
VALUES (73, 10, 1, 3, 1, DEFAULT, '2023-03-22 11:12:49', '2023-03-23 16:01:40');
INSERT INTO `course_resource`
VALUES (74, 10, 2, 3, 2, DEFAULT, '2023-03-22 11:13:16', '2023-03-23 16:01:41');
INSERT INTO `course_resource`
VALUES (75, 10, 3, 3, 3, DEFAULT, '2023-03-22 11:13:22', '2023-03-23 16:01:43');
INSERT INTO `course_resource`
VALUES (76, 10, 10, 2, 0, DEFAULT, '2023-03-23 17:18:48', '2023-03-23 17:18:48');
INSERT INTO `course_resource`
VALUES (77, 10, 11, 2, 0, DEFAULT, '2023-03-30 17:45:23', '2023-04-06 16:56:50');
INSERT INTO `course_resource`
VALUES (100, 100, 100, 1, 0, DEFAULT, '2023-03-30 17:45:23', '2023-04-06 16:56:50');
INSERT INTO `course_resource`
VALUES (101, 100, 101, 3, 1, DEFAULT, '2023-03-30 17:45:23', '2023-04-06 16:56:50');
INSERT INTO `course_resource`
VALUES (102, 100, 102, 3, 2, DEFAULT, '2023-03-30 17:45:23', '2023-04-06 16:56:50');
INSERT INTO `course_resource`
VALUES (103, 100, 103, 3, 3, DEFAULT, '2023-03-30 17:45:23', '2023-04-06 16:56:50');
INSERT INTO `course_resource`
VALUES (104, 100, 104, 3, 4, DEFAULT, '2023-03-30 17:45:23', '2023-04-06 16:56:50');
INSERT INTO `course_resource`
VALUES (105, 100, 105, 3, 5, DEFAULT, '2023-03-30 17:45:23', '2023-04-06 16:56:50');
INSERT INTO `course_resource`
VALUES (106, 100, 106, 3, 6, DEFAULT, '2023-03-30 17:45:23', '2023-04-06 16:56:50');
INSERT INTO `course_resource`
VALUES (107, 100, 107, 3, 7, DEFAULT, '2023-03-30 17:45:23', '2023-04-06 16:56:50');
INSERT INTO `course_resource`
VALUES (108, 100, 108, 3, 8, DEFAULT, '2023-03-30 17:45:23', '2023-04-06 16:56:50');
INSERT INTO `course_resource`
VALUES (109, 100, 118, 2, 0, DEFAULT, '2023-03-30 17:45:23', '2023-04-06 16:56:50');
INSERT INTO `course_resource`
VALUES (120, 1, 202, 1, 0, DEFAULT, '2023-03-30 17:45:23', '2023-04-06 16:56:50');
INSERT INTO `course_resource`
VALUES (130, 2, 201, 1, 0, DEFAULT, '2023-03-30 17:45:23', '2023-04-06 16:56:50');
INSERT INTO `course_resource`
VALUES (140, 3, 200, 1, 0, DEFAULT, '2023-03-30 17:45:23', '2023-04-06 16:56:50');
INSERT INTO `course_resource`
VALUES (150, 4, 203, 1, 0, DEFAULT, '2023-03-30 17:45:23', '2023-04-06 16:56:50');
INSERT INTO `course_resource`
VALUES (160, 5, 315, 1, 0, DEFAULT, '2023-03-30 17:45:23', '2023-04-06 16:56:50');
INSERT INTO `course_resource`
VALUES (161, 6, 316, 1, 0, DEFAULT, '2023-03-30 17:45:23', '2023-04-06 16:56:50');
INSERT INTO `course_resource`
VALUES (162, 7, 317, 1, 0, DEFAULT, '2023-03-30 17:45:23', '2023-04-06 16:56:50');

-- ----------------------------
-- Table structure for course_student
-- ----------------------------
DROP TABLE IF EXISTS `course_student`;
CREATE TABLE `course_student`
(
    `id`         int(11)      NOT NULL AUTO_INCREMENT COMMENT '主键',
    `student_id` int(11)      NOT NULL DEFAULT 0 COMMENT '学生id',
    `course_id`  int(11)      NOT NULL DEFAULT 0 COMMENT '课程id',
    `created_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
    `updated_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT = '学生-课程关系表'
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of course_student
-- ----------------------------

-- ----------------------------
-- Table structure for course_teacher
-- ----------------------------
DROP TABLE IF EXISTS `course_teacher`;
CREATE TABLE `course_teacher`
(
    `id`         int(11)      NOT NULL AUTO_INCREMENT COMMENT '主键',
    `course_id`  int(11)      NOT NULL DEFAULT 0 COMMENT '课程id',
    `teacher_id` int(11)      NOT NULL DEFAULT 0 COMMENT '教师id',
    `created_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
    `updated_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 101
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT = '课程-教师关系表'
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of course_teacher
-- ----------------------------
INSERT INTO `course_teacher`
VALUES (1, 1, 12, '2023-03-22 17:33:21', '2023-03-22 17:33:21');
INSERT INTO `course_teacher`
VALUES (2, 2, 12, '2023-03-22 18:32:43', '2023-03-22 18:32:43');
INSERT INTO `course_teacher`
VALUES (3, 3, 12, '2023-03-22 18:33:08', '2023-03-22 18:33:08');
INSERT INTO `course_teacher`
VALUES (4, 4, 12, '2023-03-23 15:17:01', '2023-03-23 15:17:01');
INSERT INTO `course_teacher`
VALUES (5, 5, 12, '2023-03-23 15:17:07', '2023-03-23 15:17:07');
INSERT INTO `course_teacher`
VALUES (6, 6, 12, '2023-03-23 15:17:07', '2023-03-23 15:17:07');
INSERT INTO `course_teacher`
VALUES (7, 7, 12, '2023-03-23 15:17:07', '2023-03-23 15:17:07');
INSERT INTO `course_teacher`
VALUES (100, 100, 12, '2023-03-23 15:17:07', '2023-03-23 15:17:07');

-- ----------------------------
-- Table structure for department
-- ----------------------------
DROP TABLE IF EXISTS `department`;
CREATE TABLE `department`
(
    `id`         int(11)                                                      NOT NULL AUTO_INCREMENT COMMENT '主键',
    `name`       varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '院系名字',
    `created_at` timestamp(0)                                                 NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
    `updated_at` timestamp(0)                                                 NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT = '院系表'
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of department
-- ----------------------------

-- ----------------------------
-- Table structure for favorite
-- ----------------------------
DROP TABLE IF EXISTS `favorite`;
CREATE TABLE `favorite`
(
    `id`          int(11)      NOT NULL AUTO_INCREMENT COMMENT '主键',
    `user_id`     int(11)      NOT NULL DEFAULT 0 COMMENT '所属用户id',
    `resource_id` int(11)      NOT NULL DEFAULT 0 COMMENT '资源id',
    `usage`       tinyint(4)   NOT NULL DEFAULT 0 COMMENT '资源所属类型，如资源，课程。',
    `created_at`  timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
    `updated_at`  timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT = '收藏表'
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of favorite
-- ----------------------------

-- ----------------------------
-- Table structure for file
-- ----------------------------
DROP TABLE IF EXISTS `file`;
CREATE TABLE `file`
(
    `id`            int(11)                                                        NOT NULL AUTO_INCREMENT COMMENT '主键',
    `original_name` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '原始名字',
    `storage_name`  varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '存储名字',
    `created_at`    timestamp(0)                                                   NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
    `updated_at`    timestamp(0)                                                   NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 3
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT = '文件表，存储上传文件原始名字及存储路径'
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of file
-- ----------------------------


-- ----------------------------
-- Table structure for history
-- ----------------------------
DROP TABLE IF EXISTS `history`;
CREATE TABLE `history`
(
    `id`            int(11)      NOT NULL AUTO_INCREMENT COMMENT '主键',
    `user_id`       int(11)      NOT NULL DEFAULT 0 COMMENT '所属用户id',
    `resource_id`   int(11)      NOT NULL DEFAULT 0 COMMENT '资源id',
    `resource_type` tinyint(4)   NOT NULL DEFAULT 0,
    `usage`         tinyint(4)   NOT NULL DEFAULT 0 COMMENT '资源所属类型，如资源，课程。',
    `created_at`    timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
    `updated_at`    timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT = '历史记录表'
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of history
-- ----------------------------

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
INSERT INTO `learning`
VALUES (4, '有限单元法与Simdroid介绍',
        '以结构有限元仿真为例，介绍有限单元法仿真，并对仿真平台Simdroid进行系统介绍。通过实际操作的形式，从Simdroid软件的安装到仿真建模、有限元网格以及边界与载荷，为初学者介绍Simdroid仿真的完整流程。最后，以齿轮的有限元分析为例，演示讲解建模、材料属性、网格剖分、边界载荷、计算和后处理等仿真环节。',
        '云道教育', 403, 2, '2023-04-25 17:31:31', '2023-04-25 19:27:06');
INSERT INTO `learning`
VALUES (5, '数字化仿真赋能高校人才培养',
        '当前，数字化转型正在推动未来社会全面变革。数字教育已成为推动世界教育变革的重要力量。其中，仿真的身影活跃在越来越多的课堂上。\r\n针对院校教学中面临的“高投入、高耗材、高危险、难实施、难观摩、难再现”等问题，仿真集数字资源、信息技术与学科实践为一体，在实训教学中发挥着愈加重要的作用。\r\n云道智造深耕仿真领域多年，让自主仿真软件进课堂，支撑数字仿真人才培养，是实现“普惠仿真”的重要方式。\r\n基于深厚行业经验和丰富合作案例，云道智造提出了数字仿真人才培养的“云道方案”，以数字仿真技术赋能高校人才培养。\r\n那么，当数字仿真遇上教育会碰撞出怎样的火花？\r\n“云道方案”如何破解数字仿真技术人才培养密码？\r\n云道智造教育解决方案专家郭隽主讲，和您聊一聊仿真的来龙去脉、仿真产业应用案例、数字仿真人才培养“云道方案”等相关的内容。',
        '云道教育', 404, 3, '2023-04-27 18:53:20', '2023-04-27 19:02:31');
INSERT INTO `learning`
VALUES (6, 'Simdroid足球侧向力影响因素分及APP制作教程',
        '足球比赛中，当你听到解说员说到“香蕉球”、“落叶球”、“电梯球”时，你能说出他们的区别吗？更关键的是，怎样才能踢出这些匪夷所思的球呢？本APP通过二维流场计算，可以分析球的前进速度，旋转速度等因素对足球侧向力的影响，从而分析出“香蕉球”、“落叶球”、“电梯球”运行轨迹的关键因素。',
        '云道教育', 405, 4, '2023-04-19 19:04:52', '2023-04-17 19:20:30');
INSERT INTO `learning`
VALUES (7, '电磁、结构、流体工程实践及算例开发培训课程',
        '“基于Simdroid的CAE仿真师资培训”是面向2021年6月批立项北京云道智造科技有限公司教育部产学合作协同育人项目的高校主持及参研教师举办的“云道智造2021年师资培训”，目的是培训合作院校教师掌握工业仿真软件Simdroid的使用。同时培训对象可扩展到有兴趣了解工业仿真软件的其他院校教师，可扩展到有一定CAE仿真基础的博士、硕士。',
        '云道教育', 406, 4, '2023-04-19 19:04:52', '2023-04-19 19:20:30');
INSERT INTO `learning`
VALUES (8, '仿真APP开发实训-流体仿真', '本课程系统讲解流体技术的原理，流体仿真技术在工程领域的应用。并在此基础上应用Simdroid平台，指导同学进行流体仿真APP案例的设计与开发。', '云道教育', 407, 4,
        '2023-04-19 19:04:52', '2023-04-19 19:20:30');

-- ----------------------------
-- Table structure for learning_category
-- ----------------------------
DROP TABLE IF EXISTS `learning_category`;
CREATE TABLE `learning_category`
(
    `id`         int(11)                                                       NOT NULL AUTO_INCREMENT COMMENT '主键',
    `title`      varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT 'learning_category标题',
    `desc`       text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci         NULL COMMENT 'learning_category描述,可以为空',
    `created_at` timestamp(0)                                                  NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
    `updated_at` timestamp(0)                                                  NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 14
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT = '学习栏目分类表'
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of learning_category
-- ----------------------------
INSERT INTO `learning_category`
VALUES (1, '基础视频', '基础视频基础视频基础视频', '2023-04-12 17:18:22', '2023-04-13 16:02:21');
INSERT INTO `learning_category`
VALUES (2, '提高视频', '基础视频基础视频基础视频', '2023-04-12 17:19:29', '2023-04-13 16:02:24');
INSERT INTO `learning_category`
VALUES (3, '进阶视频', '进阶视频进阶视频进阶视频', '2023-04-12 17:19:37', '2023-04-13 16:02:30');
INSERT INTO `learning_category`
VALUES (4, '高级视频', '高级视频高级视频高级视频', '2023-04-12 17:19:37', '2023-04-13 16:02:30');

-- ----------------------------
-- Table structure for learning_resource
-- ----------------------------
DROP TABLE IF EXISTS `learning_resource`;
CREATE TABLE `learning_resource`
(
    `id`          int(11)                                                       NOT NULL AUTO_INCREMENT COMMENT '主键',
    `title`       varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '学习资源标题',
    `desc`        text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci         NULL COMMENT '学习资源描述',
    `resource_id` int(11)                                                       NOT NULL DEFAULT 0 COMMENT '资源id，一般为视频',
    `learning_id` int(11)                                                       NOT NULL DEFAULT 0 COMMENT '所属learning id',
    `index`       int(11)                                                       NOT NULL DEFAULT 0 COMMENT '资源排序',
    `created_at`  timestamp(0)                                                  NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
    `updated_at`  timestamp(0)                                                  NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 108
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT = '学习资源条目表'
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of learning_resource
-- ----------------------------
INSERT INTO `learning_resource`
VALUES (100, 'CAE行业发展及自主仿真之路', 'CAE行业发展及自主仿真之路', 506, 1, 0, '2023-04-19 19:22:58', '2023-04-19 19:22:58');
INSERT INTO `learning_resource`
VALUES (101, '云仿真平台SimCapsule平台简介', '云仿真平台SimCapsule平台简介', 502, 2, 0, '2023-04-19 19:22:58', '2023-04-19 19:22:58');
INSERT INTO `learning_resource`
VALUES (102, '电磁铁仿真案例', '电磁铁仿真案例', 500, 3, 0, '2023-04-19 19:22:58', '2023-04-19 19:22:58');
INSERT INTO `learning_resource`
VALUES (103, '有限单元法与Simdroid介绍', '有限单元法与Simdroid介绍', 501, 4, 0, '2023-04-19 19:22:58', '2023-04-19 19:22:58');
INSERT INTO `learning_resource`
VALUES (104, '数字化仿真赋能高校人才培养', '数字化仿真赋能高校人才培养', 505, 5, 0, '2023-04-19 19:22:58', '2023-04-19 19:22:58');
INSERT INTO `learning_resource`
VALUES (105, '足球侧向力影响因素分析', '足球侧向力影响因素分析', 504, 6, 0, '2023-04-19 19:22:58', '2023-04-19 19:22:58');
INSERT INTO `learning_resource`
VALUES (106, '电磁、结构、流体工程实践及算例开发培训课程', '电磁、结构、流体工程实践及算例开发培训课程', 507, 7, 0, '2023-04-19 19:22:58', '2023-04-19 19:22:58');
INSERT INTO `learning_resource`
VALUES (107, '课程8-仿真APP开发实训-流体仿真', '课程8-仿真APP开发实训-流体仿真', 503, 8, 0, '2023-04-19 19:22:58', '2023-04-19 19:22:58');

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
INSERT INTO `lesson`
VALUES (101, '实验一：输电线路铁塔组立抱杆工程简介及仿真APP体验', '课程章节', DEFAULT, '2023-03-23 16:36:33', '2023-03-23 16:36:33');
INSERT INTO `lesson`
VALUES (102, '实验二：Simdroid软件梁壳基本操作学习实训指导书', '课程章节', DEFAULT, '2023-03-23 16:36:33', '2023-03-23 16:36:33');
INSERT INTO `lesson`
VALUES (103, '实验三：塔身吊装-参数化建模实训指导书', '课程章节', DEFAULT, '2023-03-23 16:36:33', '2023-03-23 16:36:33');
INSERT INTO `lesson`
VALUES (104, '实验四：塔身吊装-网格剖分实训指导书模板', '课程章节', DEFAULT, '2023-03-23 16:36:33', '2023-03-23 16:36:33');
INSERT INTO `lesson`
VALUES (105, '实验五：塔身吊装-边界与求解设置实训指导书', '课程章节', DEFAULT, '2023-03-23 16:36:33', '2023-03-23 16:36:33');
INSERT INTO `lesson`
VALUES (106, '实验六：塔身吊装-结果后处理实训指导书', '课程章节', DEFAULT, '2023-03-23 16:36:33', '2023-03-23 16:36:33');
INSERT INTO `lesson`
VALUES (107, '实验七：塔身吊装-APP封装实训指导书', '课程章节', DEFAULT, '2023-03-23 16:36:33', '2023-03-23 16:36:33');
INSERT INTO `lesson`
VALUES (108, '实验八：塔身吊装- APP实操', '课程章节', DEFAULT, '2023-03-23 16:36:33', '2023-03-23 16:36:33');

-- ----------------------------
-- Table structure for lesson_section
-- ----------------------------
DROP TABLE IF EXISTS `lesson_section`;
CREATE TABLE `lesson_section`
(
    `id`         int(11)                                                       NOT NULL AUTO_INCREMENT COMMENT '主键',
    `lesson_id`  int(11)                                                       NOT NULL DEFAULT 0 COMMENT '课堂id',
    `title`      varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '课堂子部分标题',
    `index`      tinyint(4)                                                    NOT NULL,
    `created_at` timestamp(0)                                                  NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
    `updated_at` timestamp(0)                                                  NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 704
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT = '课堂的各个组成部分标题'
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of lesson_section
-- ----------------------------
INSERT INTO `lesson_section`
VALUES (1, 1, '教学内容', 1, '2023-03-31 15:44:50', '2023-03-31 15:47:29');
INSERT INTO `lesson_section`
VALUES (2, 1, '参考资源', 2, '2023-03-31 15:45:48', '2023-03-31 15:47:58');
INSERT INTO `lesson_section`
VALUES (3, 1, '做实验', 3, '2023-03-31 15:46:34', '2023-03-31 15:49:28');
INSERT INTO `lesson_section`
VALUES (4, 1, '作业', 4, '2023-03-31 15:46:37', '2023-03-31 15:49:31');
INSERT INTO `lesson_section`
VALUES (5, 2, '教学内容', 1, '2023-03-31 15:46:39', '2023-03-31 15:49:39');
INSERT INTO `lesson_section`
VALUES (6, 2, '参考资源', 2, '2023-03-31 15:46:39', '2023-03-31 15:47:44');
INSERT INTO `lesson_section`
VALUES (7, 2, '做实验', 3, '2023-03-31 15:46:40', '2023-03-31 15:49:43');
INSERT INTO `lesson_section`
VALUES (8, 3, '教学内容', 1, '2023-03-31 15:46:40', '2023-03-31 16:22:29');
INSERT INTO `lesson_section`
VALUES (9, 4, 'XX', 1, '2023-04-03 13:01:34', '2023-04-04 10:46:55');
INSERT INTO `lesson_section`
VALUES (10, 4, '作业', 2, '2023-04-03 13:02:07', '2023-04-04 10:46:53');
INSERT INTO `lesson_section`
VALUES (200, 101, '教学内容', 1, '2023-03-31 15:44:50', '2023-03-31 15:47:29');
INSERT INTO `lesson_section`
VALUES (201, 101, '参考资源', 2, '2023-03-31 15:44:50', '2023-03-31 15:47:29');
INSERT INTO `lesson_section`
VALUES (202, 101, '做实验', 3, '2023-03-31 15:44:50', '2023-03-31 15:47:29');
INSERT INTO `lesson_section`
VALUES (203, 101, '作业', 4, '2023-03-31 15:44:50', '2023-03-31 15:47:29');
INSERT INTO `lesson_section`
VALUES (300, 102, '教学内容', 1, '2023-03-31 15:44:50', '2023-03-31 15:47:29');
INSERT INTO `lesson_section`
VALUES (301, 102, '参考资源', 2, '2023-03-31 15:44:50', '2023-03-31 15:47:29');
INSERT INTO `lesson_section`
VALUES (302, 102, '做实验', 3, '2023-03-31 15:44:50', '2023-03-31 15:47:29');
INSERT INTO `lesson_section`
VALUES (303, 102, '作业', 4, '2023-03-31 15:44:50', '2023-03-31 15:47:29');
INSERT INTO `lesson_section`
VALUES (400, 103, '教学内容', 1, '2023-03-31 15:44:50', '2023-03-31 15:47:29');
INSERT INTO `lesson_section`
VALUES (401, 103, '参考资源', 2, '2023-03-31 15:44:50', '2023-03-31 15:47:29');
INSERT INTO `lesson_section`
VALUES (402, 103, '做实验', 3, '2023-03-31 15:44:50', '2023-03-31 15:47:29');
INSERT INTO `lesson_section`
VALUES (403, 103, '作业', 4, '2023-03-31 15:44:50', '2023-03-31 15:47:29');
INSERT INTO `lesson_section`
VALUES (500, 104, '教学内容', 1, '2023-03-31 15:44:50', '2023-03-31 15:47:29');
INSERT INTO `lesson_section`
VALUES (501, 104, '参考资源', 2, '2023-03-31 15:44:50', '2023-03-31 15:47:29');
INSERT INTO `lesson_section`
VALUES (502, 104, '做实验', 3, '2023-03-31 15:44:50', '2023-03-31 15:47:29');
INSERT INTO `lesson_section`
VALUES (503, 104, '作业', 4, '2023-03-31 15:44:50', '2023-03-31 15:47:29');
INSERT INTO `lesson_section`
VALUES (600, 105, '教学内容', 1, '2023-03-31 15:44:50', '2023-03-31 15:47:29');
INSERT INTO `lesson_section`
VALUES (601, 105, '参考资源', 2, '2023-03-31 15:44:50', '2023-03-31 15:47:29');
INSERT INTO `lesson_section`
VALUES (602, 105, '做实验', 3, '2023-03-31 15:44:50', '2023-03-31 15:47:29');
INSERT INTO `lesson_section`
VALUES (603, 105, '作业', 4, '2023-03-31 15:44:50', '2023-03-31 15:47:29');
INSERT INTO `lesson_section`
VALUES (700, 106, '教学内容', 1, '2023-03-31 15:44:50', '2023-03-31 15:47:29');
INSERT INTO `lesson_section`
VALUES (701, 106, '参考资源', 2, '2023-03-31 15:44:50', '2023-03-31 15:47:29');
INSERT INTO `lesson_section`
VALUES (702, 106, '做实验', 3, '2023-03-31 15:44:50', '2023-03-31 15:47:29');
INSERT INTO `lesson_section`
VALUES (703, 106, '作业', 4, '2023-03-31 15:44:50', '2023-03-31 15:47:29');

-- ----------------------------
-- Table structure for lesson_section_resource
-- ----------------------------
DROP TABLE IF EXISTS `lesson_section_resource`;
CREATE TABLE `lesson_section_resource`
(
    `id`                int(11)      NOT NULL AUTO_INCREMENT COMMENT '主键',
    `lesson_section_id` int(11)      NOT NULL DEFAULT 0 COMMENT 'lesson_section_id id',
    `resource_id`       int(11)      NOT NULL DEFAULT 0 COMMENT '资源id',
    `index`             tinyint(4)   NOT NULL DEFAULT 0 COMMENT '排序.对于作业来说,类型含义如下：
        41作业要求：富文本格式要求;	 3X作业要求：作业文件;	43作业提交的文件;	44作业提交附言',
    `deleted_at`        timestamp(0) NULL COMMENT '软删',
    `created_at`        timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
    `updated_at`        timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 152
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT = '课堂部分-资源关系表'
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of lesson_section_resource
-- ----------------------------
INSERT INTO `lesson_section_resource`
VALUES (1, 1, 14, 1, DEFAULT, '2023-03-28 18:16:11', '2023-03-28 18:16:11');
INSERT INTO `lesson_section_resource`
VALUES (2, 1, 16, 2, DEFAULT, '2023-03-28 18:25:50', '2023-03-28 18:25:50');
INSERT INTO `lesson_section_resource`
VALUES (3, 1, 17, 3, DEFAULT, '2023-03-28 18:26:14', '2023-03-28 18:26:14');
INSERT INTO `lesson_section_resource`
VALUES (4, 2, 10, 1, DEFAULT, '2023-03-28 18:26:27', '2023-04-03 12:06:07');
INSERT INTO `lesson_section_resource`
VALUES (7, 1, 19, 4, DEFAULT, '2023-03-28 18:44:30', '2023-03-31 18:34:21');
INSERT INTO `lesson_section_resource`
VALUES (8, 2, 11, 2, DEFAULT, '2023-03-28 18:44:37', '2023-03-28 18:44:37');
INSERT INTO `lesson_section_resource`
VALUES (9, 2, 12, 3, DEFAULT, '2023-03-28 18:44:42', '2023-03-28 18:44:42');
INSERT INTO `lesson_section_resource`
VALUES (10, 3, 5, 3, DEFAULT, '2023-03-28 18:46:21', '2023-03-31 18:08:53');
INSERT INTO `lesson_section_resource`
VALUES (11, 4, 18, 41, DEFAULT, '2023-03-28 18:46:28', '2023-04-10 11:18:16');
INSERT INTO `lesson_section_resource`
VALUES (13, 4, 3, 42, DEFAULT, '2023-04-03 12:32:16', '2023-04-04 10:50:14');
INSERT INTO `lesson_section_resource`
VALUES (100, 200, 190, 1, DEFAULT, '2023-04-03 12:32:16', '2023-04-04 10:50:14');
INSERT INTO `lesson_section_resource`
VALUES (101, 200, 191, 2, DEFAULT, '2023-04-03 12:32:16', '2023-04-04 10:50:14');
INSERT INTO `lesson_section_resource`
VALUES (102, 200, 192, 3, DEFAULT, '2023-04-03 12:32:16', '2023-04-04 10:50:14');
INSERT INTO `lesson_section_resource`
VALUES (103, 200, 193, 4, DEFAULT, '2023-04-03 12:32:16', '2023-04-04 10:50:14');
INSERT INTO `lesson_section_resource`
VALUES (110, 201, 15, 1, DEFAULT, '2023-04-03 12:32:16', '2023-04-04 10:50:14');
INSERT INTO `lesson_section_resource`
VALUES (111, 301, 11, 1, DEFAULT, '2023-04-03 12:32:16', '2023-04-04 10:50:14');
INSERT INTO `lesson_section_resource`
VALUES (112, 202, 183, 5, DEFAULT, '2023-04-03 12:32:16', '2023-04-04 10:50:14');
INSERT INTO `lesson_section_resource`
VALUES (130, 203, 29, 41, DEFAULT, '2023-04-03 12:32:16', '2023-04-04 10:50:14');
INSERT INTO `lesson_section_resource`
VALUES (131, 203, 3, 42, DEFAULT, '2023-04-03 12:32:16', '2023-04-04 10:50:14');
INSERT INTO `lesson_section_resource`
VALUES (140, 300, 210, 1, DEFAULT, '2023-04-03 12:32:16', '2023-04-04 10:50:14');
INSERT INTO `lesson_section_resource`
VALUES (141, 300, 211, 2, DEFAULT, '2023-04-03 12:32:16', '2023-04-04 10:50:14');
INSERT INTO `lesson_section_resource`
VALUES (142, 300, 212, 3, DEFAULT, '2023-04-03 12:32:16', '2023-04-04 10:50:14');
INSERT INTO `lesson_section_resource`
VALUES (143, 300, 213, 4, DEFAULT, '2023-04-03 12:32:16', '2023-04-04 10:50:14');
INSERT INTO `lesson_section_resource`
VALUES (144, 300, 214, 5, DEFAULT, '2023-04-03 12:32:16', '2023-04-04 10:50:14');
INSERT INTO `lesson_section_resource`
VALUES (145, 302, 184, 5, DEFAULT, '2023-04-03 12:32:16', '2023-04-04 10:50:14');
INSERT INTO `lesson_section_resource`
VALUES (150, 303, 29, 41, DEFAULT, '2023-04-03 12:32:16', '2023-04-04 10:50:14');
INSERT INTO `lesson_section_resource`
VALUES (151, 303, 3, 42, DEFAULT, '2023-04-03 12:32:16', '2023-04-04 10:50:14');

-- ----------------------------
-- Table structure for mark
-- ----------------------------
DROP TABLE IF EXISTS `mark`;
CREATE TABLE `mark`
(
    `id`               int(11)      NOT NULL AUTO_INCREMENT COMMENT '主键',
    `lesson_id`        int(11)      NOT NULL DEFAULT 0 COMMENT '课堂lesson id',
    `score_item_id`    int(11)      NOT NULL DEFAULT 0 COMMENT '打分项目id',
    `score_subitem_id` int(11)      NOT NULL DEFAULT 0 COMMENT '打分子项id',
    `point`            int(11)      NOT NULL DEFAULT 0 COMMENT '子项分',
    `created_at`       timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
    `updated_at`       timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT = '打分表'
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of mark
-- ----------------------------

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
INSERT INTO `news`
VALUES (6, '新闻通知5',
        '<p style=\"text-align: center;\"><strong>第3条新闻</strong></p><p style=\"text-align: center;\"><span style=\"color: #e03e2d;\"><em><strong>09:34:03 2023-03-14</strong></em></span></p><p>&nbsp; &nbsp;新闻内容段落1</p><p style=\"padding-left: 40px;\"><span style=\"font-family: \'arial black\', sans-serif;\">情况1</span></p><p style=\"padding-left: 40px;\"><span style=\"font-family: \'arial black\', sans-serif;\">情况2</span></p><p style=\"padding-left: 40px;\"><span style=\"font-family: \'arial black\', sans-serif;\">情况3</span></p><p>&nbsp;<span style=\"text-decoration: line-through;\"> &nbsp;新闻内容段落2</span></p><table style=\"border-collapse: collapse; width: 100%;\" border=\"1\"><tbody><tr><td style=\"width: 23.2395%;\">1</td><td style=\"width: 23.2395%;\">2</td><td style=\"width: 23.2395%;\">&nbsp;</td><td style=\"width: 23.2395%;\">&nbsp;</td></tr><tr><td style=\"width: 23.2395%;\">&nbsp;</td><td style=\"width: 23.2395%;\">3</td><td style=\"width: 23.2395%;\">4</td><td style=\"width: 23.2395%;\">4</td></tr><tr><td style=\"width: 23.2395%;\">&nbsp;</td><td style=\"width: 23.2395%;\">5</td><td style=\"width: 23.2395%;\">5</td><td style=\"width: 23.2395%;\">5</td></tr></tbody></table><p><span style=\"text-decoration: underline;\">&nbsp; &nbsp;新闻内容段落3</span></p><p><span style=\"text-decoration: underline;\"><img src=\"http://172.30.100.23:8060/resource/image/1.png\" alt=\"存储在nginx里面的图片\" width=\"300\" height=\"200\" /></span></p><p>&nbsp;</p><p><span style=\"text-decoration: underline;\">图片2</span></p><p><a href=\"http://www.baidu.com\"><strong><u>百度网站</u></strong></a></p>',
        7, 'admin', 1, '2023-03-15 09:11:03', '2023-03-30 15:31:46');
INSERT INTO `news`
VALUES (7, '新闻通知6',
        '<p style=\"text-align: center;\"><strong>第4条新闻</strong></p><p style=\"text-align: center;\"><span style=\"color: #e03e2d;\"><em><strong>09:34:03 2023-03-14</strong></em></span></p><p>&nbsp; &nbsp;新闻内容段落1</p><p style=\"padding-left: 40px;\"><span style=\"font-family: \'arial black\', sans-serif;\">情况1</span></p><p style=\"padding-left: 40px;\"><span style=\"font-family: \'arial black\', sans-serif;\">情况2</span></p><p style=\"padding-left: 40px;\"><span style=\"font-family: \'arial black\', sans-serif;\">情况3</span></p><p>&nbsp;<span style=\"text-decoration: line-through;\"> &nbsp;新闻内容段落2</span></p><table style=\"border-collapse: collapse; width: 100%;\" border=\"1\"><tbody><tr><td style=\"width: 23.2395%;\">1</td><td style=\"width: 23.2395%;\">2</td><td style=\"width: 23.2395%;\">&nbsp;</td><td style=\"width: 23.2395%;\">&nbsp;</td></tr><tr><td style=\"width: 23.2395%;\">&nbsp;</td><td style=\"width: 23.2395%;\">3</td><td style=\"width: 23.2395%;\">4</td><td style=\"width: 23.2395%;\">4</td></tr><tr><td style=\"width: 23.2395%;\">&nbsp;</td><td style=\"width: 23.2395%;\">5</td><td style=\"width: 23.2395%;\">5</td><td style=\"width: 23.2395%;\">5</td></tr></tbody></table><p><span style=\"text-decoration: underline;\">&nbsp; &nbsp;新闻内容段落3</span></p><p><span style=\"text-decoration: underline;\"><img src=\"http://172.30.100.23:8060/resource/image/1.png\" alt=\"存储在nginx里面的图片\" width=\"300\" height=\"200\" /></span></p><p>&nbsp;</p><p><span style=\"text-decoration: underline;\">图片2</span></p><p><a href=\"http://www.baidu.com\"><strong><u>百度网站</u></strong></a></p>',
        1, 'admin', 1, '2023-03-15 09:11:08', '2023-03-16 17:15:04');
INSERT INTO `news`
VALUES (8, 'news', 'abcdefghijklmnopqrstuvwxyz', 6, 'admin', 1, '2023-03-17 10:01:00', '2023-03-17 10:01:00');
INSERT INTO `news`
VALUES (100, '我校虚拟仿真实验室建成启动',
        '<p><strong>本网讯 &nbsp;</strong>我校虚拟仿真实验室建成启用，新建成的虚拟仿真实验室是目前全国同类院校中建设标准和规格最高的实验室。3月22日，我校校长仲伟合一行参观了该实验室并就实验室应用、口译信息化教学等问题进行了深入的研讨。</p>\r\n<p>&nbsp;</p>\r\n<p><img src=\"http://172.30.100.23:8060/resource/image/33D.jpg\" alt=\"\" width=\"1120\" height=\"630\" /></p>\r\n<p>&nbsp;</p>\r\n<p>实验室利用先进的虚拟仿真和虚拟现实技术与多媒体设备相结合，根据外语教学的特点，形成了独具特色的外语虚拟仿真和虚拟现实的实践中心。实验室有四大功能：情景教学、情景微课录制、课堂录播、虚拟现实实训，全方位满足学生实践技能提升和教学资源建设两大方面的需求，同时对于创新课堂评估方式、提高学生积极性以及丰富翻转课堂实践等方面都有很大帮助。国家级同声传译实验教学中心副主任古煜奎和厂家代表冯大治对实验室的功能和应用进行了介绍。</p>\r\n<p>&nbsp;</p>\r\n<p><img src=\"http://172.30.100.23:8060/resource/image/34.jpg\" alt=\"\" width=\"1120\" height=\"630\" /></p>\r\n',
        101, 'admin', 1, '2023-04-17 11:06:09', '2023-04-17 17:39:40');
INSERT INTO `news`
VALUES (101, '构建未来：建设国家虚拟仿真实验室的重要性与挑战',
        '<p id=\"1LB4TVGA\">随着科技的不断发展，虚拟仿真技术逐渐成为一个备受瞩目的领域。国家虚拟仿真实验室的建设，也成为了现代化国家的必然选择。</p>\r\n<p id=\"1LB4TVGB\">首先，国家虚拟仿真实验室的建设具有重要意义。虚拟仿真技术可以模拟现实场景，使人们在虚拟环境中体验真实世界的各种情境，从而提高人们的认知水平。这项技术在军事、航空、医疗、教育等领域具有广泛的应用前景，可以有效地提高生产效率，降低事故风险，缩短研发周期，降低成本等方面的优势。</p>\r\n<p><p><img src=\"http://172.30.100.23:8060/resource/image/新闻图片1.jpg\" alt=\"\" width=\"660\" height=\"468\" /></p></p>\r\n<p id=\"1LB4TVGE\">其次，建设国家虚拟仿真实验室也面临着许多挑战。其中最大的挑战是技术方面的问题。虚拟仿真技术涉及到许多领域的知识，需要各种专业人才的参与。此外，虚拟仿真技术的成本较高，需要大量的投入。建设国家虚拟仿真实验室需要有政府的大力支持和专业的人才支持。</p>\r\n<p id=\"1LB4TVGF\">在面对这些挑战时，政府需要采取积极的措施。政府可以制定相关政策，支持虚拟仿真技术的研发和应用。政府还可以通过投资，吸引专业人才的参与，提高技术水平。此外，政府还可以积极引导企业和机构参与虚拟仿真技术的研究和应用，形成产学研合作的良好机制。</p>\r\n<img src=\"http://172.30.100.23:8060/resource/image/新闻图片2.jpg\" alt=\"\" width=\"660\" height=\"365\" />\r\n<p>2019年1月31号，教育部高等教育司发布了《关于2018年度国家虚拟仿真实验教学项目名单》的公示，共有298个项目入选，来自全国186所高校。这也是继2018年6月教育部公布首批105个国家级虚拟现实仿真实验教学项目后，教育部再次认定这一重量级教学项目。其中北京欧倍尔助力30所高校成功申报2018年度国家虚拟仿真实验教学项目，在同行业中居于领先水平，彰显了公司雄厚的技术实力。</p>\r\n<img src=\"http://172.30.100.23:8060/resource/image/新闻图片3.jpg\" alt=\"\" width=\"660\" height=\"359\" />\r\n<p id=\"1LB4TVGL\">从入选名单来看，本次共有298项目进入公示，共涉及23个学科领域。机械类教学项目最多，共有30个，数量最多。临床医学类共有29个，位居第二位。生物科学类共有20个，位居第三位。化工与制药类共有18个，数量位居第四。医学基础类、药学类、动物类、能源动力类、食品科学与工程类、土木类等数量也较多。</p>\r\n<p id=\"1LB4TVGM\">建设国家虚拟仿真实验室是一个具有重要意义的事业。在克服技术方面的困难和面对各种挑战的同时，政府和社会各界应积极支持和参与，共同推进虚拟仿真技术的发展和应用，为实现现代化建设和经济发展做出贡献。</p>',
        102, 'admin', 1, '2023-04-17 11:06:09', '2023-04-17 17:39:40');
INSERT INTO `news`
VALUES (102, '虚拟仿真实验项目助力线上实验教学',
        '<p>&nbsp; &nbsp; &nbsp; &nbsp;5月3日晚，材料学部董超芳与实验中心黄鹏、商北雁使用材料科学与工程学院实验中心建设的材料虚拟仿真项目，为2020级材料化学专业25名学生在线进行了《材料科学基础实验C》课程线上虚拟实验教学，内容包括Fe-Fe3C相图与铁碳合金平衡组织显微观察、金相制样与显微观察等。学生课后表示，虚拟仿真实验很新奇也很好操作，还能在需要注意的实验操作上给我一些提示，并且可以让我很直观地看到实验结果。教师们积极使用虚拟仿真实验项目，以保障疫情防控期间实验教学顺利开展。</p>\r\n<p>图片</p>\r\n<p>&nbsp;</p>\r\n<p>&nbsp; &nbsp; &nbsp; &nbsp;材料科学与工程学院实验中心自2007年获批国家级实验教学示范中心和2014年国家级虚拟仿真实验教学中心以来，在教育部与学校支持下，建成了虚拟仿真实验项目1.0版和2.0版。学校材料科学与工程学院、能源与环境工程学院、高等工程师学院等学院都已在课程教学、专业实习、学科竞赛中使用该项目，累计使用超过1万余人次。2017年，国内首次开设校级公选课程《材料虚拟实验》，并编写了课程讲义。2020年疫情期间项目免费向学校和国内高校开放，有力地支持了疫情期间的线上实验教学。同时，在大学生金相实验技能大赛中，创建了线上预赛、线下决赛，虚实结合比赛的新形式，为疫情期间学科竞赛提供了新赛道。其中虚拟仿真实验项目1.0版《系列材料虚拟实验》于2021年荣获第六届全国高等学校教师创新大赛一等奖。虚拟仿真实验项目2.0版&ldquo;金属材料成分、工艺、组织和性能一体化设计虚拟仿真实验&rdquo;于2021年6月申报国家级虚拟仿真一流课程，并建立了虚拟仿真沉浸式体验实验室。2021年8月，实验中心在国内率先牵头成立材料虚仿实验教学虚拟教研室，目前已有80余所国内高校加入，开展的教研活动获得了良好效果，带动了各高校推进融合式实验教学改革。2021年12月，项目获批教育部产学合作协同育人项目&ldquo;建设实验教学虚拟教研室&rdquo;，打造了材料类虚拟仿真项目共享应用云基地。​</p>\r\n<p>图片</p>\r\n<p>&nbsp; &nbsp; &nbsp; &nbsp;截至2022年5月3日，虚拟仿真实验1.0版已共享到国内国防科技大学、中山大学、青海大学等28所高校。2020年疫情出现以来，该版本支持了哈尔滨工业大学（威海）、烟台大学、江西科技师范大学、北华航天工业学院等高校开展线上实验教学。随着疫情防控常态化，山东大学、湖南大学、安徽大学等高校也开始采用虚拟仿真实验2.0版，通过虚拟仿真一流课开设了8门实验课。</p>\r\n<p>&nbsp; &nbsp; &nbsp; &nbsp;材料科学与工程学院实验中心发挥双国家级实验教学中心示范引领优势，积极探索虚实结合实验教学新模式，深度挖掘我校优质实验教学资源，丰富虚拟仿真实验项目内容，进一步拓展推广应用途径，有力支持了我校乃至国内高校材料及相关专业的本科实验教学，推动了材料专业跨校、跨地区的虚拟仿真实验教学项目共建共享，开拓了疫情防控条件下实验教学新模式。</p>',
        103, 'admin', 1, '2023-04-17 11:06:09', '2023-04-17 17:39:40');

-- ----------------------------
-- Table structure for organization
-- ----------------------------
DROP TABLE IF EXISTS `organization`;
CREATE TABLE `organization`
(
    `id`            int(11)      NOT NULL AUTO_INCREMENT COMMENT '主键',
    `school_id`     int(11)      NOT NULL DEFAULT 0 COMMENT 'school id',
    `department_id` int(11)      NOT NULL DEFAULT 0 COMMENT 'department id',
    `class_id`      int(11)      NOT NULL DEFAULT 0 COMMENT 'class id',
    `year`          int(11)      NOT NULL DEFAULT 0 COMMENT '学年',
    `created_at`    timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
    `updated_at`    timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT = '学校组织表'
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of organization
-- ----------------------------

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
INSERT INTO `resource`
VALUES (9, '第5个图片', '图片', 1, 'resource/image/5.webp', DEFAULT, '2023-03-17 11:30:37', '2023-04-19 19:17:02');
INSERT INTO `resource`
VALUES (12, '第3个个shipin', '视频z', 8, 'resource/vedio/3.mp4', DEFAULT, '2023-03-22 11:04:00', '2023-04-19 19:16:57');
INSERT INTO `resource`
VALUES (14, '实验目的', '富文本1', 7, '2', DEFAULT, '2023-03-28 18:07:30', '2023-04-06 17:56:24');
INSERT INTO `resource`
VALUES (15, '富文本2', '富文本2', 7, '3', DEFAULT, '2023-03-28 18:08:40', '2023-04-06 17:54:35');
INSERT INTO `resource`
VALUES (16, '实验要求', '富文本3', 7, '4', DEFAULT, '2023-03-28 18:08:45', '2023-04-06 17:56:43');
INSERT INTO `resource`
VALUES (17, '实验步骤和要点', '富文本4', 7, '5', DEFAULT, '2023-03-28 18:08:52', '2023-04-06 17:56:52');
INSERT INTO `resource`
VALUES (18, '注意事项', '富文本5', 7, '6', DEFAULT, '2023-03-28 18:08:57', '2023-04-06 17:56:59');
INSERT INTO `resource`
VALUES (19, '富文本6', '富文本6', 7, '7', DEFAULT, '2023-03-28 18:09:04', '2023-04-06 17:54:46');
INSERT INTO `resource`
VALUES (20, '作业要求1', '作业要求1', 7, '8', DEFAULT, '2023-04-10 11:15:01', '2023-04-10 11:15:01');
INSERT INTO `resource`
VALUES (21, '作业提交内容', '作业提交内容', 7, '9', DEFAULT, '2023-04-10 14:48:01', '2023-04-10 14:53:21');
INSERT INTO `resource`
VALUES (22, '未使用', '作业要求1', 7, '10', DEFAULT, '2023-04-10 14:48:03', '2023-04-10 14:53:26');
INSERT INTO `resource`
VALUES (23, '未使用', '作业要求1', 7, '11', DEFAULT, '2023-04-10 14:48:04', '2023-04-10 14:53:32');
INSERT INTO `resource`
VALUES (24, 'ibe_edu20230329.sql', '作业文件', 0, 'uploads/202304111117122714521008e20f8', DEFAULT, '2023-04-11 11:18:31',
        '2023-04-11 11:18:31');
INSERT INTO `resource`
VALUES (25, '11.png', '作业文件', 0, 'uploads/2023041111171227090160059b290', DEFAULT, '2023-04-11 11:18:37',
        '2023-04-11 11:18:37');
INSERT INTO `resource`
VALUES (26, '富文本', '富文本', 7, '13', '2023-04-11 15:46:31', DEFAULT, '2023-04-11 15:46:31');
INSERT INTO `resource`
VALUES (27, '作业要求a', '作业要求a', 7, '15', '2023-04-11 15:49:37', DEFAULT, '2023-04-11 15:49:37');
INSERT INTO `resource`
VALUES (28, '作业要求b', '作业要求b', 7, '16', '2023-04-11 15:50:12', DEFAULT, '2023-04-11 15:50:12');
INSERT INTO `resource`
VALUES (29, '作业要求b', '作业要求b', 7, '18', '2023-04-11 15:52:46', DEFAULT, '2023-04-11 15:52:46');
INSERT INTO `resource`
VALUES (30, '富文本', '富文本', 7, '19', '2023-04-11 16:17:56', DEFAULT, '2023-04-11 16:17:56');
INSERT INTO `resource`
VALUES (31, '富文本', '富文本', 7, '20', '2023-04-11 16:20:15', DEFAULT, '2023-04-11 16:20:15');
INSERT INTO `resource`
VALUES (32, 'ibe_edu20230329.sql', '作业文件', 0, 'uploads/202304121755165063266008e20f8', DEFAULT, '2023-04-12 17:55:19',
        '2023-04-12 17:55:19');
INSERT INTO `resource`
VALUES (33, 'ibe_edu20230329.sql', '作业文件', 0, 'uploads/202304121755403778440008e20f8', DEFAULT, '2023-04-12 17:55:41',
        '2023-04-12 17:55:41');
INSERT INTO `resource`
VALUES (34, '11.png', '作业文件', 0, 'uploads/2023041218005058953050059b290', DEFAULT, '2023-04-12 18:00:52',
        '2023-04-12 18:00:52');
INSERT INTO `resource`
VALUES (35, 'ibe_edu20230329.sql', '作业文件', 0, 'uploads/202304121800505901451008e20f8', DEFAULT, '2023-04-12 18:00:52',
        '2023-04-12 18:00:52');
INSERT INTO `resource`
VALUES (36, '11.png', '作业文件', 0, 'uploads/2023041409330312179420059b290', DEFAULT, '2023-04-14 09:33:03',
        '2023-04-14 09:33:03');
INSERT INTO `resource`
VALUES (100, '抱杆展示图片', '图片', 1, 'resource/image/抱杆.png', DEFAULT, '2023-04-12 18:00:52', '2023-04-19 19:16:43');
INSERT INTO `resource`
VALUES (101, 'banner展示图片', '图片', 1, 'resource/image/banner1.png', DEFAULT, '2023-04-12 18:00:52',
        '2023-04-19 19:16:45');
INSERT INTO `resource`
VALUES (102, 'banner展示图片', '图片', 1, 'resource/image/banner2.png', DEFAULT, '2023-04-12 18:00:52',
        '2023-04-19 19:16:47');
INSERT INTO `resource`
VALUES (103, 'banner展示图片', '图片', 1, 'resource/image/banner3.png', DEFAULT, '2023-04-12 18:00:52',
        '2023-04-19 19:16:49');
INSERT INTO `resource`
VALUES (104, 'para.txt', '作业文件', 0, 'uploads/20230418152922514970600cabfef', DEFAULT, '2023-04-18 15:29:23',
        '2023-04-18 15:29:23');
INSERT INTO `resource`
VALUES (105, '702.zip', '作业文件', 0, 'uploads/20230419160749548498300b1b4ca', DEFAULT, '2023-04-19 16:07:52',
        '2023-04-19 16:07:52');
INSERT INTO `resource`
VALUES (106, '702.zip', '作业文件', 0, 'uploads/20230419164514677148300b1b4ca', DEFAULT, '2023-04-19 16:45:17',
        '2023-04-19 16:45:17');
INSERT INTO `resource`
VALUES (107, '702.zip', '作业文件', 0, 'uploads/20230419164548380835700b1b4ca', DEFAULT, '2023-04-19 16:45:48',
        '2023-04-19 16:45:48');
INSERT INTO `resource`
VALUES (114, '11.png', '作业文件', 0, 'uploads/2023042009530936762330059b290', DEFAULT, '2023-04-20 09:53:09',
        '2023-04-20 09:53:09');
INSERT INTO `resource`
VALUES (115, '11.png', '作业文件', 0, 'uploads/2023042009533155826010059b290.png', DEFAULT, '2023-04-20 09:53:33',
        '2023-04-20 09:53:33');
INSERT INTO `resource`
VALUES (118, '视频内容', '引导视频', 8, 'resource/vedio/抱杆app工程应用和实训课程展示视频.mp4', DEFAULT, '2023-04-20 09:53:33',
        '2023-04-20 09:53:33');
INSERT INTO `resource`
VALUES (183, '实验三塔身吊装', 'ibe文件', 5, 'resource/ibe/实验三塔身吊装.ibe', DEFAULT, '2023-04-20 09:53:33', '2023-04-20 09:53:33');
INSERT INTO `resource`
VALUES (184, '实验四塔身吊装', 'ibe文件', 5, 'resource/ibe/实验四塔身吊装.ibe', DEFAULT, '2023-04-20 09:53:33', '2023-04-20 09:53:33');
INSERT INTO `resource`
VALUES (185, '实验五塔身吊装', 'ibe文件', 5, 'resource/ibe/实验五塔身吊装.ibe', DEFAULT, '2023-04-20 09:53:33', '2023-04-20 09:53:33');
INSERT INTO `resource`
VALUES (186, '实验六塔身吊装', 'ibe文件', 5, 'resource/ibe/实验六塔身吊装.ibe', DEFAULT, '2023-04-20 09:53:33', '2023-04-20 09:53:33');
INSERT INTO `resource`
VALUES (187, '实验七塔身吊装', 'ibe文件', 5, 'resource/ibe/实验七塔身吊装.ibe', DEFAULT, '2023-04-20 09:53:33', '2023-04-20 09:53:33');
INSERT INTO `resource`
VALUES (188, '塔身吊装app', 'app文件', 4, 'resource/ibe/塔身吊装app.app', DEFAULT, '2023-04-20 09:53:33', '2023-04-20 09:53:33');
INSERT INTO `resource`
VALUES (190, '一.	实验目的', '课程内容富文本', 7, '101', DEFAULT, '2023-04-20 09:53:33', '2023-04-20 09:53:33');
INSERT INTO `resource`
VALUES (191, '二.	输电线路铁塔组立抱杆工程简介', '课程内容富文本', 7, '102', DEFAULT, '2023-04-20 09:53:33', '2023-04-20 09:53:33');
INSERT INTO `resource`
VALUES (192, '三.	输电线路铁塔组立抱杆四种工况简介', '课程内容富文本', 7, '103', DEFAULT, '2023-04-20 09:53:33', '2023-04-20 09:53:33');
INSERT INTO `resource`
VALUES (193, '四.	Simdroid仿真APP体验流程', '课程内容富文本', 7, '104', DEFAULT, '2023-04-20 09:53:33', '2023-04-20 09:53:33');
INSERT INTO `resource`
VALUES (200, '左心室辅助装置', '图片', 1, 'resource/image/左心室辅助装置.png', DEFAULT, '2023-04-12 18:00:52', '2023-04-19 19:16:45');
INSERT INTO `resource`
VALUES (201, '厂房除湿除尘', '图片', 1, 'resource/image/厂房除湿除尘.png', DEFAULT, '2023-04-12 18:00:52', '2023-04-19 19:16:45');
INSERT INTO `resource`
VALUES (202, '传动轴', '图片', 1, 'resource/image/传动轴.png', DEFAULT, '2023-04-12 18:00:52', '2023-04-19 19:16:45');
INSERT INTO `resource`
VALUES (203, '绞缆机卷筒', '图片', 1, 'resource/image/绞缆机卷筒.png', DEFAULT, '2023-04-12 18:00:52', '2023-04-19 19:16:45');
INSERT INTO `resource`
VALUES (210, '一.	实验目的', '课程内容富文本', 7, '110', DEFAULT, '2023-04-20 09:53:33', '2023-04-20 09:53:33');
INSERT INTO `resource`
VALUES (211, '二.	实验要求', '课程内容富文本', 7, '111', DEFAULT, '2023-04-20 09:53:33', '2023-04-20 09:53:33');
INSERT INTO `resource`
VALUES (212, '三.	实验步骤与要点', '课程内容富文本', 7, '112', DEFAULT, '2023-04-20 09:53:33', '2023-04-20 09:53:33');
INSERT INTO `resource`
VALUES (213, '四.	注意事项', '课程内容富文本', 7, '113', DEFAULT, '2023-04-20 09:53:33', '2023-04-20 09:53:33');
INSERT INTO `resource`
VALUES (214, '五.	实操步骤', '课程内容富文本', 7, '114', DEFAULT, '2023-04-20 09:53:33', '2023-04-20 09:53:33');
INSERT INTO `resource`
VALUES (234, '富文本', '富文本', 7, '115', DEFAULT, '2023-04-23 19:49:49', '2023-04-23 19:49:49');
INSERT INTO `resource`
VALUES (235, '富文本', '富文本', 7, '116', DEFAULT, '2023-04-23 20:04:57', '2023-04-23 20:04:57');
INSERT INTO `resource`
VALUES (236, '富文本', '富文本', 7, '117', DEFAULT, '2023-04-23 20:06:24', '2023-04-23 20:06:24');
INSERT INTO `resource`
VALUES (237, '富文本', '富文本', 7, '118', DEFAULT, '2023-04-23 20:12:02', '2023-04-23 20:12:02');
INSERT INTO `resource`
VALUES (238, '富文本', '富文本', 7, '119', DEFAULT, '2023-04-23 20:12:38', '2023-04-23 20:12:38');
INSERT INTO `resource`
VALUES (239, '富文本', '富文本', 7, '120', DEFAULT, '2023-04-23 20:16:31', '2023-04-23 20:16:31');
INSERT INTO `resource`
VALUES (315, '课程图片', '课程图片', 1, 'resource/image/电子机柜三维建模与散热仿真分析.png', DEFAULT, '2023-04-23 20:16:31',
        '2023-04-23 20:16:31');
INSERT INTO `resource`
VALUES (316, '课程图片', '课程图片', 1, 'resource/image/Simetherm芯片级 IGBT模块仿真计算方案.png', DEFAULT, '2023-04-23 20:16:31',
        '2023-04-23 20:16:31');
INSERT INTO `resource`
VALUES (317, '课程图片', '课程图片', 1, 'resource/image/Simetherm电子设备机箱散热分析方案.png', DEFAULT, '2023-04-23 20:16:31',
        '2023-04-23 20:16:31');
INSERT INTO `resource`
VALUES (400, 'CAE行业发展及自主仿真之路', 'CAE行业发展及自主仿真之路', 1, 'resource/image/CAE行业发展及自主仿真之路.png', DEFAULT, '2023-04-19 19:03:46',
        '2023-04-19 19:11:30');
INSERT INTO `resource`
VALUES (401, '云仿真平台SimCapsule平台简介', '云仿真平台SimCapsule平台简介', 1, 'resource/image/云仿真平台SimCapsule平台简介.png', DEFAULT,
        '2023-04-19 19:03:46', '2023-04-19 19:11:30');
INSERT INTO `resource`
VALUES (402, '电磁铁仿真案例', '电磁铁仿真案例', 1, 'resource/image/电磁铁仿真案例.png', DEFAULT, '2023-04-19 19:03:46',
        '2023-04-19 19:11:30');
INSERT INTO `resource`
VALUES (403, '有限单元法与Simdroid介绍', '有限单元法与Simdroid介绍', 1, 'resource/image/有限单元法与Simdroid介绍.png', DEFAULT,
        '2023-04-19 19:03:46',
        '2023-04-19 19:11:30');
INSERT INTO `resource`
VALUES (404, '数字化仿真赋能高校人才培养', '数字化仿真赋能高校人才培养', 1, 'resource/image/数字化仿真赋能高校人才培养.png', DEFAULT, '2023-04-19 19:03:46',
        '2023-04-19 19:11:30');
INSERT INTO `resource`
VALUES (405, '足球侧向力影响因素分析', '足球侧向力影响因素分析', 1, 'resource/image/足球侧向力影响因素分析.png', DEFAULT, '2023-04-19 19:03:46',
        '2023-04-19 19:11:30');
INSERT INTO `resource`
VALUES (406, '电磁、结构、流体工程实践及算例开发培训课程', '电磁、结构、流体工程实践及算例开发培训课程', 1, 'resource/image/电磁、结构、流体工程实践及算例开发培训课程.png', DEFAULT,
        '2023-04-19 19:03:46', '2023-04-19 19:11:30');
INSERT INTO `resource`
VALUES (407, '课程8-仿真APP开发实训-流体仿真', '课程8-仿真APP开发实训-流体仿真', 1, 'resource/image/课程8-仿真APP开发实训-流体仿真.png', DEFAULT,
        '2023-04-19 19:03:46', '2023-04-19 19:11:30');
INSERT INTO `resource`
VALUES (500, '电磁铁仿真案例', '电磁铁仿真案例', 8, 'resource/vedio/电磁铁仿真案例.mp4', DEFAULT, '2023-04-19 19:03:46',
        '2023-04-19 19:11:30');
INSERT INTO `resource`
VALUES (501, '有限单元法与Simdroid介绍', '有限单元法与Simdroid介绍', 8, 'resource/vedio/有限单元法与Simdroid介绍.mp4', DEFAULT,
        '2023-04-19 19:03:46',
        '2023-04-19 19:11:30');
INSERT INTO `resource`
VALUES (502, '云仿真平台SimCapsule平台简介', '云仿真平台SimCapsule平台简介', 8, 'resource/vedio/云仿真平台SimCapsule平台简介.mp4', DEFAULT,
        '2023-04-19 19:03:46', '2023-04-19 19:11:30');
INSERT INTO `resource`
VALUES (503, '课程8-仿真APP开发实训-流体仿真', '课程8-仿真APP开发实训-流体仿真', 8, 'resource/vedio/课程8-仿真APP开发实训-流体仿真.mp4', DEFAULT,
        '2023-04-19 19:03:46', '2023-04-19 19:11:30');
INSERT INTO `resource`
VALUES (504, '足球侧向力影响因素分析', '足球侧向力影响因素分析', 8, 'resource/vedio/足球侧向力影响因素分析.mp4', DEFAULT, '2023-04-19 19:03:46',
        '2023-04-19 19:11:30');
INSERT INTO `resource`
VALUES (505, '数字化仿真赋能高校人才培养', '数字化仿真赋能高校人才培养', 8, 'resource/vedio/数字化仿真赋能高校人才培养.mp4', DEFAULT, '2023-04-19 19:03:46',
        '2023-04-19 19:11:30');
INSERT INTO `resource`
VALUES (506, 'CAE行业发展及自主仿真之路', 'CAE行业发展及自主仿真之路', 8, 'resource/vedio/CAE行业发展及自主仿真之路.mp4', DEFAULT, '2023-04-19 19:03:46',
        '2023-04-19 19:11:30');
INSERT INTO `resource`
VALUES (507, '电磁、结构、流体工程实践及算例开发培训课程', '电磁、结构、流体工程实践及算例开发培训课程', 8, 'resource/vedio/电磁、结构、流体工程实践及算例开发培训课程.mp4', DEFAULT,
        '2023-04-19 19:03:46', '2023-04-19 19:11:30');
INSERT INTO `resource`
VALUES (517, '富文本', '富文本', 7, '127', DEFAULT, '2023-04-28 15:51:34', '2023-04-28 15:51:34');
INSERT INTO `resource`
VALUES (519, '富文本', '富文本', 7, '129', DEFAULT, '2023-04-28 16:02:09', '2023-04-28 16:02:09');
INSERT INTO `resource`
VALUES (520, '实验二.ibe', '作业文件', 0, 'uploads/202304281655121174747693aecf9.ibe', DEFAULT, '2023-04-28 16:55:12',
        '2023-04-28 16:55:12');
INSERT INTO `resource`
VALUES (521, '富文本', '富文本', 7, '130', DEFAULT, '2023-04-28 16:55:32', '2023-04-28 16:55:32');

-- ----------------------------
-- Table structure for resource_richtext
-- ----------------------------
DROP TABLE IF EXISTS `resource_richtext`;
CREATE TABLE `resource_richtext`
(
    `id`         int(11)                                               NOT NULL AUTO_INCREMENT COMMENT '主键',
    `content`    text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '基础资源富文本内容',
    `created_at` timestamp(0)                                          NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
    `updated_at` timestamp(0)                                          NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 131
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT = '存储系统中用到的富文本资源'
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of resource_richtext
-- ----------------------------
INSERT INTO `resource_richtext`
VALUES (2,
        '<p style=\"text-align: center;\"><strong>富文本1</strong></p><p style=\"text-align: center;\"><span style=\"color: #e03e2d;\"><em><strong>09:34:03 2023-03-14</strong></em></span></p><p>&nbsp; &nbsp;新闻内容段落1</p><p style=\"padding-left: 40px;\"><span style=\"font-family: \'arial black\', sans-serif;\">情况1</span></p><p style=\"padding-left: 40px;\"><span style=\"font-family: \'arial black\', sans-serif;\">情况2</span></p><p style=\"padding-left: 40px;\"><span style=\"font-family: \'arial black\', sans-serif;\">情况3</span></p><p>&nbsp;<span style=\"text-decoration: line-through;\"> &nbsp;新闻内容段落2</span></p><table style=\"border-collapse: collapse; width: 100%;\" border=\"1\"><tbody><tr><td style=\"width: 23.2395%;\">1</td><td style=\"width: 23.2395%;\">2</td><td style=\"width: 23.2395%;\">&nbsp;</td><td style=\"width: 23.2395%;\">&nbsp;</td></tr><tr><td style=\"width: 23.2395%;\">&nbsp;</td><td style=\"width: 23.2395%;\">3</td><td style=\"width: 23.2395%;\">4</td><td style=\"width: 23.2395%;\">4</td></tr><tr><td style=\"width: 23.2395%;\">&nbsp;</td><td style=\"width: 23.2395%;\">5</td><td style=\"width: 23.2395%;\">5</td><td style=\"width: 23.2395%;\">5</td></tr></tbody></table><p><span style=\"text-decoration: underline;\">&nbsp; &nbsp;新闻内容段落3</span></p><p><span style=\"text-decoration: underline;\"><img src=\"http://172.30.100.23:8060/resource/image/1.png\" alt=\"存储在nginx里面的图片\" width=\"300\" height=\"200\" /></span></p><p>&nbsp;</p><p><span style=\"text-decoration: underline;\">图片2</span></p><p><a href=\"http://www.baidu.com\"><strong><u>百度网站</u></strong></a></p>',
        '2023-03-28 18:07:30', '2023-03-28 18:07:30');
INSERT INTO `resource_richtext`
VALUES (3,
        '<p style=\"text-align: center;\"><strong>富文本2</strong></p><p style=\"text-align: center;\"><span style=\"color: #e03e2d;\"><em><strong>09:34:03 2023-03-14</strong></em></span></p><p>&nbsp; &nbsp;新闻内容段落1</p><p style=\"padding-left: 40px;\"><span style=\"font-family: \'arial black\', sans-serif;\">情况1</span></p><p style=\"padding-left: 40px;\"><span style=\"font-family: \'arial black\', sans-serif;\">情况2</span></p><p style=\"padding-left: 40px;\"><span style=\"font-family: \'arial black\', sans-serif;\">情况3</span></p><p>&nbsp;<span style=\"text-decoration: line-through;\"> &nbsp;新闻内容段落2</span></p><table style=\"border-collapse: collapse; width: 100%;\" border=\"1\"><tbody><tr><td style=\"width: 23.2395%;\">1</td><td style=\"width: 23.2395%;\">2</td><td style=\"width: 23.2395%;\">&nbsp;</td><td style=\"width: 23.2395%;\">&nbsp;</td></tr><tr><td style=\"width: 23.2395%;\">&nbsp;</td><td style=\"width: 23.2395%;\">3</td><td style=\"width: 23.2395%;\">4</td><td style=\"width: 23.2395%;\">4</td></tr><tr><td style=\"width: 23.2395%;\">&nbsp;</td><td style=\"width: 23.2395%;\">5</td><td style=\"width: 23.2395%;\">5</td><td style=\"width: 23.2395%;\">5</td></tr></tbody></table><p><span style=\"text-decoration: underline;\">&nbsp; &nbsp;新闻内容段落3</span></p><p><span style=\"text-decoration: underline;\"><img src=\"http://172.30.100.23:8060/resource/image/1.png\" alt=\"存储在nginx里面的图片\" width=\"300\" height=\"200\" /></span></p><p>&nbsp;</p><p><span style=\"text-decoration: underline;\">图片2</span></p><p><a href=\"http://www.baidu.com\"><strong><u>百度网站</u></strong></a></p>',
        '2023-03-28 18:08:40', '2023-03-28 18:08:40');
INSERT INTO `resource_richtext`
VALUES (4,
        '<p style=\"text-align: center;\"><strong>富文本2</strong></p><p style=\"text-align: center;\"><span style=\"color: #e03e2d;\"><em><strong>09:34:03 2023-03-14</strong></em></span></p><p>&nbsp; &nbsp;新闻内容段落1</p><p style=\"padding-left: 40px;\"><span style=\"font-family: \'arial black\', sans-serif;\">情况1</span></p><p style=\"padding-left: 40px;\"><span style=\"font-family: \'arial black\', sans-serif;\">情况2</span></p><p style=\"padding-left: 40px;\"><span style=\"font-family: \'arial black\', sans-serif;\">情况3</span></p><p>&nbsp;<span style=\"text-decoration: line-through;\"> &nbsp;新闻内容段落2</span></p><table style=\"border-collapse: collapse; width: 100%;\" border=\"1\"><tbody><tr><td style=\"width: 23.2395%;\">1</td><td style=\"width: 23.2395%;\">2</td><td style=\"width: 23.2395%;\">&nbsp;</td><td style=\"width: 23.2395%;\">&nbsp;</td></tr><tr><td style=\"width: 23.2395%;\">&nbsp;</td><td style=\"width: 23.2395%;\">3</td><td style=\"width: 23.2395%;\">4</td><td style=\"width: 23.2395%;\">4</td></tr><tr><td style=\"width: 23.2395%;\">&nbsp;</td><td style=\"width: 23.2395%;\">5</td><td style=\"width: 23.2395%;\">5</td><td style=\"width: 23.2395%;\">5</td></tr></tbody></table><p><span style=\"text-decoration: underline;\">&nbsp; &nbsp;新闻内容段落3</span></p><p><span style=\"text-decoration: underline;\"><img src=\"http://172.30.100.23:8060/resource/image/1.png\" alt=\"存储在nginx里面的图片\" width=\"300\" height=\"200\" /></span></p><p>&nbsp;</p><p><span style=\"text-decoration: underline;\">图片2</span></p><p><a href=\"http://www.baidu.com\"><strong><u>百度网站</u></strong></a></p>',
        '2023-03-28 18:08:45', '2023-03-28 18:08:45');
INSERT INTO `resource_richtext`
VALUES (5,
        '<p style=\"text-align: center;\"><strong>富文本4</strong></p><p style=\"text-align: center;\"><span style=\"color: #e03e2d;\"><em><strong>09:34:03 2023-03-14</strong></em></span></p><p>&nbsp; &nbsp;新闻内容段落1</p><p style=\"padding-left: 40px;\"><span style=\"font-family: \'arial black\', sans-serif;\">情况1</span></p><p style=\"padding-left: 40px;\"><span style=\"font-family: \'arial black\', sans-serif;\">情况2</span></p><p style=\"padding-left: 40px;\"><span style=\"font-family: \'arial black\', sans-serif;\">情况3</span></p><p>&nbsp;<span style=\"text-decoration: line-through;\"> &nbsp;新闻内容段落2</span></p><table style=\"border-collapse: collapse; width: 100%;\" border=\"1\"><tbody><tr><td style=\"width: 23.2395%;\">1</td><td style=\"width: 23.2395%;\">2</td><td style=\"width: 23.2395%;\">&nbsp;</td><td style=\"width: 23.2395%;\">&nbsp;</td></tr><tr><td style=\"width: 23.2395%;\">&nbsp;</td><td style=\"width: 23.2395%;\">3</td><td style=\"width: 23.2395%;\">4</td><td style=\"width: 23.2395%;\">4</td></tr><tr><td style=\"width: 23.2395%;\">&nbsp;</td><td style=\"width: 23.2395%;\">5</td><td style=\"width: 23.2395%;\">5</td><td style=\"width: 23.2395%;\">5</td></tr></tbody></table><p><span style=\"text-decoration: underline;\">&nbsp; &nbsp;新闻内容段落3</span></p><p><span style=\"text-decoration: underline;\"><img src=\"http://172.30.100.23:8060/resource/image/1.png\" alt=\"存储在nginx里面的图片\" width=\"300\" height=\"200\" /></span></p><p>&nbsp;</p><p><span style=\"text-decoration: underline;\">图片2</span></p><p><a href=\"http://www.baidu.com\"><strong><u>百度网站</u></strong></a></p>',
        '2023-03-28 18:08:52', '2023-03-28 18:10:25');
INSERT INTO `resource_richtext`
VALUES (6,
        '<p style=\"text-align: center;\"><strong>富文本5</strong></p><p style=\"text-align: center;\"><span style=\"color: #e03e2d;\"><em><strong>09:34:03 2023-03-14</strong></em></span></p><p>&nbsp; &nbsp;新闻内容段落1</p><p style=\"padding-left: 40px;\"><span style=\"font-family: \'arial black\', sans-serif;\">情况1</span></p><p style=\"padding-left: 40px;\"><span style=\"font-family: \'arial black\', sans-serif;\">情况2</span></p><p style=\"padding-left: 40px;\"><span style=\"font-family: \'arial black\', sans-serif;\">情况3</span></p><p>&nbsp;<span style=\"text-decoration: line-through;\"> &nbsp;新闻内容段落2</span></p><table style=\"border-collapse: collapse; width: 100%;\" border=\"1\"><tbody><tr><td style=\"width: 23.2395%;\">1</td><td style=\"width: 23.2395%;\">2</td><td style=\"width: 23.2395%;\">&nbsp;</td><td style=\"width: 23.2395%;\">&nbsp;</td></tr><tr><td style=\"width: 23.2395%;\">&nbsp;</td><td style=\"width: 23.2395%;\">3</td><td style=\"width: 23.2395%;\">4</td><td style=\"width: 23.2395%;\">4</td></tr><tr><td style=\"width: 23.2395%;\">&nbsp;</td><td style=\"width: 23.2395%;\">5</td><td style=\"width: 23.2395%;\">5</td><td style=\"width: 23.2395%;\">5</td></tr></tbody></table><p><span style=\"text-decoration: underline;\">&nbsp; &nbsp;新闻内容段落3</span></p><p><span style=\"text-decoration: underline;\"><img src=\"http://172.30.100.23:8060/resource/image/1.png\" alt=\"存储在nginx里面的图片\" width=\"300\" height=\"200\" /></span></p><p>&nbsp;</p><p><span style=\"text-decoration: underline;\">图片2</span></p><p><a href=\"http://www.baidu.com\"><strong><u>百度网站</u></strong></a></p>',
        '2023-03-28 18:08:57', '2023-03-28 18:10:17');
INSERT INTO `resource_richtext`
VALUES (7,
        '<p style=\"text-align: center;\"><strong>富文本6</strong></p><p style=\"text-align: center;\"><span style=\"color: #e03e2d;\"><em><strong>09:34:03 2023-03-14</strong></em></span></p><p>&nbsp; &nbsp;新闻内容段落1</p><p style=\"padding-left: 40px;\"><span style=\"font-family: \'arial black\', sans-serif;\">情况1</span></p><p style=\"padding-left: 40px;\"><span style=\"font-family: \'arial black\', sans-serif;\">情况2</span></p><p style=\"padding-left: 40px;\"><span style=\"font-family: \'arial black\', sans-serif;\">情况3</span></p><p>&nbsp;<span style=\"text-decoration: line-through;\"> &nbsp;新闻内容段落2</span></p><table style=\"border-collapse: collapse; width: 100%;\" border=\"1\"><tbody><tr><td style=\"width: 23.2395%;\">1</td><td style=\"width: 23.2395%;\">2</td><td style=\"width: 23.2395%;\">&nbsp;</td><td style=\"width: 23.2395%;\">&nbsp;</td></tr><tr><td style=\"width: 23.2395%;\">&nbsp;</td><td style=\"width: 23.2395%;\">3</td><td style=\"width: 23.2395%;\">4</td><td style=\"width: 23.2395%;\">4</td></tr><tr><td style=\"width: 23.2395%;\">&nbsp;</td><td style=\"width: 23.2395%;\">5</td><td style=\"width: 23.2395%;\">5</td><td style=\"width: 23.2395%;\">5</td></tr></tbody></table><p><span style=\"text-decoration: underline;\">&nbsp; &nbsp;新闻内容段落3</span></p><p><span style=\"text-decoration: underline;\"><img src=\"images/1.png\" alt=\"存储在nginx里面的图片\" width=\"300\" height=\"200\" /></span></p><p>&nbsp;</p><p><span style=\"text-decoration: underline;\">图片2</span></p><p><a href=\"http://www.baidu.com\"><strong><u>百度网站</u></strong></a></p>',
        '2023-03-28 18:09:04', '2023-03-28 18:09:04');
INSERT INTO `resource_richtext`
VALUES (8,
        '作业提交说明,\n要求文件名清晰明确，作业的提交形式包括：\n1）截图（*.jpg,*.png）\n2）word文档(*.doc,*.docx)\n3）simdroid工程文件（*.ibe）\n4)simdroid apps(*.app)',
        '2023-04-10 11:15:01', '2023-04-10 11:15:01');
INSERT INTO `resource_richtext`
VALUES (9, '这是一段学生提交的作业说明', '2023-04-10 14:48:01', '2023-04-10 14:49:37');
INSERT INTO `resource_richtext`
VALUES (10, '作业提交说明,\n要未使用', '2023-04-10 14:48:03', '2023-04-10 14:53:48');
INSERT INTO `resource_richtext`
VALUES (11, '作业提交说明,\n未使用', '2023-04-10 14:48:04', '2023-04-10 14:53:54');
INSERT INTO `resource_richtext`
VALUES (12, '作业提交 作业提交 这是一段富文本', '2023-04-11 15:12:10', '2023-04-11 15:12:10');
INSERT INTO `resource_richtext`
VALUES (18,
        '<p>要求提交的文件名清晰明确,作业的提交形式包括：<br />1）截图（*.jpg,*.png）；<br />2）文档(*.doc,*.docx,*.pdf)；<br />3）simdroid工程文件（*.ibe）；<br />4）simdroid apps(*.app)。</p>',
        '2023-04-11 15:52:36', '2023-04-11 15:52:36');
INSERT INTO `resource_richtext`
VALUES (19, '作业提交2 作业提交 这是一段富文本', '2023-04-11 16:17:55', '2023-04-11 16:17:55');
INSERT INTO `resource_richtext`
VALUES (20, '这是一段文字，用于作业提交', '2023-04-23 19:08:54', '2023-04-23 19:08:54');
INSERT INTO `resource_richtext`
VALUES (101, '<h2>1.1 实验目的</h2><ul>\r\n<li>熟悉输电线路铁塔组立抱杆四种工况；</li>\r\n<li>熟悉Simdroid仿真APP操作流程。</li>\r\n</ul>',
        '2023-04-11 16:20:15', '2023-04-11 16:20:15');
INSERT INTO `resource_richtext`
VALUES (102,
        '<h2>2.1 抱杆工程简介</h2><p>&nbsp; &nbsp; &nbsp; &nbsp;在电力建设施工过程中，抱杆是不可或缺的起重装置。它的主要功能是依托在建铁塔，将塔材提升至需要的位置进行施工安装，因此广泛应用于电力建设施工尤其是吊装施工并且是主要的受力单元。在施工过程中抱杆受力是否均衡合理，是否有超载使用，直接影响着施工的质量和安全。在传统的抱杆受力计算中，一般采用简化的方式，利用三角函数进行计算。然而这种简化方式只能粗略计算简单受力工况，当采用反向拉线等复杂工况或现场地形复杂时，不易分析各个结构的受力结果。对于现场施工人员，抱杆的力学计算所需的专业知识要求较高，繁琐并且无法保证可靠性，因此在施工中主要依靠工作经验来控制抱杆受力。</p>\r\n<p>&nbsp; &nbsp; &nbsp; &nbsp;近年来，随着特高压工程的不断兴建，抱杆的使用场景越来越恶劣，已多次出现由于抱杆弯折导致的工程安全事故，损失严重。综上所述，虽然抱杆在实际施工中应用比较广泛，但是它的安全性评估方法大部分集中在设计阶段。对于施工状态下，抱杆的安全性评估方法还没有形成系统化的研究成果或理论。因此，开发一款使用简单、计算准确的抱杆的施工过程安全评估APP，使得现场施工人员能够随时评估抱杆的状态，对于工程安全而言意义重大。</p>\r\n<p>&nbsp; &nbsp; &nbsp; &nbsp;针对传统计算方法复杂，不精确，效率低以及对现场施工人员知识背景要求过高以及无法面对多变的施工实况的问题，基于第三代仿真系统-Simdroid（云道智造具有完全自主知识产权）多物理场耦合仿真软件平台，采用图形交互式APP开发环境，建立了输电线路铁塔组立抱杆受力仿真APP，实现抱杆仿真APP开发与部署。</p>\r\n<p>&nbsp; &nbsp; &nbsp; &nbsp;抱杆仿真APP采用数值仿真手段对系统进行仿真计算，相对于传统方法，计算效率高，计算准确，可考虑情况完善。基于自主仿真数值平台进行仿真APP开发，模型实现了全参数化三维建模，终端用户通过手机或者其他电子设备访问仿真APP，驱动仿真APP计算并反馈计算结果。现场施工人员即可根据实际施工状况收集、调整、填写数据，经过APP计算得到抱杆的受力情况，及时对危险工况进行调整，为安全施工保驾护航。</p>\r\n<p>&nbsp;</p>\r\n<div style=\"text-align: center;\"><img src=\"/uploads/抱杆/1/图片1.png\" alt=\"抱杆图片\" width=\"508\" height=\"498\" />\r\n<p>图 1 抱杆系统构成</p></div>',
        '2023-04-11 16:20:15', '2023-04-11 16:20:15');
INSERT INTO `resource_richtext`
VALUES (103,
        '<h2><a name=\"_Toc124760266\"></a>3.1&nbsp;&nbsp; 抱杆塔身吊装工况</h2>\r\n<p>指大型输电铁塔的塔身吊装过程。结合抱杆，将塔身的各部件通过抱杆起吊，将铁塔塔身组装完成。</p>\r\n<p>&nbsp;</p>\r\n<div style=\"text-align: center;\"><img src=\"/uploads/抱杆/1/图片2.png\" alt=\"抱杆图片\" width=\"368\" height=\"319\" />\r\n<p>图 2抱杆塔身吊装</p></div>\r\n\r\n<p>&nbsp;</p>\r\n<h2><a name=\"_Toc124760267\"></a>3.2&nbsp;&nbsp; 抱杆横担吊装</h2>\r\n<p>横担是指组成输电铁塔的一部分，常见的横担是电线杆顶部横向固定的角铁，上面安装有瓷瓶来支撑架空电线，此处指大型铁塔横向的塔架。横担吊装可分为带人字杆和无人字杆。</p>\r\n<p>&nbsp;</p>\r\n<div style=\"text-align: center;\"><img src=\"/uploads/抱杆/1/图片3.png\" alt=\"抱杆图片\" width=\"350\" height=\"303\" />\r\n<p>图 3横担吊装</p></div>\r\n\r\n<p>&nbsp;</p>\r\n<h2><a name=\"_Toc124760268\"></a>3.3&nbsp;&nbsp; 抱杆屈曲工况</h2>\r\n<p>屈曲分析主要用于研究结构在特定载荷下的稳定性以及确定结构失稳的临界载荷。抱杆在吊装过程中很容易发生失稳现象，因此需要进行屈曲分析。</p>\r\n<p>&nbsp;</p>\r\n<div style=\"text-align: center;\"><img src=\"/uploads/抱杆/1/图片4.png\" alt=\"抱杆图片\" width=\"357\" height=\"293\" />\r\n<p>图 4抱杆屈曲分析</p></div>\r\n\r\n<p>&nbsp;</p>\r\n<h2><a name=\"_Toc124760269\"></a>3.4&nbsp;&nbsp; 抱杆起立工况</h2>\r\n<p>大型输电铁塔使用的抱杆一般比较高，大概20米左右，因此将其竖起来是一个大型的施工过程。在抱杆起立过程中要保证抱杆结构不会受到损坏。起立抱杆有三种方法，分别是：用人字抱杆整体起立抱杆；用汽车起立机整体起立抱杆；用塔腿起立抱杆。</p>\r\n<p>&nbsp;</p>\r\n<div style=\"text-align: center;\"><img src=\"/uploads/抱杆/1/图片5.png\" alt=\"抱杆图片\" width=\"339\" height=\"211\" />\r\n<p>图 5抱杆起立</p></div>\r\n\r\n<p>&nbsp;</p>',
        '2023-04-11 16:20:15', '2023-04-11 16:20:15');
INSERT INTO `resource_richtext`
VALUES (104,
        '<h2><a name=\"_Toc124760271\"></a>4.1 &nbsp;&nbsp;&nbsp;&nbsp; APP制作流程简介</h2>\r\n<p>抱杆仿真APP是将需要人工设置的建模、边界设置、求解设置和后处理分析步骤进行参数化和自动化，形成一个针对抱杆施工过程安全评估的仿真APP。可以对影响抱杆结构性能的参数，如抱杆尺寸参数、材料参数、绳索参数、工况参数等进行修改，快速得到拉线轴力、承托绳轴力、抱杆应力、抱杆屈曲模态、抱杆变形情况等的云图显示。</p>\r\n<p>&nbsp;</p>\r\n<p>&nbsp;</p>\r\n<div style=\"text-align: center;\"><img src=\"/uploads/抱杆/1/图片6.png\" alt=\"抱杆图片\" width=\"386\" height=\"367\" />\r\n<p>图 6仿真APP开发步骤</p></div>\r\n\r\n<p>&nbsp;</p>\r\n<h2><a name=\"_Toc124760272\"></a>4.2 &nbsp;&nbsp;&nbsp;&nbsp; APP体验流程简介</h2>\r\n<p>以&ldquo;22m-塔身吊装&rdquo;的操作流程为例，如下：</p>\r\n<ul>\r\n<li>在Simdroid 4.0平台的本地APP环境加载&ldquo;22m-塔身吊装&rdquo;，如下图所示：</li>\r\n</ul>\r\n<p>&nbsp;</p>\r\n<p>&nbsp;</p>\r\n<div style=\"text-align: center;\"><img src=\"/uploads/抱杆/1/图片7.png\" alt=\"抱杆图片\" width=\"567\" height=\"307\" />\r\n<p>图 7 Simdroid平台的本地APP环境加载</p></div>\r\n\r\n<p>&nbsp;</p>\r\n<p>（2）双击&ldquo;22m-塔身吊装&rdquo;图标，打开软件，如下图所示：</p>\r\n<p>&nbsp;</p>\r\n<div style=\"text-align: center;\"><img src=\"/uploads/抱杆/1/图片8.png\" alt=\"抱杆图片\" width=\"567\" height=\"301\" />\r\n<p>图 8 &ldquo;22m-塔身吊装&rdquo;APP</p></div>\r\n\r\n<p>&nbsp;</p>\r\n<p>（3）&ldquo;22m-塔身吊装&rdquo;软件启动后，即进入设置界面，此APP输入参数包含：抱杆参数、抱杆材料参数、绳索参数、工况参数四个部分。如下图所示，用户可以在此界面进行参数的更改。</p>\r\n<p>&nbsp;</p>\r\n<div style=\"text-align: center;\"><img src=\"/uploads/抱杆/1/图片9.png\" alt=\"抱杆图片\" width=\"500\" height=\"455\" />\r\n<p>图 9 &ldquo;22m-塔身吊装&rdquo;参数界面</p></div>\r\n\r\n<p>&nbsp;</p>\r\n<p>（4）页面菜单栏介绍：</p>\r\n<div style=\"text-align: center;\"><img src=\"/uploads/抱杆/1/图片10.png\" alt=\"抱杆图片\" width=\"567\" height=\"259\" />\r\n<p>图 10 &ldquo;22m-塔身吊装&rdquo;页面菜单介绍</p></div>\r\n\r\n<ul>\r\n<li><strong>生成几何</strong>：用户可以根据自己的工程应用场景设置封装的参数。设置完成后，点击&ldquo;生成几何&rdquo;命令，实现几何建模。</li>\r\n<li><strong>清除几何</strong>：如果生成的几何需要更改，用户可以清除掉刚生产的几何，重新设置生成几何；</li>\r\n<li><strong>生成网格</strong>：点击&ldquo;生成网格&rdquo;，按照设定的网格尺寸，生成网格模型；</li>\r\n<li><strong>计算</strong>：点击&ldquo;计算&rdquo;按钮，APP进入求解计算，用户等待即可，计算完成后，可以查看后处理云图；</li>\r\n<li><strong>显示日志：</strong>点击&ldquo;显示日志&rdquo;按钮，界面弹出该工程文件计算的所有流程及信息；</li>\r\n<li><strong>退出</strong>：APP应用结束后，点击&ldquo;退出&rdquo;，退出APP界面。</li>\r\n</ul>\r\n<p>（5）显示界面介绍</p>\r\n<p>&nbsp;</p>\r\n<div style=\"text-align: center;\"><img src=\"/uploads/抱杆/1/图片11.png\" alt=\"抱杆图片\" width=\"567\" height=\"268\" />\r\n<p>图 11 &ldquo;22m-塔身吊装&rdquo;显示界面介绍</p></div>\r\n\r\n<p>显示界面包含：角度计算器、抱杆几何模型、网格、拉线1轴力、拉线2轴力、拉线3轴力、拉线4轴力、承托绳轴力云图、抱杆应力云图、结果汇总。在以此执行&ldquo;生成几何&rdquo;、&ldquo;生成网格&rdquo;、&ldquo;计算&rdquo;命令后，可以查看显示界面所有窗口。</p>\r\n<p>（6）其他APP的体验流程相同，在此不赘述。</p>\r\n<p>&nbsp;</p>\r\n<p>&nbsp;</p>',
        '2023-04-11 16:20:15', '2023-04-11 16:20:15');
INSERT INTO `resource_richtext`
VALUES (110, '<h2>1.1 实验目的</h2><ul>\r\n<li>熟练掌握Simdroid梁、壳单元相关的各种工具。</li>\r\n<li>熟悉掌握Simdroid完整仿真算例计算流程。</li>\r\n</ul>',
        '2023-04-11 16:20:15', '2023-04-11 16:20:15');
INSERT INTO `resource_richtext`
VALUES (111,
        '<h2>2.1 实验要求</h2><ul>\r\n<li>学习Simdroid工具栏各种工具，包含截面定义、赋予截面属性；</li>\r\n<li>学习Simdroid完整计算流程，包含模型建立、材料属性、加载和约束设置、后处理等。</li>\r\n</ul>',
        '2023-04-11 16:20:15', '2023-04-11 16:20:15');
INSERT INTO `resource_richtext`
VALUES (112,
        '<ul>\r\n<li>\r\n<h2><a name=\"_Toc122939793\"></a>3.1 实验步骤</h2>\r\n<ul>\r\n<li>新建分析</li>\r\n<li>几何建模</li>\r\n<li>创建截面</li>\r\n<li>创建材料</li>\r\n<li>创建网格</li>\r\n<li>设置截面属性</li>\r\n<li>施加边界和连接</li>\r\n<li>设置分析和加载</li>\r\n<li>计算和后处理</li>\r\n</ul>\r\n<h2><a name=\"_Toc122939794\"></a>3.2 要点总结</h2>\r\n<ul>\r\n<li>创建截面、赋予截面属性、指定截面方向是一系列的连续操作，缺一不可。</li>\r\n<li>掌握各类梁单元截面形式的含义。</li>\r\n</ul>\r\n</li>\r\n</ul>',
        '2023-04-11 16:20:15', '2023-04-11 16:20:15');
INSERT INTO `resource_richtext`
VALUES (113, '<p>无</p>', '2023-04-11 16:20:15', '2023-04-11 16:20:15');
INSERT INTO `resource_richtext`
VALUES (114,
        '<h2><a name=\"_Toc122939797\"></a>5.1 案例介绍</h2>\r\n<p>本案例建立了梁支架模型，详细介绍了结构通用静力仿真分析过程。</p>\r\n<h2><a name=\"_Toc122939798\"></a>5.2 模型说明</h2>\r\n<div style=\"text-align: center;\"><img src=\"/uploads/抱杆/2/图片1.png\" alt=\"抱杆图片\" width=\"316\" height=\"187\" />\r\n<p>图 1支架模型</p></div>\r\n\r\n<p>几何：梁支架高度为1m，截面为圆形，半径0.01m。如图所示，单块板长为1m，板厚为0.005m。</p>\r\n<p>材料：所有几何材料均为铝材，密度：2700kg/m^3，弹性模量：7E10Pa，泊松比：0.3。</p>\r\n<p>边界与载荷：固定约束所有梁支架的下端点。中间板施加-Y向的压力250Pa。</p>\r\n<p>&nbsp;</p>\r\n<h2><a name=\"_Toc122939799\"></a>5.3 仿真及开发过程</h2>\r\n<h3><a name=\"_Toc122939800\"></a>5.3.1 &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 新建分析</h3>\r\n<ul>\r\n<li>启动Simdroid。</li>\r\n<li>在【新建】对话框，分析类型选择【结构分析】，维度选择【三维】，名称输入&ldquo;梁支架通用静力仿真分析&rdquo;，选择工作路径，单击【确定】。</li>\r\n</ul>\r\n<div style=\"text-align: center;\"><img src=\"/uploads/抱杆/2/图片2.png\" alt=\"抱杆图片\" width=\"692\" height=\"497\" />\r\n<p>图 2新建项目</p></div>\r\n\r\n<h3><a name=\"_Toc122939801\"></a>5.3.2 &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 几何建模</h3>\r\n<ul>\r\n<li>选择【三维建模】&gt;【创建2D草图】，选择XY平面，进入草图环境。</li>\r\n<li>建立梁支架草图，选择&ldquo;线&rdquo;，在草图视图区创建如下图所示的模型，单击【完成草图】。</li>\r\n</ul>\r\n<p>&nbsp;</p>\r\n<div style=\"text-align: center;\"><img src=\"/uploads/抱杆/2/图片3.png\" alt=\"抱杆图片\" width=\"1269\" height=\"588\" />\r\n<p>图 3单根梁草图</p></div>\r\n\r\n<ul>\r\n<li>选中模型树上刚创建的&ldquo;草图&rdquo;，单击【三维建模】&gt;【生成线】，将草图变成线体。</li>\r\n<li>选中模型树上刚创建的&ldquo;线&rdquo;，单击【三维建模】&gt;【阵列】，输入下图界面中参数，单击&ldquo;&radic;&rdquo;，模型树上生成&ldquo;阵列&rdquo;节点，将其重命名为&ldquo;梁&rdquo;。</li>\r\n</ul>\r\n<p>&nbsp;</p>\r\n<div style=\"text-align: center;\"><img src=\"/uploads/抱杆/2/图片4.png\" alt=\"抱杆图片\" width=\"297\" height=\"251\" /><p>&nbsp;</p><img src=\"/uploads/抱杆/2/图片5.png\" alt=\"抱杆图片\" width=\"1269\" height=\"671\" />\r\n<p>图 4阵列生成梁支架</p></div>\r\n\r\n<ul>\r\n<li>再次选择【三维建模】&gt;【创建2D草图】，选择YZ平面，进入草图环境。</li>\r\n<li>建立支架上薄板的草图，选择&ldquo;线&rdquo;，在草图视图区创建如下图所示的模型，单击【完成草图】。</li>\r\n</ul>\r\n<p>&nbsp;</p>\r\n<div style=\"text-align: center;\"><img src=\"/uploads/抱杆/2/图片6.png\" alt=\"抱杆图片\" width=\"1269\" height=\"605\" />\r\n<p>图 5薄板草图</p></div>\r\n\r\n<ul>\r\n<li>选中模型树上刚创建的&ldquo;草图1&rdquo;，单击【三维建模】&gt;【拉伸】，在任务面板修改&ldquo;正向长度&rdquo;为1m。单击&ldquo;&radic;&rdquo;，模型树上生成&ldquo;拉伸&rdquo;节点，将其重命名为&ldquo;薄板&rdquo;。</li>\r\n</ul>\r\n<p>&nbsp;</p>\r\n<div style=\"text-align: center;\"><img src=\"/uploads/抱杆/2/图片7.png\" alt=\"抱杆图片\" width=\"1269\" height=\"654\" />\r\n<p>图 6拉伸生成薄板</p></div>\r\n\r\n<p>&nbsp;</p>\r\n<h3><a name=\"_Toc122939802\"></a>5.3.3 &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 创建截面</h3>\r\n<ul>\r\n<li>单击【三维建模】&gt;【截面】，&ldquo;类型&rdquo;选择为&ldquo;壳&rdquo;，&ldquo;名称&rdquo;改为&ldquo;板截面&rdquo;，&ldquo;厚度(t)&rdquo;输入&ldquo;005m&rdquo;。单击&ldquo;&radic;&rdquo;完成壳截面的设置。</li>\r\n</ul>\r\n<p>&nbsp;</p>\r\n<div style=\"text-align: center;\"><img src=\"/uploads/抱杆/2/图片8.png\" alt=\"抱杆图片\" width=\"296\" height=\"292\" />\r\n<p>图 7设置壳截面</p></div>\r\n\r\n<ul>\r\n<li>单击【三维建模】&gt;【截面】，&ldquo;类型&rdquo;选择为&ldquo;圆形&rdquo;，&ldquo;名称&rdquo;改为&ldquo;梁截面&rdquo;，&ldquo;半径(r)&rdquo;输入&ldquo;01m&rdquo;。单击&ldquo;&radic;&rdquo;完成梁截面的设置。</li>\r\n</ul>\r\n<div style=\"text-align: center;\"><img src=\"/uploads/抱杆/2/图片9.png\" alt=\"抱杆图片\" width=\"281\" height=\"339\" />\r\n<p>图 8设置梁截面</p></div>\r\n\r\n<h3><a name=\"_Toc122939803\"></a>5.3.4 &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 创建材料</h3>\r\n<ul>\r\n<li>选择【材料属性】&gt;【创建材料】，修改&ldquo;标签&rdquo;为&ldquo;板材料&rdquo;。在【目标】中修改&ldquo;拾取类型&rdquo;为&ldquo;面(壳)&rdquo;，视图区点选所有薄板，在【属性】中双击选择【基本属性】&gt;【基本材料】&gt;【密度】和【固体属性】&gt;【线弹性】&gt;【各向同性】，设置密度为2700kg/m^3、弹性模量为7e10Pa、泊松比为3，单击【&radic;】，完成薄板材料的施加。</li>\r\n</ul>\r\n<div style=\"text-align: center;\"><img src=\"/uploads/抱杆/2/图片10.png\" alt=\"抱杆图片\" width=\"279\" height=\"734\" />\r\n<p>图 9创建薄板材料</p></div>\r\n\r\n<ul>\r\n<li>同样的步骤设置梁支架的材料参数。</li>\r\n</ul>\r\n<h3><a name=\"_Toc122939804\"></a>5.3.5 &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 创建网格</h3>\r\n<ul>\r\n<li>选中【几何模型】&gt;【薄板】节点，选择【网格划分】&gt;【单体剖分】，设置&ldquo;最大尺寸&rdquo;为1m，&ldquo;单元类型&rdquo;选择为&ldquo;TMQ(TMT)&rdquo;，单击&ldquo;&radic;&rdquo;，完成薄板的网格剖分。</li>\r\n</ul>\r\n<div style=\"text-align: center;\"><img src=\"/uploads/抱杆/2/图片11.png\" alt=\"抱杆图片\" width=\"302\" height=\"307\" /><p>&nbsp;</p><img src=\"/uploads/抱杆/2/图片12.png\" alt=\"抱杆图片\" width=\"691\" height=\"333\" />\r\n<p>图 10薄板网格剖分</p></div>\r\n\r\n<ul>\r\n<li>选中【几何模型】&gt;【梁】节点，选择【网格划分】&gt;【单体剖分】，设置&ldquo;最大尺寸&rdquo;为1m，&ldquo;单元类型&rdquo;选择为&ldquo;Beam2&rdquo;，单击&ldquo;&radic;&rdquo;，完成梁支架的网格剖分。</li>\r\n</ul>\r\n<div style=\"text-align: center;\"><img src=\"/uploads/抱杆/2/图片13.png\" alt=\"抱杆图片\" width=\"301\" height=\"310\" /><p>&nbsp;</p><img src=\"/uploads/抱杆/2/图片14.png\" alt=\"抱杆图片\" width=\"691\" height=\"334\" />\r\n<p>图 11梁支架网格剖分</p></div>\r\n\r\n<p>&nbsp;</p>\r\n<h3><a name=\"_Toc122939805\"></a>5.3.6 &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 设置截面属性</h3>\r\n<ul>\r\n<li>选择【分析】&gt;【指定截面】，将&ldquo;选择截面&rdquo;选择为&ldquo;板截面&rdquo;，&ldquo;拾取类型&rdquo;选择为&ldquo;壳&rdquo;，视图区选择三个薄板，单击&ldquo;&radic;&rdquo;完成薄板的截面设置。</li>\r\n</ul>\r\n<p>&nbsp;</p>\r\n<div style=\"text-align: center;\"><img src=\"/uploads/抱杆/2/图片15.png\" alt=\"抱杆图片\" width=\"1269\" height=\"651\" />\r\n<p>图 12指定薄板截面</p></div>\r\n\r\n<p>&nbsp;</p>\r\n<ul>\r\n<li>选择【分析】&gt;【指定截面】，将&ldquo;选择截面&rdquo;选择为&ldquo;梁截面&rdquo;，&ldquo;拾取类型&rdquo;选择为&ldquo;梁/杆&rdquo;，视图区选择八个梁支架，单击&ldquo;&radic;&rdquo;完成梁支架的截面设置。</li>\r\n</ul>\r\n<p>&nbsp;</p>\r\n<div style=\"text-align: center;\"><img src=\"/uploads/抱杆/2/图片16.png\" alt=\"抱杆图片\" width=\"1269\" height=\"659\" />\r\n<p>图 13指定梁支架截面</p></div>\r\n\r\n<p>&nbsp;</p>\r\n<ul>\r\n<li>选择【分析】&gt;【指定截面方向】，视图区选择八个梁支架，方向为默认，单击&ldquo;&radic;&rdquo;完成梁支架的截面方向设置。</li>\r\n</ul>\r\n<p>&nbsp;</p>\r\n<p>&nbsp;</p>\r\n<div style=\"text-align: center;\"><img src=\"/uploads/抱杆/2/图片17.png\" alt=\"抱杆图片\" width=\"1269\" height=\"645\" />\r\n<p>图 14指定梁支架截面方向</p></div>\r\n\r\n<p>&nbsp;</p>\r\n<h3><a name=\"_Toc122939806\"></a>5.3.7 &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 施加边界和连接</h3>\r\n<h4>5.3.7.1 &nbsp;设置耦合连接</h4>\r\n<ul>\r\n<li>选择【分析】&gt;【耦合连接】打开任务面板（单击【耦合连接】前先选择模型树上【几何模型】&gt;【薄板】节点，右键单击&ldquo;切换可见性 空格&rdquo;来隐藏【薄板】的几何模型，或通过空格键来隐藏/显示）。</li>\r\n<li>【目标-主边界-点】窗口默认处于激活状态，视图区选择如图所示点为耦合连接的主边界点。</li>\r\n</ul>\r\n<table>\r\n<tbody>\r\n<tr>\r\n<td width=\"156\">&nbsp;</td>\r\n</tr>\r\n<tr>\r\n<td>&nbsp;</td>\r\n<td>&nbsp;</td>\r\n</tr>\r\n</tbody>\r\n</table>\r\n<p>&nbsp;</p>\r\n<p>&nbsp;</p>\r\n<div style=\"text-align: center;\"><img src=\"/uploads/抱杆/2/图片18.png\" alt=\"抱杆图片\" width=\"1269\" height=\"656\" />\r\n<p>图 15选择主边界点</p></div>\r\n\r\n<ul>\r\n<li>鼠标单击【目标-从边界-体/面/边/点】窗口来激活从边界拾取窗口，&ldquo;拾取类型&rdquo;选择&ldquo;点&rdquo;。</li>\r\n<li>单击视图区工具栏的【切换几何模型可见性】图标，切换视图区的几何模型显示，当前视图区显示为薄板模型。</li>\r\n<li>选择薄板上与主边界点对应的薄板上的点。单击&ldquo;&radic;&rdquo;，完成第一对耦合连接的设置。</li>\r\n</ul>\r\n<table>\r\n<tbody>\r\n<tr>\r\n<td width=\"161\">&nbsp;</td>\r\n</tr>\r\n<tr>\r\n<td>&nbsp;</td>\r\n<td>&nbsp;</td>\r\n</tr>\r\n</tbody>\r\n</table>\r\n<p>&nbsp;</p>\r\n<p>&nbsp;</p>\r\n<div style=\"text-align: center;\"><img src=\"/uploads/抱杆/2/图片19.png\" alt=\"抱杆图片\" width=\"1269\" height=\"654\" />\r\n<p>图 16选择从边界点</p></div>\r\n\r\n<ul>\r\n<li>同样的步骤设置剩余7对耦合连接。</li>\r\n</ul>\r\n<h4>5.3.7.2 &nbsp;设置约束边界</h4>\r\n<ul>\r\n<li>单击视图区工具栏的【显示全部几何模型】图标，显示所有的几何模型。</li>\r\n<li>选择【分析】&gt;【常规约束】，在任务面板将&ldquo;拾取类型&rdquo;选择为&ldquo;点&rdquo;，依次选择视图区梁支架的8个下端点。勾选约束6个自由度，单击&ldquo;&radic;&rdquo;完成约束边界的设置。</li>\r\n</ul>\r\n<p>&nbsp;</p>\r\n<div style=\"text-align: center;\"><img src=\"/uploads/抱杆/2/图片20.png\" alt=\"抱杆图片\" width=\"1269\" height=\"655\" />\r\n<p>图 17施加常规约束</p></div>\r\n\r\n<h3><a name=\"_Toc122939807\"></a>5.3.8 &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 设置分析和加载</h3>\r\n<ul>\r\n<li>选择【分析】&gt;【创建分析】，选择【通用静力分析】，单击【添加】单击&ldquo;&radic;&rdquo;。</li>\r\n</ul>\r\n<div style=\"text-align: center;\"><img src=\"/uploads/抱杆/2/图片21.png\" alt=\"抱杆图片\" width=\"374\" height=\"580\" />\r\n<p>图 18创建分析</p></div>\r\n\r\n<ul>\r\n<li>选择【分析】&gt;【压力】，拾取中间的薄板，输入【荷载设置】&gt;【数值】为250Pa，单击&ldquo;&radic;&rdquo;完成载荷的设置。</li>\r\n</ul>\r\n<p>&nbsp;</p>\r\n<div style=\"text-align: center;\"><img src=\"/uploads/抱杆/2/图片22.png\" alt=\"抱杆图片\" width=\"692\" height=\"444\" />\r\n<p>图 19施加载荷</p></div>\r\n\r\n<h3><a name=\"_Toc122939808\"></a>5.3.9 &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 计算和后处理</h3>\r\n<ul>\r\n<li>选择【分析】&gt;【计算】，开始计算，在报告浏览器中可查看求解过程。</li>\r\n<li>选中模型树上【结果展示】，在工具栏中选择【云图】，物理场选择【Disp】，单击&ldquo;&radic;&rdquo;。可以查看位移变形云图。</li>\r\n</ul>\r\n<p>&nbsp;</p>\r\n<div style=\"text-align: center;\"><img src=\"/uploads/抱杆/2/图片23.png\" alt=\"抱杆图片\" width=\"691\" height=\"295\" />\r\n<p>图 20查看位移云图</p></div>\r\n\r\n<p>&nbsp;</p>\r\n<p>&nbsp;</p>\r\n<ul>\r\n<li>同样的步骤可以查看应力云图。</li>\r\n</ul>\r\n<div style=\"text-align: center;\"><img src=\"/uploads/抱杆/2/图片24.png\" alt=\"抱杆图片\" width=\"691\" height=\"298\" />\r\n<p>图 21 查看位移云图</p></div>\r\n',
        '2023-04-11 16:20:15', '2023-04-11 16:20:15');
INSERT INTO `resource_richtext`
VALUES (115, 'some message from class for lesson work', '2023-04-23 19:49:49', '2023-04-23 19:49:49');
INSERT INTO `resource_richtext`
VALUES (116, '', '2023-04-23 20:04:57', '2023-04-23 20:04:57');
INSERT INTO `resource_richtext`
VALUES (117, '', '2023-04-23 20:06:24', '2023-04-23 20:06:24');
INSERT INTO `resource_richtext`
VALUES (118, '一段作业提交message1111', '2023-04-23 20:12:01', '2023-04-23 20:12:01');
INSERT INTO `resource_richtext`
VALUES (119, '一段作业提交message4444444444441', '2023-04-23 20:12:38', '2023-04-23 20:12:38');
INSERT INTO `resource_richtext`
VALUES (120, '一段作业提交message4444444444441', '2023-04-23 20:16:31', '2023-04-23 20:16:31');
INSERT INTO `resource_richtext`
VALUES (127, '老师好...a1 32222', '2023-04-28 15:51:34', '2023-04-28 15:51:34');
INSERT INTO `resource_richtext`
VALUES (129, '老师您好，希望重点解释第一节内容。第二节内容比较简单。因此希望第一节进行问题解答', '2023-04-28 16:02:09', '2023-04-28 16:02:09');
INSERT INTO `resource_richtext`
VALUES (130, '此为实验二ibe文件', '2023-04-28 16:55:32', '2023-04-28 16:55:32');

-- ----------------------------
-- Table structure for school
-- ----------------------------
DROP TABLE IF EXISTS `school`;
CREATE TABLE `school`
(
    `id`         int(11)                                                      NOT NULL AUTO_INCREMENT COMMENT '主键',
    `name`       varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '学校名字',
    `created_at` timestamp(0)                                                 NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
    `updated_at` timestamp(0)                                                 NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT = '学校表'
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of school
-- ----------------------------

-- ----------------------------
-- Table structure for score
-- ----------------------------
DROP TABLE IF EXISTS `score`;
CREATE TABLE `score`
(
    `id`         int(11)      NOT NULL AUTO_INCREMENT COMMENT '主键',
    `student_id` int(11)      NOT NULL DEFAULT 0 COMMENT '学生id',
    `course_id`  int(11)      NOT NULL DEFAULT 0 COMMENT '课程id',
    `work_id`    int(11)      NOT NULL DEFAULT 0 COMMENT '课堂作业id',
    `mark_id`    int(11)      NOT NULL DEFAULT 0 COMMENT '打分id',
    `score`      int(11)      NOT NULL DEFAULT 0 COMMENT '课堂得分',
    `created_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
    `updated_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT = '课堂成绩关系表'
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of score
-- ----------------------------

-- ----------------------------
-- Table structure for score_item
-- ----------------------------
DROP TABLE IF EXISTS `score_item`;
CREATE TABLE `score_item`
(
    `id`         int(11)      NOT NULL AUTO_INCREMENT COMMENT '主键',
    `item`       int(11)      NOT NULL DEFAULT 0 COMMENT '打分项',
    `created_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
    `updated_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT = '打分项表'
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of score_item
-- ----------------------------

-- ----------------------------
-- Table structure for score_subitem
-- ----------------------------
DROP TABLE IF EXISTS `score_subitem`;
CREATE TABLE `score_subitem`
(
    `id`         int(11)      NOT NULL AUTO_INCREMENT COMMENT '主键',
    `subitem`    int(11)      NOT NULL DEFAULT 0 COMMENT '打分子项',
    `created_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
    `updated_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT = '打分子项表'
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of score_subitem
-- ----------------------------

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
-- Table structure for teacher
-- ----------------------------
DROP TABLE IF EXISTS `teacher`;
CREATE TABLE `teacher`
(
    `id`              int(11)                                                      NOT NULL AUTO_INCREMENT COMMENT '主键',
    `login_name`      varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '登录id，唯一',
    `name`            varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '实际姓名',
    `introduce`       text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci        NULL COMMENT '介绍',
    `password`        varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '密码',
    `organization_id` int(11)                                                      NOT NULL DEFAULT 0 COMMENT '组织id，包括学校院系班级',
    `phone_number`    varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '电话号码',
    `is_deleted`      tinyint(4)                                                   NOT NULL DEFAULT 1 COMMENT '是否删除 1-否 2-是',
    `created_at`      timestamp(0)                                                 NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
    `updated_at`      timestamp(0)                                                 NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 20
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT = '教师表，编写或者讲解课程的作者'
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of teacher
-- ----------------------------
INSERT INTO `teacher`
VALUES (12, '云道教育', '云道教育', NULL, '', 0, '', 1, '2023-04-25 16:03:03', '2023-04-25 16:03:39');
INSERT INTO `teacher`
VALUES (13, '田靖军', '田靖军', NULL, '', 0, '', 1, '2023-04-25 16:03:07', '2023-04-25 16:03:32');
INSERT INTO `teacher`
VALUES (14, '刘娇', '刘娇', NULL, '', 0, '', 1, '2023-04-25 16:03:07', '2023-04-25 16:03:32');


-- ----------------------------
-- Table structure for teacher_class
-- ----------------------------
DROP TABLE IF EXISTS `teacher_class`;
CREATE TABLE `teacher_class`
(
    `id`              int(11)      NOT NULL AUTO_INCREMENT COMMENT '主键',
    `organization_id` int(11)      NOT NULL DEFAULT 0 COMMENT '组织id，包括学校院系班级',
    `is_deleted`      tinyint(4)   NOT NULL DEFAULT 1 COMMENT '是否删除 1-否 2-是',
    `created_at`      timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
    `updated_at`      timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 20
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT = '教师表所教授的班级'
  ROW_FORMAT = Dynamic;


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

-- ----------------------------
-- Table structure for work_commit
-- ----------------------------
DROP TABLE IF EXISTS `work_commit`;
CREATE TABLE `work_commit`
(
    `id`         int(11)      NOT NULL AUTO_INCREMENT COMMENT '主键',
    `lesson_id`  int(11)      NOT NULL DEFAULT 0 COMMENT '所属课堂id',
    `user_id`    int(11)      NOT NULL DEFAULT 0 COMMENT '所属用户id',
    `created_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
    `updated_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 27
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT = '作业提交表'
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of work_commit
-- ----------------------------
INSERT INTO `work_commit`
VALUES (8, 1, 3, '2023-04-11 16:20:13', '2023-04-23 18:58:06');
INSERT INTO `work_commit`
VALUES (9, 1, 0, '2023-04-23 19:32:38', '2023-04-23 19:32:38');
INSERT INTO `work_commit`
VALUES (10, 1, 3, '2023-04-23 19:34:45', '2023-04-23 19:34:45');
INSERT INTO `work_commit`
VALUES (11, 1, 3, '2023-04-23 19:49:13', '2023-04-23 19:49:13');
INSERT INTO `work_commit`
VALUES (14, 3, 0, '2023-04-23 20:11:27', '2023-04-23 20:11:27');
INSERT INTO `work_commit`
VALUES (15, 4, 0, '2023-04-23 20:12:35', '2023-04-23 20:12:35');
INSERT INTO `work_commit`
VALUES (16, 5, 3, '2023-04-23 20:16:29', '2023-04-23 20:16:29');
INSERT INTO `work_commit`
VALUES (23, 101, 3, '2023-04-28 15:51:34', '2023-04-28 15:51:34');
INSERT INTO `work_commit`
VALUES (25, 101, 1, '2023-04-28 16:02:09', '2023-04-28 16:02:09');
INSERT INTO `work_commit`
VALUES (26, 101, 4, '2023-04-28 16:55:32', '2023-04-28 16:55:32');

-- ----------------------------
-- Table structure for work_commit_resource
-- ----------------------------
DROP TABLE IF EXISTS `work_commit_resource`;
CREATE TABLE `work_commit_resource`
(
    `id`             int(11)      NOT NULL AUTO_INCREMENT COMMENT '主键',
    `work_commit_id` int(11)      NOT NULL DEFAULT 0 COMMENT '作业提交id',
    `resource_id`    int(11)      NOT NULL DEFAULT 0 COMMENT '资源id',
    `type`           tinyint(4)   NOT NULL DEFAULT 0 COMMENT '资源所属类型，如，提交作业图片，文件，留言',
    `can_modify`     tinyint(4)   NOT NULL DEFAULT 1 COMMENT '是否可以再次修改，1 可以，2 不可以',
    `created_at`     timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
    `updated_at`     timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 54
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT = '作业提交-资源表'
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of work_commit_resource
-- ----------------------------
INSERT INTO `work_commit_resource`
VALUES (10, 6, 12, 44, 0, '2023-04-11 15:12:40', '2023-04-11 15:12:40');
INSERT INTO `work_commit_resource`
VALUES (11, 7, 1, 43, 0, '2023-04-11 16:17:41', '2023-04-11 16:17:41');
INSERT INTO `work_commit_resource`
VALUES (12, 7, 3, 43, 0, '2023-04-11 16:17:44', '2023-04-11 16:17:44');
INSERT INTO `work_commit_resource`
VALUES (13, 7, 30, 44, 0, '2023-04-11 16:17:58', '2023-04-11 16:17:58');
INSERT INTO `work_commit_resource`
VALUES (14, 8, 1, 43, 0, '2023-04-11 16:20:14', '2023-04-11 16:20:14');
INSERT INTO `work_commit_resource`
VALUES (15, 8, 3, 43, 0, '2023-04-11 16:20:14', '2023-04-11 16:20:14');
INSERT INTO `work_commit_resource`
VALUES (16, 8, 31, 44, 0, '2023-04-11 16:20:15', '2023-04-11 16:20:15');
INSERT INTO `work_commit_resource`
VALUES (17, 11, 25, 43, 0, '2023-04-23 19:49:16', '2023-04-23 19:49:16');
INSERT INTO `work_commit_resource`
VALUES (18, 11, 24, 43, 0, '2023-04-23 19:49:18', '2023-04-23 19:49:18');
INSERT INTO `work_commit_resource`
VALUES (19, 11, 234, 44, 0, '2023-04-23 19:49:49', '2023-04-23 19:49:49');
INSERT INTO `work_commit_resource`
VALUES (20, 12, 235, 44, 0, '2023-04-23 20:04:57', '2023-04-23 20:04:57');
INSERT INTO `work_commit_resource`
VALUES (21, 13, 236, 44, 0, '2023-04-23 20:06:24', '2023-04-23 20:06:24');
INSERT INTO `work_commit_resource`
VALUES (22, 14, 24, 43, 0, '2023-04-23 20:11:35', '2023-04-23 20:11:35');
INSERT INTO `work_commit_resource`
VALUES (23, 14, 25, 43, 0, '2023-04-23 20:11:37', '2023-04-23 20:11:37');
INSERT INTO `work_commit_resource`
VALUES (24, 14, 237, 44, 0, '2023-04-23 20:12:07', '2023-04-23 20:12:07');
INSERT INTO `work_commit_resource`
VALUES (25, 15, 24, 43, 0, '2023-04-23 20:12:36', '2023-04-23 20:12:36');
INSERT INTO `work_commit_resource`
VALUES (26, 15, 25, 43, 0, '2023-04-23 20:12:38', '2023-04-23 20:12:38');
INSERT INTO `work_commit_resource`
VALUES (27, 15, 238, 44, 0, '2023-04-23 20:12:38', '2023-04-23 20:12:38');
INSERT INTO `work_commit_resource`
VALUES (28, 16, 24, 43, 0, '2023-04-23 20:16:29', '2023-04-23 20:16:29');
INSERT INTO `work_commit_resource`
VALUES (29, 16, 25, 43, 0, '2023-04-23 20:16:31', '2023-04-23 20:16:31');
INSERT INTO `work_commit_resource`
VALUES (30, 16, 239, 44, 0, '2023-04-23 20:16:31', '2023-04-23 20:16:31');
INSERT INTO `work_commit_resource`
VALUES (45, 23, 513, 43, 0, '2023-04-28 15:51:34', '2023-04-28 15:51:34');
INSERT INTO `work_commit_resource`
VALUES (46, 23, 515, 43, 0, '2023-04-28 15:51:34', '2023-04-28 15:51:34');
INSERT INTO `work_commit_resource`
VALUES (47, 23, 517, 44, 0, '2023-04-28 15:51:34', '2023-04-28 15:51:34');
INSERT INTO `work_commit_resource`
VALUES (50, 25, 510, 43, 0, '2023-04-28 16:02:09', '2023-04-28 16:02:09');
INSERT INTO `work_commit_resource`
VALUES (51, 25, 519, 44, 0, '2023-04-28 16:02:09', '2023-04-28 16:02:09');
INSERT INTO `work_commit_resource`
VALUES (52, 26, 520, 43, 0, '2023-04-28 16:55:32', '2023-04-28 16:55:32');
INSERT INTO `work_commit_resource`
VALUES (53, 26, 521, 44, 0, '2023-04-28 16:55:32', '2023-04-28 16:55:32');

SET FOREIGN_KEY_CHECKS = 1;
