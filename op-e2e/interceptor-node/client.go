package interceptornode

import (
	"context"
	"fmt"

	"github.com/ethereum-optimism/optimism/op-node/node"
	"github.com/ethereum-optimism/optimism/op-service/client"
	"github.com/ethereum-optimism/optimism/op-service/testlog"
	"github.com/ethereum/go-ethereum/log"
)

// CosmosTxResult is the response of the 'cosmos_sendTransaction' endpoint
type CosmosTxResult struct{}

// CosmosRPCClient should implement any 'cosmos_' prefixed RPC client methods in the interceptor
type cosmosRPCClient struct {
	logger log.Logger
	client client.RPC
}

// Creates a new RPC client that connects to the interceptor node.
func CreateCosmosClient(t testlog.Testing, l2Endpoint node.L2EndpointSetup) (*cosmosRPCClient, error) {
	interceptorAddress := l2Endpoint.(*node.L2EndpointConfig).L2EngineAddr

	logger := testlog.Logger(t, log.LvlInfo).New("interceptor")
	clientRPC, err := client.NewRPC(context.TODO(), logger, interceptorAddress)
	if err != nil {
		return nil, fmt.Errorf("could not create interceptor client: %w", err)
	}

	cosmosRPCClient := &cosmosRPCClient{
		logger: logger,
		client: clientRPC,
	}

	return cosmosRPCClient, nil
}

// Close the RPC client.
func (c *cosmosRPCClient) Close() {
	c.client.Close()
}

// TODO(jim): wrap tx in a better type?
// SendCosmosTx sends a transaction to the 'cosmos_sendTx' endpoint of the interceptor
func (c *cosmosRPCClient) SendCosmosTx(tx []byte) (*CosmosTxResult, error) {
	c.logger.Info("Sending Cosmos Tx", "tx", tx)

	var result CosmosTxResult
	err := c.client.CallContext(context.TODO(), &result, "cosmos_sendTransaction", tx)
	if err != nil {
		return nil, fmt.Errorf("could not send cosmos tx: %w", err)
	}

	c.logger.Info("cosmos_sendTransaction result", "result", result)
	return &result, nil
}
