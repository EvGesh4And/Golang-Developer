package main

import (
	"fmt"
	"os"
)

func main() {

	env := os.Environ() // слайс строк "key=value"
	fmt.Println(env)    // USER=rob

	user, exists := os.LookupEnv("USER")
	if exists {
		fmt.Println(user) // rob
	}

	city, cityExists := os.LookupEnv("CITY")
	fmt.Println(city, cityExists)

	//os.Setenv("CITY", "Ottawa") // установить
	//os.Unsetenv("CITY")                                 // удалить
	fmt.Println(os.ExpandEnv("$USER lives in ${CITY}")) // "шаблонизация"
}

// CITY=Ekaterinburg
// export CITY
// unset CITY
