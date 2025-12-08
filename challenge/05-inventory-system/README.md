# Go Database Challenge: High-Performance Inventory System

## Overview
Design a concurrent inventory management system that handles high-volume writes, complex queries, and maintains strict transactional consistency.

## Requirements

### 1. Database Schema
```sql
-- Products table
CREATE TABLE products (
    id UUID PRIMARY KEY,
    sku VARCHAR(50) UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL,
    category VARCHAR(100) NOT NULL,
    price DECIMAL(10,2) NOT NULL,
    metadata JSONB,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_category (category),
    INDEX idx_sku (sku)
);

-- Inventory table (heavily written)
CREATE TABLE inventory (
    id UUID PRIMARY KEY,
    product_id UUID NOT NULL REFERENCES products(id) ON DELETE CASCADE,
    warehouse_id UUID NOT NULL,
    quantity INT NOT NULL CHECK (quantity >= 0),
    reserved_quantity INT NOT NULL DEFAULT 0 CHECK (reserved_quantity >= 0),
    available_quantity INT GENERATED ALWAYS AS (quantity - reserved_quantity) STORED,
    location_code VARCHAR(20),
    last_updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(product_id, warehouse_id),
    INDEX idx_warehouse (warehouse_id),
    INDEX idx_available_quantity (available_quantity)
);

-- Inventory transactions (append-only ledger)
CREATE TABLE inventory_transactions (
    id UUID PRIMARY KEY,
    product_id UUID NOT NULL,
    warehouse_id UUID NOT NULL,
    transaction_type VARCHAR(20) NOT NULL, -- 'RECEIPT', 'SALE', 'ADJUSTMENT', 'RESERVATION'
    quantity_change INT NOT NULL,
    reference_id VARCHAR(100),
    metadata JSONB,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_product_created (product_id, created_at),
    INDEX idx_warehouse_created (warehouse_id, created_at),
    INDEX idx_type_created (transaction_type, created_at)
);

-- Sales orders
CREATE TABLE sales_orders (
    id UUID PRIMARY KEY,
    customer_id UUID NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'PENDING', -- PENDING, PROCESSING, COMPLETED, CANCELLED
    total_amount DECIMAL(10,2),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_status_created (status, created_at),
    INDEX idx_customer (customer_id)
);

-- Order items with inventory reservation
CREATE TABLE order_items (
    id UUID PRIMARY KEY,
    order_id UUID NOT NULL REFERENCES sales_orders(id) ON DELETE CASCADE,
    product_id UUID NOT NULL,
    warehouse_id UUID NOT NULL,
    quantity INT NOT NULL CHECK (quantity > 0),
    unit_price DECIMAL(10,2) NOT NULL,
    reserved_quantity INT NOT NULL DEFAULT 0,
    status VARCHAR(20) DEFAULT 'RESERVED', -- RESERVED, FULFILLED, CANCELLED
    INDEX idx_product_warehouse (product_id, warehouse_id),
    INDEX idx_order_status (order_id, status)
);
```

### 2. Core Implementation Tasks

```go
package main

import (
    "context"
    "database/sql"
    "fmt"
    "sync"
    "time"
)

type InventoryService struct {
    db *sql.DB
    // Add any necessary fields
}

// Task 1: High-Concurrency Inventory Update
// Implement a method that handles concurrent inventory updates with deadlock prevention
func (s *InventoryService) BatchUpdateInventory(ctx context.Context, updates []InventoryUpdate) error {
    // Requirements:
    // 1. Handle 1000+ concurrent updates per second
    // 2. Prevent deadlocks using appropriate isolation levels or locking strategies
    // 3. Ensure atomicity across multiple inventory rows
    // 4. Log all transactions in inventory_transactions table
    // 5. Implement retry logic for deadlock victims
}

// Task 2: Complex Analytical Query
func (s *InventoryService) GetInventoryAnalytics(ctx context.Context, filters AnalyticsFilters) (*InventoryAnalytics, error) {
    // Requirements:
    // 1. Calculate real-time inventory metrics across multiple dimensions
    // 2. Include: turnover rate, stockout risk, category performance
    // 3. Use window functions for moving averages
    // 4. Optimize for read performance with appropriate indexes
    // 5. Handle pagination for large result sets
}

// Task 3: Distributed Inventory Reservation
func (s *InventoryService) ReserveInventoryForOrder(ctx context.Context, order *SalesOrder) error {
    // Requirements:
    // 1. Reserve inventory across multiple warehouses (nearest-first strategy)
    // 2. Implement two-phase commit for cross-warehouse reservations
    // 3. Set reservation timeout (TTL) with automatic rollback
    // 4. Handle partial fulfillment scenarios
    // 5. Prevent overselling with strict consistency
}

// Task 4: Bulk Data Processing Pipeline
func (s *InventoryService) ProcessInventoryFeed(ctx context.Context, feed <-chan InventoryFeedItem) (*ProcessingStats, error) {
    // Requirements:
    // 1. Process stream of inventory updates (CSV/JSON lines)
    // 2. Batch inserts/updates for efficiency
    // 3. Handle schema validation and data cleansing
    // 4. Implement circuit breaker for database load
    // 5. Provide real-time progress reporting
}

// Task 5: Data Consistency Monitor
func (s *InventoryService) VerifyDataConsistency(ctx context.Context) ([]ConsistencyIssue, error) {
    // Requirements:
    // 1. Cross-table consistency checks (e.g., sum of transactions = current inventory)
    // 2. Detect phantom inventory (available < 0)
    // 3. Identify orphaned reservations
    // 4. Implement checksum verification for critical data
    // 5. Generate repair scripts for detected issues
}
```

