# SHA1 HTTP Service

[![GitHub release](https://img.shields.io/github/v/release/yourname/sha1-service)](https://github.com/yourname/sha1-service/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yourname/sha1-service)](https://goreportcard.com/report/github.com/yourname/sha1-service)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![CI](https://github.com/yourname/sha1-service/actions/workflows/ci.yml/badge.svg)](https://github.com/yourname/sha1-service/actions/workflows/ci.yml)

基于 Go 语言实现的高性能 SHA1 哈希计算 HTTP 服务，提供 RESTful API 接口，支持字符串和文件哈希计算。

## 🚀 功能特性

- 🔑 字符串 SHA1 哈希计算
- 📁 文件 SHA1 校验计算
- ⚡ 高性能并发处理
- 📦 轻量级 Docker 镜像支持
- 🔒 HTTPS 安全支持
- 📊 Prometheus 监控指标
- ✅ 完整的单元测试覆盖

## 📦 快速开始

### 前提条件
- Go 1.20+
- Docker 20.10+ (可选)

### 安装运行

#### 方式1：从源码运行
```bash
# 克隆仓库
git clone https://github.com/yourname/sha1-service.git
cd sha1-service

# 启动服务（开发模式）
go run main.go

# 生产模式运行
ENV=production PORT=8080 go run main.go
