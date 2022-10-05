import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethCrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func MakeKey() (string, string, error) {
	privateKey, err := ethCrypto.GenerateKey()
	if err != nil {
		return "", "", err
	}

	publicKey := privateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := ethCrypto.FromECDSAPub(publicKeyECDSA)
	privateKeyBytes := ethCrypto.FromECDSA(privateKey)

	publicKeyHex := hexutil.Encode(publicKeyBytes)[4:]   // 0x와 04 제거
	privateKeyHex := hexutil.Encode(privateKeyBytes)[2:] // 0x 제거

	fmt.Printf("공개키: %s, 개인키: %s\n", publicKeyHex, privateKeyHex)

	return publicKeyHex, privateKeyHex, nil
}
