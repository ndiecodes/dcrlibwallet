package dcrlibwallet

import (
	"github.com/decred/dcrwallet/wallet"
	"github.com/raedahgroup/dcrlibwallet/txhelper"
)

type Account struct {
	Number           int32
	Name             string
	Balance          *Balance
	TotalBalance     int64
	ExternalKeyCount int32
	InternalKeyCount int32
	ImportedKeyCount int32
}

type Balance struct {
	Total                   int64
	Spendable               int64
	ImmatureReward          int64
	ImmatureStakeGeneration int64
	LockedByTickets         int64
	VotingAuthority         int64
	UnConfirmed             int64
}

type Accounts struct {
	Count              int
	ErrorMessage       string
	ErrorCode          int
	ErrorOccurred      bool
	Acc                []*Account
	CurrentBlockHash   []byte
	CurrentBlockHeight int32
}

type UnspentOutput struct {
	TransactionHash []byte
	OutputIndex     uint32
	OutputKey       string
	ReceiveTime     int64
	Amount          int64
	FromCoinbase    bool
	Tree            int32
	PkScript        []byte
}

type UnsignedTransaction struct {
	UnsignedTransaction       []byte
	EstimatedSignedSize       int
	ChangeIndex               int
	TotalOutputAmount         int64
	TotalPreviousOutputAmount int64
}

type Transaction struct {
	Hash          string `storm:"id,unique"`
	Raw           string
	Hex           string
	Confirmations int32
	Fee           int64
	Timestamp     int64
	Type          string
	Amount        int64
	BlockHeight   int32
	Direction     txhelper.TransactionDirection
	Debits        []*TransactionDebit
	Credits       []*TransactionCredit
}

type TransactionDebit struct {
	Index           int32
	PreviousAccount int32
	PreviousAmount  int64
	AccountName     string
}

type TransactionCredit struct {
	Index    int32
	Account  int32
	Internal bool
	Amount   int64
	Address  string
}

type PurchaseTicketsRequest struct {
	Passphrase            []byte
	Account               uint32
	RequiredConfirmations uint32
	TicketAddress         string
	NumTickets            uint32
	PoolAddress           string
	PoolFees              float64
	Expiry                uint32
	TxFee                 int64
	TicketFee             int64
}

type GetTicketsRequest struct {
	StartingBlockHash   []byte
	StartingBlockHeight int32
	EndingBlockHash     []byte
	EndingBlockHeight   int32
	TargetTicketCount   int32
}

type GetTicketsResponse struct {
	Ticket       *wallet.TicketSummary
	BlockHeight  uint32
	TicketStatus TicketStatus
}

type TicketPriceResponse struct {
	TicketPrice int64
	Height      int32
}
