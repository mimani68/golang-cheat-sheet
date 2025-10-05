// string_cheatsheet.go
// Go String Data Type Cheat Sheet with Advanced Examples and Edge Cases

package main

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	fmt.Println("=== GO STRING CHEAT SHEET ===\n")

	// 1. STRING BASICS
	fmt.Println("1. STRING BASICS")
	fmt.Println("================")

	// String declaration and initialization
	var s1 string         // empty string ""
	s2 := "Hello, World!" // string literal
	s3 := `Raw string literal
with newlines` // raw string literal
	s4 := "Hello" + " " + "World" // concatenation

	fmt.Printf("s1: %q (len: %d)\n", s1, len(s1))
	fmt.Printf("s2: %q\n", s2)
	fmt.Printf("s3: %q\n", s3)
	fmt.Printf("s4: %q\n", s4)

	// String immutability
	// s2[0] = 'h' // ERROR: cannot assign to s2[0]

	// 2. STRING INDEXING AND SLICING
	fmt.Println("\n2. STRING INDEXING AND SLICING")
	fmt.Println("==============================")

	str := "Hello, 世界"
	fmt.Printf("String: %q\n", str)
	fmt.Printf("len(str): %d bytes\n", len(str))
	fmt.Printf("utf8.RuneCountInString(str): %d runes\n", utf8.RuneCountInString(str))

	// Byte indexing (not safe for UTF-8)
	fmt.Printf("str[0]: %c (byte: %d)\n", str[0], str[0])
	fmt.Printf("str[7]: %c (byte: %d)\n", str[7], str[7]) // First byte of 世

	// Proper rune iteration
	fmt.Print("Runes: ")
	for i, r := range str {
		fmt.Printf("[%d]=%c ", i, r)
	}
	fmt.Println()

	// Slicing (byte-based, be careful with UTF-8)
	fmt.Printf("str[0:5]: %q\n", str[0:5])
	fmt.Printf("str[7:]: %q\n", str[7:])

	// 3. STRING COMPARISON
	fmt.Println("\n3. STRING COMPARISON")
	fmt.Println("====================")

	a, b := "apple", "banana"
	fmt.Printf("%q == %q: %v\n", a, b, a == b)
	fmt.Printf("%q < %q: %v\n", a, b, a < b)
	fmt.Printf("%q > %q: %v\n", a, b, a > b)

	// Case-insensitive comparison
	s5, s6 := "GoLang", "golang"
	fmt.Printf("strings.EqualFold(%q, %q): %v\n", s5, s6, strings.EqualFold(s5, s6))

	// 4. STRINGS PACKAGE FUNCTIONS
	fmt.Println("\n4. STRINGS PACKAGE FUNCTIONS")
	fmt.Println("============================")

	text := "  Go is awesome!  "

	// Trimming
	fmt.Printf("strings.TrimSpace(%q): %q\n", text, strings.TrimSpace(text))
	fmt.Printf("strings.Trim(%q, \" !\"): %q\n", text, strings.Trim(text, " !"))
	fmt.Printf("strings.TrimPrefix(\"Hello, World\", \"Hello\"): %q\n",
		strings.TrimPrefix("Hello, World", "Hello"))
	fmt.Printf("strings.TrimSuffix(\"Hello.txt\", \".txt\"): %q\n",
		strings.TrimSuffix("Hello.txt", ".txt"))

	// Case conversion
	fmt.Printf("strings.ToUpper(\"hello\"): %q\n", strings.ToUpper("hello"))
	fmt.Printf("strings.ToLower(\"HELLO\"): %q\n", strings.ToLower("HELLO"))
	fmt.Printf("strings.Title(\"hello world\"): %q\n", strings.Title("hello world"))

	// Searching
	haystack := "The quick brown fox jumps over the lazy dog"
	fmt.Printf("strings.Contains(%q, \"fox\"): %v\n", haystack, strings.Contains(haystack, "fox"))
	fmt.Printf("strings.HasPrefix(%q, \"The\"): %v\n", haystack, strings.HasPrefix(haystack, "The"))
	fmt.Printf("strings.HasSuffix(%q, \"dog\"): %v\n", haystack, strings.HasSuffix(haystack, "dog"))
	fmt.Printf("strings.Index(%q, \"fox\"): %d\n", haystack, strings.Index(haystack, "fox"))
	fmt.Printf("strings.LastIndex(%q, \"o\"): %d\n", haystack, strings.LastIndex(haystack, "o"))
	fmt.Printf("strings.Count(%q, \"o\"): %d\n", haystack, strings.Count(haystack, "o"))

	// Splitting and Joining
	parts := strings.Split("a,b,c,d", ",")
	fmt.Printf("strings.Split(\"a,b,c,d\", \",\"): %v\n", parts)
	fmt.Printf("strings.Join(parts, \"-\"): %q\n", strings.Join(parts, "-"))

	fields := strings.Fields("  multiple   spaces   between   words  ")
	fmt.Printf("strings.Fields(\"  multiple   spaces   between   words  \"): %v\n", fields)

	// Replacing
	fmt.Printf("strings.Replace(\"foo bar foo\", \"foo\", \"baz\", 1): %q\n",
		strings.Replace("foo bar foo", "foo", "baz", 1))
	fmt.Printf("strings.ReplaceAll(\"foo bar foo\", \"foo\", \"baz\"): %q\n",
		strings.ReplaceAll("foo bar foo", "foo", "baz"))

	// 5. STRING BUILDER (EFFICIENT CONCATENATION)
	fmt.Println("\n5. STRING BUILDER")
	fmt.Println("=================")

	var builder strings.Builder
	for i := 0; i < 5; i++ {
		builder.WriteString(fmt.Sprintf("Line %d\n", i))
	}
	result := builder.String()
	fmt.Printf("Builder result:\n%s", result)

	// 6. BYTE OPERATIONS
	fmt.Println("\n6. BYTE OPERATIONS")
	fmt.Println("==================")

	// String to bytes and back
	original := "Hello, 世界"
	byteSlice := []byte(original)
	fmt.Printf("[]byte(%q): %v\n", original, byteSlice)
	fmt.Printf("string(byteSlice): %q\n", string(byteSlice))

	// bytes package operations
	fmt.Printf("bytes.Equal([]byte(\"abc\"), []byte(\"abc\")): %v\n",
		bytes.Equal([]byte("abc"), []byte("abc")))
	fmt.Printf("bytes.Contains([]byte(\"seafood\"), []byte(\"foo\")): %v\n",
		bytes.Contains([]byte("seafood"), []byte("foo")))

	// 7. UNICODE OPERATIONS
	fmt.Println("\n7. UNICODE OPERATIONS")
	fmt.Println("=====================")

	unicodeStr := "Hello, 世界! 123"

	// Rune operations
	runes := []rune(unicodeStr)
	fmt.Printf("[]rune(%q): %v\n", unicodeStr, runes)
	fmt.Printf("string(runes[7:9]): %q\n", string(runes[7:9]))

	// Unicode properties
	for _, r := range "Hello123!世界" {
		fmt.Printf("'%c': Letter=%v, Digit=%v, Space=%v, Punct=%v\n",
			r, unicode.IsLetter(r), unicode.IsDigit(r),
			unicode.IsSpace(r), unicode.IsPunct(r))
	}

	// 8. STRING CONVERSION
	fmt.Println("\n8. STRING CONVERSION")
	fmt.Println("====================")

	// String to numbers
	intStr := "42"
	intVal, err := strconv.Atoi(intStr)
	// intVal, err := strconv.ParseInt(a, 10, 64)  // Alternative
	fmt.Printf("strconv.Atoi(%q): %d, error: %v\n", intStr, intVal, err)

	floatStr := "3.14159"
	floatVal, err := strconv.ParseFloat(floatStr, 64)
	fmt.Printf("strconv.ParseFloat(%q, 64): %f, error: %v\n", floatStr, floatVal, err)

	boolStr := "true"
	boolVal, err := strconv.ParseBool(boolStr)
	fmt.Printf("strconv.ParseBool(%q): %v, error: %v\n", boolStr, boolVal, err)

	// Numbers to string
	fmt.Printf("strconv.Itoa(42): %q\n", strconv.Itoa(42))
	fmt.Printf("strconv.FormatFloat(3.14159, 'f', 2, 64): %q\n",
		strconv.FormatFloat(3.14159, 'f', 2, 64))
	fmt.Printf("strconv.FormatBool(true): %q\n", strconv.FormatBool(true))

	// fmt.Sprintf for complex formatting
	formatted := fmt.Sprintf("Name: %s, Age: %d, Score: %.2f", "John", 30, 95.5)
	fmt.Printf("fmt.Sprintf result: %q\n", formatted)

	// 9. REGULAR EXPRESSIONS
	fmt.Println("\n9. REGULAR EXPRESSIONS")
	fmt.Println("======================")

	pattern := `\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Z|a-z]{2,}\b`
	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Printf("Regex compile error: %v\n", err)
	}

	testStr := "Contact us at info@example.com or support@test.org"
	matches := re.FindAllString(testStr, -1)
	fmt.Printf("Email matches in %q: %v\n", testStr, matches)

	// Replace with regex
	result = re.ReplaceAllString(testStr, "[REDACTED]")
	fmt.Printf("After redacting emails: %q\n", result)

	// 10. EDGE CASES AND GOTCHAS
	fmt.Println("\n10. EDGE CASES AND GOTCHAS")
	fmt.Println("==========================")

	// Empty string checks
	empty := ""
	fmt.Printf("Empty string: %q, len: %d, == \"\": %v\n", empty, len(empty), empty == "")

	// Nil vs empty string
	var nilStr string
	fmt.Printf("Uninitialized string: %q, == \"\": %v, == nilStr: %v\n",
		nilStr, nilStr == "", nilStr == empty)

	// UTF-8 edge cases
	emoji := "Hello 👋 World 🌍"
	fmt.Printf("Emoji string: %q\n", emoji)
	fmt.Printf("len(emoji): %d bytes\n", len(emoji))
	fmt.Printf("utf8.RuneCountInString(emoji): %d runes\n", utf8.RuneCountInString(emoji))

	// Invalid UTF-8
	invalidUTF8 := string([]byte{0xff, 0xfe, 0xfd})
	fmt.Printf("Invalid UTF-8: %q\n", invalidUTF8)
	fmt.Printf("utf8.ValidString(invalidUTF8): %v\n", utf8.ValidString(invalidUTF8))

	// String interning behavior
	str1 := "hello"
	str2 := "hello"
	str3 := "hel" + "lo"
	fmt.Printf("str1 == str2: %v (both literals)\n", str1 == str2)
	fmt.Printf("str1 == str3: %v (concatenation)\n", str1 == str3)

	// Substring sharing (strings are immutable but share memory)
	bigString := "This is a very long string with lots of content"
	substring := bigString[10:20]
	fmt.Printf("Substring %q shares memory with original\n", substring)

	// 11. PERFORMANCE TIPS
	fmt.Println("\n11. PERFORMANCE TIPS")
	fmt.Println("====================")

	// Use strings.Builder for multiple concatenations
	var inefficient string
	var efficient strings.Builder

	// Inefficient way (creates new strings each time)
	for i := 0; i < 5; i++ {
		inefficient += fmt.Sprintf("Item %d, ", i)
	}

	// Efficient way
	for i := 0; i < 5; i++ {
		efficient.WriteString(fmt.Sprintf("Item %d, ", i))
	}

	fmt.Printf("Inefficient result: %q\n", inefficient)
	fmt.Printf("Efficient result: %q\n", efficient.String())

	// Use byte slices for mutable string operations
	mutable := []byte("Hello")
	mutable[0] = 'h'
	fmt.Printf("Mutable string operations: %q\n", string(mutable))

	// 12. ADVANCED STRING OPERATIONS
	fmt.Println("\n12. ADVANCED STRING OPERATIONS")
	fmt.Println("==============================")

	// Custom split function
	customSplit := func(r rune) bool {
		return r == ',' || r == ';' || r == '|'
	}
	result2 := strings.FieldsFunc("a,b;c|d", customSplit)
	fmt.Printf("Custom split result: %v\n", result2)

	// String reader for IO operations
	reader := strings.NewReader("Hello, Reader!")
	buffer := make([]byte, 5)
	n, _ := reader.Read(buffer)
	fmt.Printf("Read %d bytes: %q\n", n, string(buffer[:n]))

	// Map function to transform runes
	rot13 := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return 'A' + (r-'A'+13)%26
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+13)%26
		}
		return r
	}
	encoded := strings.Map(rot13, "Hello, World!")
	fmt.Printf("ROT13 encoded: %q\n", encoded)
}

