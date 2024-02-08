package op_e2e

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"

	"github.com/ethereum-optimism/optimism/op-e2e/external"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

type interceptorSession struct {
	session   *gexec.Session
	Endpoints *external.Endpoints
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
	matcher := gbytes.Say("Execution engine rpc server enabled")
	var httpUrl, wsUrl string
	urlRE := regexp.MustCompile(`Execution engine rpc server enabled\s+http=(.+)\sws=(.+)`)
	for httpUrl == "" && wsUrl == "" {
		match, err := matcher.Match(sess.Out)
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

		for _, line := range strings.Split(string(sess.Out.Contents()), "\n") {
			found := urlRE.FindStringSubmatch(line)
			if len(found) == 3 {
				httpUrl = found[1]
				wsUrl = found[2]
				continue
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
