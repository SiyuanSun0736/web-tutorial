-- ================================
-- 数据库管理脚本 - 数据维护和检查
-- ================================

-- 1. 检查数据完整性
SELECT '=== 数据完整性检查 ===' AS info;

-- 检查是否有空值
SELECT 
    'ID为空的记录数' AS check_type,
    COUNT(*) AS count
FROM company WHERE id IS NULL OR id = ''
UNION ALL
SELECT 
    '名称为空的记录数' AS check_type,
    COUNT(*) AS count
FROM company WHERE name IS NULL OR name = ''
UNION ALL
SELECT 
    '昵称为空的记录数' AS check_type,
    COUNT(*) AS count
FROM company WHERE nickname IS NULL OR nickname = '';

-- 2. 检查重复数据
SELECT '=== 重复数据检查 ===' AS info;

-- 检查重复的公司名称
SELECT 
    name AS "重复的公司名称",
    COUNT(*) AS "重复次数"
FROM company 
GROUP BY name 
HAVING COUNT(*) > 1;

-- 检查重复的昵称
SELECT 
    nickname AS "重复的昵称",
    COUNT(*) AS "重复次数"
FROM company 
GROUP BY nickname 
HAVING COUNT(*) > 1;

-- 3. 数据长度分析
SELECT '=== 数据长度分析 ===' AS info;
SELECT 
    'ID' AS field_name,
    MIN(LENGTH(id)) AS min_length,
    MAX(LENGTH(id)) AS max_length,
    AVG(LENGTH(id))::NUMERIC(5,2) AS avg_length
FROM company
UNION ALL
SELECT 
    'name' AS field_name,
    MIN(LENGTH(name)) AS min_length,
    MAX(LENGTH(name)) AS max_length,
    AVG(LENGTH(name))::NUMERIC(5,2) AS avg_length
FROM company
UNION ALL
SELECT 
    'nickname' AS field_name,
    MIN(LENGTH(nickname)) AS min_length,
    MAX(LENGTH(nickname)) AS max_length,
    AVG(LENGTH(nickname))::NUMERIC(5,2) AS avg_length
FROM company;

-- 4. 查找异常数据
SELECT '=== 异常数据检查 ===' AS info;

-- 查找ID格式异常的记录
SELECT 
    id,
    name,
    'ID格式异常' AS issue
FROM company 
WHERE id NOT SIMILAR TO 'COMP[0-9]{3}';

-- 查找创建时间晚于更新时间的记录（异常情况）
SELECT 
    id,
    name,
    created_at,
    updated_at,
    '创建时间晚于更新时间' AS issue
FROM company 
WHERE created_at > updated_at;

-- 5. 按时间段统计
SELECT '=== 按时间段统计 ===' AS info;
SELECT 
    CASE 
        WHEN created_at >= CURRENT_DATE THEN '今天'
        WHEN created_at >= CURRENT_DATE - INTERVAL '7 days' THEN '本周'
        WHEN created_at >= CURRENT_DATE - INTERVAL '30 days' THEN '本月'
        ELSE '更早'
    END AS time_period,
    COUNT(*) AS count
FROM company 
GROUP BY CASE 
    WHEN created_at >= CURRENT_DATE THEN '今天'
    WHEN created_at >= CURRENT_DATE - INTERVAL '7 days' THEN '本周'
    WHEN created_at >= CURRENT_DATE - INTERVAL '30 days' THEN '本月'
    ELSE '更早'
END
ORDER BY 
    CASE 
        WHEN time_period = '今天' THEN 1
        WHEN time_period = '本周' THEN 2
        WHEN time_period = '本月' THEN 3
        ELSE 4
    END;

-- 6. 导出数据的查询（CSV格式友好）
SELECT '=== 导出格式查询 ===' AS info;
COPY (
    SELECT 
        id,
        name,
        nickname,
        created_at::text,
        updated_at::text
    FROM company 
    ORDER BY created_at
) TO STDOUT WITH CSV HEADER;
