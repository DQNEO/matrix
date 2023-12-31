package main

import (
	"fmt"

	"github.com/DQNEO/matrix"
)

func main() {
	// Create a new 3x3 matrix
	m1 := matrix.NewMatrix(3, 3, []float64{
		0, 2, 1,
		2, 0, 0,
		0, 3, 2,
	})

	// Create another 3x3 matrix
	m2 := matrix.NewMatrix(3, 3, []float64{
		10, 11, 12,
		13, 14, 15,
		16, 17, 18,
	})

	// Get the size of m1 as string
	fmt.Println("\nSize of Matrix 1:")
	fmt.Println(m1.Type())

	// Print the matrices
	fmt.Println("Matrix 1:")
	fmt.Println(m1.String())
	fmt.Println("\nMatrix 2:")
	fmt.Println(m2.String())

	// Add m1 and m2
	fmt.Println("\nSum of Matrix 1 and 2:")
	sum := matrix.Add(m1, m2)
	fmt.Println(sum.String())

	// Multiply m1 and m2
	fmt.Println("\nProduct of Matrix 1 and 2:")
	product := matrix.Mul(m1, m2)
	fmt.Println(product.String())

	// Calculate the inverse of the matrix
	inv := matrix.Inv(m1)

	// Print the inverse matrix
	fmt.Println("\nInverse Matrix:")
	fmt.Println(inv.String())

	// Create a new Identity matrix with size 3x3
	identity := matrix.NewIdentityMatrix(3)
	fmt.Println("\nNew 3x3 Identity Matrix: ")
	fmt.Println(identity.String())

	// Check if the original matrix x the inverse matrix matches identity matrix
	isEqual := matrix.Eq(matrix.Mul(m1, inv), identity)
	if isEqual {
		fmt.Println("\n m1 * inv == identity")
	}
	// Transpose m1
	transposed := m1.Tr()
	fmt.Println("\nTransposed Matrix 1:")
	fmt.Println(transposed.String())

	// Get determinant of m1
	det := m1.Det()
	fmt.Println("\nDeterminant of Matrix 1:")
	fmt.Println(det)

	// Scale m1 by 2
	scaled := matrix.Scale(2, m1)
	fmt.Println("\nMatrix 1 scaled by 2:")
	fmt.Println(scaled.String())

	// Get a specific element from m1
	elm := m1.GetElm(2, 3)
	fmt.Println("\nElement at (2,3) in Matrix 1:")
	fmt.Println(elm)

	// Create a new zero matrix with size 3x2
	zero := matrix.NewZeroMatrix(3, 2)
	fmt.Println("\nNew 3x2 Zero Matrix: ")
	fmt.Println(zero.String())
}
