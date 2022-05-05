/*
Your task is to split the chocolate bar of given dimension n x m into small squares. Each square is of size 1x1 and unbreakable. Implement a function that will return minimum number of breaks needed.

For example if you are given a chocolate bar of size 2 x 1 you can split it to single squares in just one break, but for size 3 x 1 you must do two breaks.

If input data is invalid you should return 0 (as in no breaks are needed if we do not have any chocolate to split). Input will always be a non-negative integer.

Problem
- Given two non-negative integers, n and m, which represent the dimensions of a bar of chocolate, return an integer representing the minimum number of breaks to reduce the bar of chocolate to 1x1 pieces.

- If either m or n is 0, the function should return 0.

Example
BreakChocolate(2, 1) // 1
BreakChocolate(3, 1) // 2
BreakChocolate(3, 2) // 5

Let's say n is width and m is height.

2x1
-- --
-- -- <-- must be broken on the vertical axis

3x1
-- -- --
-- -- -- <-- must be broken twice, along vertical axes

3x2
-- -- --
-- -- --
-- -- -- <-- should first be broken along horizontal axis

-- -- --
-- -- -- <-- now each piece is 3x1 and needs two breaks

-- -- --
-- -- -- <-- needs two more breaks

Once the piece of chocolate is reduced to either n x 1 or 1 x m, the breaks needed for that piece of chocolate are n - 1 or m - 1, respectively.

Until that point, the breaks should be made in such a way that the lesser number is decremented by 1

Let's try 5x5

-- -- -- -- --
-- -- -- -- --
-- -- -- -- --
-- -- -- -- --
-- -- -- -- --
-- -- -- -- --

In this case, n and m are the same, so we should cut either dimension in half. We'll choose m.

We get two pieces, one 5x3 piece and one 5x2 piece. We've consumed one break.
-- -- -- -- --
-- -- -- -- --
-- -- -- -- --
-- -- -- -- --

-- -- -- -- --
-- -- -- -- --
-- -- -- -- --

Let's work on the 5x3 piece first
-- -- -- -- --
-- -- -- -- --
-- -- -- -- --
-- -- -- -- --

To decrement 3, we then get a 5x2 and a 5x1 piece. This consumes another break.
-- -- -- -- --
-- -- -- -- --
-- -- -- -- --

-- -- -- -- --
-- -- -- -- -- <-- this is equivalent to 4 breaks

Let's now work on the 5x2... it should be broken into two 5x1 pieces, which are worth 4 breaks each, plus the break consumed to produce the two pieces.

So the 5x3 piece is woth 15 breaks.

Back to our 5x2 piece.
We know that to break that piece into two, we consume one break, plus the four known breaks per piece, for another 9 breaks.

15 + 9 = 24

Data Structure
What comes to mind first is an array of arrays, where each array contains two integers representing m and n.
This could also definitely be done with recursion

Algorithm (recursive)

If either m or n equals 1, return the other dimension minus 1
Otherwise, identify the smaller of m and n
	- Divide that smaller dimension in two (integer division)
	- Return the sum:
		- 1 +
		- Invocation of the function with (larger, result of division of smaller) +
		- Invocation of (larger, smaller - result of division of smaller)

e.g. 3x2
neither dimension is 1
we identify 2 as the smaller of m and n
	- 2/2 is 1
	- we'll return 1 + fn(3, 1) + fn(3, (2-1))

this will recursively call 3x1
since one of the dimensions is 1, this call returns (3-1) = 2

we call 3x1 again, which again returns 2

Our final answer is 1 + 2 + 2 = 5
*/

package main

import "fmt"

func idString(n, m int) string {
	return fmt.Sprintf("%dx%d", n, m)
}

func BreakChocolate(n, m int, memo map[string]int) int {
	if totalBreaks, ok := memo[idString(n, m)]; ok {
		return totalBreaks
	}

	switch {
	case n == 1:
		memo[idString(n, m)] = m - 1
		return m - 1
	case m == 1:
		memo[idString(n, m)] = n - 1
		return n - 1
	}

	smaller, larger := n, m
	if n > m {
		larger, smaller = m, n
	}

	half := smaller / 2
	remainder := smaller - half

	memo[idString(n, m)] = 1 + BreakChocolate(larger, half, memo) + BreakChocolate(larger, remainder, memo)
	return memo[idString(n, m)]
}

func main() {
	fmt.Println(BreakChocolate(2, 1, make(map[string]int))) // 1
	fmt.Println(BreakChocolate(3, 1, make(map[string]int))) // 2
	fmt.Println(BreakChocolate(3, 2, make(map[string]int))) // 5
	fmt.Println(BreakChocolate(5, 5, make(map[string]int))) // 24
	fmt.Println(BreakChocolate(7, 5, make(map[string]int))) // 34
}
