package main

func swap(n, m int) (int, int) {
	return m, n
}
func main() {
	n, m := swap(10, 20)
	println(n, m)
}
