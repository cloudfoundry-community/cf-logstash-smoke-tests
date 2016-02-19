package main

import (
	"fmt"
	"time"

	"github.com/starkandwayne/cf-logstash-smoke-tests/Godeps/_workspace/src/github.com/cloudfoundry-incubator/cf-test-helpers/cf"
	"github.com/starkandwayne/cf-logstash-smoke-tests/Godeps/_workspace/src/github.com/cloudfoundry-incubator/cf-test-helpers/runner"
	. "github.com/starkandwayne/cf-logstash-smoke-tests/Godeps/_workspace/src/github.com/onsi/ginkgo"
	. "github.com/starkandwayne/cf-logstash-smoke-tests/Godeps/_workspace/src/github.com/onsi/gomega"
	. "github.com/starkandwayne/cf-logstash-smoke-tests/Godeps/_workspace/src/github.com/onsi/gomega/gbytes"
	. "github.com/starkandwayne/cf-logstash-smoke-tests/Godeps/_workspace/src/github.com/onsi/gomega/gexec"
	"github.com/starkandwayne/cf-logstash-smoke-tests/Godeps/_workspace/src/github.com/pborman/uuid"
)

var _ = Describe("CfLogstashSmokeTests", func() {
	var timeout = time.Second * 1200
	var cfUser string
	var cfPass string
	var cfAPI string
	var testOrg string
	var testSpace string
	var testService string
	var testPlan string
	var appName string
	var domain string
	var appURI string
	appPath := "/Users/vannguyen/CloudFoundry/cf-env"

	assertAppIsRunning := func(appName string) {
		pingUri := appURI + "/some-error"
		fmt.Println("Checking that the app is responding at url: ", pingUri)
		Eventually(runner.Curl(pingUri, "-k"), timeout*3, time.Second*1).Should(Say(""))
		//Curl should return nothing which means it has no errors
	}

	BeforeEach(func() {
		cfUser = "admin"
		cfPass = "c989071b5e64a0a4361a"
		cfAPI = "https://api.run.vsphere.starkandwayne.com"
		testOrg = "test"
		testSpace = "test"
		testService = "test"
		testPlan = "test"
		domain = "apps.vsphere.starkandwayne.com"
	})

	Context("Credentials Tests", func() {
		It("CF Creds", func() {
			Expect("admin").To(Equal(cfUser))
			Expect("c989071b5e64a0a4361a").To(Equal(cfPass))
			Expect("https://api.run.vsphere.starkandwayne.com").To(Equal(cfAPI))
		})

		It("Test Creds", func() {
			Expect("test").To(Equal(testOrg))
			Expect("test").To(Equal(testSpace))
			Expect("test").To(Equal(testService))
			Expect("test").To(Equal(testPlan))
		})
	})

	Context("Example App Tests", func() {
		BeforeEach(func() {
			appName = uuid.New()
			appURI = "https://" + appName + "." + domain
			Eventually(cf.Cf("login", "-a", cfAPI, "-u", cfUser, "-p", cfPass, "--skip-ssl-validation"), timeout).Should(Exit(0))
			Eventually(cf.Cf("create-org", testOrg), timeout).Should(Exit(0))
			Eventually(cf.Cf("target", "-o", testOrg), timeout).Should(Exit(0))
			Eventually(cf.Cf("create-space", testSpace), timeout).Should(Exit(0))
			Eventually(cf.Cf("target", "-s", testSpace), timeout).Should(Exit(0))
			Eventually(cf.Cf("push", appName, "-m", "126M", "-p", appPath, "-no-start"), timeout).Should(Exit(0))
		})

		AfterEach(func() {
			Eventually(cf.Cf("delete", appName, "-f"), timeout).Should(Exit(0))
			Eventually(cf.Cf("delete-space", testSpace, "-f"), timeout).Should(Exit(0))
			Eventually(cf.Cf("delete-org", testOrg, "-f"), timeout).Should(Exit(0))
		})

		It("Pushing app and see if it running with no errors", func() {
			Eventually(cf.Cf("start", appName), 2*time.Minute).Should(Exit(0))
			assertAppIsRunning(appName)
		})

	})
})
