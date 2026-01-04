# Inventory Management System - Testing & Troubleshooting Guide

## 🚀 Quick Start

### Prerequisites
```bash
# Ensure Go is installed
go version

# Install dependencies
go mod tidy

# Set up database (if using PostgreSQL/MySQL)
# Update your database connection in config files
```

### Start the Application
```bash
# Start API server
cd api
go run .

# Start RPC server (in another terminal)
cd rpc
go run .

# Or use docker-compose if available
docker-compose up -d
```

### Run Tests
```bash
# Make test script executable
chmod +x test-inventory-system.sh

# Run comprehensive tests
./test-inventory-system.sh
```

---

## 🧪 Manual Testing Guide

### 1. Health Check
```bash
# Test API server
curl http://localhost:9100/health

# Expected response: {"status": "ok"}
```

### 2. Warehouse Operations

#### Create Warehouse
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

#### Get Warehouse
```bash
curl http://localhost:9100/inventory/warehouse/1
```

#### Update Warehouse
```bash
curl -X PUT http://localhost:9100/inventory/warehouse/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Main Warehouse Updated",
    "capacity": 15000
  }'
```

#### List Warehouses
```bash
curl http://localhost:9100/inventory/warehouse
```

#### Delete Warehouse
```bash
curl -X DELETE http://localhost:9100/inventory/warehouse/1
```

### 3. Product Operations

#### Create Product
```bash
curl -X POST http://localhost:9100/inventory/product \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Laptop Computer",
    "sku": "LT-001",
    "description": "High-performance laptop",
    "price": 1299.99,
    "category": "Electronics",
    "brand": "TechCorp"
  }'
```

#### Get Product
```bash
curl http://localhost:9100/inventory/product/1
```

#### List Products
```bash
curl http://localhost:9100/inventory/product
```

### 4. Inventory Operations

#### Create Inventory Record
```bash
curl -X POST http://localhost:9100/inventory/stock \
  -H "Content-Type: application/json" \
  -d '{
    "product_id": 1,
    "warehouse_id": 1,
    "quantity": 100,
    "min_stock_level": 10,
    "max_stock_level": 500
  }'
```

#### Update Stock Quantity
```bash
curl -X PUT http://localhost:9100/inventory/stock/1 \
  -H "Content-Type: application/json" \
  -d '{
    "quantity": 150
  }'
```

### 5. Stock Movement Operations

#### Record Stock Movement (IN)
```bash
curl -X POST http://localhost:9100/inventory/movement \
  -H "Content-Type: application/json" \
  -d '{
    "product_id": 1,
    "warehouse_id": 1,
    "movement_type": "IN",
    "quantity": 25,
    "reference": "PO-2024-001",
    "notes": "Received new stock"
  }'
```

#### Record Stock Movement (OUT)
```bash
curl -X POST http://localhost:9100/inventory/movement \
  -H "Content-Type: application/json" \
  -d '{
    "product_id": 1,
    "warehouse_id": 1,
    "movement_type": "OUT",
    "quantity": 10,
    "reference": "SO-2024-001",
    "notes": "Sold to customer"
  }'
```

---

## 🔧 Troubleshooting Guide

### Common Issues & Solutions

#### 1. API Server Not Starting
```bash
# Check for compilation errors
go build ./api/...

# Check database connection
# Verify your database credentials in config files

# Check port availability
lsof -i :9100
```

#### 2. Database Connection Issues
```bash
# Test database connection
# For PostgreSQL:
psql -h localhost -U your_user -d your_db

# For MySQL:
mysql -h localhost -u your_user -p your_db

# Check Ent schema generation
go generate ./rpc/ent
```

#### 3. gRPC Service Issues
```bash
# Test gRPC service connectivity
grpcurl -plaintext localhost:9090 list

# Check protobuf compilation
protoc --go_out=. --go-grpc_out=. rpc/desc/*.proto
```

#### 4. Ent ORM Issues
```bash
# Regenerate Ent entities
go generate ./rpc/ent

# Check for schema conflicts
go run rpc/ent/migrate/main.go
```

#### 5. Test Failures

##### Warehouse Tests Failing
```bash
# Check warehouse table schema
# Verify foreign key constraints
# Check database permissions
```

##### Product Tests Failing
```bash
# Verify product table structure
# Check unique constraints on SKU
# Validate data types
```

##### Inventory Tests Failing
```bash
# Ensure warehouse and product exist before inventory creation
# Check referential integrity constraints
# Verify quantity validation logic
```

##### Stock Movement Tests Failing
```bash
# Ensure inventory record exists
# Check movement type validation
# Verify quantity doesn't go negative
```

#### 6. Performance Issues
```bash
# Add database indexes
# Check query performance
# Monitor memory usage
```

---

## 🐛 Debugging Steps

### 1. Enable Debug Logging
```go
// Add to your main.go or server files
import "log"
log.SetFlags(log.LstdFlags | log.Lshortfile)

// Enable verbose logging
os.Setenv("LOG_LEVEL", "debug")
```

### 2. Database Query Debugging
```bash
# Enable SQL query logging in Ent
client := ent.Open("postgres", dsn,
    ent.Log(func(a ...any) {
        log.Println(a...)
    }),
)
```

### 3. API Request Debugging
```bash
# Use curl with verbose output
curl -v http://localhost:9100/inventory/warehouse

# Check request/response headers
curl -I http://localhost:9100/inventory/warehouse
```

### 4. gRPC Debugging
```bash
# Install grpcurl
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest

# List services
grpcurl -plaintext localhost:9090 list

# Call service methods
grpcurl -plaintext -d '{"id": 1}' localhost:9090 core.WarehouseService/GetWarehouse
```

---

## 📊 Test Coverage Checklist

### Warehouse Management ✅
- [x] Create warehouse
- [x] Get warehouse by ID
- [x] Update warehouse
- [x] List warehouses
- [x] Delete warehouse
- [x] Validation (required fields, capacity limits)

### Product Management ✅
- [x] Create product
- [x] Get product by ID
- [x] Update product
- [x] List products
- [x] Delete product
- [x] SKU uniqueness validation
- [x] Price validation

### Inventory Management ✅
- [x] Create inventory record
- [x] Update stock levels
- [x] Get inventory by ID
- [x] List inventory records
- [x] Delete inventory
- [x] Stock level alerts (min/max)
- [x] Referential integrity

### Stock Movement Tracking ✅
- [x] Record stock movements (IN/OUT)
- [x] Movement audit trail
- [x] Reference tracking
- [x] Quantity validation
- [x] Movement history

### Integration Tests ✅
- [x] End-to-end workflows
- [x] Cross-entity operations
- [x] Error handling
- [x] Data consistency

---

## 🚀 Production Deployment Checklist

### Pre-Deployment
- [ ] Database backup
- [ ] Schema migrations tested
- [ ] Performance benchmarks
- [ ] Security audit
- [ ] API documentation updated

### Deployment
- [ ] Zero-downtime deployment
- [ ] Database connection pooling
- [ ] Health checks configured
- [ ] Monitoring setup
- [ ] Rollback plan ready

### Post-Deployment
- [ ] Smoke tests passed
- [ ] Performance monitoring
- [ ] Error tracking
- [ ] User acceptance testing

---

## 📞 Support

If you encounter issues:

1. **Check the troubleshooting section above**
2. **Run the automated test script**: `./test-inventory-system.sh`
3. **Review application logs**
4. **Verify database connectivity**
5. **Check network configurations**

For additional help, please provide:
- Error messages
- Application logs
- Database schema
- Test output
- System information
