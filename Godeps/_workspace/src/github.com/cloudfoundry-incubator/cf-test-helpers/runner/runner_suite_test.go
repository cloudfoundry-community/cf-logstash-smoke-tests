package runner_test

import (
	. "github.com/starkandwayne/cf-logstash-smoke-tests/Godeps/_workspace/src/github.com/onsi/ginkgo"
	. "github.com/starkandwayne/cf-logstash-smoke-tests/Godeps/_workspace/src/github.com/onsi/gomega"
	"testing"
)

func TestRunner(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Runner Suite")
}
