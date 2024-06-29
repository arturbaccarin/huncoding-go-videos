package main

type str struct {
	word string
}

func main() {
	values := [6]string{"one", "two", "3", "4", "five", "six"}
	DefineMapFirstWay(values)

	DefineMapSecWay(values)

	DefineMapThirdWay(values)
}

func DefineMapFirstWay(value [6]string) {
	words := make(map[int]interface{})

	for i := 0; i < len(value); i++ {
		words[i] = value[i]
	}
}

// simulating a real struct
func DefineMapSecWay(value [6]string) {
	var words = map[int]str{}

	for i := 0; i < len(value); i++ {
		words[i] = str{value[i]}
	}
}

func DefineMapThirdWay(value [6]string) {
	words := make(map[int]string)

	for i := 0; i < len(value); i++ {
		words[i] = value[i]
	}
}
