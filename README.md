Here’s a README file for your `mathskills` package and main program. It includes an overview, installation instructions, usage, and a description of the functionalities.
made by {{@fahdaguenouz}}

```markdown
# Linear Stats

## Overview

The **Linear Stats** program calculates the Linear Regression Line and the Pearson Correlation Coefficient for a set of numerical data read from a file. The data is represented as \(y\) values, while the \(x\) values are inferred from the line numbers (0, 1, 2, ...).

## Features

- **Average**: Calculates the mean of a slice of float64 numbers.
- **Variance**: Computes the variance of the data.
- **Standard Deviation**: Derives the standard deviation from the variance.
- **Linear Regression Line**: Determines the slope and intercept for a linear fit of the data.
- **Pearson Correlation Coefficient**: Measures the strength and direction of association between two variables.


## Usage

1. Prepare your data file with each number on a new line. For example:

   ```
   189
   113
   121
   114
   145
   110
   ```

2. Run the program using the following command:

   ```bash
   go run main.go <data.txt>
   ```


## Output Format

After running the program, it will display:

- The Linear Regression Line in the format:
  
  ```
  Linear Regression Line: y = <value>x + <value>
  ```

- The Pearson Correlation Coefficient:

  ```
  Pearson Correlation Coefficient: <value>
  ```

The Linear Regression coefficients will have 6 decimal places, while the Pearson Correlation Coefficient will have 10 decimal places.

## Example

Given the data in `data.txt`:

```
189
113
121
114
145
110
```

The output might look like:

```
Linear Regression Line: y = 0.250000x + 114.666667
Pearson Correlation Coefficient: 0.8500000000
```

## Testing

This program can be tested with a provided random data set generator that compares the output of your program with expected results.

## Learning Objectives

- Understand statistical calculations involving averages, variance, standard deviation, and correlation.
- Gain experience with linear regression analysis.
- Implement data reading and processing in Go.

```

### Notes:

<https://github.com/fahdaguenouz/linear-stats.git>