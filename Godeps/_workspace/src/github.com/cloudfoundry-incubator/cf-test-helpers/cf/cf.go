package cf

import (
	"github.com/starkandwayne/cf-logstash-smoke-tests/Godeps/_workspace/src/github.com/cloudfoundry-incubator/cf-test-helpers/runner"
	"github.com/starkandwayne/cf-logstash-smoke-tests/Godeps/_workspace/src/github.com/onsi/gomega/gexec"
)

var Cf = func(args ...string) *gexec.Session {
	return runner.Run("cf", args...)
}
