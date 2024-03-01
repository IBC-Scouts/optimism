package interceptornode

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/ethereum-optimism/optimism/op-node/node"
	"github.com/ethereum-optimism/optimism/op-service/testlog"
	"github.com/ethereum/go-ethereum/log"
)

// CosmosTxResult is the response of the 'cosmos_sendTransaction' endpoint
type CosmosTxResult struct{}

// interceptorClient allows us to invoke all of the rpc methods defined on the interceptor.
// It wraps the ethclient.Client bound to our engine addr to allow intercepting of calls
// made via `CallContract`.
type interceptorClient struct {
	client *ethclient.Client
	logger log.Logger
}

// EthClient returns the ethclient.Client.
func (c *interceptorClient) EthClient() *ethclient.Client {
	return c.client
}

// RpcClient returns the rpc.Client.
func (c *interceptorClient) RpcClient() *rpc.Client {
	return c.client.Client()
}

// Close the RPC client.
func (c *interceptorClient) Close() {
	c.client.Close()
}

// Creates a new RPC client that connects to the interceptor node.
func NewInterceptorClient(t testlog.Testing, l2Endpoint node.L2EndpointSetup) (*interceptorClient, error) {
	interceptorAddress := l2Endpoint.(*node.L2EndpointConfig).L2EngineAddr

	rpcClient, err := rpc.DialWebsocket(context.TODO(), interceptorAddress, "")
	if err != nil {
		return nil, fmt.Errorf("could not create rpc client: %w", err)
	}

	ethClient := ethclient.NewClient(rpcClient)
	logger := testlog.Logger(t, log.LvlInfo).New("interceptor")

	cosmosRPCClient := &interceptorClient{
		ethClient,
		logger,
	}

	return cosmosRPCClient, nil
}

// SendCosmosTx sends a transaction to the 'cosmos_sendTx' endpoint of the interceptor
func (c *interceptorClient) SendCosmosTx(tx []byte) (*CosmosTxResult, error) {
	c.logger.Info("Sending Cosmos Tx", "tx", tx)

	var result CosmosTxResult
	err := c.RpcClient().CallContext(context.TODO(), &result, "cosmos_sendTransaction", tx)
	if err != nil {
		return nil, fmt.Errorf("could not send cosmos tx: %w", err)
	}

	c.logger.Info("cosmos_sendTransaction result", "result", result)
	return &result, nil
}
