package main

func swap2(n, m *int) {

}
func main() {
	n, m := 10, 20
	swap2(&n, &m)
	println(n, m)
}
