package workpool_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestWorkpool(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Workpool Suite")
}
