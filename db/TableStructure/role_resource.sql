/*
 Navicat Premium Data Transfer

 Source Server         : HRmanager
 Source Server Type    : MySQL
 Source Server Version : 80405
 Source Host           : localhost:3306
 Source Schema         : NewProject

 Target Server Type    : MySQL
 Target Server Version : 80405
 File Encoding         : 65001

 Date: 04/09/2025 16:43:10
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for role_resource
-- ----------------------------
DROP TABLE IF EXISTS `role_resource`;
CREATE TABLE `role_resource`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `role_id` int(0) UNSIGNED NOT NULL COMMENT '角色ID',
  `resource_id` int(0) UNSIGNED NOT NULL COMMENT '资源ID',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '状态：1启用/2禁用',
  `created_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '创建人',
  `created_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
  `updated_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '更新人',
  `updated_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 13 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '角色资源权限表' ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
