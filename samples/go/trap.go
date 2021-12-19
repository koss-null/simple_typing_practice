package main

import "fmt"

// import "fmt"

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func maximum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Pair struct {
	val, idx int
}

func solveIter(height []Pair) (int, int, []Pair) {
	leftHeight := 0
	i := 0
	sum := 0
	subSum := 0
	height = append(height, Pair{0, height[len(height)-1].idx + 1})
	packed := make([]Pair, 0)

	// going up at the beginning
	for i+1 < len(height) && height[i].val < height[i+1].val {
		i++
	}

	packed = append(packed, height[i])
	for i+1 < len(height) {
		// up -> down
		leftHeight = height[i].val
		start := i
		sub := 0
		length, lastIdx := 0, height[i].idx
		for i+1 < len(height) && height[i].val >= height[i+1].val {
			sub += height[i+1].val
			length += height[i+1].idx - lastIdx
			lastIdx = height[i+1].idx
			i++
			fmt.Println("i", i, sub, leftHeight, height[i], sum)
		}
		// down -> up
		didStep := false
		for (i+1 < len(height) && height[i].val <= height[i+1].val) &&
			(height[i].val <= leftHeight) {
			sub += height[i+1].val
			length += height[i+1].idx - lastIdx
			lastIdx = height[i+1].idx
			i++
			didStep = true
			fmt.Println("ii", i, sub, leftHeight, height[i], sum)
		}
		sub -= height[i].val

		if i < len(height) {
			fmt.Println("iii", i, "sub", sub, "lh", leftHeight, "heighti", height[i], "sum", sum)
			fmt.Println("len", length)
			length = maximum(length-1, 0)
			fmt.Println("len", length, "min_h", min(height[start].val, height[i].val), "sub", sub)
			if didStep && height[i].val < height[start].val {
				// need to find the first greater element from the right
				newStart := i - 1
				for newStart >= 0 && height[newStart].val < height[i].val {
					newStart--
				}
				if newStart > -1 {
					length -= (height[newStart].idx - height[start].idx)
					for j := newStart; j > start; j-- {
						fmt.Println("increasing sum ", height[j].val)
						sum += height[j].val
					}
				}
				start = newStart
				fmt.Println("NEWlen", length, "sub", sub)
			}
			if didStep {
				packed = append(packed, height[i])
				subSum += sub
			}
			fmt.Println("adding ", maximum(length*min(height[start].val, height[i].val)-sub, 0), subSum)
			sum += length * min(height[start].val, height[i].val)
		}
	}
	fmt.Println("sum", sum, "ss", subSum)
	return sum, subSum, packed
}

func toPairs(a []int) []Pair {
	b := make([]Pair, 0, len(a))
	for i := range a {
		b = append(b, Pair{a[i], i})
	}
	return b
}

func trap(height []int) int {
	lastSum, lastSub := 0, 0

	pairedHeight := toPairs(height)
	for sum, sub, packed := solveIter(pairedHeight); lastSum < sum-sub-lastSub; sum, sub, packed = solveIter(pairedHeight) {
		fmt.Println("SUM", sum, "SUB", sub)
		lastSub += sub
		lastSum = sum - lastSub
		pairedHeight = packed
		fmt.Println(pairedHeight, lastSum)
		fmt.Println("----------")
	}

	return lastSum
}

func main() {
	fmt.Println(trap([]int{5, 5, 1, 7, 1, 1, 5, 2, 7, 6}))
}
