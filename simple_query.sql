-- ================================
-- 简化查询脚本 - 快速查看表信息
-- ================================

-- 显示所有数据
SELECT * FROM company ORDER BY created_at;

-- 显示记录总数
SELECT COUNT(*) AS total_companies FROM company;

-- 显示表结构
\d company;

-- 显示最近5条记录
SELECT * FROM company ORDER BY created_at DESC LIMIT 5;
