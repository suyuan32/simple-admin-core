<!-- Thanks for sending a pull request! Here are some tips for you:

1. If you want **faster** PR reviews, read how: https://github.com/kubesphere/community/blob/master/developer-guide/development/the-pr-author-guide-to-getting-through-code-review.md
2. In case you want to know how your PR got reviewed, read: https://github.com/kubesphere/community/blob/master/developer-guide/development/code-review-guide.md
3. Here are some coding conventions followed by KubeSphere community: https://github.com/kubesphere/community/blob/master/developer-guide/development/coding-conventions.md
-->

### What type of PR is this?
<!--
Add one of the following kinds:
/kind bug
/kind cleanup
/kind documentation
/kind feature
/kind design
/kind dependencies
/kind test

Optionally add one or more of the following kinds if applicable:
/kind api-change
/kind deprecation
/kind failing-test
/kind flake
/kind regression
-->
/kind feature
/kind api-change

### What this PR does / why we need it:

This PR implements a comprehensive inventory management system for the simple-admin-core project. The implementation includes:

#### **Core Features Added:**
- **Warehouse Management**: Full CRUD operations for warehouse entities
- **Product Management**: Complete product catalog with inventory tracking
- **Inventory Management**: Stock level monitoring and management
- **Stock Movement Tracking**: Audit trail for all inventory movements

#### **Technical Implementation:**
- **Database Layer**: Ent ORM integration with generated entities
- **API Layer**: RESTful API definitions with proper request/response models
- **gRPC Services**: Protobuf definitions and service implementations
- **Business Logic**: Clean architecture with separate logic handlers
- **Data Validation**: Input validation and error handling

### Which issue(s) this PR fixes:
<!--
Usage: `Fixes #<issue number>`, or `Fixes (paste link of issue)`.
_If PR is about `failing-tests or flakes`, please post the related issues/tests in a comment and do not use `Fixes`_*
-->
Fixes #

### Special notes for reviewers:

#### **Architecture Overview:**
```
API Layer (REST) → gRPC Services → Business Logic → Database (Ent ORM)
```

#### **Key Components:**
1. **Ent Entities**: Warehouse, Product, Inventory, StockMovement
2. **Protobuf Services**: CRUD operations for all entities
3. **Logic Handlers**: Clean separation of business logic
4. **API Definitions**: RESTful endpoints with OpenAPI specs

#### **Database Schema:**
- **Warehouse**: Location and capacity management
- **Product**: Catalog with pricing and descriptions
- **Inventory**: Stock levels per product per warehouse
- **StockMovement**: Audit trail for stock changes

### Does this PR introduced a user-facing change?
<!--
If no, just write "None" in the release-note block below.
If yes, a release note is required:
Enter your extended release note in the block below. If the PR requires additional action from users switching to the new release, include the string "action required".

For more information on release notes see: https://github.com/kubernetes/community/blob/master/contributors/guide/release-notes.md
-->
```release-note
Added comprehensive inventory management system including warehouse, product, inventory, and stock movement management with full CRUD operations via REST APIs and gRPC services.
```

### Additional documentation, usage docs, etc.:
<!--
This section can be blank if this pull request does not require a release note.
Please use the following format for linking documentation or pass the
section below:
- [KEP]: <link>
- [Usage]: <link>
- [Other doc]: <link>
-->
```docs

```

### Files Changed Summary:

#### **New Files Added (80+ files):**

**Entity Layer (Ent ORM):**
```
├── rpc/ent/inventory.go
├── rpc/ent/inventory/inventory.go
├── rpc/ent/inventory/where.go
├── rpc/ent/inventory_create.go
├── rpc/ent/inventory_delete.go
├── rpc/ent/inventory_query.go
├── rpc/ent/inventory_update.go
├── rpc/ent/product.go
├── rpc/ent/product/product.go
├── rpc/ent/product/where.go
├── rpc/ent/product_create.go
├── rpc/ent/product_delete.go
├── rpc/ent/product_query.go
├── rpc/ent/product_update.go
├── rpc/ent/stockmovement.go
├── rpc/ent/stockmovement/stockmovement.go
├── rpc/ent/stockmovement/where.go
├── rpc/ent/stockmovement_create.go
├── rpc/ent/stockmovement_delete.go
├── rpc/ent/stockmovement_query.go
├── rpc/ent/stockmovement_update.go
├── rpc/ent/warehouse.go
├── rpc/ent/warehouse/warehouse.go
├── rpc/ent/warehouse/where.go
├── rpc/ent/warehouse_create.go
├── rpc/ent/warehouse_delete.go
├── rpc/ent/warehouse_query.go
├── rpc/ent/warehouse_update.go
├── rpc/ent/schema/inventory.go
├── rpc/ent/schema/product.go
├── rpc/ent/schema/stock_movement.go
├── rpc/ent/schema/warehouse.go
```

