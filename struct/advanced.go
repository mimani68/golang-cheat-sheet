package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"
	"unsafe"
	"sync"
	"errors"
	"strconv"
)

// =============================================================================
// GOLANG STRUCT ADVANCED USE CASES CHEAT SHEET
// =============================================================================

// 1. STRUCT EMBEDDING & COMPOSITION
// Use Case: Create complex types through composition instead of inheritance
func structEmbedding() {
	fmt.Println("=== 1. STRUCT EMBEDDING & COMPOSITION ===")
	
	// Base types
	type Address struct {
		Street  string
		City    string
		Country string
	}
	
	type Contact struct {
		Email string
		Phone string
	}
	
	// Embedded struct with composition
	type Person struct {
		Name string
		Age  int
		Address        // Embedded struct
		Contact        // Embedded struct
		WorkAddress Address // Named field
	}
	
	// Method on embedded struct
	func (a Address) FullAddress() string {
		return fmt.Sprintf("%s, %s, %s", a.Street, a.City, a.Country)
	}
	
	// Method on main struct
	func (p Person) Profile() string {
		return fmt.Sprintf("%s (%d years old)", p.Name, p.Age)
	}
	
	// Create instance
	person := Person{
		Name: "John Doe",
		Age:  30,
		Address: Address{
			Street:  "123 Main St",
			City:    "New York",
			Country: "USA",
		},
		Contact: Contact{
			Email: "john@example.com",
			Phone: "+1-555-0123",
		},
		WorkAddress: Address{
			Street:  "456 Business Ave",
			City:    "New York",
			Country: "USA",
		},
	}
	
	// Access embedded fields directly
	fmt.Printf("Name: %s\n", person.Name)
	fmt.Printf("Street: %s\n", person.Street) // From embedded Address
	fmt.Printf("Email: %s\n", person.Email)   // From embedded Contact
	
	// Access methods from embedded structs
	fmt.Printf("Home Address: %s\n", person.Address.FullAddress())
	fmt.Printf("Work Address: %s\n", person.WorkAddress.FullAddress())
	fmt.Printf("Profile: %s\n", person.Profile())
	
	// Method promotion - embedded methods are promoted
	fmt.Printf("Direct call: %s\n", person.FullAddress()) // Calls Address.FullAddress()
}

// 2. STRUCT TAGS & REFLECTION
// Use Case: Use struct tags for serialization, validation, and metadata
func structTags() {
	fmt.Println("\n=== 2. STRUCT TAGS & REFLECTION ===")
	
	// Struct with various tags
	type User struct {
		ID       int       `json:"id" xml:"id" db:"user_id" validate:"required"`
		Name     string    `json:"name" xml:"name" db:"full_name" validate:"required,min=2,max=50"`
		Email    string    `json:"email" xml:"email" db:"email_address" validate:"required,email"`
		Age      int       `json:"age" xml:"age" db:"age" validate:"min=18,max=120"`
		Password string    `json:"-" xml:"-" db:"password_hash" validate:"required,min=8"`
		IsActive bool      `json:"is_active" xml:"is_active" db:"is_active" validate:""`
		Tags     []string  `json:"tags,omitempty" xml:"tags" db:"tags"`
		Created  time.Time `json:"created_at" xml:"created_at" db:"created_at"`
	}
	
	user := User{
		ID:       1,
		Name:     "Alice Johnson",
		Email:    "alice@example.com",
		Age:      25,
		Password: "secretpass123",
		IsActive: true,
		Tags:     []string{"admin", "developer"},
		Created:  time.Now(),
	}
	
	// JSON serialization (respects json tags)
	jsonData, _ := json.MarshalIndent(user, "", "  ")
	fmt.Printf("JSON output:\n%s\n", jsonData)
	
	// Reflection to read struct tags
	t := reflect.TypeOf(user)
	fmt.Println("\nStruct tag analysis:")
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		jsonTag := field.Tag.Get("json")
		dbTag := field.Tag.Get("db")
		validateTag := field.Tag.Get("validate")
		
		fmt.Printf("Field: %s\n", field.Name)
		fmt.Printf("  JSON tag: %s\n", jsonTag)
		fmt.Printf("  DB tag: %s\n", dbTag)
		fmt.Printf("  Validate tag: %s\n", validateTag)
		fmt.Println()
	}
	
	// Custom tag processor
	processCustomTags(user)
}

