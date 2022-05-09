package types

const (
	// ModuleName defines the module name
	ModuleName = "dex"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_dex"

	// BooksStoreKey defines the books store key
	BooksStoreKey = "books"

	// AccountOrdersStoreKey defines the account orders store key
	AccountOrdersStoreKey = "account_orders"

	// AccountOrdersCountKey defines the account orders count key
	AccountOrdersCountKey = "account_orders_count"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