### 3. Performance Requirements

1. **Write Throughput**: Sustain 5,000 writes/sec on inventory table
2. **Query Latency**: < 100ms for complex analytical queries on 100M+ records
3. **Concurrency**: Handle 10,000 concurrent connections
4. **Data Consistency**: Zero tolerance for inventory overselling
5. **Availability**: 99.99% uptime during peak loads

### 4. Advanced Challenges

```go
// Challenge 1: Sharded Inventory Architecture
// Implement sharding strategy for global inventory across multiple regions
type ShardedInventoryService struct {
    shards map[string]*sql.DB // Region -> Database connection
}

func (s *ShardedInventoryService) GlobalInventoryLookup(ctx context.Context, productID string) ([]WarehouseInventory, error) {
    // Query across all regional shards in parallel
    // Aggregate results with consistency guarantees
    // Handle partial failures gracefully
}

// Challenge 2: Time-Series Inventory Forecasting
func (s *InventoryService) ForecastInventoryDemand(ctx context.Context, productID string, horizonDays int) (*DemandForecast, error) {
    // Use historical transaction data to predict future demand
    // Implement seasonal adjustment and trend detection
    // Generate reorder recommendations
}

// Challenge 3: Real-time Inventory Cache Coherence
type CachedInventoryService struct {
    db     *sql.DB
    cache  *redis.Client
    pubsub *redis.Client
}

func (s *CachedInventoryService) GetCachedInventory(ctx context.Context, productID string) (*InventorySnapshot, error) {
    // Implement read-through caching with write-behind
    // Handle cache invalidation across multiple instances
    // Prevent cache stampede with probabilistic early expiration
}
```

### 5. Evaluation Criteria

1. **Database Design** (20%)
   - Appropriate indexing strategy
   - Normalization vs denormalization balance
   - Partitioning strategy for large tables

2. **Transaction Management** (25%)
   - Proper isolation levels
   - Deadlock prevention and handling
   - Rollback strategies for complex operations

3. **Performance Optimization** (25%)
   - Query optimization and explain plan analysis
   - Connection pooling and resource management
   - Batch processing efficiency

4. **Concurrency Control** (20%)
   - Race condition prevention
   - Optimistic vs pessimistic locking choices
   - Distributed transaction coordination

5. **Production Readiness** (10%)
   - Monitoring and metrics
   - Error handling and logging
   - Migration strategies

### 6. Bonus Points

- Implement using **PostgreSQL** with its advanced features (CTEs, window functions, partial indexes)
- Add **prometheus metrics** for query performance monitoring
- Create **database migration scripts** using goose or migrate
- Implement **connection pooling** with pgbouncer configuration
- Add **load testing** scripts using k6 or vegeta
- Demonstrate **EXPLAIN ANALYZE** optimizations for all major queries

## Submission Guidelines

1. Provide complete, runnable Go code
2. Include comprehensive unit tests with >80% coverage
3. Add performance benchmarks for critical paths
4. Include SQL migration files
5. Provide a docker-compose setup for local testing
6. Document all design decisions and trade-offs
7. Include query execution plans for complex queries

## Expected Deliverables

- Complete Go implementation with all interfaces
- Database schema with indexes and constraints
- Migration scripts
- Load testing results
- Performance optimization report
- Production deployment considerations

This challenge tests deep understanding of database internals, Go concurrency patterns, and production-ready system design. Candidates should focus on data consistency, performance under load, and handling edge cases in distributed systems.