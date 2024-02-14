package op_e2e

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/wait"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"
)

func TestSendCosmosTx(t *testing.T) {
	InitParallel(t)
	cfg := DefaultSystemConfig(t)
	delete(cfg.Nodes, "verifier")
	sys, err := cfg.Start(t)
	require.Nil(t, err, "Error starting up system")
	defer sys.Close()

	l1Client := sys.Clients["l1"]
	l2Verif := sys.Clients["sequencer"]

	cosmosClient, err := CreateCosmosClient(sys)
	require.Nil(t, err, "Error creating cosmos client")
	defer cosmosClient.client.Close()

	// invoke sendTx with random data
	res, err := cosmosClient.SendCosmosTx([]byte("blob"))
	require.Nil(t, err, "Error sending cosmos tx")
	require.NotNil(t, res, "Expected a response")

	// create signer
	aliceKey := cfg.Secrets.Alice
	opts, err := bind.NewKeyedTransactorWithChainID(aliceKey, cfg.L1ChainIDBig())
	require.Nil(t, err)
	fromAddr := opts.From

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	startBalance, err := l2Verif.BalanceAt(ctx, fromAddr, nil)
	cancel()
	require.Nil(t, err)

	ctx, cancel = context.WithTimeout(context.Background(), 1*time.Second)
	startNonce, err := l2Verif.NonceAt(ctx, fromAddr, nil)
	require.NoError(t, err)
	cancel()

	toAddr := common.Address{0xff, 0xff}
	mintAmount := big.NewInt(9_000_000)
	opts.Value = mintAmount
	SendDepositTx(t, cfg, l1Client, l2Verif, opts, func(l2Opts *DepositTxOpts) {
		l2Opts.ToAddr = toAddr
		// trigger a revert by transferring more than we have available
		l2Opts.Value = new(big.Int).Mul(common.Big2, startBalance)
		l2Opts.ExpectedStatus = types.ReceiptStatusFailed
	})

	// Confirm balance
	ctx, cancel = context.WithTimeout(context.Background(), 15*time.Second)
	endBalance, err := wait.ForBalanceChange(ctx, l2Verif, fromAddr, startBalance)
	cancel()
	require.Nil(t, err)

	ctx, cancel = context.WithTimeout(context.Background(), 1*time.Second)
	toAddrBalance, err := l2Verif.BalanceAt(ctx, toAddr, nil)
	cancel()
	require.NoError(t, err)

	diff := new(big.Int)
	diff = diff.Sub(endBalance, startBalance)
	require.Equal(t, mintAmount, diff, "Did not get expected balance change")
	require.Equal(t, common.Big0.Int64(), toAddrBalance.Int64(), "The recipient account balance should be zero")

	ctx, cancel = context.WithTimeout(context.Background(), 1*time.Second)
	endNonce, err := l2Verif.NonceAt(ctx, fromAddr, nil)
	require.NoError(t, err)
	cancel()
	require.Equal(t, startNonce+1, endNonce, "Nonce of deposit sender should increment on L2, even if the deposit fails")
}
