
### **Advanced Golang Concurrency & Channel Design Challenge**  

**Task:** Implement a **fault-tolerant concurrent data processing pipeline** in Go that:  
1. Reads integers from an input channel.  
2. Processes them through **three sequential stages** (defined below).  
3. Outputs the final result via an output channel.  
4. Handles errors gracefully and supports **clean shutdown**.  

---

### **Pipeline Stages**  
1. **Filter Stage**:  
   - Receives integers from `inputChan`.  
   - Sends **only even integers** to the next stage.  
   - **Error Case**: If a non-integer value is detected (e.g., via type assertion on a `interface{}` input), log an error and skip the value.  

2. **Transform Stage**:  
   - Receives even integers.  
   - Squares each value and sends it to the next stage.  
   - **Error Case**: If the squared value exceeds `1e18`, log an error and skip the value.  

3. **Aggregate Stage**:  
   - Receives squared values.  
   - Maintains a **running sum** of all valid values.  
   - Sends the **final sum** to `outputChan` when the input is closed.  

---

### **Requirements**  
#### **Core Functionality**  
- Use **channels** to connect all stages.  
- Each stage must run in its **own goroutine**.  
- Implement **backpressure handling** (e.g., use buffered channels).  
- Ensure **no goroutine leaks** during shutdown.  

#### **Error Handling**  
- If **any stage encounters an error**, it must:  
  - Log the error (using `log.Printf`).  
  - **Stop processing further data** in that stage.  
  - **Propagate the error upstream** (e.g., close the input channel of the previous stage).  
- The pipeline should **not crash** due to errors.  

#### **Graceful Shutdown**  
- Support termination via a `context.Context` (e.g., `context.Cancel`).  
- All goroutines must exit cleanly when shutdown is requested.  

#### **Testing & Validation**  
- Write **unit tests** covering:  
  - Normal operation (valid input).  
  - Error cases (e.g., invalid input, overflow).  
  - Graceful shutdown.  
- Include a `main` function demonstrating usage.  

#### **Code Quality**  
- Follow Go best practices (e.g., idiomatic concurrency, `range` loops, `defer`).  
- Add **documentation** explaining key design choices.  
- Use **named return values** and **error handling** in functions.  

---

### **Input/Output Channels**  
- **Input**: `inputChan <-chan interface{}` (values may include non-integers for testing).  
- **Output**: `outputChan chan int64` (emits the final sum or an error code).  

---

### **Submission Guidelines**  
1. **Source Code**:  
   - Implement the pipeline in a single Go module.  
   - Include a `main` function to demonstrate:  
     - Starting the pipeline.  
     - Sending input values.  
     - Triggering shutdown.  
   - Structure code into logical functions (e.g., `filterStage`, `transformStage`).  

2. **Tests**:  
   - Write tests using Go’s `testing` package.  
   - Cover edge cases (e.g., empty input, concurrent input, errors).  

3. **Documentation**:  
   - Add comments explaining:  
     - How errors are propagated.  
     - How shutdown is coordinated.  
     - Why specific concurrency patterns (e.g., buffered channels) were chosen.  

---

### **Scoring Criteria**  
- **Concurrency Expertise**: Proper use of channels, goroutines, and `select` (if applicable).  
- **Error Handling**: Robust propagation and logging without panics.  
- **Graceful Shutdown**: Context integration and resource cleanup.  
- **Testing**: Comprehensive test coverage.  
- **Code Quality**: Readability, idiomatic Go, and documentation.  

---

### **Example Scenario**  
**Input**: `[2, 3, "error", 4, 5]`  
- **Filter Stage**: Drops `"error"` and `3`, `5` (odd).  
- **Transform Stage**: Squares `2` → `4`, `4` → `16`.  
- **Aggregate Stage**: Sum = `4 + 16 = 20`.  
**Output**: `20`  

**Error Scenario**:  
If `Transform Stage` receives `1e9` (squared becomes `1e18` → allowed), but `2e9` would exceed `1e18` → log error and skip.  

---

**Goal**: Assess the candidate’s ability to design **robust, concurrent systems** in Go while adhering to best practices.  

--- 

This prompt evaluates **advanced concepts** like:  
- Channel directionality and buffer sizing.  
- Context-driven cancellation.  
- Error propagation across goroutines.  
- Testing concurrent code.  
- System design for fault tolerance.  

Let me know if you need further refinements!