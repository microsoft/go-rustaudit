package rustaudit

import (
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"os/exec"
	"testing"
)

func TestLinuxRustDependencies(t *testing.T) {
	//Build a Rust binary with audit information
	cmd := exec.Command("docker", "build", "-f", "test/Dockerfile", "-o", ".", ".")
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

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
