select version();
select database();

show databases;
create database ibe_edu;
use ibe_edu;


DROP TABLE IF EXISTS user;
# 1 用户表
CREATE TABLE `user` (
`id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
`role` tinyint(4)   NOT NULL DEFAULT '1' COMMENT '账户类型：1学生，2教师，3管理员',
`name` varchar(50)  NOT NULL DEFAULT '' COMMENT '姓名',
`password` varchar(50)  NOT NULL DEFAULT '' COMMENT '密码',
`creator_id` int(11)  NOT NULL DEFAULT '0' COMMENT '创建人id',
`is_deleted` tinyint(4)  NOT NULL DEFAULT '1' COMMENT '是否删除 1-否 2-是',
`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '用户表';

# 2 学校表 院系表 班级表 学校组织表
 DROP TABLE IF EXISTS `school`;
#insert into school (user_id, product_id) values (1,1),(1,2);
#insert into school (user_id, product_id) values (2,1);
#select *from school;
CREATE TABLE `school` (
`id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
`name` varchar(50)  NOT NULL DEFAULT '' COMMENT  '学校名字',
`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '学校表';

 
DROP TABLE IF EXISTS `department`;

CREATE TABLE `department` (
`id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
`name` varchar(50)  NOT NULL DEFAULT '' COMMENT  '院系名字',
`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '院系表';


DROP TABLE IF EXISTS `class`;

CREATE TABLE `class` (
`id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
`name` varchar(50)  NOT NULL DEFAULT '' COMMENT  '班级名字',
`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '班级表';


DROP TABLE IF EXISTS `organization`;

CREATE TABLE `organization` (
`id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
`school_id` int(11)  NOT NULL DEFAULT '0' COMMENT 'school id',
`department_id` int(11)  NOT NULL DEFAULT '0' COMMENT 'department id',
`class_id` int(11)  NOT NULL DEFAULT '0' COMMENT 'class id',
`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '学校组织表';



# 3 基础资源表
DROP TABLE IF EXISTS `resource`;

CREATE TABLE `resource` (
`id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
`title` varchar(100)  NOT NULL DEFAULT '0' COMMENT '资源标题',
`desc` TINYTEXT   COMMENT '资源描述',
`type` tinyint(4)  NOT NULL DEFAULT '0' COMMENT '资源类型，1图片，2 word文档，3 pdf文档，4 apps，5 ibe，6留言7 富文本',
`content` TINYTEXT   COMMENT '资源内容，存储路径、url及少量的文字',
`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '存储系统中用到的基础资源';


DROP TABLE IF EXISTS `richtext`;
CREATE TABLE `richtext` (
`id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
`content` TEXT   COMMENT '基础资源富文本内容',
`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '存储系统中用到的富文本资源';

# 4 课程表
DROP TABLE IF EXISTS `course`;
CREATE TABLE `course` (
`id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
`title` varchar(100)  NOT NULL DEFAULT '0' COMMENT '课程标题',
`desc` TINYTEXT   COMMENT '课程描述',
`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '课程表';


DROP TABLE IF EXISTS `course_resource`;
CREATE TABLE `course_resource` (
`id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
`course_id` int(11)  NOT NULL DEFAULT '0' COMMENT '课程id',
`resource_id` int(11)  NOT NULL DEFAULT '0' COMMENT '资源id',
`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '课程-资源关系表';


DROP TABLE IF EXISTS `student_course`;
CREATE TABLE `course_student` (
`id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
`student_id` int(11)  NOT NULL DEFAULT '0' COMMENT '学生id',
`courese_id` int(11)  NOT NULL DEFAULT '0' COMMENT '课程id',
# `score` int(11)  NOT NULL DEFAULT '0' COMMENT '分数',
`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '学生-课程关系表';

DROP TABLE IF EXISTS `course_teacher`;
CREATE TABLE `course_teacher` (
`id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
`course_id` int(11)  NOT NULL DEFAULT '0' COMMENT '课程id',
`teacher_id` int(11)  NOT NULL DEFAULT '0' COMMENT '教师id',
`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '课程-教师关系表';


# 5 课堂表
DROP TABLE IF EXISTS `lesson`;
CREATE TABLE `lesson` (
`id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
`title` varchar(100)  NOT NULL DEFAULT '0' COMMENT '课堂标题',
`desc` TINYTEXT   COMMENT '课堂描述',
`course_id` int(11)  NOT NULL DEFAULT '0' COMMENT '所属课程id',
`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '课程-教师关系表';


DROP TABLE IF EXISTS `lesson_resource`;
CREATE TABLE `lesson_resource` (
`id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
`lesson_id` int(11)  NOT NULL DEFAULT '0' COMMENT '课堂id',
`resource_id` int(11)  NOT NULL DEFAULT '0' COMMENT '资源id',
`useage` tinyint(4)  NOT NULL DEFAULT '0' COMMENT '资源在课堂里面用途，1教学内容 2参考资源，3实验资源，4布置作业',
`content_type` tinyint(4)  NOT NULL DEFAULT '0' COMMENT 'useage为教学内容时，本项有效。1 实验目的 2 实验要求 3.。。',
`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '课堂-资源关系表';


DROP TABLE IF EXISTS `score`;
CREATE TABLE `score` (
`id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
`student_id` int(11)  NOT NULL DEFAULT '0' COMMENT '学生id',
`course_id` int(11)  NOT NULL DEFAULT '0' COMMENT '课程id',
`work_id` int(11)  NOT NULL DEFAULT '0' COMMENT '课堂作业id',
`mark_id` int(11)  NOT NULL DEFAULT '0' COMMENT '打分id',
`score` int(11)  NOT NULL DEFAULT '0' COMMENT '课堂得分',
`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '课堂成绩关系表';


# 6 作业表
DROP TABLE IF EXISTS `work`;
CREATE TABLE `work` (
`id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
`lesson_id` int(11)  NOT NULL DEFAULT '0' COMMENT '所属课堂id',
`user_id` int(11)  NOT NULL DEFAULT '0' COMMENT '所属用户id',

`message` TINYTEXT  COMMENT '学生提交说明',
`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '作业表';

DROP TABLE IF EXISTS `work_resource`;
CREATE TABLE `work_resource` (
`id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
`work_id` int(11)  NOT NULL DEFAULT '0' COMMENT '作业id',
`resource_id` int(11)  NOT NULL DEFAULT '0' COMMENT '资源id',
`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '作业-资源表';


# 7 打分表
DROP TABLE IF EXISTS `score_item`;
CREATE TABLE `score_item` (
`id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
`item` int(11)  NOT NULL DEFAULT '0' COMMENT '打分项',
`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '打分项表';


DROP TABLE IF EXISTS `score_subitem`;
CREATE TABLE `score_subitem` (
`id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
`subitem` int(11)  NOT NULL DEFAULT '0' COMMENT '打分子项',
`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '打分子项表';


DROP TABLE IF EXISTS `mark`;
CREATE TABLE `mark` (
`id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
`lesson_id` int(11)  NOT NULL DEFAULT '0' COMMENT '课堂lesson id',
`score_item_id` int(11)  NOT NULL DEFAULT '0' COMMENT '打分项目id',
`score_subitem_id` int(11)  NOT NULL DEFAULT '0' COMMENT '打分子项id',
`point` int(11)  NOT NULL DEFAULT '0' COMMENT '子项分',
`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '打分表';

# 7 新闻动态表
DROP TABLE IF EXISTS `news`;
CREATE TABLE `news` (
`id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
`title` varchar(100)  NOT NULL DEFAULT '0' COMMENT '新闻标题',
`content` TEXT  COMMENT '新闻内容',
`publisher` varchar(100)  NOT NULL DEFAULT '0' COMMENT '发布人',
`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '新闻动态表';


# 8 历史记录表
DROP TABLE IF EXISTS `history`;
CREATE TABLE `history` (
`id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
`user_id` int(11)  NOT NULL DEFAULT '0' COMMENT '所属用户id',
`resource_id` int(11)  NOT NULL DEFAULT '0' COMMENT '资源id',
`resource_type` tinyint(4)  NOT NULL DEFAULT '0' COMMENT '',
`useage` tinyint(4)  NOT NULL DEFAULT '0' COMMENT '资源所属类型，如资源，课程。',

`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '历史记录表';


# 9 收藏表
DROP TABLE IF EXISTS `favorite`;
CREATE TABLE `favorite` (
`id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
`user_id` int(11)  NOT NULL DEFAULT '0' COMMENT '所属用户id',
`resource_id` int(11)  NOT NULL DEFAULT '0' COMMENT '资源id',
`useage` tinyint(4)  NOT NULL DEFAULT '0' COMMENT '资源所属类型，如资源，课程。',
`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '收藏表';

