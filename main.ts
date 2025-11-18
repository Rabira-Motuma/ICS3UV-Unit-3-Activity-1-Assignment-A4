/**
* @author Rabira Motuma
* @version 1.0.0
* @date 2025-11-11
* @fileoverview This program finds and displays the sidelengths of a cube with a volume of 1000 mm³.
*/

// INPUTS - number data type numbers
let numbers : number = [56.9, 89.7, 90.2];

// PROCESS
// calculate answers
// https://www.geeksforgeeks.org/javascript/how-to-compute-the-sum-and-average-of-elements-in-an-array-in-javascript/
let sum : number = numbers[0] + numbers[1] + numbers[2];
let average : number = sum / numbers.length;

// OUTPUT
// display average in conclusion sentences
console.log("The average of 56.9, 89.7, and 90.2 is " + average)

console.log("\nDone.")
