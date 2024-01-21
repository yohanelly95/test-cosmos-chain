package types

const (
	// ModuleName defines the module name
	ModuleName = "razorchain"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_razorchain"
)

var (
	ParamsKey = []byte("p_razorchain")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
