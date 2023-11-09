package pointer

import "fmt"

// pointer
func double(n *int) {
	*n = *n * 2
}

func appendGreeting(s *string) {
	*s = fmt.Sprintf("Hi, %s", *s)
}

func main() {
	// pointer
	n := 2
	double(&n)
	fmt.Println(n)

	s := "Bob"
	appendGreeting(&s)
	fmt.Println(s)
}
