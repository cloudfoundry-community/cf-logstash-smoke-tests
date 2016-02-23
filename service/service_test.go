package service_test

import (
	"encoding/json"
	"os"
	"time"

	"github.com/pborman/uuid"

	"github.com/cloudfoundry-incubator/cf-test-helpers/cf"
	"github.com/cloudfoundry-incubator/cf-test-helpers/runner"
	"github.com/cloudfoundry-incubator/cf-test-helpers/services"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"
)

type logstashConfig struct {
	services.Config

	Space   string `json:"space"`
	Plan    string `json:"plan"`
	Service string `json:"service"`
}

func loadConfig() (testConfig logstashConfig) {
	path := os.Getenv("CONFIG_PATH")
	configFile, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&testConfig)
	if err != nil {
		panic(err)
	}

	return testConfig
}

var config = loadConfig()

var _ = Describe("CfLogstashSmokeTests", func() {
	var timeout = time.Second * 60
	var config = loadConfig()
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
	var serviceName = "test-service"
	appPath := "../assets/cf-env"

	assertAppIsRunning := func(appName string) {
		pingURI := appURI + "/some-error"
		//fmt.Println("Checking that the app is responding at url: ", pingUri)
		Eventually(runner.Curl(pingURI, "-k"), config.ScaledTimeout(timeout*3), time.Second*1).Should(Say(""))
		//Curl should return nothing which means it has no errors
	}

	BeforeSuite(func() {
		config.TimeoutScale = 3
		cfUser = config.AdminUser
		cfPass = config.AdminPassword
		cfAPI = config.ApiEndpoint
		testOrg = config.OrgName
		testSpace = config.Space
		testService = config.Service
		testPlan = config.Plan
		domain = config.AppsDomain
		pluginPath := os.Getenv("PLUGIN_PATH")
		Eventually(cf.Cf("login", "-a", cfAPI, "-u", cfUser, "-p", cfPass, "--skip-ssl-validation"), config.ScaledTimeout(timeout)).Should(Exit(0))
		Eventually(cf.Cf("create-org", testOrg), config.ScaledTimeout(timeout)).Should(Exit(0))
		Eventually(cf.Cf("target", "-o", testOrg), config.ScaledTimeout(timeout)).Should(Exit(0))
		Eventually(cf.Cf("create-space", testSpace), config.ScaledTimeout(timeout)).Should(Exit(0))
		Eventually(cf.Cf("target", "-s", testSpace), config.ScaledTimeout(timeout)).Should(Exit(0))
		Eventually(cf.Cf("install-plugin", pluginPath), config.ScaledTimeout(timeout)).Should(Exit(0))
	})

	AfterSuite(func() {
		Eventually(cf.Cf("delete-space", testSpace, "-f"), config.ScaledTimeout(timeout)).Should(Exit(0))
		Eventually(cf.Cf("delete-org", testOrg, "-f"), config.ScaledTimeout(timeout)).Should(Exit(0))
	})

	Context("Example App Tests", func() {
		BeforeEach(func() {
			appName = uuid.New()
			appURI = "https://" + appName + "." + domain
			Eventually(cf.Cf("push", appName, "-m", "126M", "-p", appPath, "-no-start"), config.ScaledTimeout(timeout)).Should(Exit(0))
		})

		AfterEach(func() {
			Eventually(cf.Cf("delete", appName, "-f"), config.ScaledTimeout(timeout)).Should(Exit(0))
		})

		It("Pushing app and see if it running with no errors", func() {
			Eventually(cf.Cf("create-service", config.Service, config.Plan, serviceName), config.ScaledTimeout(timeout)).Should(Exit(0))
			Eventually(cf.Cf("bind-service", appName, serviceName), config.ScaledTimeout(timeout)).Should(Exit(0))
			Eventually(cf.Cf("start", appName), config.ScaledTimeout(3*time.Minute)).Should(Exit(0))
			Eventually(cf.Cf("start", appName), config.ScaledTimeout(3*time.Minute)).Should(Exit(0))
			assertAppIsRunning(appName)
		})

	})
})
