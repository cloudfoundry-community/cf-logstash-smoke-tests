package main

import (
	. "github.com/starkandwayne/cf-logstash-smoke-tests/Godeps/_workspace/src/github.com/onsi/ginkgo"
	. "github.com/starkandwayne/cf-logstash-smoke-tests/Godeps/_workspace/src/github.com/onsi/gomega"
	"testing"
)

func TestCfLogstashSmokeTests(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CfLogstashSmokeTests Suite")
}
