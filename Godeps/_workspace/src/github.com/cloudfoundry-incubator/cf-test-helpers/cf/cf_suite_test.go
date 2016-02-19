package cf_test

import (
	"testing"

	"github.com/starkandwayne/cf-logstash-smoke-tests/Godeps/_workspace/src/github.com/cloudfoundry-incubator/cf-test-helpers/cf"
	"github.com/starkandwayne/cf-logstash-smoke-tests/Godeps/_workspace/src/github.com/cloudfoundry-incubator/cf-test-helpers/runner"
	. "github.com/starkandwayne/cf-logstash-smoke-tests/Godeps/_workspace/src/github.com/onsi/ginkgo"
	. "github.com/starkandwayne/cf-logstash-smoke-tests/Godeps/_workspace/src/github.com/onsi/gomega"
)

var originalCf = cf.Cf
var originalCommandInterceptor = runner.CommandInterceptor

var _ = AfterEach(func() {
	cf.Cf = originalCf
	runner.CommandInterceptor = originalCommandInterceptor
})

func TestCf(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cf Suite")
}
