package rustaudit

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinuxRustDependencies(t *testing.T) {
	// Generate this with `DOCKER_BUILDKIT=1 docker build -f test/Dockerfile -o . .`
	r, err := os.Open("crate_with_features_bin")
	if err != nil {
		log.Fatal(err)
	}
	versionInfo, err := GetDependencyInfo(r)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, 2, len(versionInfo.Packages))
	assert.Equal(t, Package{Name: "crate_with_features", Version: "0.1.0", Source: "local", Kind: "runtime", Features: []string{"default", "library_crate"}, Dependencies: []uint{1}}, versionInfo.Packages[0])
}
