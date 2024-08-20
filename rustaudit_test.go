package rustaudit

import (
	"bytes"
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
	assert.Equal(t, Package{Name: "crate_with_features", Version: "0.1.0", Source: "local", Kind: "runtime", Dependencies: []uint{1}, Root: true}, versionInfo.Packages[0])
	assert.Equal(t, false, versionInfo.Packages[1].Root)
}

func TestWasmRustDependencies(t *testing.T) {
	// Generate this with `DOCKER_BUILDKIT=1 docker build -f test/Dockerfile -o . .`
	r, err := os.Open("wasm_crate.wasm")
	if err != nil {
		log.Fatal(err)
	}
	versionInfo, err := GetDependencyInfo(r)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, 18, len(versionInfo.Packages))
	assert.Equal(t, Package{Name: "bumpalo", Version: "3.16.0", Source: "crates.io", Kind: "runtime", Dependencies: nil, Root: false}, versionInfo.Packages[0])
	assert.Equal(t, false, versionInfo.Packages[1].Root)
}

func FuzzWasm(f *testing.F) {
	// Use the test fixture as a seed
	data, err := os.ReadFile("wasm_crate.wasm")
	if err != nil {
		log.Fatal(err)
	}
	f.Add(data)
	f.Fuzz(func(t *testing.T, input []byte) {
		r := bytes.NewReader(input)
		w := wasmReader{r}

		_, err := w.ReadRustDepSection()
		if err != ErrNoRustDepInfo && err != ErrUnknownFileFormat && err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
	})
}
