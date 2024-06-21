package skyalert_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestSkyalertGo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SkyalertGo Suite")
}
