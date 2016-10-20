// Copyright(c) 2016 Ethan Zhuang <zhuangwj@gmail.com>.

package goraph

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestProxy(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Goraph Suite")
}

var _ = Describe("Test initialization", func() {
	Context("Register suite setup and teardown function", func() {
		BeforeSuite(func() {
		})

		AfterSuite(func() {
		})
	})
})

type myVertex struct {
	id     Id
	outTo  map[Id]float64
	inFrom map[Id]float64
}

func (vertex *myVertex) Id() Id {
	return vertex.id
}

func (vertex *myVertex) Out() map[Id]float64 {
	return vertex.outTo
}

func (vertex *myVertex) In() map[Id]float64 {
	return vertex.inFrom
}