func processCustomTags(v interface{}) {
	t := reflect.TypeOf(v)
	val := reflect.ValueOf(v)
	
	fmt.Println("Custom tag processing:")
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldVal := val.Field(i)
		
		// Check if field is required
		validateTag := field.Tag.Get("validate")
		if strings.Contains(validateTag, "required") {
			if fieldVal.IsZero() {
				fmt.Printf("  WARNING: Required field '%s' is empty\n", field.Name)
			} else {
				fmt.Printf("  OK: Required field '%s' has value\n", field.Name)
			}
		}
	}
}

// 3. MEMORY LAYOUT & ALIGNMENT
// Use Case: Optimize memory usage by understanding struct layout
func memoryLayout() {
	fmt.Println("\n=== 3. MEMORY LAYOUT & ALIGNMENT ===")
	
	// Poorly aligned struct
	type BadStruct struct {
		A bool   // 1 byte
		B int64  // 8 bytes
		C bool   // 1 byte
		D int32  // 4 bytes
		E bool   // 1 byte
	}
	
	// Well-aligned struct
	type GoodStruct struct {
		B int64  // 8 bytes
		D int32  // 4 bytes
		A bool   // 1 byte
		C bool   // 1 byte
		E bool   // 1 byte
		// implicit padding: 1 byte
	}
	
	// Analyze sizes
	fmt.Printf("BadStruct size: %d bytes\n", unsafe.Sizeof(BadStruct{}))
	fmt.Printf("GoodStruct size: %d bytes\n", unsafe.Sizeof(GoodStruct{}))
	
	// Show field offsets
	bad := BadStruct{}
	good := GoodStruct{}
	
	fmt.Println("\nBadStruct field offsets:")
	fmt.Printf("  A: %d\n", unsafe.Offsetof(bad.A))
	fmt.Printf("  B: %d\n", unsafe.Offsetof(bad.B))
	fmt.Printf("  C: %d\n", unsafe.Offsetof(bad.C))
	fmt.Printf("  D: %d\n", unsafe.Offsetof(bad.D))
	fmt.Printf("  E: %d\n", unsafe.Offsetof(bad.E))
	
	fmt.Println("\nGoodStruct field offsets:")
	fmt.Printf("  B: %d\n", unsafe.Offsetof(good.B))
	fmt.Printf("  D: %d\n", unsafe.Offsetof(good.D))
	fmt.Printf("  A: %d\n", unsafe.Offsetof(good.A))
	fmt.Printf("  C: %d\n", unsafe.Offsetof(good.C))
	fmt.Printf("  E: %d\n", unsafe.Offsetof(good.E))
	
	// Zero-sized structs
	type Empty struct{}
	fmt.Printf("\nEmpty struct size: %d bytes\n", unsafe.Sizeof(Empty{}))
	
	// Struct with zero-sized field
	type WithEmpty struct {
		Data  int
		Empty struct{}
	}
	fmt.Printf("WithEmpty struct size: %d bytes\n", unsafe.Sizeof(WithEmpty{}))
}

