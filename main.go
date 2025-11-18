/*
 * Author: Rabira Motuma
 * Version: 1.0.0
 * Date 2025-11-14
 * This program displays a multiplication table from 0 to 12, with the number 9.
 */

package main

import "fmt"

func main() {
	// INPUTS - float32 data type numbers
	var numbers []float64 = []float64{56.9, 89.7, 90.2}

  // PROCESS
  // calculate answers
  // https://www.geeksforgeeks.org/javascript/how-to-compute-the-sum-and-average-of-elements-in-an-array-in-javascript/
  var sum float64 = numbers[0] + numbers[1] + numbers[2]
  var average float64 = sum / float64(len(numbers))

	// OUTPUT
	// display average in conclusion statement
	fmt.Println("The average of 56.9, 89.7, and 90.2 is", average, ".")

	fmt.Println("\nDone.")
}
