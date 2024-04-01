/*
 Navicat Premium Data Transfer

 Source Server         : local-postgres
 Source Server Type    : PostgreSQL
 Source Server Version : 160002 (160002)
 Source Host           : localhost:5432
 Source Catalog        : distsys
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 160002 (160002)
 File Encoding         : 65001

 Date: 01/04/2024 21:49:57
*/


-- ----------------------------
-- Sequence structure for casbin_rule_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."casbin_rule_id_seq";
CREATE SEQUENCE "public"."casbin_rule_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
ALTER SEQUENCE "public"."casbin_rule_id_seq" OWNER TO "distsys";

-- ----------------------------
-- Sequence structure for privilege_resource_role_map_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."privilege_resource_role_map_id_seq";
CREATE SEQUENCE "public"."privilege_resource_role_map_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
ALTER SEQUENCE "public"."privilege_resource_role_map_id_seq" OWNER TO "distsys";

-- ----------------------------
-- Sequence structure for privilege_resource_roles_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."privilege_resource_roles_id_seq";
CREATE SEQUENCE "public"."privilege_resource_roles_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
ALTER SEQUENCE "public"."privilege_resource_roles_id_seq" OWNER TO "distsys";

-- ----------------------------
-- Sequence structure for privilege_rest_action_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."privilege_rest_action_id_seq";
CREATE SEQUENCE "public"."privilege_rest_action_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
ALTER SEQUENCE "public"."privilege_rest_action_id_seq" OWNER TO "distsys";

-- ----------------------------
-- Sequence structure for privilege_user_resorce_role_map_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."privilege_user_resorce_role_map_id_seq";
CREATE SEQUENCE "public"."privilege_user_resorce_role_map_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
ALTER SEQUENCE "public"."privilege_user_resorce_role_map_id_seq" OWNER TO "distsys";

-- ----------------------------
-- Sequence structure for privilege_user_role_map_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."privilege_user_role_map_id_seq";
CREATE SEQUENCE "public"."privilege_user_role_map_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
ALTER SEQUENCE "public"."privilege_user_role_map_id_seq" OWNER TO "distsys";

-- ----------------------------
-- Sequence structure for privilege_user_roles_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."privilege_user_roles_id_seq";
CREATE SEQUENCE "public"."privilege_user_roles_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
ALTER SEQUENCE "public"."privilege_user_roles_id_seq" OWNER TO "distsys";

-- ----------------------------
-- Sequence structure for user_address_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."user_address_id_seq";
CREATE SEQUENCE "public"."user_address_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
ALTER SEQUENCE "public"."user_address_id_seq" OWNER TO "distsys";

-- ----------------------------
-- Sequence structure for user_group_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."user_group_id_seq";
CREATE SEQUENCE "public"."user_group_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
ALTER SEQUENCE "public"."user_group_id_seq" OWNER TO "distsys";

-- ----------------------------
-- Sequence structure for user_group_member_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."user_group_member_id_seq";
CREATE SEQUENCE "public"."user_group_member_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
ALTER SEQUENCE "public"."user_group_member_id_seq" OWNER TO "distsys";

-- ----------------------------
-- Sequence structure for user_info_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."user_info_id_seq";
CREATE SEQUENCE "public"."user_info_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
ALTER SEQUENCE "public"."user_info_id_seq" OWNER TO "distsys";

-- ----------------------------
-- Sequence structure for user_level_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."user_level_id_seq";
CREATE SEQUENCE "public"."user_level_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
ALTER SEQUENCE "public"."user_level_id_seq" OWNER TO "distsys";

