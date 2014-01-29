package fake_backend

import (
	"testing"

	"github.com/onsi-experimental/garden/backend"
)

func FunctionTakingBackend(backend.Backend) {

}

func FunctionTakingContainer(backend.Container) {

}

func TestFakeBackendConformsToInterface(*testing.T) {
	FunctionTakingBackend(&FakeBackend{})
}

func TestFakeContainerConformsToInterface(*testing.T) {
	FunctionTakingContainer(&FakeContainer{})
}
