package day08

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type node struct {
	id    string
	left  string
	right string
}

type network map[string]node

type direction string

const (
	left  direction = "L"
	right direction = "R"
)

func getDirectionFromString(s string) direction {
	if s == "L" {
		return left
	}
	if s == "R" {
		return right
	}
	return ""
}

type instructions []direction

// Part1 solves the first part of the exercise
func Part1(input []string) string {
	nav := parseInstructions(input[0])
	net := buildNetwork(input[2:])

	steps := traverseNetwork(nav, net)

	return strconv.Itoa(steps)
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	nav := parseInstructions(input[0])
	net := buildNetwork(input[2:])

	startingNodes := getStartingNodes(net)
	cycleLengths := determineCycleLengths(startingNodes, net, nav)

	steps := calculateSteps(cycleLengths)

	return strconv.Itoa(steps)
}

func parseInstructions(line string) instructions {
	chars := strings.Split(line, "")

	inst := instructions{}
	for _, char := range chars {
		inst = append(inst, getDirectionFromString(char))
	}
	return inst
}

func buildNetwork(lines []string) network {
	net := make(network)

	for _, line := range lines {
		lineParts := strings.Split(line, " = ")

		id := lineParts[0]
		leftID := strings.Split(strings.Trim(lineParts[1], "()"), ", ")[0]
		rightID := strings.Split(strings.Trim(lineParts[1], "()"), ", ")[1]

		node := node{id: id, left: leftID, right: rightID}
		net[id] = node
	}
	return net
}

func traverseNetwork(nav instructions, net network) int {
	steps := 0

	currentNode := net["AAA"]
	navIndex := 0

	for currentNode.id != "ZZZ" {
		instruction := nav[navIndex]
		navIndex++
		if navIndex == len(nav) {
			navIndex = 0
		}

		if instruction == left {
			currentNode = net[currentNode.left]
		} else {
			currentNode = net[currentNode.right]
		}
		steps++
	}
	return steps
}

// func traverseNetworkGhost(nav instructions, net network) int {
// 	steps := 0

// 	currentNodes := getStartingNodes(net)
// 	navIndex := 0

// 	for !allNodesEndOnZ(currentNodes) {
// 		instruction := nav[navIndex]
// 		navIndex++
// 		if navIndex == len(nav) {
// 			navIndex = 0
// 		}

// 		if instruction == left {
// 			for i, node := range currentNodes {
// 				currentNodes[i] = net[node.left]
// 			}
// 		} else {
// 			for i, node := range currentNodes {
// 				currentNodes[i] = net[node.right]
// 			}
// 		}
// 		steps++
// 	}
// 	return steps
// }

func getStartingNodes(net network) []node {
	startingNodes := []node{}

	for id, node := range net {
		if strings.HasSuffix(id, "A") {
			startingNodes = append(startingNodes, node)
		}
	}
	return startingNodes
}

func determineCycleLengths(startingNodes []node, net network, nav instructions) []int {
	cycleLengths := []int{}

	for _, node := range startingNodes {

		cycleLength := 0
		navIndex := 0
		currentNode := node

		for !endsOnZ(currentNode) {
			instruction := nav[navIndex]
			navIndex++
			if navIndex == len(nav) {
				navIndex = 0
			}

			if instruction == left {
				currentNode = net[currentNode.left]
			} else {
				currentNode = net[currentNode.right]
			}

			cycleLength++
		}

		cycleLengths = append(cycleLengths, cycleLength)
	}
	return cycleLengths
}

func endsOnZ(node node) bool {
	return strings.HasSuffix(node.id, "Z")
}

func calculateSteps(cycleLengths []int) int {
	maxFactors := make(map[int]int)

	for _, length := range cycleLengths {
		primeFactors := calculatePrimeFactors(length)

		for factor, count := range primeFactors {
			if count > maxFactors[factor] {
				maxFactors[factor] = count
			}
		}
	}

	steps := 1
	for factor, count := range maxFactors {
		for i := 0; i < count; i++ {
			steps *= factor
		}
	}

	return steps
}

// Function to calculate prime factors of a number
func calculatePrimeFactors(n int) map[int]int {
	factors := make(map[int]int)

	for i := 2; i*i <= n; i++ {
		for n%i == 0 {
			factors[i]++
			n /= i
		}
	}

	if n > 1 {
		factors[n]++
	}

	return factors
}

// func allNodesEndOnZ(nodes []node) bool {
// 	for _, node := range nodes {
// 		if !strings.HasSuffix(node.id, "Z") {
// 			return false
// 		}
// 	}
// 	return true
// }

// Run function of the daily challenge
func Run(input []string, mode int) {
	var time1 time.Duration
	var time2 time.Duration

	if mode == 1 || mode == 3 {
		start1 := time.Now()
		res1 := Part1(input)
		time1 = time.Since(start1)

		fmt.Printf("Part one - Result: %v, Time: %v\n", res1, time1)
	}
	if mode == 2 || mode == 3 {
		start2 := time.Now()
		res2 := Part2(input)
		time2 = time.Since(start2)

		fmt.Printf("Part two - Result: %v, Time: %v\n", res2, time2)
	}
	if mode == 3 {
		fmt.Printf("Time Sum: %v\n", (time1 + time2))
	}
}
