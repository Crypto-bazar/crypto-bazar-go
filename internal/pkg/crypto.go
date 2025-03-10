package crypto

import (
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/storyicon/sigverify"
)

func VerifySignature(message, signature, expectedAddress string) (bool, error) {
	valid, err := sigverify.VerifyEllipticCurveHexSignatureEx(
		ethcommon.HexToAddress(expectedAddress),
		[]byte(message),
		signature,
	)

	return valid, err
}