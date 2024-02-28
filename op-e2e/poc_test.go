package op_e2e

import (
	"context"
	"math/big"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	transfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v8/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v8/modules/core/04-channel/types"
	ibcexported "github.com/cosmos/ibc-go/v8/modules/core/exported"

	e2eutils "github.com/ethereum-optimism/optimism/indexer/e2e_tests/utils"
	"github.com/ethereum-optimism/optimism/op-bindings/bindings"
	"github.com/ethereum-optimism/optimism/op-bindings/predeploys"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/transactions"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/wait"
	interceptornode "github.com/ethereum-optimism/optimism/op-e2e/interceptor-node"
	"github.com/ethereum-optimism/optimism/op-service/testlog"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/require"
)

const (
	keyName = "mykey"
	// This is a test account on the cosmoshub with some uatom. You should use your own via an env variable such as `TestMnemonic := os.Getenv("MY_CLIENT_MNEMONIC")`
	// https://dev.mintscan.io/cosmos/account/cosmos165smcg7d2fqtwewj3uf33rc33dh8h46yns3sm5
	TestMnemonic = "pencil surprise brave age old level saddle because olive find winter auto develop spin milk tunnel make demand tattoo wasp primary save bubble keep"

	testRandomAddr  = "cosmos1vapwvcsr0m32ptal6z6g9hjctywrw4yzyf6y6v"
	testRandomAddr2 = "cosmos1vapwvcsr0m32ptal6z6g9hjctywrw4yzyf6y6y"
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

	cosmosClient, err := interceptornode.CreateCosmosClient(sys.t, sys.Cfg.Nodes["sequencer"].L2)
	require.Nil(t, err, "Error creating cosmos client")
	defer cosmosClient.Close()

	// create cosmos signer
	w := interceptornode.SetupWalletSigner()
	_, acc := w.LoadKeyFromMnemonic(keyName, TestMnemonic, "1234567890")

	// Create a bank send message from our account (acc) -> some other account address
	randomAcc, err := sdk.AccAddressFromBech32(testRandomAddr)
	require.Nil(t, err, "Error converting random address")
	msg1 := banktypes.NewMsgSend(acc, randomAcc, sdk.NewCoins(sdk.NewInt64Coin("uatom", 1)))
	if err := w.TxBuilder.SetMsgs(msg1); err != nil {
		panic(err)
	}

	// Set the transaction information (do before signing the Tx)
	w.TxBuilder.SetGasLimit(100_000)
	w.TxBuilder.SetMemo("my test memo")
	//w.TxBuilder.SetFeeAmount(sdk.NewCoins(sdk.NewInt64Coin("uatom", 750)))

	// // THIS IS NOT DOING ANYTHING AT THE MOMENT
	if err := w.SignTx(keyName, 1, 1); err != nil {
		panic(err)
	}

	// broadcast the Tx that has been signed
	txBytes := w.GetTxBytes()

	// invoke sendTx with random data
	res, err := cosmosClient.SendCosmosTx(txBytes)
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

func TestIBCTransfer(t *testing.T) {
	InitParallel(t)

	cfg := DefaultSystemConfig(t)

	sys, err := cfg.Start(t)
	require.Nil(t, err, "Error starting up system")
	defer sys.Close()

	log := testlog.Logger(t, log.LvlInfo)
	log.Info("genesis", "l2", sys.RollupConfig.Genesis.L2, "l1", sys.RollupConfig.Genesis.L1, "l2_time", sys.RollupConfig.Genesis.L2Time)

	//l1Client := sys.Clients["l1"]
	l2Client := sys.Clients["interceptor"]

	l2Opts, err := bind.NewKeyedTransactorWithChainID(sys.Cfg.Secrets.Alice, cfg.L2ChainIDBig())
	require.NoError(t, err)

	cosmosClient, err := interceptornode.CreateCosmosClient(sys.t, sys.Cfg.Nodes["sequencer"].L2)
	require.Nil(t, err, "Error creating cosmos client")
	defer cosmosClient.Close()

	// tx signing code start here ----------
	// create cosmos signer
	w := interceptornode.SetupWalletSigner()
	//_, acc := w.LoadKeyFromMnemonic(keyName, TestMnemonic, "1234567890")

	proof := []byte{0x01}
	// Create a bank send message from our account (acc) -> some other account address
	randomAcc, err := sdk.AccAddressFromBech32(testRandomAddr)
	require.Nil(t, err, "Error converting random address")
	msg1 := channeltypes.NewMsgChannelOpenInit(transfertypes.PortID, transfertypes.Version, channeltypes.UNORDERED, []string{ibcexported.LocalhostConnectionID}, transfertypes.PortID, randomAcc.String())
	msg2 := channeltypes.NewMsgChannelOpenTry(transfertypes.PortID, transfertypes.Version, channeltypes.UNORDERED, []string{ibcexported.LocalhostConnectionID}, transfertypes.PortID, "channel-0", transfertypes.Version, proof, clienttypes.Height{}, randomAcc.String())
	msg3 := channeltypes.NewMsgChannelOpenAck(transfertypes.PortID, "channel-0", "channel-1", transfertypes.Version, proof, clienttypes.Height{}, randomAcc.String())
	msg4 := channeltypes.NewMsgChannelOpenConfirm(transfertypes.PortID, "channel-1", proof, clienttypes.Height{}, randomAcc.String())

	transfer_msg := transfertypes.NewMsgTransfer(
		transfertypes.PortID,
		"channel-0",
		sdk.NewInt64Coin("uatom", 1),
		randomAcc.String(),
		randomAcc.String(),
		clienttypes.Height{0, 10},
		0,
		"",
	)

	if err := w.TxBuilder.SetMsgs(msg1, msg2, msg3, msg4, transfer_msg); err != nil {
		panic(err)
	}

	// Set the transaction information (do before signing the Tx)
	w.TxBuilder.SetGasLimit(100_000)
	w.TxBuilder.SetMemo("my test memo")
	//w.TxBuilder.SetFeeAmount(sdk.NewCoins(sdk.NewInt64Coin("uatom", 750)))

	// // THIS IS NOT DOING ANYTHING AT THE MOMENT
	//	if err := w.SignTx(keyName, 1, 1); err != nil {
	//		panic(err)
	//	}

	// broadcast the Tx that has been signed
	txBytes := w.GetTxBytes()

	// invoke sendTx with random data
	res, err := cosmosClient.SendCosmosTx(txBytes)
	require.Nil(t, err, "Error sending cosmos tx")
	require.NotNil(t, res, "Expected a response")

	// --------------

	// invoke cross domain messenger (just to test setup of the cross domain messenger)
	ibcMessenger, err := bindings.NewIBCCrossDomainMessenger(predeploys.IBCCrossDomainMessengerAddr, l2Client)
	require.NoError(t, err)
	tx, err := transactions.PadGasEstimate(l2Opts, 1.1, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		return ibcMessenger.SendMessage(l2Opts, l2Opts.From, []byte("hello cosmos"), 100000)
	})
	require.NoError(t, err)
	ibcMsgReceipt, err := wait.ForReceiptOK(context.Background(), l2Client, tx.Hash())
	require.NoError(t, err)

	t.Log("Message sent through IBCCrossDomainMessenger", "gas used", ibcMsgReceipt.GasUsed)

	crossDomainMsg, err := e2eutils.ParseCrossDomainMessage(ibcMsgReceipt)
	require.NoError(t, err)

	t.Log("cross chain messenger event:", "sender", crossDomainMsg.Sender, "to:", crossDomainMsg.Target, "gas limit:", crossDomainMsg.GasLimit, "message:", string(crossDomainMsg.Message), "amount:", crossDomainMsg.Value)

	// initiate IBC transfer
	l2Opts.Value = big.NewInt(params.Ether)
	ibcEscrowContract, err := bindings.NewIBCStandardBridge(predeploys.IBCStandardBridgeAddr, l2Client)
	require.NoError(t, err)

	messengerAddr, err := ibcEscrowContract.Messenger(&bind.CallOpts{Context: context.Background()})
	require.NoError(t, err)
	require.Equal(t, predeploys.IBCCrossDomainMessengerAddr, messengerAddr)

	tx, err = ibcEscrowContract.Withdraw(l2Opts, predeploys.LegacyERC20ETHAddr, l2Opts.Value, 200_000, []byte{byte(1)})
	require.NoError(t, err)
	transferReceipt, err := wait.ForReceiptOK(context.Background(), l2Client, tx.Hash())
	require.NoError(t, err)

	t.Log("Message sent through IBCStandardBridge", "gas used", ibcMsgReceipt.GasUsed)

	crossDomainMsg, err = e2eutils.ParseCrossDomainMessage(transferReceipt)
	require.NoError(t, err)

	t.Log("cross chain messenger event:", "sender", crossDomainMsg.Sender, "to:", crossDomainMsg.Target, "gas limit:", crossDomainMsg.GasLimit, "message:", string(crossDomainMsg.Message), "amount:", crossDomainMsg.Value)

	//
	//	w = interceptornode.SetupWalletSigner()
	//_, acc := w.LoadKeyFromMnemonic(keyName, TestMnemonic, "1234567890")

	// Create a bank send message from our account (acc) -> some other account address
	//	msgTransfer := transfertypes.NewMsgTransfer(transfertypes.PortID, "channel-0", transfertypes.Version, channeltypes.UNORDERED, []string{ibcexported.LocalhostConnectionID}, transfertypes.PortID, randomAcc.String())
	//	if err := w.TxBuilder.SetMsgs(msgTransfer); err != nil {
	//		panic(err)
	//	}

	// Set the transaction information (do before signing the Tx)
	//	w.TxBuilder.SetGasLimit(100_000)
	//	w.TxBuilder.SetMemo("my test memo")
	//w.TxBuilder.SetFeeAmount(sdk.NewCoins(sdk.NewInt64Coin("uatom", 750)))

	// // THIS IS NOT DOING ANYTHING AT THE MOMENT
	//	if err := w.SignTx(keyName, 1, 1); err != nil {
	//		panic(err)
	//	}

	// broadcast the Tx that has been signed
	//	txBytes := w.GetTxBytes()

	// invoke sendTx with random data
	//	res, err := cosmosClient.SendCosmosTx(txBytes)
	//	require.Nil(t, err, "Error sending cosmos tx")
	//	require.NotNil(t, res, "Expected a response")
}
