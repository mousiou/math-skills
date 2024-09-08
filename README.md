# Math Skills Project

## Objectives

The purpose of this project is to calculate and display the following statistical measures based on data read from a file:

- Average
- Median
- Variance
- Standard Deviation

## Concepts

### Average

The average (or mean) of a set of numbers is found by adding all the numbers together and then dividing by the total number of values. 

**How to Calculate:**
1. Sum all the numbers.
2. Divide the sum by the count of numbers.

For example, if you have numbers `2, 4, 6`, the average is `(2 + 4 + 6) / 3 = 12 / 3 = 4`.

### Median

The median is the middle number in a list of numbers arranged in ascending order. If the list has an even number of values, the median is the average of the two middle numbers.

**How to Calculate:**
1. Sort the numbers in ascending order.
2. If there is an odd number of values, the median is the middle number.
3. If there is an even number of values, the median is the average of the two middle numbers.

For example, in the sorted list `1, 3, 5`, the median is `3`. For the list `1, 2, 3, 4`, the median is `(2 + 3) / 2 = 2.5`.

### Variance

Variance measures how much the numbers in a data set differ from the average. It is calculated by averaging the squared differences between each number and the mean.

**How to Calculate:**
1. Find the average of the numbers.
2. Subtract the average from each number and square the result.
3. Find the average of these squared differences.

For example, for the numbers `1, 2, 3`, first find the average which is `2`. Then calculate the squared differences from the average: `(1-2)^2 = 1`, `(2-2)^2 = 0`, and `(3-2)^2 = 1`. The variance is `(1 + 0 + 1) / 3 = 2 / 3 = 0.67`.

### Standard Deviation

Standard deviation is the square root of the variance and gives an indication of the spread of the numbers around the mean.

**How to Calculate:**
1. Find the variance.
2. Take the square root of the variance.

For example, if the variance is `0.67`, then the standard deviation is `√0.67 ≈ 0.82`.

## Instructions

### Running the Program

1. Ensure that you have a Go environment set up.
2. Save the code provided in `main.go`.
3. Prepare a data file (`data.txt`) with each value on a new line.
4. Run the program with the following command:

   ```bash
   go run main.go data.txt
   ```

### Example Output

The program will output the calculated statistics in the following format:

```
Average: 35
Median: 4
Variance: 5
Standard Deviation: 65
```

Note that the values are rounded integers.
