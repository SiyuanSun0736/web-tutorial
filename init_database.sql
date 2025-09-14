-- 创建 company 表
CREATE TABLE IF NOT EXISTS company (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    nickname VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 创建更新时间戳的触发器函数
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

-- 创建触发器，在更新记录时自动更新 updated_at 字段
DROP TRIGGER IF EXISTS update_company_updated_at ON company;
CREATE TRIGGER update_company_updated_at
    BEFORE UPDATE ON company
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

-- 插入一些示例数据
INSERT INTO company (id, name, nickname) VALUES 
    ('COMP001', 'Google Inc.', 'Google'),
    ('COMP002', 'Microsoft Corporation', 'Microsoft'),
    ('COMP003', 'Apple Inc.', 'Apple'),
    ('COMP004', 'Amazon.com Inc.', 'Amazon'),
    ('COMP005', '阿里巴巴集团控股有限公司', '阿里巴巴')
ON CONFLICT (id) DO NOTHING;

-- 查看创建的表结构
\d company;

-- 查看插入的数据
SELECT * FROM company;