// === GO STRING CHEAT SHEET ===

// 1. STRING BASICS
// ================
// s1: "" (len: 0)
// s2: "Hello, World!"
// s3: "Raw string literal\nwith newlines"
// s4: "Hello World"

// 2. STRING INDEXING AND SLICING
// ==============================
// String: "Hello, 世界"
// len(str): 13 bytes
// utf8.RuneCountInString(str): 9 runes
// str[0]: H (byte: 72)
// str[7]: ä (byte: 228)
// Runes: [0]=H [1]=e [2]=l [3]=l [4]=o [5]=, [6]=  [7]=世 [10]=界
// str[0:5]: "Hello"
// str[7:]: "世界"

// 3. STRING COMPARISON
// ====================
// "apple" == "banana": false
// "apple" < "banana": true
// "apple" > "banana": false
// strings.EqualFold("GoLang", "golang"): true

// 4. STRINGS PACKAGE FUNCTIONS
// ============================
// strings.TrimSpace("  Go is awesome!  "): "Go is awesome!"
// strings.Trim("  Go is awesome!  ", " !"): "Go is awesome"
// strings.TrimPrefix("Hello, World", "Hello"): ", World"
// strings.TrimSuffix("Hello.txt", ".txt"): "Hello"
// strings.ToUpper("hello"): "HELLO"
// strings.ToLower("HELLO"): "hello"
// strings.Title("hello world"): "Hello World"
// strings.Contains("The quick brown fox jumps over the lazy dog", "fox"): true
// strings.HasPrefix("The quick brown fox jumps over the lazy dog", "The"): true
// strings.HasSuffix("The quick brown fox jumps over the lazy dog", "dog"): true
// strings.Index("The quick brown fox jumps over the lazy dog", "fox"): 16
// strings.LastIndex("The quick brown fox jumps over the lazy dog", "o"): 41
// strings.Count("The quick brown fox jumps over the lazy dog", "o"): 4
// strings.Split("a,b,c,d", ","): [a b c d]
// strings.Join(parts, "-"): "a-b-c-d"
// strings.Fields("  multiple   spaces   between   words  "): [multiple spaces between words]
// strings.Replace("foo bar foo", "foo", "baz", 1): "baz bar foo"
// strings.ReplaceAll("foo bar foo", "foo", "baz"): "baz bar baz"

