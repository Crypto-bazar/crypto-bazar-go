package models

type User struct {
	ID         uint   `db:"id"`
	EthAddress string `db:"eth_address"`
}
