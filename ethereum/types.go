package ethereum

import "github.com/ethereum/go-ethereum/ethclient"

// EthService ethereum service
type EthService struct {
	// ks     *keystore.KeyStore
	Client *ethclient.Client
}