// 4. ANONYMOUS STRUCTS & FIELDS
// Use Case: Use anonymous structs for temporary data structures
func anonymousStructs() {
	fmt.Println("\n=== 4. ANONYMOUS STRUCTS & FIELDS ===")
	
	// Anonymous struct for configuration
	config := struct {
		Host     string
		Port     int
		Database string
		Options  struct {
			MaxConnections int
			Timeout        time.Duration
		}
	}{
		Host:     "localhost",
		Port:     5432,
		Database: "myapp",
		Options: struct {
			MaxConnections int
			Timeout        time.Duration
		}{
			MaxConnections: 100,
			Timeout:        30 * time.Second,
		},
	}
	
	fmt.Printf("Config: %+v\n", config)
	
	// Anonymous struct for API response
	response := struct {
		Status  string `json:"status"`
		Data    interface{} `json:"data"`
		Message string `json:"message,omitempty"`
	}{
		Status: "success",
		Data: map[string]interface{}{
			"users": []string{"Alice", "Bob", "Charlie"},
			"total": 3,
		},
		Message: "Users fetched successfully",
	}
	
	responseJSON, _ := json.MarshalIndent(response, "", "  ")
	fmt.Printf("API Response:\n%s\n", responseJSON)
	
	// Anonymous struct slice for table data
	tableData := []struct {
		Name  string
		Age   int
		Score float64
	}{
		{"Alice", 25, 95.5},
		{"Bob", 30, 87.2},
		{"Charlie", 22, 91.8},
	}
	
	fmt.Println("\nTable Data:")
	for _, row := range tableData {
		fmt.Printf("  %s: %d years old, score: %.1f\n", row.Name, row.Age, row.Score)
	}
	
	// Anonymous fields in struct
	type Person struct {
		string      // Anonymous field
		int         // Anonymous field
		Address     // Named embedded struct
	}
	
	person := Person{
		string:  "John Doe",
		int:     30,
		Address: Address{Street: "123 Main St", City: "NYC", Country: "USA"},
	}
	
	fmt.Printf("\nPerson with anonymous fields: %+v\n", person)
	fmt.Printf("Name (anonymous string): %s\n", person.string)
	fmt.Printf("Age (anonymous int): %d\n", person.int)
}

// 5. STRUCT COMPARISON & EQUALITY
// Use Case: Understand struct comparison rules and implement custom equality
func structComparison() {
	fmt.Println("\n=== 5. STRUCT COMPARISON & EQUALITY ===")
	
	// Comparable struct
	type Point struct {
		X, Y int
	}
	
	p1 := Point{1, 2}
	p2 := Point{1, 2}
	p3 := Point{2, 3}
	
	fmt.Printf("p1 == p2: %v\n", p1 == p2)
	fmt.Printf("p1 == p3: %v\n", p1 == p3)
	
	// Non-comparable struct (contains slice)
	type NonComparable struct {
		Name string
		Data []int
	}
	
	nc1 := NonComparable{Name: "test", Data: []int{1, 2, 3}}
	nc2 := NonComparable{Name: "test", Data: []int{1, 2, 3}}
	
	// This would cause compile error:
	// fmt.Printf("nc1 == nc2: %v\n", nc1 == nc2)
	
	// Custom equality function
	fmt.Printf("nc1 equals nc2: %v\n", nonComparableEqual(nc1, nc2))
	
	// Using reflection for deep equality
	fmt.Printf("nc1 deep equal nc2: %v\n", reflect.DeepEqual(nc1, nc2))
	
	// Struct with comparable and non-comparable fields
	type Mixed struct {
		ID   int
		Name string
		Tags []string
	}
	
	m1 := Mixed{ID: 1, Name: "Alice", Tags: []string{"admin"}}
	m2 := Mixed{ID: 1, Name: "Alice", Tags: []string{"admin"}}
	
	fmt.Printf("m1 deep equal m2: %v\n", reflect.DeepEqual(m1, m2))
	
	// Custom equality with specific logic
	fmt.Printf("m1 custom equal m2: %v\n", mixedEqual(m1, m2))
}

func nonComparableEqual(a, b NonComparable) bool {
	if a.Name != b.Name {
		return false
	}
	if len(a.Data) != len(b.Data) {
		return false
	}
	for i, v := range a.Data {
		if v != b.Data[i] {
			return false
		}
	}
	return true
}

func mixedEqual(a, b Mixed) bool {
	return a.ID == b.ID && 
		   a.Name == b.Name && 
		   reflect.DeepEqual(a.Tags, b.Tags)
}