// 5. STRING BUILDER
// =================
// Builder result:
// Line 0
// Line 1
// Line 2
// Line 3
// Line 4

// 6. BYTE OPERATIONS
// ==================
// []byte("Hello, 世界"): [72 101 108 108 111 44 32 228 184 150 231 149 140]
// string(byteSlice): "Hello, 世界"
// bytes.Equal([]byte("abc"), []byte("abc")): true
// bytes.Contains([]byte("seafood"), []byte("foo")): true

// 7. UNICODE OPERATIONS
// =====================
// []rune("Hello, 世界! 123"): [72 101 108 108 111 44 32 19990 30028 33 32 49 50 51]
// string(runes[7:9]): "世界"
// 'H': Letter=true, Digit=false, Space=false, Punct=false
// 'e': Letter=true, Digit=false, Space=false, Punct=false
// 'l': Letter=true, Digit=false, Space=false, Punct=false
// 'l': Letter=true, Digit=false, Space=false, Punct=false
// 'o': Letter=true, Digit=false, Space=false, Punct=false
// '1': Letter=false, Digit=true, Space=false, Punct=false
// '2': Letter=false, Digit=true, Space=false, Punct=false
// '3': Letter=false, Digit=true, Space=false, Punct=false
// '!': Letter=false, Digit=false, Space=false, Punct=true
// '世': Letter=true, Digit=false, Space=false, Punct=false
// '界': Letter=true, Digit=false, Space=false, Punct=false

