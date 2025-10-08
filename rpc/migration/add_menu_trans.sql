-- Add trans column to sys_menus table for i18n support
-- Created: 2025-10-08
-- Purpose: Support multi-language menu titles

-- Add trans column
ALTER TABLE sys_menus ADD COLUMN IF NOT EXISTS trans VARCHAR(255) DEFAULT '';

-- Update existing menu records with i18n keys
UPDATE sys_menus SET trans = 'sys.route.systemManagementTitle' WHERE name = 'SystemManagement';
UPDATE sys_menus SET trans = 'sys.route.menuManagementTitle' WHERE name = 'MenuManagement';
UPDATE sys_menus SET trans = 'sys.route.roleManagementTitle' WHERE name = 'RoleManagement';
UPDATE sys_menus SET trans = 'sys.route.apiManagementTitle' WHERE name = 'ApiManagement';
UPDATE sys_menus SET trans = 'sys.route.userManagementTitle' WHERE name = 'UserManagement';
UPDATE sys_menus SET trans = 'sys.route.fileManagementTitle' WHERE name = 'FileManagement';
UPDATE sys_menus SET trans = 'sys.route.userProfileTitle' WHERE name = 'Profile';
UPDATE sys_menus SET trans = 'sys.route.dictionaryManagementTitle' WHERE name = 'DictionaryManagement';
UPDATE sys_menus SET trans = 'sys.route.dictionaryDetailManagementTitle' WHERE name = 'DictionaryDetailManagement';
UPDATE sys_menus SET trans = 'sys.route.oauthManagement' WHERE name = 'OauthManagement';
UPDATE sys_menus SET trans = 'sys.route.tokenManagement' WHERE name = 'TokenManagement';
UPDATE sys_menus SET trans = 'sys.route.positionManagement' WHERE name = 'PositionManagement';
UPDATE sys_menus SET trans = 'sys.route.taskManagement' WHERE name = 'TaskManagement';
UPDATE sys_menus SET trans = 'sys.route.departmentManagement' WHERE name = 'DepartmentManagement';
UPDATE sys_menus SET trans = 'sys.route.configurationManagement' WHERE name = 'ConfigurationManagement';

COMMENT ON COLUMN sys_menus.trans IS 'i18n key for menu title | 菜单标题国际化key';
