package quad
package main

import "os"

func QuadA(x, y int) {
	// If x or y are not positive, print nothing
	if x <= 0 || y <= 0 {
		return
	}

	// Handle special case where width is 1
	if x == 1 {
		for i := 0; i < y; i++ {
			os.Stdout.WriteString("o\n")
		}
		return
	}

	// Handle special case where height is 1
	if y == 1 {
		os.Stdout.WriteString("o")
		for i := 1; i < x-1; i++ {
			os.Stdout.WriteString("-")
		}
		os.Stdout.WriteString("o\n")
		return
	}

	// Draw the rectangle
	for row := 0; row < y; row++ {
		for col := 0; col < x; col++ {
			if row == 0 || row == y-1 {
				// Top and bottom rows
				if col == 0 || col == x-1 {
					os.Stdout.WriteString("o") // Corners
				} else {
					os.Stdout.WriteString("-") // Horizontal edges
				}
			} else {
				// Middle rows
				if col == 0 || col == x-1 {
					os.Stdout.WriteString("|") // Vertical edges
				} else {
					os.Stdout.WriteString(" ") // Interior space
				}
			}
		}
		os.Stdout.WriteString("\n") // New line after each row
	}
}

// Example usage and test cases
func main() {
	os.Stdout.WriteString("QuadA(5, 3):\n")
	QuadA(5, 3)

	os.Stdout.WriteString("\nQuadA(5, 1):\n")
	QuadA(5, 1)

	os.Stdout.WriteString("\nQuadA(1, 1):\n")
	QuadA(1, 1)

	os.Stdout.WriteString("\nQuadA(1, 5):\n")
	QuadA(1, 5)

	os.Stdout.WriteString("\nQuadA(4, 4):\n")
	QuadA(4, 4)

	os.Stdout.WriteString("\nQuadA(0, 5): (should print nothing)\n")
	QuadA(0, 5)

	os.Stdout.WriteString("\nQuadA(-1, 3): (should print nothing)\n")
	QuadA(-1, 3)
}
