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

 Date: 04/09/2025 16:43:30
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user_role
-- ----------------------------
DROP TABLE IF EXISTS `user_role`;
CREATE TABLE `user_role`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `user_id` int(0) UNSIGNED NOT NULL COMMENT '用户ID',
  `role_id` int(0) UNSIGNED NOT NULL COMMENT '角色ID',
  `status` tinyint(0) UNSIGNED NOT NULL DEFAULT 1 COMMENT '状态：1启用/2禁用',
  `created_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '创建人',
  `created_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
  `updated_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '更新人',
  `updated_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '用户角色绑定表' ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
