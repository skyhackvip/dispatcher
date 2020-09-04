package main

import (
	"fmt"
	"github.com/skyhackvip/dispatcher/loadbalance"
)

func main() {
	randomLb := loadbalance.LoadBalanceFactory(loadbalance.Random)
	randomLb.Add("127.0.0.1:1001")
	randomLb.Add("127.0.0.1:1002")
	randomLb.Add("127.0.0.1:1003")
	randomLb.Add("127.0.0.1:1004")
	randomLb.Add("127.0.0.1:1005")

	roundLb := loadbalance.LoadBalanceFactory(loadbalance.RoundRobin)
	roundLb.Add("127.0.0.1:1001")
	roundLb.Add("127.0.0.1:1002")
	roundLb.Add("127.0.0.1:1003")
	roundLb.Add("127.0.0.1:1004")
	roundLb.Add("127.0.0.1:1005")

	weightLb := loadbalance.LoadBalanceFactory(loadbalance.Weight)
	weightLb.Add("127.0.0.1:1001", "1")
	weightLb.Add("127.0.0.1:1002", "2")
	weightLb.Add("127.0.0.1:1003", "3")
	weightLb.Add("127.0.0.1:1004", "4")
	weightLb.Add("127.0.0.1:1005", "5")

	var randomCount = make(map[string]int)
	var roundCount = make(map[string]int)
	var weightCount = make(map[string]int)
	for i := 0; i < 100000; i++ {
		randomRs, _ := randomLb.Get()
		randomCount[randomRs]++
		roundRs, _ := roundLb.Get()
		roundCount[roundRs]++
		weightRs, _ := weightLb.Get()
		weightCount[weightRs]++
	}
	fmt.Println(randomCount)
	fmt.Println(roundCount)
	fmt.Println(weightCount)
}
