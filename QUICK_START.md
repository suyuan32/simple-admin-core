# 🚀 Quick Start Guide - Inventory Management System

## ✅ Issue Fixed!
The `/home/data` directory error has been resolved! Your servers can now start.

## 🔧 Prerequisites Setup

### 1. Start Redis (Required)
```bash
# Using Docker
docker run -d -p 6379:6379 redis:alpine

# Or using Homebrew (macOS)
brew install redis
brew services start redis

# Or manually
redis-server
```

### 2. Setup Database (MySQL)
```bash
# Using Docker
docker run -d -p 3306:3306 \
  -e MYSQL_ROOT_PASSWORD=root \
  -e MYSQL_DATABASE=simple_admin \
  mysql:8.0

# Or install MySQL locally and create database
mysql -u root -p
CREATE DATABASE simple_admin;
```

### 3. Update Configuration
Edit the database credentials in both config files:
- `api/etc/core.yaml`
- `rpc/etc/core.yaml`

```yaml
DatabaseConf:
  Username: root  # or your username
  Password: root  # or your password
```

## 🚀 Start Your Servers

### Terminal 1 - API Server
```bash
cd api
go run .
```
Expected output: Server starts on `http://localhost:9100`

### Terminal 2 - RPC Server
```bash
cd rpc
go run .
```
Expected output: Server starts on `localhost:9101`

## 🧪 Test Your Inventory System

### Quick Health Check
```bash
curl http://localhost:9100/health
# Expected: {"status": "ok"}
```

### Create a Warehouse
```bash
curl -X POST http://localhost:9100/inventory/warehouse \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Main Warehouse",
    "location": "123 Industrial St",
    "capacity": 10000,
    "description": "Primary storage facility"
  }'
```

### Create a Product
```bash
curl -X POST http://localhost:9100/inventory/product \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Laptop Computer",
    "sku": "LT-001",
    "price": 1299.99,
    "category": "Electronics"
  }'
```

### Run Automated Tests
```bash
# Run comprehensive tests
./test-inventory-system.sh
```

## 🔧 Troubleshooting

### Redis Connection Issues
```bash
# Check if Redis is running
redis-cli ping
# Expected: PONG

# Start Redis service
brew services start redis
# or
docker start your-redis-container
```

### Database Connection Issues
```bash
# Test MySQL connection
mysql -h 127.0.0.1 -P 3306 -u root -p simple_admin

# Check MySQL service
brew services list | grep mysql
```

### Port Conflicts
```bash
# Check what's using ports 9100, 9101
lsof -i :9100
lsof -i :9101

# Change ports in config if needed
# api/etc/core.yaml: Port: 9102
# rpc/etc/core.yaml: ListenOn: 0.0.0.0:9103
```

### Configuration Issues
```bash
# Re-run the fix script
./fix-macos-issue.sh

# Check configuration files
cat api/etc/core.yaml | grep -A 5 Log
cat rpc/etc/core.yaml | grep -A 5 Log
```

## 📊 What Was Fixed

1. ✅ **macOS Compatibility**: Changed `/home/data/logs/*` to `./logs/*`
2. ✅ **Directory Creation**: Auto-created `logs/core/api` and `logs/core/rpc`
3. ✅ **Configuration Updates**: Fixed both API and RPC config files

## 🎯 Next Steps

1. **Start Redis and MySQL**
2. **Update database credentials in config files**
3. **Run both servers**
4. **Test the inventory APIs**
5. **Run the comprehensive test suite**

Your inventory management system is now ready! 🎉
