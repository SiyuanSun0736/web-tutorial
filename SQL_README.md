# 数据库查询脚本说明

本项目包含了多个SQL脚本文件，用于查询和管理数据库中的表信息。

## 文件说明

### 1. `init_database.sql`
- **用途**: 初始化数据库，创建表结构和示例数据
- **包含内容**: 
  - 创建company表
  - 创建触发器
  - 插入示例数据
  - 基本查询

### 2. `simple_query.sql` 
- **用途**: 简单快速的查询脚本
- **包含内容**:
  - 显示所有数据
  - 显示记录总数
  - 显示表结构
  - 显示最近5条记录

### 3. `query_database.sql`
- **用途**: 详细的查询脚本，提供各种统计和分析
- **包含内容**:
  - 显示所有公司信息
  - 统计公司数量
  - 按创建时间分组统计
  - 显示最近创建/更新的公司
  - 关键词搜索
  - 名称长度统计
  - 表结构信息
  - 基本统计信息
  - 按昵称首字母分组
  - ID格式统计
  - 详细时间信息

### 4. `admin_queries.sql`
- **用途**: 管理和维护脚本，用于数据完整性检查
- **包含内容**:
  - 数据完整性检查
  - 重复数据检查
  - 数据长度分析
  - 异常数据检查
  - 按时间段统计
  - 数据导出查询

## 使用方法

### 1. 连接到数据库
```bash
# 使用psql连接PostgreSQL数据库
psql -h localhost -U your_username -d your_database
```

### 2. 执行SQL脚本
```bash
# 初始化数据库
\i init_database.sql

# 简单查询
\i simple_query.sql

# 详细查询
\i query_database.sql

# 管理查询
\i admin_queries.sql
```

### 3. 在Go应用中使用
可以通过Go应用的数据库连接执行这些查询：

```go
// 读取SQL文件并执行
sqlFile, err := ioutil.ReadFile("query_database.sql")
if err != nil {
    log.Fatal(err)
}

_, err = db.Exec(string(sqlFile))
if err != nil {
    log.Fatal(err)
}
```

## 常用查询示例

### 查看所有公司
```sql
SELECT * FROM company ORDER BY created_at;
```

### 统计公司数量
```sql
SELECT COUNT(*) FROM company;
```

### 搜索特定公司
```sql
SELECT * FROM company WHERE name LIKE '%Google%';
```

### 按创建时间排序
```sql
SELECT * FROM company ORDER BY created_at DESC LIMIT 10;
```

## 注意事项

1. 执行脚本前请确保数据库连接正常
2. 管理脚本包含数据导出功能，请谨慎使用
3. 某些查询可能需要根据实际数据量调整LIMIT参数
4. 建议在生产环境使用前先在测试环境验证

## 扩展建议

可以根据需要添加更多查询：
- 按业务需求的自定义查询
- 性能分析查询
- 数据备份脚本
- 索引分析脚本
