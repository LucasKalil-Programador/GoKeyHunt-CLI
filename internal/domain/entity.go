package domain

type Range struct {
	Min    string `json:"min"`
	Max    string `json:"max"`
	Status int    `json:"status"`
}

type Ranges struct {
	Ranges []Range `json:"ranges"`
}

type Wallets struct {
	Addresses [][]byte `json:"wallets"`
}

type Parameters struct {
	WorkerCount    int
	TargetWallet   int
	UpdateInterval int
	BatchCount     int
	BatchSize      int64
	Rng            bool
}
