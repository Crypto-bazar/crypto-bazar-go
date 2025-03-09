package crypto

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// VerifySignature проверяет подлинность подписи
func VerifySignature(address, message, signature string) bool {
	// Добавляем Ethereum Signed Message (EIP-191)
	prefix := fmt.Sprintf("\x19Ethereum Signed Message:\n%s%s", strconv.Itoa(len(message)), message)
	msgHash := crypto.Keccak256Hash([]byte(prefix))

	// Декодируем подпись
	sigBytes, err := hexutil.Decode(signature)
	if err != nil || len(sigBytes) != 65 {
		return false
	}

	// Корректируем v (Recovery ID)
	if sigBytes[64] < 27 {
		sigBytes[64] += 27
	}

	// Получаем публичный ключ из подписи
	pubKey, err := crypto.SigToPub(msgHash.Bytes(), sigBytes)
	if err != nil {
		return false
	}

	// Получаем адрес подписанта
	signerAddress := crypto.PubkeyToAddress(*pubKey).Hex()

	// Проверяем без учета регистра
	return strings.EqualFold(signerAddress, address)
}
