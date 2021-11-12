package main

import (
	"fmt"
	"math"
)

func main() {
	var number float32 = 0.085

	fmt.Printf("Starting Number: %f\n", number)

	bits := math.Float32bits(number)

	binary := fmt.Sprintf("%.32b", bits)
	fmt.Printf("Bit Pattern: %s | %s %s | %s %s %s %s %s %s\n",
		binary[0:1],
		binary[1:5],
		binary[5:9],
		binary[9:12],
		binary[12:16],
		binary[16:20],
		binary[20:24],
		binary[24:28],
		binary[28:32],
	)

	fmt.Println(binary)

	bias := 127
	sign := bits & (1 << 31)
	exponentRaw := int(bits >> 23)
	exponent := exponentRaw - bias
	var mantissa float64
	for i, bit := range binary[9:32] {
		if bit == 49 {
			position := i + 1
			bitValue := math.Pow(2, float64(position))
			fractional := 1 / bitValue
			mantissa += fractional
		}
	}
	value := (1 + mantissa) * math.Pow(2, float64(exponent))

	fmt.Printf("Sign : %d, Exponent : %d (%d), Mantissa : %f, Value : %f\n",
		sign,
		exponentRaw,
		exponent,
		mantissa,
		value,
	)
}
