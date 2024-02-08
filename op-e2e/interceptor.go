package op_e2e

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/ethereum-optimism/optimism/op-e2e/external"
	"github.com/onsi/gomega/gexec"
)

type interceptorSession struct {
	session   *gexec.Session
	endpoints *external.Endpoints
}

func start(binPath, gethEngineAddr string) (*interceptorSession, error) {
	cmd := exec.Command(
		binPath,
		"--geth-engine-addr", gethEngineAddr,
		"start",
	)
	sess, err := gexec.Start(cmd, os.Stdout, os.Stderr)
	if err != nil {
		return nil, fmt.Errorf("could not start interceptor session: %w", err)
	}

	return &interceptorSession{
		session: sess,
		// 		endpoints: &external.Endpoints{
		//			HTTPEndpoint:     fmt.Sprintf("http://127.0.0.1:%d/", httpPort),
		//			WSEndpoint:       fmt.Sprintf("ws://127.0.0.1:%d/", httpPort),
		//			HTTPAuthEndpoint: fmt.Sprintf("http://127.0.0.1:%d/", enginePort),
		//			WSAuthEndpoint:   fmt.Sprintf("ws://127.0.0.1:%d/", enginePort),
		// },
	}, nil
}
