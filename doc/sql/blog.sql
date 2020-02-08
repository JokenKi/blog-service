/*
Navicat MariaDB Data Transfer

Source Server         : 127.0.0.1
Source Server Version : 100314
Source Host           : 127.0.0.1:3306
Source Database       : blog

Target Server Type    : MariaDB
Target Server Version : 100314
File Encoding         : 65001

Date: 2020-02-08 19:12:05
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for r_intelligent_que
-- ----------------------------
DROP TABLE IF EXISTS `r_intelligent_que`;
CREATE TABLE `r_intelligent_que` (
  `ID` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键(自增)',
  `PROJECT_CODE` bigint(20) NOT NULL COMMENT '项目编码',
  `ORG_CODE` varchar(20) NOT NULL COMMENT '单位编码',
  `AREA_CODE` varchar(6) NOT NULL COMMENT '地区',
  `QUE_ID` varchar(20) DEFAULT NULL COMMENT '问题id',
  `MAC_ID` varchar(50) NOT NULL COMMENT '硬件ID',
  `CHANNEL_NO` bigint(20) DEFAULT NULL COMMENT '渠道编码',
  `QUE_STATUS` int(1) NOT NULL COMMENT '问题回复状态(0:直接回答、1:引导回答2:未回答、3:未知状态)',
  `CLASSIFY_CODE` varchar(20) NOT NULL COMMENT '分类编码',
  `DEVICE_CODE` varchar(5) NOT NULL COMMENT '设备类型',
  `USER_ID` varchar(20) DEFAULT NULL COMMENT '用户ID',
  `CREATE_TIME` datetime NOT NULL COMMENT '创建时间',
  `EXT_1` int(1) DEFAULT NULL COMMENT '预留',
  `EXT_2` varchar(10) DEFAULT NULL COMMENT '预留',
  `EXT_3` varchar(10) DEFAULT NULL COMMENT '预留',
  `EXT_4` varchar(10) DEFAULT NULL COMMENT '预留',
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of r_intelligent_que
-- ----------------------------

-- ----------------------------
-- Table structure for tb_blog
-- ----------------------------
DROP TABLE IF EXISTS `tb_blog`;
CREATE TABLE `tb_blog` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `customer_id` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '用户ID',
  `type_id` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '类型ID',
  `blog_title` varchar(255) NOT NULL DEFAULT '' COMMENT '博客标题',
  `content` text NOT NULL COMMENT '博客内容',
  `read_num` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '阅读数',
  `status` tinyint(3) unsigned NOT NULL DEFAULT 5 COMMENT '0:不可用;5:可用',
  `time_create` bigint(20) unsigned NOT NULL DEFAULT 0,
  `time_update` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of tb_blog
-- ----------------------------
INSERT INTO `tb_blog` VALUES ('1', '1', '1', '1', '1', '0', '5', '0', '0');
INSERT INTO `tb_blog` VALUES ('2', '3', '1', '一天', '记录', '0', '5', '1566372718', '1566372718');
INSERT INTO `tb_blog` VALUES ('3', '3', '1', '一天', '记录', '0', '5', '1566442332', '1566442332');
INSERT INTO `tb_blog` VALUES ('4', '3', '1', '一天', '记录', '0', '5', '1567390863', '1567390863');

-- ----------------------------
-- Table structure for tb_blog_type
-- ----------------------------
DROP TABLE IF EXISTS `tb_blog_type`;
CREATE TABLE `tb_blog_type` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `type_name` varchar(255) NOT NULL DEFAULT '' COMMENT '分类名称',
  `status` tinyint(3) unsigned NOT NULL DEFAULT 5 COMMENT '0:不可用;5:可用',
  `time_create` bigint(20) unsigned NOT NULL DEFAULT 0,
  `time_update` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of tb_blog_type
-- ----------------------------
INSERT INTO `tb_blog_type` VALUES ('1', '分类1', '5', '1566372564', '1566372564');

-- ----------------------------
-- Table structure for tb_blog_type_tag
-- ----------------------------
DROP TABLE IF EXISTS `tb_blog_type_tag`;
CREATE TABLE `tb_blog_type_tag` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `blog_id` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '博客ID',
  `tag_id` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '标签ID',
  `status` tinyint(3) unsigned DEFAULT 5 COMMENT '0:不可用;5:可用',
  `time_create` bigint(20) unsigned NOT NULL DEFAULT 0,
  `time_update` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of tb_blog_type_tag
-- ----------------------------

-- ----------------------------
-- Table structure for tb_customer
-- ----------------------------
DROP TABLE IF EXISTS `tb_customer`;
CREATE TABLE `tb_customer` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `name` varchar(64) NOT NULL DEFAULT '' COMMENT '用户名',
  `nick_name` varchar(50) NOT NULL DEFAULT '' COMMENT '用户昵称',
  `passwd` varchar(32) NOT NULL DEFAULT '' COMMENT '登陆密码',
  `salt` int(6) unsigned DEFAULT 0 COMMENT '盐值',
  `phone` varchar(11) DEFAULT '' COMMENT '手机号',
  `account_type` tinyint(1) unsigned DEFAULT 0 COMMENT '账户类型; 0:普通用户 5:管理员',
  `status` tinyint(3) unsigned DEFAULT 5 COMMENT '0:不可用;5:可用',
  `time_create` bigint(20) unsigned NOT NULL DEFAULT 0,
  `time_update` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '更新时间',
  `time_latest_login` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '最近登陆时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name_index` (`name`) USING HASH
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of tb_customer
-- ----------------------------
INSERT INTO `tb_customer` VALUES ('1', 'admin', '超级管理员', '8462C8732F15FC9558C688BA200D9C67', '1234', '18612802880', '5', '5', '1557733831', '1557733831', '0');
INSERT INTO `tb_customer` VALUES ('3', '18612802880', '18612802880', '6E3FC24D2FDA7ED39C52B0409A8C9A64', '480961', '18612802880', '1', '5', '1561358580', '1566454173', '0');

-- ----------------------------
-- Table structure for tb_customer_collection
-- ----------------------------
DROP TABLE IF EXISTS `tb_customer_collection`;
CREATE TABLE `tb_customer_collection` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `customer_id` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '用户ID',
  `blog_id` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '博客ID',
  `status` tinyint(3) unsigned DEFAULT 5 COMMENT '0:不可用;5:可用',
  `time_create` bigint(20) unsigned NOT NULL DEFAULT 0,
  `time_update` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户收藏';

-- ----------------------------
-- Records of tb_customer_collection
-- ----------------------------

-- ----------------------------
-- Table structure for tb_customer_tag
-- ----------------------------
DROP TABLE IF EXISTS `tb_customer_tag`;
CREATE TABLE `tb_customer_tag` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `customer_id` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '用户ID',
  `tag_id` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '标签ID',
  `status` tinyint(3) unsigned NOT NULL DEFAULT 5 COMMENT '0:不可用;5:可用',
  `time_create` bigint(20) unsigned NOT NULL DEFAULT 0,
  `time_update` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of tb_customer_tag
-- ----------------------------

-- ----------------------------
-- Table structure for tb_tag
-- ----------------------------
DROP TABLE IF EXISTS `tb_tag`;
CREATE TABLE `tb_tag` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `tag_name` varchar(10) CHARACTER SET utf8 NOT NULL DEFAULT '' COMMENT '博客标题',
  `status` tinyint(3) unsigned NOT NULL DEFAULT 5 COMMENT '0:不可用;5:可用',
  `time_create` bigint(20) unsigned NOT NULL DEFAULT 0,
  `time_update` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='标签';

-- ----------------------------
-- Records of tb_tag
-- ----------------------------
