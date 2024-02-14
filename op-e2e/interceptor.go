package op_e2e

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"

	"github.com/ethereum-optimism/optimism/op-e2e/external"
	rollupNode "github.com/ethereum-optimism/optimism/op-node/node"
	"github.com/ethereum-optimism/optimism/op-service/client"
	"github.com/ethereum-optimism/optimism/op-service/testlog"
	"github.com/ethereum/go-ethereum/log"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

// Regular expression to match the http and ws endpoints in the interceptor's logs
const rgxp = `(http|ws)=\S+`

// CosmosTxResult is the response of the 'cosmos_sendTransaction' endpoint
type CosmosTxResult struct{}

// CosmosRPCClient implements the 'cosmos_' prefixed RPC client methods in the interceptor
type cosmosRPCClient struct {
	logger log.Logger
	client client.RPC
}

// Creates a new RPC client that connects to the interceptor
func CreateCosmosClient(sys *System) (*cosmosRPCClient, error) {
	interceptorAddress := sys.Cfg.Nodes["sequencer"].L2.(*rollupNode.L2EndpointConfig).L2EngineAddr

	logger := testlog.Logger(sys.t, log.LvlInfo).New("interceptor")
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

type interceptorSession struct {
	session   *gexec.Session
	Endpoints *external.Endpoints
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

func start(binPath, configPath, gethEngineAddr string) (*interceptorSession, error) {
	cmd := exec.Command(
		binPath,
		"--geth-engine-addr", gethEngineAddr,
		"--config", configPath,
		"start",
	)
	sess, err := gexec.Start(cmd, os.Stdout, os.Stderr)
	if err != nil {
		return nil, fmt.Errorf("could not start interceptor session: %w", err)
	}

	// code copied from https://github.com/polymerdao/optimism-dev/blob/518341f3e2dc7bf88eb06513a740fc9ced1ccf39/op-e2e/e2eutils/external_polymer/main.go#L150
	// Modified to look in stderr since our logger logs there!
	matcher := gbytes.Say("Execution engine rpc server enabled")
	var httpUrl, wsUrl string
	urlRE := regexp.MustCompile(rgxp)
	for httpUrl == "" && wsUrl == "" {
		match, err := matcher.Match(sess.Err)
		if err != nil {
			return nil, fmt.Errorf("could not execute matcher")
		}
		if !match {
			if sess.Out.Closed() {
				return nil, fmt.Errorf("interceptor exited before announcing http ports")
			}
			// Wait for a bit more output, then try again
			time.Sleep(10 * time.Millisecond)
			continue
		}

		for _, line := range strings.Split(string(sess.Err.Contents()), "\n") {
			found := urlRE.FindAllString(line, -1)
			if len(found) == 2 {
				httpUrl, _ = strings.CutPrefix(found[0], "http=")
				wsUrl, _ = strings.CutPrefix(found[1], "ws=")
				break
			}
		}
	}

	return &interceptorSession{
		session: sess,
		Endpoints: &external.Endpoints{
			HTTPEndpoint:     httpUrl,
			WSEndpoint:       wsUrl,
			HTTPAuthEndpoint: httpUrl,
			WSAuthEndpoint:   wsUrl,
		},
	}, nil
}
