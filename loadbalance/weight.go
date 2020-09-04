package loadbalance

import (
	"errors"
	"strconv"
)

//node list
type WeightRoundRobinBalance struct {
	curIdx   int
	allNodes []*WeightNode
}

//weight node
type WeightNode struct {
	node          string
	weight        int //init weight
	currentWeight int //every round weight
}

//add node
func (r *WeightRoundRobinBalance) Add(params ...string) error {
	if len(params) != 2 {
		return errors.New("param len need 2")
	}
	parInt, err := strconv.ParseInt(params[1], 10, 64)
	if err != nil {
		return err
	}
	node := &WeightNode{node: params[0], weight: int(parInt)}
	r.allNodes = append(r.allNodes, node)
	return nil
}

//get node
func (r *WeightRoundRobinBalance) Get() (string, error) {
	totalWeight := 0
	var bestNode *WeightNode
	for i := 0; i < len(r.allNodes); i++ {
		curNode := r.allNodes[i]
		totalWeight += curNode.weight
		curNode.currentWeight += curNode.weight

		//choose the largest weight
		if bestNode == nil || curNode.currentWeight > bestNode.currentWeight {
			bestNode = curNode
		}
	}
	if bestNode == nil {
		return "", errors.New("get error")
	}
	bestNode.currentWeight -= totalWeight
	return bestNode.node, nil
}
