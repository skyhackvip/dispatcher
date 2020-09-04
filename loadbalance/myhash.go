package loadbalance

type HashBalance struct {
	addr []string
	idx  int
}

func (h *HashBalance) Add(param ...string) error {
	return nil
}

func (h *HashBalance) Get() (string, error) {
	return "", nil
}

func (h *HashBalance) Hao() string {
	return "aahao"
}
