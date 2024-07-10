package domain

// Range represents a range with a minimum and maximum value and a status.
//
// Fields:
// - Min: The minimum value of the range (string).
// - Max: The maximum value of the range (string).
// - Status: The status of the range (integer).
type Range struct {
	Min    string `json:"min"`
	Max    string `json:"max"`
	Status int    `json:"status"`
}

// Ranges represents a collection of Range objects.
//
// Fields:
// - Ranges: A slice of Range objects.
type Ranges struct {
	Ranges []Range `json:"ranges"`
}

// Wallets represents a collection of wallet addresses.
//
// Fields:
// - Addresses: A slice of byte slices, where each byte slice represents a wallet address.
type Wallets struct {
	Addresses [][]byte `json:"wallets"`
}

// Parameters represents the configuration parameters for the application.
//
// Fields:
// - WorkerCount: Number of worker threads (integer).
// - TargetWallet: Index of the target wallet (integer).
// - UpdateInterval: Interval for progress updates in seconds (integer).
// - BatchCount: Number of batches (integer).
// - BatchSize: Size of each batch (int64).
// - Rng: Flag to indicate if a random start location should be generated (boolean).
// - VerboseSummary: Flag to enable or disable verbose summary output (boolean).
// - VerboseProgress: Flag to enable or disable verbose progress output (boolean).
// - VerboseKeyFind: Flag to enable or disable verbose key find output (boolean).
//
// Note: The Parameters struct layout is designed with memory alignment considerations,
// so the boolean fields are followed by 4 bytes of padding.
type Parameters struct {
	WorkerCount     int   // 4 bytes
	TargetWallet    int   // 4 bytes
	UpdateInterval  int   // 4 bytes
	BatchCount      int   // 4 bytes
	BatchSize       int64 // 8 bytes
	Rng             bool  // 1 byte
	VerboseSummary  bool  // 1 byte
	VerboseProgress bool  // 1 byte
	VerboseKeyFind  bool  // 1 byte + 4 bytes padding
}
