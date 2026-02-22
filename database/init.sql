-- database/init.sql

-- 1. 创建用户信息表（对应个人界面功能）
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    role VARCHAR(20) DEFAULT 'member' -- user/admin 权限隔离
);

-- 2. 创建收集表（对应核心痛点功能）
CREATE TABLE IF NOT EXISTS collection_forms (
    id SERIAL PRIMARY KEY,
    title VARCHAR(200) NOT NULL,
    content JSONB, -- 存储动态的表单数据
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 3. 创建任务单表
CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL PRIMARY KEY,
    task_name TEXT NOT NULL,
    status VARCHAR(20) DEFAULT 'pending'
);