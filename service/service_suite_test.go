package service_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCfLogstashSmokeTests(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CfLogstashSmokeTests Suite")
}
