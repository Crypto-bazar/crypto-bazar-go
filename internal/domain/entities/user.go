package entities

// User представляет пользователя.
type User struct {
    // ID уникальный идентификатор пользователя
    // swagger:model
    ID uint `db:"id"`

    // EthAddress Ethereum адрес пользователя
    // swagger:model
    EthAddress string `db:"eth_address"`
}