-- ----------------------------
-- Sequence structure for user_score_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."user_score_id_seq";
CREATE SEQUENCE "public"."user_score_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
ALTER SEQUENCE "public"."user_score_id_seq" OWNER TO "distsys";

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS "public"."casbin_rule";
CREATE TABLE "public"."casbin_rule" (
  "id" int8 NOT NULL DEFAULT nextval('casbin_rule_id_seq'::regclass),
  "ptype" varchar(200) COLLATE "pg_catalog"."default",
  "v0" varchar(200) COLLATE "pg_catalog"."default",
  "v1" varchar(200) COLLATE "pg_catalog"."default",
  "v2" varchar(200) COLLATE "pg_catalog"."default",
  "v3" varchar(200) COLLATE "pg_catalog"."default",
  "v4" varchar(200) COLLATE "pg_catalog"."default",
  "v5" varchar(200) COLLATE "pg_catalog"."default"
)
;
ALTER TABLE "public"."casbin_rule" OWNER TO "distsys";

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
BEGIN;
INSERT INTO "public"."casbin_rule" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (1, 'g', 'uid::1', 'group::admin', '_', '_', NULL, NULL);
INSERT INTO "public"."casbin_rule" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (2, 'g2', '/*', 'module::all', NULL, NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3, 'p', 'group::admin', 'module::all', '(GET)|(POST)|(PUT)|(DELETE)', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (5, 'g2', '/privilege/*', 'module::privilege', NULL, NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (4, 'g2', '/user/*', 'module::user', NULL, NULL, NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for privilege_resource_role_map
-- ----------------------------
DROP TABLE IF EXISTS "public"."privilege_resource_role_map";
CREATE TABLE "public"."privilege_resource_role_map" (
  "id" int8 NOT NULL DEFAULT nextval('privilege_resource_role_map_id_seq'::regclass),
  "created_at" timestamptz(6),
  "updated_at" timestamptz(6),
  "deleted_at" timestamptz(6),
  "resource" text COLLATE "pg_catalog"."default",
  "action" text COLLATE "pg_catalog"."default",
  "role_id" int8,
  "is_active" bool,
  "is_deleted" bool
)
;
ALTER TABLE "public"."privilege_resource_role_map" OWNER TO "distsys";

-- ----------------------------
-- Records of privilege_resource_role_map
-- ----------------------------
BEGIN;
INSERT INTO "public"."privilege_resource_role_map" ("id", "created_at", "updated_at", "deleted_at", "resource", "action", "role_id", "is_active", "is_deleted") VALUES (1, NULL, NULL, NULL, '/*', NULL, 1, 't', 'f');
INSERT INTO "public"."privilege_resource_role_map" ("id", "created_at", "updated_at", "deleted_at", "resource", "action", "role_id", "is_active", "is_deleted") VALUES (2, NULL, NULL, NULL, '/privilege/*', NULL, 2, 't', 'f');
INSERT INTO "public"."privilege_resource_role_map" ("id", "created_at", "updated_at", "deleted_at", "resource", "action", "role_id", "is_active", "is_deleted") VALUES (3, NULL, NULL, NULL, '/public/*', NULL, 3, 't', 'f');
INSERT INTO "public"."privilege_resource_role_map" ("id", "created_at", "updated_at", "deleted_at", "resource", "action", "role_id", "is_active", "is_deleted") VALUES (4, NULL, NULL, NULL, '/user/*', NULL, 4, 't', 'f');
COMMIT;

-- ----------------------------
-- Table structure for privilege_resource_roles
-- ----------------------------
DROP TABLE IF EXISTS "public"."privilege_resource_roles";
CREATE TABLE "public"."privilege_resource_roles" (
  "id" int8 NOT NULL DEFAULT nextval('privilege_resource_roles_id_seq'::regclass),
  "created_at" timestamptz(6),
  "updated_at" timestamptz(6),
  "deleted_at" timestamptz(6),
  "label" varchar(30) COLLATE "pg_catalog"."default",
  "description" varchar(255) COLLATE "pg_catalog"."default",
  "is_active" bool,
  "is_deleted" bool
)
;
ALTER TABLE "public"."privilege_resource_roles" OWNER TO "distsys";

-- ----------------------------
-- Records of privilege_resource_roles
-- ----------------------------
BEGIN;
INSERT INTO "public"."privilege_resource_roles" ("id", "created_at", "updated_at", "deleted_at", "label", "description", "is_active", "is_deleted") VALUES (1, NULL, NULL, NULL, 'module::all', 'for admin group', 't', 'f');
INSERT INTO "public"."privilege_resource_roles" ("id", "created_at", "updated_at", "deleted_at", "label", "description", "is_active", "is_deleted") VALUES (2, NULL, NULL, NULL, 'module::privilege', 'privilege module', 't', 'f');
INSERT INTO "public"."privilege_resource_roles" ("id", "created_at", "updated_at", "deleted_at", "label", "description", "is_active", "is_deleted") VALUES (3, NULL, NULL, NULL, 'module::public', 'share for all modules', 't', 'f');
INSERT INTO "public"."privilege_resource_roles" ("id", "created_at", "updated_at", "deleted_at", "label", "description", "is_active", "is_deleted") VALUES (4, NULL, NULL, NULL, 'module::user', 'user module', 't', 'f');
COMMIT;

-- ----------------------------
-- Table structure for privilege_rest_action
-- ----------------------------
DROP TABLE IF EXISTS "public"."privilege_rest_action";
CREATE TABLE "public"."privilege_rest_action" (
  "id" int8 NOT NULL DEFAULT nextval('privilege_rest_action_id_seq'::regclass),
  "label" varchar(30) COLLATE "pg_catalog"."default",
  "description" varchar(200) COLLATE "pg_catalog"."default",
  "created_at" timestamptz(6),
  "updated_at" timestamptz(6),
  "deleted_at" timestamptz(6),
  "is_active" bool,
  "is_deleted" bool
)
;
ALTER TABLE "public"."privilege_rest_action" OWNER TO "distsys";

-- ----------------------------
-- Records of privilege_rest_action
-- ----------------------------
BEGIN;
INSERT INTO "public"."privilege_rest_action" ("id", "label", "description", "created_at", "updated_at", "deleted_at", "is_active", "is_deleted") VALUES (1, 'GET', NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO "public"."privilege_rest_action" ("id", "label", "description", "created_at", "updated_at", "deleted_at", "is_active", "is_deleted") VALUES (2, 'POST', NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO "public"."privilege_rest_action" ("id", "label", "description", "created_at", "updated_at", "deleted_at", "is_active", "is_deleted") VALUES (3, 'PUT', NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO "public"."privilege_rest_action" ("id", "label", "description", "created_at", "updated_at", "deleted_at", "is_active", "is_deleted") VALUES (4, 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO "public"."privilege_rest_action" ("id", "label", "description", "created_at", "updated_at", "deleted_at", "is_active", "is_deleted") VALUES (5, '(GET)|(POST)', NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO "public"."privilege_rest_action" ("id", "label", "description", "created_at", "updated_at", "deleted_at", "is_active", "is_deleted") VALUES (6, '(GET)|({POST)|(PUT)|(DELETE)', NULL, NULL, NULL, NULL, NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for privilege_user_resorce_role_map
-- ----------------------------
DROP TABLE IF EXISTS "public"."privilege_user_resorce_role_map";
CREATE TABLE "public"."privilege_user_resorce_role_map" (
  "id" int8 NOT NULL DEFAULT nextval('privilege_user_resorce_role_map_id_seq'::regclass),
  "created_at" timestamptz(6),
  "updated_at" timestamptz(6),
  "deleted_at" timestamptz(6),
  "user_role_id" int8,
  "user_role_label" varchar(30) COLLATE "pg_catalog"."default",
  "rescource_role_id" int8,
  "resource_role_label" varchar(30) COLLATE "pg_catalog"."default",
  "action_id" int8,
  "action_label" varchar(30) COLLATE "pg_catalog"."default",
  "is_active" bool,
  "is_deleted" bool
)
;
ALTER TABLE "public"."privilege_user_resorce_role_map" OWNER TO "distsys";

-- ----------------------------
-- Records of privilege_user_resorce_role_map
-- ----------------------------
BEGIN;
INSERT INTO "public"."privilege_user_resorce_role_map" ("id", "created_at", "updated_at", "deleted_at", "user_role_id", "user_role_label", "rescource_role_id", "resource_role_label", "action_id", "action_label", "is_active", "is_deleted") VALUES (1, NULL, NULL, NULL, 1, 'group::admin', 1, 'module::all', 6, '(GET)|({POST)|(PUT)|(DELETE)', 't', 'f');
COMMIT;

-- ----------------------------
-- Table structure for privilege_user_role_map
-- ----------------------------
DROP TABLE IF EXISTS "public"."privilege_user_role_map";
CREATE TABLE "public"."privilege_user_role_map" (
  "id" int8 NOT NULL DEFAULT nextval('privilege_user_role_map_id_seq'::regclass),
  "created_at" timestamptz(6),
  "updated_at" timestamptz(6),
  "deleted_at" timestamptz(6),
  "user_id" int8,
  "role_id" int8,
  "is_active" bool,
  "is_deleted" bool
)
;
ALTER TABLE "public"."privilege_user_role_map" OWNER TO "distsys";

-- ----------------------------
-- Records of privilege_user_role_map
-- ----------------------------
BEGIN;
INSERT INTO "public"."privilege_user_role_map" ("id", "created_at", "updated_at", "deleted_at", "user_id", "role_id", "is_active", "is_deleted") VALUES (1, NULL, NULL, NULL, 1, 1, 't', 't');
COMMIT;

-- ----------------------------
-- Table structure for privilege_user_roles
-- ----------------------------
DROP TABLE IF EXISTS "public"."privilege_user_roles";
CREATE TABLE "public"."privilege_user_roles" (
  "id" int8 NOT NULL DEFAULT nextval('privilege_user_roles_id_seq'::regclass),
  "created_at" timestamptz(6),
  "updated_at" timestamptz(6),
  "deleted_at" timestamptz(6),
  "label" varchar(30) COLLATE "pg_catalog"."default",
  "description" varchar(255) COLLATE "pg_catalog"."default",
  "is_active" bool,
  "is_deleted" bool
)
;
ALTER TABLE "public"."privilege_user_roles" OWNER TO "distsys";

-- ----------------------------
-- Records of privilege_user_roles
-- ----------------------------
BEGIN;
INSERT INTO "public"."privilege_user_roles" ("id", "created_at", "updated_at", "deleted_at", "label", "description", "is_active", "is_deleted") VALUES (1, NULL, NULL, NULL, 'group::admin', 'admin group', 't', 'f');
COMMIT;

-- ----------------------------
-- Table structure for user_address
-- ----------------------------
DROP TABLE IF EXISTS "public"."user_address";
CREATE TABLE "public"."user_address" (
  "id" int8 NOT NULL DEFAULT nextval('user_address_id_seq'::regclass),
  "created_at" timestamptz(6),
  "updated_at" timestamptz(6),
  "deleted_at" timestamptz(6),
  "country_id" int8,
  "region_id" int8,
  "city_id" int8,
  "area_id" int8,
  "street" varchar(100) COLLATE "pg_catalog"."default",
  "raw_addr" varchar(255) COLLATE "pg_catalog"."default",
  "latitude" numeric,
  "longitude" numeric
)
;
ALTER TABLE "public"."user_address" OWNER TO "distsys";

-- ----------------------------
-- Records of user_address
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for user_group
-- ----------------------------
DROP TABLE IF EXISTS "public"."user_group";
CREATE TABLE "public"."user_group" (
  "id" int8 NOT NULL DEFAULT nextval('user_group_id_seq'::regclass),
  "created_at" timestamptz(6),
  "updated_at" timestamptz(6),
  "deleted_at" timestamptz(6),
  "label" varchar(30) COLLATE "pg_catalog"."default",
  "description" varchar(255) COLLATE "pg_catalog"."default",
  "owner_uid" int8,
  "is_active" bool,
  "is_deleted" bool
)
;
ALTER TABLE "public"."user_group" OWNER TO "distsys";

-- ----------------------------
-- Records of user_group
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for user_group_member
-- ----------------------------
DROP TABLE IF EXISTS "public"."user_group_member";
CREATE TABLE "public"."user_group_member" (
  "id" int8 NOT NULL DEFAULT nextval('user_group_member_id_seq'::regclass),
  "created_at" timestamptz(6),
  "updated_at" timestamptz(6),
  "deleted_at" timestamptz(6),
  "user_id" int8,
  "group_id" int8,
  "member_level" int8,
  "is_cofounder" bool,
  "is_deleted" bool
)
;
ALTER TABLE "public"."user_group_member" OWNER TO "distsys";

-- ----------------------------
-- Records of user_group_member
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for user_info
-- ----------------------------
DROP TABLE IF EXISTS "public"."user_info";
CREATE TABLE "public"."user_info" (
  "id" int8 NOT NULL DEFAULT nextval('user_info_id_seq'::regclass),
  "created_at" timestamptz(6),
  "updated_at" timestamptz(6),
  "deleted_at" timestamptz(6),
  "nickname" varchar(30) COLLATE "pg_catalog"."default",
  "wechat_uids" varchar(255) COLLATE "pg_catalog"."default",
  "portrait" varchar(100) COLLATE "pg_catalog"."default",
  "phone" varchar(20) COLLATE "pg_catalog"."default",
  "password" varchar(150) COLLATE "pg_catalog"."default",
  "age" int8,
  "addr_id" int8,
  "level_id" int8,
  "is_active" bool,
  "is_deleted" bool,
  "score" int8
)
;
ALTER TABLE "public"."user_info" OWNER TO "distsys";

-- ----------------------------
-- Records of user_info
-- ----------------------------
BEGIN;
INSERT INTO "public"."user_info" ("id", "created_at", "updated_at", "deleted_at", "nickname", "wechat_uids", "portrait", "phone", "password", "age", "addr_id", "level_id", "is_active", "is_deleted", "score") VALUES (1, '2024-04-01 13:47:49.24821+00', '2024-04-01 13:47:49.24821+00', NULL, 'admin', '20e00dd683030928c75b97d03043c87a', '', '18676706732', 'MtVmL5gTj9AkbqHQdXS2Iw==', 0, 0, 0, 't', 'f', 0);
COMMIT;

-- ----------------------------
-- Table structure for user_level
-- ----------------------------
DROP TABLE IF EXISTS "public"."user_level";
CREATE TABLE "public"."user_level" (
  "id" int8 NOT NULL DEFAULT nextval('user_level_id_seq'::regclass),
  "created_at" timestamptz(6),
  "updated_at" timestamptz(6),
  "deleted_at" timestamptz(6),
  "label" varchar(30) COLLATE "pg_catalog"."default",
  "min_score" int8,
  "max_score" int8
)
;
ALTER TABLE "public"."user_level" OWNER TO "distsys";

-- ----------------------------
-- Records of user_level
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for user_score
-- ----------------------------
DROP TABLE IF EXISTS "public"."user_score";
CREATE TABLE "public"."user_score" (
  "id" int8 NOT NULL DEFAULT nextval('user_score_id_seq'::regclass),
  "created_at" timestamptz(6),
  "updated_at" timestamptz(6),
  "deleted_at" timestamptz(6),
  "user_id" int8,
  "activity_id" int8,
  "act_type_id" int8,
  "score" int8,
  "group_id" int8,
  "from_user_id" int8,
  "to_user_id" int8,
  "is_active" bool,
  "is_calced" bool
)
;
ALTER TABLE "public"."user_score" OWNER TO "distsys";

-- ----------------------------
-- Records of user_score
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."casbin_rule_id_seq"
OWNED BY "public"."casbin_rule"."id";
SELECT setval('"public"."casbin_rule_id_seq"', 5, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."privilege_resource_role_map_id_seq"
OWNED BY "public"."privilege_resource_role_map"."id";
SELECT setval('"public"."privilege_resource_role_map_id_seq"', 4, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."privilege_resource_roles_id_seq"
OWNED BY "public"."privilege_resource_roles"."id";
SELECT setval('"public"."privilege_resource_roles_id_seq"', 4, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."privilege_rest_action_id_seq"
OWNED BY "public"."privilege_rest_action"."id";
SELECT setval('"public"."privilege_rest_action_id_seq"', 6, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."privilege_user_resorce_role_map_id_seq"
OWNED BY "public"."privilege_user_resorce_role_map"."id";
SELECT setval('"public"."privilege_user_resorce_role_map_id_seq"', 1, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."privilege_user_role_map_id_seq"
OWNED BY "public"."privilege_user_role_map"."id";
SELECT setval('"public"."privilege_user_role_map_id_seq"', 1, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."privilege_user_roles_id_seq"
OWNED BY "public"."privilege_user_roles"."id";
SELECT setval('"public"."privilege_user_roles_id_seq"', 1, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."user_address_id_seq"
OWNED BY "public"."user_address"."id";
SELECT setval('"public"."user_address_id_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."user_group_id_seq"
OWNED BY "public"."user_group"."id";
SELECT setval('"public"."user_group_id_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."user_group_member_id_seq"
OWNED BY "public"."user_group_member"."id";
SELECT setval('"public"."user_group_member_id_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."user_info_id_seq"
OWNED BY "public"."user_info"."id";
SELECT setval('"public"."user_info_id_seq"', 1, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."user_level_id_seq"
OWNED BY "public"."user_level"."id";
SELECT setval('"public"."user_level_id_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."user_score_id_seq"
OWNED BY "public"."user_score"."id";
SELECT setval('"public"."user_score_id_seq"', 1, false);

-- ----------------------------
-- Primary Key structure for table casbin_rule
-- ----------------------------
ALTER TABLE "public"."casbin_rule" ADD CONSTRAINT "casbin_rule_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table privilege_resource_role_map
-- ----------------------------
CREATE INDEX "idx_privilege_resource_role_map_deleted_at" ON "public"."privilege_resource_role_map" USING btree (
  "deleted_at" "pg_catalog"."timestamptz_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table privilege_resource_role_map
-- ----------------------------
ALTER TABLE "public"."privilege_resource_role_map" ADD CONSTRAINT "privilege_resource_role_map_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table privilege_resource_roles
-- ----------------------------
CREATE INDEX "idx_privilege_resource_roles_deleted_at" ON "public"."privilege_resource_roles" USING btree (
  "deleted_at" "pg_catalog"."timestamptz_ops" ASC NULLS LAST
);

-- ----------------------------
-- Uniques structure for table privilege_resource_roles
-- ----------------------------
ALTER TABLE "public"."privilege_resource_roles" ADD CONSTRAINT "uni_privilege_resource_roles_label" UNIQUE ("label");

-- ----------------------------
-- Primary Key structure for table privilege_resource_roles
-- ----------------------------
ALTER TABLE "public"."privilege_resource_roles" ADD CONSTRAINT "privilege_resource_roles_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table privilege_rest_action
-- ----------------------------
CREATE INDEX "idx_privilege_rest_action_deleted_at" ON "public"."privilege_rest_action" USING btree (
  "deleted_at" "pg_catalog"."timestamptz_ops" ASC NULLS LAST
);

-- ----------------------------
-- Uniques structure for table privilege_rest_action
-- ----------------------------
ALTER TABLE "public"."privilege_rest_action" ADD CONSTRAINT "uni_privilege_rest_action_label" UNIQUE ("label");

-- ----------------------------
-- Primary Key structure for table privilege_rest_action
-- ----------------------------
ALTER TABLE "public"."privilege_rest_action" ADD CONSTRAINT "privilege_rest_action_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table privilege_user_resorce_role_map
-- ----------------------------
CREATE INDEX "idx_privilege_user_resorce_role_map_deleted_at" ON "public"."privilege_user_resorce_role_map" USING btree (
  "deleted_at" "pg_catalog"."timestamptz_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table privilege_user_resorce_role_map
-- ----------------------------
ALTER TABLE "public"."privilege_user_resorce_role_map" ADD CONSTRAINT "privilege_user_resorce_role_map_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table privilege_user_role_map
-- ----------------------------
CREATE INDEX "idx_privilege_user_role_map_deleted_at" ON "public"."privilege_user_role_map" USING btree (
  "deleted_at" "pg_catalog"."timestamptz_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table privilege_user_role_map
-- ----------------------------
ALTER TABLE "public"."privilege_user_role_map" ADD CONSTRAINT "privilege_user_role_map_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table privilege_user_roles
-- ----------------------------
CREATE INDEX "idx_privilege_user_roles_deleted_at" ON "public"."privilege_user_roles" USING btree (
  "deleted_at" "pg_catalog"."timestamptz_ops" ASC NULLS LAST
);

-- ----------------------------
-- Uniques structure for table privilege_user_roles
-- ----------------------------
ALTER TABLE "public"."privilege_user_roles" ADD CONSTRAINT "uni_privilege_user_roles_label" UNIQUE ("label");

-- ----------------------------
-- Primary Key structure for table privilege_user_roles
-- ----------------------------
ALTER TABLE "public"."privilege_user_roles" ADD CONSTRAINT "privilege_user_roles_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table user_address
-- ----------------------------
CREATE INDEX "idx_user_address_deleted_at" ON "public"."user_address" USING btree (
  "deleted_at" "pg_catalog"."timestamptz_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table user_address
-- ----------------------------
ALTER TABLE "public"."user_address" ADD CONSTRAINT "user_address_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table user_group
-- ----------------------------
CREATE INDEX "idx_user_group_deleted_at" ON "public"."user_group" USING btree (
  "deleted_at" "pg_catalog"."timestamptz_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table user_group
-- ----------------------------
ALTER TABLE "public"."user_group" ADD CONSTRAINT "user_group_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table user_group_member
-- ----------------------------
CREATE INDEX "idx_user_group_member_deleted_at" ON "public"."user_group_member" USING btree (
  "deleted_at" "pg_catalog"."timestamptz_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table user_group_member
-- ----------------------------
ALTER TABLE "public"."user_group_member" ADD CONSTRAINT "user_group_member_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table user_info
-- ----------------------------
CREATE INDEX "idx_user_info_deleted_at" ON "public"."user_info" USING btree (
  "deleted_at" "pg_catalog"."timestamptz_ops" ASC NULLS LAST
);
CREATE UNIQUE INDEX "idx_user_info_phone" ON "public"."user_info" USING btree (
  "phone" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE UNIQUE INDEX "idx_user_info_wechat_uid_s" ON "public"."user_info" USING btree (
  "wechat_uids" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table user_info
-- ----------------------------
ALTER TABLE "public"."user_info" ADD CONSTRAINT "user_info_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table user_level
-- ----------------------------
CREATE INDEX "idx_user_level_deleted_at" ON "public"."user_level" USING btree (
  "deleted_at" "pg_catalog"."timestamptz_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table user_level
-- ----------------------------
ALTER TABLE "public"."user_level" ADD CONSTRAINT "user_level_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table user_score
-- ----------------------------
CREATE INDEX "idx_user_score_deleted_at" ON "public"."user_score" USING btree (
  "deleted_at" "pg_catalog"."timestamptz_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table user_score
-- ----------------------------
ALTER TABLE "public"."user_score" ADD CONSTRAINT "user_score_pkey" PRIMARY KEY ("id");
