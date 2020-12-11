package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type adapter struct {
	jo    int
	dnext int
	priors   int
}

func newAdapter(jo int) adapter {
	a := adapter{jo, 0, 0}
	return a
}

func sortedAdapters(path string) []adapter {
	file, _ := os.Open(path)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var joltOuts []int
	for scanner.Scan() {
		jo, _ := strconv.Atoi(scanner.Text())
		joltOuts = append(joltOuts, jo)
	}
	sort.Ints(joltOuts)
	var adapters []adapter
	outlet := newAdapter(0)
	adapters = append(adapters, outlet)
	for _, v := range joltOuts {
		adapters = append(adapters, newAdapter(v))
	}
	device := newAdapter(adapters[len(adapters)-1].jo + 3)
	adapters = append(adapters, device)
	return adapters
}

func adapterDistances(adapters []adapter) []adapter {
	for i := 1; i < len(adapters); i++ {
		adapters[i-1].dnext = adapters[i].jo - adapters[i-1].jo
	}
	return adapters
}

func adapterDistSums(adapters []adapter) map[int]int {
	distSums := make(map[int]int)
	for _, a := range adapters {
		switch a.dnext {
		case 1:
			distSums[1] += 1
		case 2:
			distSums[2] += 1
		case 3:
			distSums[3] += 1
		}
	}
	return distSums
}

func adapterChains(adapters []adapter) []adapter {
	for i := 1; i < len(adapters); i++ {
		sliceLow := 0
		if i - 3 < 0 {
			sliceLow = 0
		} else {
			sliceLow = i - 3
		}
		for _, posPrior := range adapters[sliceLow:i] {
			if adapters[i].jo - posPrior.jo <= 3 {
				if posPrior.priors > 1 {
					adapters[i].priors += posPrior.priors
				} else {
					adapters[i].priors += 1
				}
			}
		}
	}
	return adapters
}

func main() {
	// path := "10_small1.txt"
	// path := "10_small2.txt"
	path := "10.txt"
	adapters := adapterDistances(sortedAdapters(path))
	fmt.Println("Part 1:", adapterDistSums(adapters)[1] * adapterDistSums(adapters)[3])
	fmt.Println("Part 2:", adapterChains(adapters)[len(adapters)-1].priors)
}
