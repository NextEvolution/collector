package facebookripper_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestFacebookripper(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Facebookripper Suite")
}
