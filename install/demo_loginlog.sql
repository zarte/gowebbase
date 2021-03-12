/*
Navicat MySQL Data Transfer

Source Server         : 本地
Source Server Version : 50726
Source Host           : 127.0.0.1:3306
Source Database       : test

Target Server Type    : MYSQL
Target Server Version : 50726
File Encoding         : 65001

Date: 2021-03-12 19:05:45
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for demo_loginlog
-- ----------------------------
DROP TABLE IF EXISTS `demo_loginlog`;
CREATE TABLE `demo_loginlog` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uid` int(11) NOT NULL,
  `ctime` datetime DEFAULT NULL,
  `ip` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for demo_urole
-- ----------------------------
DROP TABLE IF EXISTS `demo_urole`;
CREATE TABLE `demo_urole` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uid` int(11) NOT NULL,
  `roleid` int(11) DEFAULT NULL,
  `status` tinyint(255) DEFAULT '0' COMMENT '1删除',
  `ctime` datetime DEFAULT NULL,
  `stime` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for demo_user
-- ----------------------------
DROP TABLE IF EXISTS `demo_user`;
CREATE TABLE `demo_user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(255) NOT NULL,
  `passwd` varchar(255) NOT NULL,
  `salt` varchar(12) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `point` decimal(20,0) DEFAULT '0',
  `money` decimal(20,0) DEFAULT '0',
  `phone` varchar(32) DEFAULT NULL,
  `state` tinyint(255) DEFAULT '0' COMMENT '1删除',
  `ctime` datetime DEFAULT NULL,
  `stime` datetime DEFAULT NULL,
  `roles` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;
