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

func getConfigPath() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	basePath := strings.SplitAfter(wd, "op-e2e/")[0]
	configPath := fmt.Sprintf("%s/%s", basePath, "interceptor-node/config.json")

	if _, err := os.Stat(configPath); err != nil {
		return "", fmt.Errorf("could not locate config.json in working directory: %w", err)
	}

	return configPath, nil
}

// BinRun starts the interceptor binary and returns the session and endpoints
func BinRun(gethEngineAddr string) (*interceptorSession, error) {
	binPath, err := getBinaryPath()
	if err != nil {
		return nil, fmt.Errorf("could not get binary path: %w", err)
	}

	configPath, err := getConfigPath()
	if err != nil {
		return nil, fmt.Errorf("could not get config path: %w", err)
	}

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
