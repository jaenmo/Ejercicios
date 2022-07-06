package main

import "fmt"

func main() {
	pairs := []Pair{
		{C1: "BNB", C2: "ETH", ExchangeRate: 7.22169, Fee: 0.25},
		{C1: "ETH", C2: "BNB", ExchangeRate: 0.138471, Fee: 0.25},
		{C1: "BNB", C2: "USDT", ExchangeRate: 0.0027379, Fee: 0.25},
		{C1: "USDT", C2: "BNB", ExchangeRate: 365.271, Fee: 0.25},
		{C1: "ETH", C2: "USDT", ExchangeRate: 0.000381627, Fee: 0.25},
		{C1: "USDT", C2: "ETH", ExchangeRate: 2620.36, Fee: 0.25},
	}

	// for {
	//    time.Sleep(time.Second * 5)
	//    var pairs []Pair
	//    pairs = append(pairs, network.GetPair(["BNB", "ETH"]))
	//    pairs = append(pairs, network.GetPair(["ETH", "BNB"]))
	//    pairs = append(pairs, network.GetPair(["BNB", "USDT"]))
	//    ...
	//    bestPair := findBestPath(pairs)
	//    if bestPath == nil {
	//    	fmt.Println("No path found")
	//    	continue
	//    }
	//    fmt.Println(bestPath) // FOUND A DEAL!
	//
	// 	  amount := network.GetBalance(bestPath.Pairs[0].C1) / 2 // The starting currency
	//
	//    for _, pair := range bestPath.Pairs { // ETH -> BNB -> USDT -> ETH
	// 		amount = network.Swap(pair.C1, pair.C2, amount)
	//    }
	//    network.ExecuteSmartContract("flashLoan", bestPath, 10,000,000)
	//
	//    fmt.Println(amount) // 100.66 ETH
	//    100 ETH -> 100.66 ETH: profit is 0.66 ETH
	// }

	var paths []*Path

	for i := 0; i < len(pairs); i++ {
		for j := 0; j < len(pairs); j++ {
			for k := 0; k < len(pairs); k++ {
				if i != j && i != k && j != k {
					if pairs[i].C2 == pairs[j].C1 && pairs[j].C2 == pairs[k].C1 && pairs[k].C2 == pairs[i].C1 {
						paths = append(paths, &Path{Pairs: []Pair{pairs[i], pairs[j], pairs[k]}})
					}
				}
			}
		}
	}

	t := findBestPath(paths)

	fmt.Println("BEST PATH: ", t)
}

type Pair struct {
	C1           string
	C2           string
	ExchangeRate float64
	Fee          float64
}

func (p Pair) Exchange(amountC1 float64) float64 {
	return amountC1 / p.ExchangeRate * (100 - p.Fee) / 100
}

type Path struct {
	Pairs  []Pair
	Amount float64
}

func (p Path) String() string {
	var (
		s, first string
	)
	for i, pair := range p.Pairs {
		if i == 0 {
			first = pair.C1
		}
		s += fmt.Sprintf("%s -> ", pair.C1)
	}
	return s + " " + fmt.Sprintf("%s: %.2f", first, p.Amount)
}

func (p *Path) Exchange(a0 float64) {
	for _, p := range p.Pairs {
		a0 = p.Exchange(a0)
	}
	p.Amount = a0
}

func findBestPath(paths []*Path) *Path {
	const inputAmount = 100

	var (
		endAmount float64
		bestPath  *Path
	)

	for i := range paths {
		paths[i].Exchange(inputAmount)
		exchangedAmount := paths[i].Amount

		fmt.Printf("EXCHANGE %d: %s \n", i+1, paths[i])

		if exchangedAmount > inputAmount && exchangedAmount > endAmount {
			endAmount = exchangedAmount
			bestPath = paths[i]
		}
	}

	return bestPath
}