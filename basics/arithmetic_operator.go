package basics

import (
	"fmt"
)

func main() {
	var a, b int = 10, 3
	var result int

	result = a + b // addition
	fmt.Println("Addition:", result)

	result = a - b // subtraction
	fmt.Println("Subtraction:", result)

	result = a * b // multiplication
	fmt.Println("Multiplication:", result)

	result = a / b // division
	fmt.Println("Division:", result)

	result = a % b // modulus
	fmt.Println("Modulus:", result)

	result = a & b // bitwise AND
	fmt.Println("Bitwise AND:", result)

	result = a | b // bitwise OR
	fmt.Println("Bitwise OR:", result)

	result = a ^ b // bitwise XOR
	fmt.Println("Bitwise XOR:", result)

	result = a << 1 // left shift
	fmt.Println("Left Shift:", result)

	result = a >> 1 // right shift
	fmt.Println("Right Shift:", result)
     //overflow of signedInt
     var maxInt int64 = 9223372036854775807 // Maximum value for int64
      fmt.Println("Maximum int64 value:", maxInt)

	 maxInt = maxInt + 1 // This will cause an overflow
	fmt.Println("After overflow, maxInt:", maxInt)
	maxInt = maxInt - 1 // This will cause an underflow
	fmt.Println("After underflow, maxInt:", maxInt)

	//overflow of unsignedInt
	var maxUint uint64 = 18446744073709551615 // Maximum value for uint64  // can't go below 0
	fmt.Println("Maximum uint64 value:", maxUint)
	maxUint = maxUint + 1 // This will cause an overflow
	fmt.Println("After overflow, maxUint:", maxUint)
	maxUint = maxUint - 1 // This will cause an underflow
	fmt.Println("After underflow, maxUint:", maxUint)



	fmt.Println("Arithmetic operations completed.")
}
