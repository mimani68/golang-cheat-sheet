
### Golang `make` Cheat Sheet

#### **1. Slices**
- **Syntax**: `make([]Type, length, [capacity])`  
  *(Capacity is optional; defaults to `length`)*  
- **Example**:  
  ```go
  s := make([]int, 5)     // Creates a slice of 5 zeros: [0, 0, 0, 0, 0] Length: 5, Capacity: 5
  s := make([]int, 0, 5)  // OPTIMIZE WAY! Creates a empty slice: [] Length: 0, Capacity: 5 (Memory is reserved )
  ```

---

#### **2. Maps**
- **Syntax**: `make(map[KeyType]ValueType, [size])`  
  *(Size is optional; hints at initial allocation)*  
- **Example**:  
  ```go
  m := make(map[string]int) // Creates an empty map (not nil)
  ```

---

#### **3. Channels**
- **Syntax**: `make(chan Type, [buffer])`  
  *(Buffer is optional; unbuffered if omitted)*  
- **Example**:  
  ```go
  ch := make(chan int, 10) // Creates a buffered channel (capacity: 10)
  ```

---

### Key Notes:
- `make` **only works for slices, maps, and channels** (not for structs, pointers, etc.).
- Unlike `new`, `make` **initializes the underlying data structure** (e.g., allocates memory for slice elements).
- `new` returns a **pointer** to zero-initialized memory; `make` returns a **value** of the specified type.