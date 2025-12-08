package main

type BaseModel struct {
	Id        int    ``
	CreatedAt string ``
	UpdatedAt string ``
}

type Product struct {
	BaseModel
	Sku      int ``
	Name     int ``
	Category int ``
	Price    int ``
	Metadata int ``
}

type InventoryUpdate struct{}

type AnalyticsFilters struct{}

type InventoryAnalytics struct{}

type SalesOrder struct{}

type ProcessingStats struct{}

type InventoryFeedItem struct{}

type ConsistencyIssue struct{}
