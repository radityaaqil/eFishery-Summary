package lib

//Public function uses uppercase on its first letter
func SayHello(name string) string {
	return "Hello public, " + name
}

//Private function uses uppercase on its first letter
func sayHi(name string) string {
	return "Hello private, " + name
}