**Business Logic Layer:**
```
├── rpc/internal/logic/inventory/create_inventory_logic.go
├── rpc/internal/logic/inventory/delete_inventory_logic.go
├── rpc/internal/logic/inventory/get_inventory_by_id_logic.go
├── rpc/internal/logic/inventory/get_inventory_list_logic.go
├── rpc/internal/logic/inventory/update_inventory_logic.go
├── rpc/internal/logic/product/create_product_logic.go
├── rpc/internal/logic/product/delete_product_logic.go
├── rpc/internal/logic/product/get_product_by_id_logic.go
├── rpc/internal/logic/product/get_product_list_logic.go
├── rpc/internal/logic/product/update_product_logic.go
├── rpc/internal/logic/stock_movement/create_stock_movement_logic.go
├── rpc/internal/logic/stock_movement/delete_stock_movement_logic.go
├── rpc/internal/logic/stock_movement/get_stock_movement_by_id_logic.go
├── rpc/internal/logic/stock_movement/get_stock_movement_list_logic.go
├── rpc/internal/logic/stock_movement/update_stock_movement_logic.go
├── rpc/internal/logic/warehouse/create_warehouse_logic.go
├── rpc/internal/logic/warehouse/delete_warehouse_logic.go
├── rpc/internal/logic/warehouse/get_warehouse_by_id_logic.go
├── rpc/internal/logic/warehouse/get_warehouse_list_logic.go
├── rpc/internal/logic/warehouse/update_warehouse_logic.go
```

**API Handler Layer:**
```
├── api/internal/handler/inventory/create_inventory_handler.go
├── api/internal/handler/inventory/delete_inventory_handler.go
├── api/internal/handler/inventory/get_inventory_by_id_handler.go
├── api/internal/handler/inventory/get_inventory_list_handler.go
├── api/internal/handler/inventory/update_inventory_handler.go
├── api/internal/handler/product/create_product_handler.go
├── api/internal/handler/product/delete_product_handler.go
├── api/internal/handler/product/get_product_by_id_handler.go
├── api/internal/handler/product/get_product_list_handler.go
├── api/internal/handler/product/update_product_handler.go
├── api/internal/handler/stock_movement/create_stock_movement_handler.go
├── api/internal/handler/stock_movement/delete_stock_movement_handler.go
├── api/internal/handler/stock_movement/get_stock_movement_by_id_handler.go
├── api/internal/handler/stock_movement/get_stock_movement_list_handler.go
├── api/internal/handler/stock_movement/update_stock_movement_handler.go
├── api/internal/handler/warehouse/create_warehouse_handler.go
├── api/internal/handler/warehouse/delete_warehouse_handler.go
├── api/internal/handler/warehouse/get_warehouse_by_id_handler.go
├── api/internal/handler/warehouse/get_warehouse_list_handler.go
├── api/internal/handler/warehouse/update_warehouse_handler.go
```

**API Definitions:**
```
├── api/desc/core/inventory.api
├── api/desc/core/product.api
├── api/desc/core/stock_movement.api
├── api/desc/core/warehouse.api
```

**Protobuf Definitions:**
```
├── rpc/desc/inventory.proto
├── rpc/desc/product.proto
├── rpc/desc/stockmovement.proto
├── rpc/desc/warehouse.proto
```

#### **Modified Files:**
```
├── api/desc/all.api                        # Updated API definitions
├── api/internal/handler/routes.go          # Added new API routes
├── api/internal/types/types.go             # Updated type definitions
├── core.json                               # Updated configuration
├── go.sum                                  # Updated dependencies
├── rpc/core.proto                          # Main protobuf schema
├── rpc/coreclient/core.go                  # Updated client code
├── rpc/ent/client.go                       # Updated Ent client
├── rpc/ent/ent.go                          # Updated Ent configuration
├── rpc/ent/hook/hook.go                    # Updated hooks
├── rpc/ent/intercept/intercept.go          # Updated interceptors
├── rpc/ent/migrate/schema.go               # Updated migration schema
├── rpc/ent/mutation.go                     # Updated mutations
├── rpc/ent/pagination.go                   # Updated pagination
├── rpc/ent/predicate/predicate.go          # Updated predicates
├── rpc/ent/runtime/runtime.go              # Updated runtime
├── rpc/ent/set_not_nil.go                  # Updated set operations
├── rpc/ent/tx.go                           # Updated transactions
├── rpc/internal/server/core_server.go      # Added new service handlers
├── rpc/types/core/core.pb.go               # Generated protobuf code
└── rpc/types/core/core_grpc.pb.go          # Generated gRPC code
```

### API Endpoints Added:

#### **Warehouse Management:**
- `POST /inventory/warehouse` - Create warehouse
- `GET /inventory/warehouse/{id}` - Get warehouse by ID
- `GET /inventory/warehouse` - List warehouses
- `PUT /inventory/warehouse/{id}` - Update warehouse
- `DELETE /inventory/warehouse/{id}` - Delete warehouse

#### **Product Management:**
- `POST /inventory/product` - Create product
- `GET /inventory/product/{id}` - Get product by ID
- `GET /inventory/product` - List products
- `PUT /inventory/product/{id}` - Update product
- `DELETE /inventory/product/{id}` - Delete product

#### **Inventory Management:**
- `POST /inventory/stock` - Create inventory record
- `GET /inventory/stock/{id}` - Get inventory by ID
- `GET /inventory/stock` - List inventory records
- `PUT /inventory/stock/{id}` - Update inventory
- `DELETE /inventory/stock/{id}` - Delete inventory

#### **Stock Movement Tracking:**
- `POST /inventory/movement` - Record stock movement
- `GET /inventory/movement/{id}` - Get movement by ID
- `GET /inventory/movement` - List stock movements
- `PUT /inventory/movement/{id}` - Update movement
- `DELETE /inventory/movement/{id}` - Delete movement

### Testing Recommendations:

1. **Unit Tests**: Test individual logic handlers
2. **Integration Tests**: Test full API workflows
3. **Database Tests**: Verify Ent schema and relationships
4. **API Tests**: Validate REST and gRPC endpoints

### Future Enhancements:
- Add inventory alerts and notifications
- Implement batch operations
- Add reporting and analytics
- Integrate with external systems
