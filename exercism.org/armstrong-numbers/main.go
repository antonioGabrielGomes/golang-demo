package armstrongnumbers

import (
	"log"
	"math"
	"strconv"
)

func sum(total *int, plus int) {
	*total = *total + plus
}

func IsArmstrongNumber(number int) int {
	number_parse := strconv.Itoa(number)
	number_length := len(number_parse)
	total := 0

	for i := 0; i < number_length; i++ {

		digit, err := strconv.Atoi(string(number_parse[i]))

		if err != nil {
			log.Fatal(err.Error())
		}

		result := int(math.Pow(float64(digit), float64(number_length)))

		sum(&total, result)
	}

	return total
}

/*
func main() {

	fmt.Println("is armstrong numbers ?")
	fmt.Println("9   = 9   ?", IsArmstrongNumber(9) == 9)     //  True
	fmt.Println("10  = 10  ?", IsArmstrongNumber(10) == 10)   // False
	fmt.Println("153 = 153 ?", IsArmstrongNumber(153) == 153) // True
	fmt.Println("154 = 154 ?", IsArmstrongNumber(154) == 154) // False
}

*/
