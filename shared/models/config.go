package models

type Config struct {
	ETH_NODE              string
	ETH_PRIVATE_KEY       string
	ARWEAVE_KEY_FILE      string
	ARWEAVE_NODE          string
	FILE_PORT             string
	ENDPOINT              string
	FEE_PER_BYTE          int64
	MIN_BOUNTY            int64
	MIN_DIGGING_FEE       int64
	MAX_RESURRECTION_TIME int64
	CONTRACT_ADDRESS      string
	TOKEN_ADDRESS         string
	ADD_TO_FREE_BOND      int64
	REMOVE_FROM_FREE_BOND int64
	PAYMENT_ADDRESS       string
	GAS_PRICE_OVERRIDE    int64
	MNEMONIC              string
}
