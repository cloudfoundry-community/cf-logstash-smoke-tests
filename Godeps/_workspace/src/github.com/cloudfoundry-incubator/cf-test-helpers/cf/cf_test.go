package cf_test

import (
	"os/exec"
	"time"

	. "github.com/starkandwayne/cf-logstash-smoke-tests/Godeps/_workspace/src/github.com/cloudfoundry-incubator/cf-test-helpers/cf"
	"github.com/starkandwayne/cf-logstash-smoke-tests/Godeps/_workspace/src/github.com/cloudfoundry-incubator/cf-test-helpers/runner"
	. "github.com/starkandwayne/cf-logstash-smoke-tests/Godeps/_workspace/src/github.com/onsi/ginkgo"
	. "github.com/starkandwayne/cf-logstash-smoke-tests/Godeps/_workspace/src/github.com/onsi/gomega"
)

var _ = Describe("Cf", func() {
	It("sends the request to current CF target", func() {
		runner.CommandInterceptor = func(cmd *exec.Cmd) *exec.Cmd {
			Expect(cmd.Path).To(Equal(exec.Command("cf").Path))
			Expect(cmd.Args).To(Equal([]string{"cf", "apps"}))

			return exec.Command("bash", "-c", `exit 42`)
		}

		runner.NewCmdRunner(Cf("apps"), 1*time.Second).WithExitCode(42).Run()
	})
})
