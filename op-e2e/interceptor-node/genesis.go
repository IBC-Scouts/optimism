package interceptornode

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum/common"
)

var defaultPeptideDir string

// abciGenesis is the genesis of the abci app
// for now, only grab the genesis block.
type abciGenesis struct {
	GenesisBlock eth.BlockID `json:"genesis_block"`
}

func UpdateL2GenesisHash(rollupCfg *rollup.Config) {
	ethHash := rollupCfg.Genesis.L2.Hash

	abciGenesisBlock := readGenesisFromFile()
	abciHash := abciGenesisBlock.GenesisBlock.Hash

	rollupCfg.Genesis.L2.Hash = compositeHash(abciHash, ethHash)
}

// Initialize the default peptide directory
func init() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	defaultPeptideDir = filepath.Join(userHomeDir, ".peptide")
}

// See CompositeBlock.Hash
func compositeHash(abciHash, ethHash common.Hash) common.Hash {
	buf := ethHash.Bytes()
	buf = append(buf, abciHash.Bytes()...)

	hash := sha256.Sum256(buf)
	return common.BytesToHash(hash[:])
}

// Reads the genesis contents from the home directory peptide uses and returns the genesis block.
// It rm's the genesis file after reading it to prevent it from causing the 'config already exists'
// error when starting the node.
func readGenesisFromFile() abciGenesis {
	genesisPath := filepath.Join(defaultPeptideDir, "config", "genesis.json")
	if _, err := os.Stat(genesisPath); err != nil {
		panic(err)
	}
	bz, err := os.ReadFile(genesisPath)
	if err != nil {
		panic(err)
	}

	genesis := abciGenesis{}
	if err := json.Unmarshal(bz, &genesis); err != nil {
		panic(err)
	}

	// Delete the genesis file after reading it
	if err := os.Remove(genesisPath); err != nil {
		panic(err)
	}

	// print the genesis block
	fmt.Printf("Genesis block: %s\n", genesis.GenesisBlock)
	return genesis
}
