package loadbalance

import "errors"

//list store all the node
type RoundRobinBalance struct {
	curIdx   int
	allNodes []string
}

//add node
func (r *RoundRobinBalance) Add(params ...string) error {
	r.allNodes = append(r.allNodes, params[0])
	return nil
}

//get node
func (r *RoundRobinBalance) Get() (string, error) {
	if len(r.allNodes) == 0 {
		return "", errors.New("addr list is empty")
	}
	lens := len(r.allNodes)
	//move to beginning node
	if r.curIdx >= lens {
		r.curIdx = 0
	}
	curNode := r.allNodes[r.curIdx]
	//move to next node
	r.curIdx = (r.curIdx + 1) % lens
	return curNode, nil
}