// 6. STRUCT SERIALIZATION & DESERIALIZATION
// Use Case: Advanced JSON/XML serialization with custom marshaling
func structSerialization() {
	fmt.Println("\n=== 6. STRUCT SERIALIZATION & DESERIALIZATION ===")
	
	// Custom marshaling struct
	type CustomTime struct {
		time.Time
	}
	
	// Custom JSON marshaling
	func (ct CustomTime) MarshalJSON() ([]byte, error) {
		return json.Marshal(ct.Format("2006-01-02 15:04:05"))
	}
	
	func (ct *CustomTime) UnmarshalJSON(data []byte) error {
		var s string
		if err := json.Unmarshal(data, &s); err != nil {
			return err
		}
		t, err := time.Parse("2006-01-02 15:04:05", s)
		if err != nil {
			return err
		}
		ct.Time = t
		return nil
	}
	
	type Event struct {
		ID          int        `json:"id"`
		Name        string     `json:"name"`
		Description string     `json:"description,omitempty"`
		StartTime   CustomTime `json:"start_time"`
		EndTime     CustomTime `json:"end_time"`
		Tags        []string   `json:"tags,omitempty"`
		Private     bool       `json:"-"`
	}
	
	event := Event{
		ID:          1,
		Name:        "Go Conference",
		Description: "Annual Go programming conference",
		StartTime:   CustomTime{time.Now()},
		EndTime:     CustomTime{time.Now().Add(8 * time.Hour)},
		Tags:        []string{"golang", "programming", "conference"},
		Private:     true,
	}
	
	// Serialize to JSON
	jsonData, _ := json.MarshalIndent(event, "", "  ")
	fmt.Printf("Serialized Event:\n%s\n", jsonData)
	
	// Deserialize from JSON
	var deserializedEvent Event
	json.Unmarshal(jsonData, &deserializedEvent)
	fmt.Printf("Deserialized Event: %+v\n", deserializedEvent)
	
	// Custom field transformation
	type TransformStruct struct {
		Name    string `json:"name"`
		Age     int    `json:"age"`
		RawData string `json:"-"`
	}
	
	// Implement custom marshaling with transformation
	func (ts TransformStruct) MarshalJSON() ([]byte, error) {
		type Alias TransformStruct
		return json.Marshal(&struct {
			*Alias
			AgeGroup string `json:"age_group"`
		}{
			Alias:    (*Alias)(&ts),
			AgeGroup: getAgeGroup(ts.Age),
		})
	}
	
	func getAgeGroup(age int) string {
		switch {
		case age < 18:
			return "minor"
		case age < 65:
			return "adult"
		default:
			return "senior"
		}
	}
	
	ts := TransformStruct{Name: "Alice", Age: 30, RawData: "sensitive"}
	tsJSON, _ := json.MarshalIndent(ts, "", "  ")
	fmt.Printf("Transformed struct:\n%s\n", tsJSON)
}

// 7. STRUCT VALIDATION & CONSTRAINTS
// Use Case: Implement comprehensive struct validation
func structValidation() {
	fmt.Println("\n=== 7. STRUCT VALIDATION & CONSTRAINTS ===")
	
	// Validatable struct
	type User struct {
		ID       int    `validate:"required,min=1"`
		Username string `validate:"required,min=3,max=20,alphanum"`
		Email    string `validate:"required,email"`
		Age      int    `validate:"min=18,max=120"`
		Website  string `validate:"url"`
	}
	
	// Validation interface
	type Validator interface {
		Validate() error
	}
	
	// Implement validation
	func (u User) Validate() error {
		var errors []string
		
		// ID validation
		if u.ID <= 0 {
			errors = append(errors, "ID must be positive")
		}
		
		// Username validation
		if len(u.Username) < 3 || len(u.Username) > 20 {
			errors = append(errors, "Username must be 3-20 characters")
		}
		
		// Email validation (simple check)
		if !strings.Contains(u.Email, "@") {
			errors = append(errors, "Invalid email format")
		}
		
		// Age validation
		if u.Age < 18 || u.Age > 120 {
			errors = append(errors, "Age must be between 18 and 120")
		}
		
		// Website validation
		if u.Website != "" && !strings.HasPrefix(u.Website, "http") {
			errors = append(errors, "Website must start with http")
		}
		
		if len(errors) > 0 {
			return fmt.Errorf("validation errors: %s", strings.Join(errors, ", "))
		}
		
		return nil
	}
	
	// Test validation
	validUser := User{
		ID:       1,
		Username: "alice123",
		Email:    "alice@example.com",
		Age:      25,
		Website:  "https://alice.dev",
	}
	
	invalidUser := User{
		ID:       0,
		Username: "al",
		Email:    "invalid-email",
		Age:      15,
		Website:  "invalid-url",
	}
	
	fmt.Printf("Valid user validation: %v\n", validUser.Validate())
	fmt.Printf("Invalid user validation: %v\n", invalidUser.Validate())
	
	// Generic validation function using reflection
	validateUsingReflection(validUser)
	validateUsingReflection(invalidUser)
}

