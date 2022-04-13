package main

func main() {
	m := make(map[int]*int)

	for i := 0; i < 6; i++ {
		// m[i] = &i // try to only change this part so that the program prints "012"

		item := new(int)
		// item = &i  // fill "item" in value by "memory address" of "i"
		*item = i // fill "memory address" of "item" by "i" int
		m[i] = item
	}

	for _, v := range m {
		print(*v)
	}
}
