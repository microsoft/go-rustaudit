package rustaudit

import (
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestLinuxRustDependencies(t *testing.T) {
	r, err := os.Open("hello-auditable")
	if err != nil {
		log.Fatal(err)
	}
	versionInfo, err := GetDependencyInfo(r)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, 19, len(versionInfo.Packages))
	assert.Equal(t, Package{Name: "adler", Version: "1.0.2", Source: "registry", Kind: "build"}, versionInfo.Packages[0])
}
