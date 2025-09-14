-- ================================
-- 数据库查询脚本 - 显示表中的信息
-- ================================

-- 1. 显示所有公司信息
SELECT '=== 所有公司信息 ===' AS info;
SELECT 
    id AS "公司ID",
    name AS "公司名称", 
    nickname AS "公司昵称",
    created_at AS "创建时间",
    updated_at AS "更新时间"
FROM company 
ORDER BY created_at;

-- 2. 统计公司数量
SELECT '=== 统计信息 ===' AS info;
SELECT COUNT(*) AS "公司总数" FROM company;

-- 3. 按创建时间分组统计
SELECT 
    DATE(created_at) AS "创建日期",
    COUNT(*) AS "当日新增公司数"
FROM company 
GROUP BY DATE(created_at)
ORDER BY DATE(created_at);

-- 4. 显示最近创建的公司
SELECT '=== 最近创建的5家公司 ===' AS info;
SELECT 
    id AS "公司ID",
    name AS "公司名称",
    nickname AS "公司昵称",
    created_at AS "创建时间"
FROM company 
ORDER BY created_at DESC 
LIMIT 5;

-- 5. 显示最近更新的公司
SELECT '=== 最近更新的5家公司 ===' AS info;
SELECT 
    id AS "公司ID",
    name AS "公司名称",
    nickname AS "公司昵称",
    updated_at AS "更新时间"
FROM company 
ORDER BY updated_at DESC 
LIMIT 5;

-- 6. 查找包含特定关键词的公司
SELECT '=== 包含"公司"关键词的企业 ===' AS info;
SELECT 
    id AS "公司ID",
    name AS "公司名称",
    nickname AS "公司昵称"
FROM company 
WHERE name LIKE '%公司%' OR name LIKE '%Corporation%' OR name LIKE '%Inc.%';

-- 7. 显示公司名称长度统计
SELECT '=== 公司名称长度统计 ===' AS info;
SELECT 
    id AS "公司ID",
    name AS "公司名称",
    LENGTH(name) AS "名称长度",
    LENGTH(nickname) AS "昵称长度"
FROM company 
ORDER BY LENGTH(name) DESC;

-- 8. 显示表结构信息
SELECT '=== 表结构信息 ===' AS info;
SELECT 
    column_name AS "列名",
    data_type AS "数据类型",
    is_nullable AS "允许为空",
    column_default AS "默认值"
FROM information_schema.columns 
WHERE table_name = 'company'
ORDER BY ordinal_position;

-- 9. 显示表的基本统计信息
SELECT '=== 表的基本统计信息 ===' AS info;
SELECT 
    COUNT(*) AS "总记录数",
    COUNT(DISTINCT name) AS "不重复公司名数量",
    COUNT(DISTINCT nickname) AS "不重复昵称数量",
    MIN(created_at) AS "最早创建时间",
    MAX(created_at) AS "最晚创建时间",
    MIN(updated_at) AS "最早更新时间",
    MAX(updated_at) AS "最晚更新时间"
FROM company;

-- 10. 按昵称首字母分组
SELECT '=== 按昵称首字母分组 ===' AS info;
SELECT 
    LEFT(nickname, 1) AS "首字母",
    COUNT(*) AS "数量",
    STRING_AGG(nickname, ', ') AS "公司昵称列表"
FROM company 
GROUP BY LEFT(nickname, 1)
ORDER BY LEFT(nickname, 1);

-- 11. 显示ID格式统计
SELECT '=== ID格式统计 ===' AS info;
SELECT 
    CASE 
        WHEN id LIKE 'COMP%' THEN 'COMP格式'
        ELSE '其他格式'
    END AS "ID格式",
    COUNT(*) AS "数量"
FROM company 
GROUP BY CASE 
    WHEN id LIKE 'COMP%' THEN 'COMP格式'
    ELSE '其他格式'
END;

-- 12. 显示详细的时间信息
SELECT '=== 详细时间信息 ===' AS info;
SELECT 
    id AS "公司ID",
    name AS "公司名称",
    TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS') AS "创建时间",
    TO_CHAR(updated_at, 'YYYY-MM-DD HH24:MI:SS') AS "更新时间",
    EXTRACT(EPOCH FROM (updated_at - created_at)) AS "创建到更新间隔(秒)"
FROM company 
ORDER BY created_at;
