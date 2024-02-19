package interceptornode

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

// Regular expression to match the http and ws endpoints in the interceptor's logs
const rgxp = `(http|ws)=\S+`

type interceptorSession struct {
	session   *gexec.Session
	Endpoints *external.Endpoints
}

func getBinaryPath() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	basePath := strings.SplitAfter(wd, "op-e2e/")[0]
	binPath := fmt.Sprintf("%s/%s", basePath, "interceptor-node/interceptor")
	fmt.Printf("Base path: %s\n", basePath)

	if _, err := os.Stat(binPath); err != nil {
		return "", fmt.Errorf("could not locate interceptor in working directory: %w", err)
	}

	return binPath, nil
}

// Run initialization command for the binary
func binInit(binPath string) error {
	// TODO(jim): wtf for --l1-hash and --l1-height?
	// ./binary init --l1-hash "0x92838392" --l1-height 1
	cmd := exec.Command(
		binPath,
		"--l1-hash", "0x92838392",
		"--l1-height", "1",
		"--override", "True",
		"init",
	)

	return cmd.Run()
}

// Run seal command of the binary
func binSeal(binPath string) error {
	// ./binary seal
	cmd := exec.Command(
		binPath,
		"seal",
	)

	return cmd.Run()
}

func binStart(binPath, gethEngineAddr string) (*interceptorSession, error) {
	// ./binary start --geth-engine-addr <geth-client-addr>
	cmd := exec.Command(
		binPath,
		"--geth-engine-addr", gethEngineAddr,
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
				return nil, fmt.Errorf("op-polymer exited before announcing http ports")
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

// Execute all the steps to start the interceptor
func BinRun(gethEngineAddr string) (*interceptorSession, error) {
	binPath, err := getBinaryPath()
	if err != nil {
		return nil, fmt.Errorf("could not get binary path: %w", err)
	}

	err = binInit(binPath)
	if err != nil {
		return nil, fmt.Errorf("could not init binary: %w", err)
	}
	fmt.Printf("Binary initialized\n")

	err = binSeal(binPath)
	if err != nil {
		return nil, fmt.Errorf("could not seal binary: %w", err)
	}
	fmt.Printf("Binary sealed\n")
	// TODO: Maybe dump genesis?

	sess, err := binStart(binPath, gethEngineAddr)
	if err != nil {
		return nil, fmt.Errorf("could not start binary: %w", err)
	}
	fmt.Printf("Binary started\n")

	return sess, nil
}
