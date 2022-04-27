package types

const (
	// ModuleName defines the module name
	ModuleName = "bridge"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_bridge"

	// DepositsKey defines the observed deposits store key
	DepositsKey = "deposits"

	// WithdrawalCountsKey defines the withdrawal counts store key
	WithdrawalCountsKey = "withdrawal_counts"

	// WithdrawalsKey defines the withdrawals store key
	WithdrawalsKey = "withdrawals"

	// WithdrawalApprovalsKey defines the withdrawal approvals store key
	WithdrawalApprovalsKey = "withdrawal_approvals"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
