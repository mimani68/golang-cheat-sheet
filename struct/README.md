# Compare JAVA, C# and Golang in term of Class paradigm 

| Feature | Java | C# (ASP.NET) | Go (Golang) |
| :--- | :--- | :--- | :--- |
| **Primary Paradigm** | Class-based Object-Oriented | Class-based Object-Oriented (Multi-paradigm) | Structural / Composition-based |
| **Core Type Keyword** | `class`, `interface`, `enum`, `record` | `class`, `struct`, `interface`, `record` | `struct`, `interface` |
| **Type Nature** | Classes are Reference Types (Heap) | Classes are Reference Types; Structs are Value Types (Stack/Inline) | Structs can be Value types or Pointer types (Reference to Heap) |
| **Inheritance** | **Single Class Inheritance** (`extends`)<br>Multiple Interface Implementation (`implements`) | **Single Class Inheritance** (`:`)<br>Multiple Interface Implementation (`:`) | **No Inheritance** (Does not extend)<br>Uses **Struct Embedding** (Composition) |
| **Interface Implementation** | **Explicit**: Must declare `implements Interface` | **Explicit**: Must declare `: Interface`<br>Supports Explicit/Implicit implementation | **Implicit** (Duck Typing): Satisfies interface if methods match signature, regardless of declaration |
| **Polymorphism** | Subtyping via Inheritance and Interfaces | Subtyping, Generics, Dynamic typing (`dynamic`) | Interface satisfaction (Duck typing), Type assertion |
| **Encapsulation / Visibility** | `public`, `private`, `protected`, package-private | `public`, `private`, `protected`, `internal`, `protected internal` | **Exported** (Uppercase) vs **Unexported** (Lowercase)<br>No `protected` keyword |
| **Properties / Accessors** | Getters and Setters (Boilerplate methods) | First-class **Properties** (`{ get; set; }`) | Exported Fields or Getter Methods (No built-in property syntax) |
| **Method Definition** | Defined inside the class body | Defined inside the class/struct body | Defined outside the struct using **Receivers** (`func (r *T) Method()`) |
| **Constructor** | Constructor name matches class name (`MyClass()`). Supports overloading. | Constructor name matches class name. Supports **Primary Constructors**. | No built-in constructor. Convention: `NewTypeName()` function. |
| **Generics** | Yes (Type Erasure) | Yes (Reified, Runtime type info) | Yes (Introduced in v1.18) |
| **Error Handling** | **Exceptions** (`try-catch`, `throws`) | **Exceptions** (`try-catch`) | **Error Values** (returned as last argument) and `panic/recover` |
| **Static Members** | `static` keyword | `static` keyword | No `static` keyword. Uses package-level variables or functions. |
| **Abstract Types** | `abstract` class/methods, `interface` | `abstract` class/methods, `interface` | Interface definition with method signatures (no bodies usually, unless Go 1.18+) |
| **Final / Sealed** | `final` keyword (prevents overriding) | `sealed` keyword (prevents inheritance) | No keyword. Preventing override is not applicable (no inheritance). |
| **Operator Overloading** | Not supported | Supported (e.g., `+`, `==`) | Not supported |
| **Extension Methods** | Not supported (Default methods in interfaces) | Supported (`this` parameter) | Not supported (but embedding provides similar behavior) |
| **Inner / Nested Classes** | Static and Non-static inner classes | Nested classes | No nested classes (only types in same package) |
| **Object Initialization** | `new Type(args)` | `new Type(args)` or `new Type { Prop = val }` (Object Initializer) | `Type{Field: val}` (Composite Literal) |
| **Concurrency Primitives** | `synchronized`, `volatile`, `java.util.concurrent` | `lock`, `Monitor`, `async/await`, `Task` | `go` keyword (Goroutines), `channels`, `sync.Mutex` |
| Feature | Java | C# (ASP.NET) | Go (Golang) |
| :--- | :--- | :--- | :--- |
| **Primary Paradigm** | Class-based Object-Oriented | Class-based Object-Oriented (Multi-paradigm) | Structural / Composition-based |
| **Core Type Keyword** | `class`, `interface`, `enum`, `record` | `class`, `struct`, `interface`, `record` | `struct`, `interface` |
| **Type Nature** | Classes are Reference Types (Heap) | Classes are Reference Types; Structs are Value Types (Stack/Inline) | Structs can be Value types or Pointer types (Reference to Heap) |
| **Inheritance** | **Single Class Inheritance** (`extends`)<br>Multiple Interface Implementation (`implements`) | **Single Class Inheritance** (`:`)<br>Multiple Interface Implementation (`:`) | **No Inheritance** (Does not extend)<br>Uses **Struct Embedding** (Composition) |
| **Interface Implementation** | **Explicit**: Must declare `implements Interface` | **Explicit**: Must declare `: Interface`<br>Supports Explicit/Implicit implementation | **Implicit** (Duck Typing): Satisfies interface if methods match signature, regardless of declaration |
| **Polymorphism** | Subtyping via Inheritance and Interfaces | Subtyping, Generics, Dynamic typing (`dynamic`) | Interface satisfaction (Duck typing), Type assertion |
| **Encapsulation / Visibility** | `public`, `private`, `protected`, package-private | `public`, `private`, `protected`, `internal`, `protected internal` | **Exported** (Uppercase) vs **Unexported** (Lowercase)<br>No `protected` keyword |
| **Properties / Accessors** | Getters and Setters (Boilerplate methods) | First-class **Properties** (`{ get; set; }`) | Exported Fields or Getter Methods (No built-in property syntax) |
| **Method Definition** | Defined inside the class body | Defined inside the class/struct body | Defined outside the struct using **Receivers** (`func (r *T) Method()`) |
| **Constructor** | Constructor name matches class name (`MyClass()`). Supports overloading. | Constructor name matches class name. Supports **Primary Constructors**. | No built-in constructor. Convention: `NewTypeName()` function. |
| **Generics** | Yes (Type Erasure) | Yes (Reified, Runtime type info) | Yes (Introduced in v1.18) |
| **Error Handling** | **Exceptions** (`try-catch`, `throws`) | **Exceptions** (`try-catch`) | **Error Values** (returned as last argument) and `panic/recover` |
| **Static Members** | `static` keyword | `static` keyword | No `static` keyword. Uses package-level variables or functions. |
| **Abstract Types** | `abstract` class/methods, `interface` | `abstract` class/methods, `interface` | Interface definition with method signatures (no bodies usually, unless Go 1.18+) |
| **Final / Sealed** | `final` keyword (prevents overriding) | `sealed` keyword (prevents inheritance) | No keyword. Preventing override is not applicable (no inheritance). |
| **Operator Overloading** | Not supported | Supported (e.g., `+`, `==`) | Not supported |
| **Extension Methods** | Not supported (Default methods in interfaces) | Supported (`this` parameter) | Not supported (but embedding provides similar behavior) |
| **Inner / Nested Classes** | Static and Non-static inner classes | Nested classes | No nested classes (only types in same package) |
| **Object Initialization** | `new Type(args)` | `new Type(args)` or `new Type { Prop = val }` (Object Initializer) | `Type{Field: val}` (Composite Literal) |
| **Concurrency Primitives** | `synchronized`, `volatile`, `java.util.concurrent` | `lock`, `Monitor`, `async/await`, `Task` | `go` keyword (Goroutines), `channels`, `sync.Mutex` |
