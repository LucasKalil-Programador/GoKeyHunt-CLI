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
	WorkerCount     int   // 4 bytes
	TargetWallet    int   // 4 bytes
	UpdateInterval  int   // 4 bytes
	BatchCount      int   // 4 bytes
	BatchSize       int64 // 8 bytes
	Rng             bool  // 1 byte
	VerboseSummary  bool  // 1 byte
	VerboseProgress bool  // 1 byte
	VerboseKeyFind  bool  // 1 byte + 4 padding
}