// 8. STRING CONVERSION
// ====================
// strconv.Atoi("42"): 42, error: <nil>
// strconv.ParseFloat("3.14159", 64): 3.141590, error: <nil>
// strconv.ParseBool("true"): true, error: <nil>
// strconv.Itoa(42): "42"
// strconv.FormatFloat(3.14159, 'f', 2, 64): "3.14"
// strconv.FormatBool(true): "true"
// fmt.Sprintf result: "Name: John, Age: 30, Score: 95.50"

// 9. REGULAR EXPRESSIONS
// ======================
// Email matches in "Contact us at info@example.com or support@test.org": [info@example.com support@test.org]
// After redacting emails: "Contact us at [REDACTED] or [REDACTED]"

// 10. EDGE CASES AND GOTCHAS
// ==========================
// Empty string: "", len: 0, == "": true
// Uninitialized string: "", == "": true, == nilStr: true
// Emoji string: "Hello 👋 World 🌍"
// len(emoji): 21 bytes
// utf8.RuneCountInString(emoji): 15 runes
// Invalid UTF-8: "\xff\xfe\xfd"
// utf8.ValidString(invalidUTF8): false
// str1 == str2: true (both literals)
// str1 == str3: true (concatenation)
// Substring "very long " shares memory with original

// 11. PERFORMANCE TIPS
// ====================
// Inefficient result: "Item 0, Item 1, Item 2, Item 3, Item 4, "
// Efficient result: "Item 0, Item 1, Item 2, Item 3, Item 4, "
// Mutable string operations: "hello"

// 12. ADVANCED STRING OPERATIONS
// ==============================
// Custom split result: [a b c d]
// Read 5 bytes: "Hello"
// ROT13 encoded: "Uryyb, Jbeyq!"