func validateUsingReflection(v interface{}) {
	val := reflect.ValueOf(v)
	typ := reflect.TypeOf(v)
	
	fmt.Printf("\nReflection validation for %T:\n", v)
	
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)
		tag := fieldType.Tag.Get("validate")
		
		if tag == "" {
			continue
		}
		
		rules := strings.Split(tag, ",")
		for _, rule := range rules {
			if rule == "required" && field.IsZero() {
				fmt.Printf("  ERROR: %s is required\n", fieldType.Name)
			}
			if strings.HasPrefix(rule, "min=") {
				min, _ := strconv.Atoi(rule[4:])
				if field.Kind() == reflect.String && field.Len() < min {
					fmt.Printf("  ERROR: %s is too short (min: %d)\n", fieldType.Name, min)
				}
				if field.Kind() == reflect.Int && int(field.Int()) < min {
					fmt.Printf("  ERROR: %s is too small (min: %d)\n", fieldType.Name, min)
				}
			}
		}
	}
}

// 8. BUILDER PATTERN WITH STRUCTS
// Use Case: Implement builder pattern for complex struct construction
func builderPattern() {
	fmt.Println("\n=== 8. BUILDER PATTERN WITH STRUCTS ===")
	
	// Product struct
	type Database struct {
		Host     string
		Port     int
		Name     string
		Username string
		Password string
		SSLMode  string
		Timeout  time.Duration
		MaxConns int
	}
	
	// Builder struct
	type DatabaseBuilder struct {
		db Database
	}
	
	// Constructor
	func NewDatabaseBuilder() *DatabaseBuilder {
		return &DatabaseBuilder{
			db: Database{
				Port:     5432,
				SSLMode:  "prefer",
				Timeout:  30 * time.Second,
				MaxConns: 10,
			},
		}
	}
	
	// Builder methods
	func (b *DatabaseBuilder) Host(host string) *DatabaseBuilder {
		b.db.Host = host
		return b
	}
	
	func (b *DatabaseBuilder) Port(port int) *DatabaseBuilder {
		b.db.Port = port
		return b
	}
	
	func (b *DatabaseBuilder) Name(name string) *DatabaseBuilder {
		b.db.Name = name
		return b
	}
	
	func (b *DatabaseBuilder) Credentials(username, password string) *DatabaseBuilder {
		b.db.Username = username
		b.db.Password = password
		return b
	}
	
	func (b *DatabaseBuilder) SSLMode(mode string) *DatabaseBuilder {
		b.db.SSLMode = mode
		return b
	}
	
	func (b *DatabaseBuilder) Timeout(timeout time.Duration) *DatabaseBuilder {
		b.db.Timeout = timeout
		return b
	}
	
	func (b *DatabaseBuilder) MaxConnections(max int) *DatabaseBuilder {
		b.db.MaxConns = max
		return b
	}
	
	func (b *DatabaseBuilder) Build() Database {
		return b.db
	}
	
	// Usage
	db := NewDatabaseBuilder().
		Host("localhost").
		Port(5432).
		Name("myapp").
		Credentials("user", "pass").
		SSLMode("require").
		Timeout(60 * time.Second).
		MaxConnections(50).
		Build()
	
	fmt.Printf("Built database config: %+v\n", db)
	
	// Functional options pattern (alternative to builder)
	type ServerOption func(*Server)
	
	type Server struct {
		Host string
		Port int
		TLS  bool
	}
	
	func WithHost(host string) ServerOption {
		return func(s *Server) {
			s.Host = host
		}
	}
	
	func WithPort(port int) ServerOption {
		return func(s *Server) {
			s.Port = port
		}
	}
	
	func WithTLS(enabled bool) ServerOption {
		return func(s *Server) {
			s.TLS = enabled
		}
	}
	
	func NewServer(opts ...ServerOption) *Server {
		s := &Server{
			Host: "localhost",
			Port: 8080,
			TLS:  false,
		}
		
		for _, opt := range opts {
			opt(s)
		}
		
		return s
	}
	
	// Usage
	server := NewServer(
		WithHost("example.com"),
		WithPort(443),
		WithTLS(true),
	)
	
	fmt.Printf("Built server config: %+v\n", server)
}

