package loadbalance

import (
	"errors"
	"math/rand"
)

//list store all the node
type RandomBalance struct {
	curIdx   int
	allNodes []string
}

//add node
func (r *RandomBalance) Add(params ...string) error {
	if len(params) == 0 {
		return errors.New("param len 1 at least")
	}
	addr := params[0]
	r.allNodes = append(r.allNodes, addr)
	return nil
}

//get node
func (r *RandomBalance) Get() (string, error) {
	if len(r.allNodes) == 0 {
		return "", errors.New("allNodes is empty")
	}
	r.curIdx = rand.Intn(len(r.allNodes)) //use rand to generate random
	return r.allNodes[r.curIdx], nil
}
