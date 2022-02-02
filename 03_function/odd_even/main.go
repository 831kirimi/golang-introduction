package main

func IsEven(n int) bool {
	if n%2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	for i := 1; i <= 100; i++ {
		print(i)
		// if i%2 == 0 {
		if IsEven(i) {
			println("-偶数")
		} else {
			println("-奇数")
		}
	}
}