// 9. FACTORY PATTERN WITH STRUCTS
// Use Case: Create different struct types based on parameters
func factoryPattern() {
	fmt.Println("\n=== 9. FACTORY PATTERN WITH STRUCTS ===")
	
	// Shape interface
	type Shape interface {
		Area() float64
		Perimeter() float64
		String() string
	}
	
	// Rectangle struct
	type Rectangle struct {
		Width  float64
		Height float64
	}
	
	func (r Rectangle) Area() float64 {
		return r.Width * r.Height
	}
	
	func (r Rectangle) Perimeter() float64 {
		return 2 * (r.Width + r.Height)
	}
	
	func (r Rectangle) String() string {
		return fmt.Sprintf("Rectangle(%.2f x %.2f)", r.Width, r.Height)
	}
	
	// Circle struct
	type Circle struct {
		Radius float64
	}
	
	func (c Circle) Area() float64 {
		return 3.14159 * c.Radius * c.Radius
	}
	
	func (c Circle) Perimeter() float64 {
		return 2 * 3.14159 * c.Radius
	}
	
	func (c Circle) String() string {
		return fmt.Sprintf("Circle(radius: %.2f)", c.Radius)
	}
	
	// Factory function
	func CreateShape(shapeType string, params ...float64) (Shape, error) {
		switch shapeType {
		case "rectangle":
			if len(params) != 2 {
				return nil, errors.New("rectangle requires width and height")
			}
			return Rectangle{Width: params[0], Height: params[1]}, nil
		case "circle":
			if len(params) != 1 {
				return nil, errors.New("circle requires radius")
			}
			return Circle{Radius: params[0]}, nil
		default:
			return nil, fmt.Errorf("unknown shape type: %s", shapeType)
		}
	}
	
	// Usage
	shapes := []struct {
		shapeType string
		params    []float64
	}{
		{"rectangle", []float64{10, 5}},
		{"circle", []float64{3}},
		{"rectangle", []float64{7, 7}},
	}
	
	for _, s := range shapes {
		shape, err := CreateShape(s.shapeType, s.params...)
		if err != nil {
			fmt.Printf("Error creating shape: %v\n", err)
			continue
		}
		
		fmt.Printf("Created: %s\n", shape.String())
		fmt.Printf("  Area: %.2f\n", shape.Area())
		fmt.Printf("  Perimeter: %.2f\n", shape.Perimeter())
	}
	
	// Registry-based factory
	type ShapeFactory struct {
		creators map[string]func([]float64) (Shape, error)
	}
	
	func NewShapeFactory() *ShapeFactory {
		return &ShapeFactory{
			creators: make(map[string]func([]float64) (Shape, error)),
		}
	}
	
	func (sf *ShapeFactory) Register(name string, creator func([]float64) (Shape, error)) {
		sf.creators[name] = creator
	}
	
	func (sf *ShapeFactory) Create(name string, params []float64) (Shape, error) {
		creator, exists := sf.creators[name]
		if !exists {
			return nil, fmt.Errorf("unknown shape: %s", name)
		}
		return creator(params)
	}
	
	// Setup factory
	factory := NewShapeFactory()
	factory.Register("rectangle", func(params []float64) (Shape, error) {
		if len(params) != 2 {
			return nil, errors.New("rectangle requires width and height")
		}
		return Rectangle{Width: params[0], Height: params[1]}, nil
	})
	factory.Register("circle", func(params []float64) (Shape, error) {
		if len(params) != 1 {
			return nil, errors.New("circle requires radius")
		}
		return Circle{Radius: params[0]}, nil
	})
	
	// Use registry factory
	fmt.Println("\nUsing registry factory:")
	rect, _ := factory.Create("rectangle", []float64{4, 6})
	fmt.Printf("Created via registry: %s, Area: %.2f\n", rect.String(), rect.Area())
}

