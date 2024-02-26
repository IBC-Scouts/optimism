package interceptornode

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	"github.com/cosmos/cosmos-sdk/x/auth"
	// authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/consensus"
	"github.com/cosmos/cosmos-sdk/x/mint"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking"

	// cosmostx "cosmossdk.io/api/cosmos/tx/v1beta1"
	signing "github.com/cosmos/cosmos-sdk/types/tx/signing"
	xauthsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
)

const (
	// https://github.com/cosmos/chain-registry/blob/master/cosmoshub/chain.json
	ChainId  = "cosmoshub-4"
	// GrpcAddr = "cosmos-grpc.polkachu.com:14990"
)

// WalletSigner is a struct that holds the basics for signing a transaction.
// This is not required, but is nice to have instead of passing through tons of method arguments.
type WalletSigner struct {
	Ctx       context.Context
	TxBuilder client.TxBuilder
	EncCfg    moduletestutil.TestEncodingConfig
	Keyring   keyring.Keyring
}

// SetupWalletSigner sets up the wallet signer basics and returns a pointer to
// the WalletSigner struct.
func SetupWalletSigner() *WalletSigner {
	// To sign a transaction, the AppModuleBasic must be provided here. This
	// is for the protobuf (so we can encode/decode the transaction bytes)
	encCfg := moduletestutil.MakeTestEncodingConfig(
		auth.AppModuleBasic{},
		bank.AppModuleBasic{},
		staking.AppModuleBasic{},
		mint.AppModuleBasic{},
		params.AppModuleBasic{},
		slashing.AppModuleBasic{},
		consensus.AppModuleBasic{},
	)

	// Setup a struct to share data with helper methods.
	w := &WalletSigner{
		Ctx:      context.Background(),
		EncCfg:   encCfg,
		Keyring:  keyring.NewInMemory(encCfg.Codec),
	}

	// Set the TxBuilder to be empty
	w.ResetTxBuilder()

	return w
}

// ResetTxBuilder resets the TxBuilder to a new TxBuilder.
func (w *WalletSigner) ResetTxBuilder() {
	w.TxBuilder = w.EncCfg.TxConfig.NewTxBuilder()
}

// LoadKeyFromMnemonic loads a key from a mnemonic and returns the keyring record and the address.
func (w *WalletSigner) LoadKeyFromMnemonic(keyName, mnemonic, password string) (*keyring.Record, sdk.AccAddress) {
	path := sdk.GetConfig().GetFullBIP44Path()
	r, err := w.Keyring.NewAccount(keyName, mnemonic, password, path, hd.Secp256k1)
	if err != nil {
		panic(err)
	}

	a, err := r.GetAddress()
	if err != nil {
		panic(err)
	}

	return r, a
}

func (w *WalletSigner) GetTxBytes() []byte {
	// Generated Protobuf-encoded bytes.
	txBytes, err := w.EncCfg.TxConfig.TxEncoder()(w.TxBuilder.GetTx())
	if err != nil {
		panic(err)
	}

	return txBytes
}

func (w *WalletSigner) SignTx(keyName string, account_num, sequence uint64,) error {
	k, err := w.Keyring.Key(keyName)
	if err != nil {
		return err
	}

	krAcc, err := k.GetAddress()
	if err != nil {
		return err
	}

	pubKey, err := k.GetPubKey()
	if err != nil {
		return err
	}

	// Get the base Tx bytes
	txBytes, err := w.EncCfg.TxConfig.TxEncoder()(w.TxBuilder.GetTx())
	if err != nil {
		return err
	}

	// First round: we gather all the signer infos. We use the "set empty
	// signature" hack to do that.
	if err := w.TxBuilder.SetSignatures(signing.SignatureV2{
		PubKey: pubKey,
		Data: &signing.SingleSignatureData{
			SignMode:  signing.SignMode(w.EncCfg.TxConfig.SignModeHandler().DefaultMode()),
			Signature: nil,
		},
		Sequence: sequence,
	}); err != nil {
		panic(err)
	}

	_, _, err = w.Keyring.Sign(keyName, txBytes, signing.SignMode(w.EncCfg.TxConfig.SignModeHandler().DefaultMode()))
	if err != nil {
		panic(err)
	}

	// Second round: all signer infos are set, so each signer can sign.
	_ = xauthsigning.SignerData{
		Address:       krAcc.String(),
		ChainID:       ChainId,
		AccountNumber: account_num,
		Sequence:      sequence,
		PubKey:        pubKey,
	}

	return nil
}