// 10. THREAD-SAFE STRUCT OPERATIONS
// Use Case: Implement thread-safe struct operations with sync primitives
func threadSafeStructs() {
	fmt.Println("\n=== 10. THREAD-SAFE STRUCT OPERATIONS ===")
	
	// Thread-safe counter
	type SafeCounter struct {
		mu    sync.RWMutex
		value int
	}
	
	func (sc *SafeCounter) Increment() {
		sc.mu.Lock()
		defer sc.mu.Unlock()
		sc.value++
	}
	
	func (sc *SafeCounter) Decrement() {
		sc.mu.Lock()
		defer sc.mu.Unlock()
		sc.value--
	}
	
	func (sc *SafeCounter) Value() int {
		sc.mu.RLock()
		defer sc.mu.RUnlock()
		return sc.value
	}
	
	// Thread-safe map
	type SafeMap struct {
		mu   sync.RWMutex
		data map[string]interface{}
	}
	
	func NewSafeMap() *SafeMap {
		return &SafeMap{
			data: make(map[string]interface{}),
		}
	}
	
	func (sm *SafeMap) Set(key string, value interface{}) {
		sm.mu.Lock()
		defer sm.mu.Unlock()
		sm.data[key] = value
	}
	
	func (sm *SafeMap) Get(key string) (interface{}, bool) {
		sm.mu.RLock()
		defer sm.mu.RUnlock()
		val, exists := sm.data[key]
		return val, exists
	}
	
	func (sm *SafeMap) Delete(key string) {
		sm.mu.Lock()
		defer sm.mu.Unlock()
		delete(sm.data, key)
	}
	
	func (sm *SafeMap) Keys() []string {
		sm.mu.RLock()
		defer sm.mu.RUnlock()
		keys := make([]string, 0, len(sm.data))
		for k := range sm.data {
			keys = append(keys, k)
		}
		return keys
	}
	
	// Usage with goroutines
	counter := &SafeCounter{}
	safeMap := NewSafeMap()
	
	var wg sync.WaitGroup
	
	// Concurrent counter operations
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				counter.Increment()
				safeMap.Set(fmt.Sprintf("key_%d_%d", id, j), fmt.Sprintf("value_%d_%d", id, j))
			}
		}(i)
	}
	
	wg.Wait()
	
	fmt.Printf("Final counter value: %d\n", counter.Value())
	fmt.Printf("SafeMap keys count: %d\n", len(safeMap.Keys()))
	
	// Example with sync.Once for initialization
	type SingletonConfig struct {
		once sync.Once
		data map[string]string
	}
	
	var config SingletonConfig
	
	func (sc *SingletonConfig) Initialize() {
		sc.once.Do(func() {
			fmt.Println("Initializing configuration...")
			sc.data = map[string]string{
				"host": "localhost",
				"port": "8080",
			}
		})
	}
	
	func (sc *SingletonConfig) Get(key string) string {
		sc.Initialize()
		return sc.data[key]
	}
	
	// Test singleton
	fmt.Printf("Config host: %s\n", config.Get("host"))
	fmt.Printf("Config port: %s\n", config.Get("port"))
	
	// Atomic operations for simple cases
	type AtomicCounter struct {
		count int64
	}
	
	func (ac *AtomicCounter) Add(delta int64) {
		// In real code, use sync/atomic
		// atomic.AddInt64(&ac.count, delta)
		// For demonstration:
		ac.count += delta
	}
	
	func (ac *AtomicCounter) Load() int64 {
		// In real code, use sync/atomic
		// return atomic.LoadInt64(&ac.count)
		// For demonstration:
		return ac.count
	}
	
	atomicCounter := &AtomicCounter{}
	atomicCounter.Add(42)
	fmt.Printf("Atomic counter value: %d\n", atomicCounter.Load())
}

// Helper structs for embedding example
type Address struct {
	Street  string
	City    string
	Country string
}

// =============================================================================
// MAIN FUNCTION - DEMONSTRATES ALL USE CASES
// =============================================================================

func main() {
	fmt.Println("GOLANG STRUCT ADVANCED USE CASES")
	fmt.Println("=" + strings.Repeat("=", 50))
	
	// Run all examples
	structEmbedding()
	structTags()
	memoryLayout()
	anonymousStructs()
	structComparison()
	structSerialization()
	structValidation()
	builderPattern()
	factoryPattern()
	threadSafeStructs()
	
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("All struct examples completed!")
}
