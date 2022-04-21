package healthcheckmodule

import (
	"fmt"
	"os"
	"strings"

	"github.com/silverspell/commonlibs"
)

func DiscoverEnv() []string {
	discoveredEnvs := commonlibs.Filter(os.Environ(), func(item string) bool { return strings.HasPrefix(item, "MRT_") })
	fmt.Printf("Discovered envs %+v\n", discoveredEnvs)
	return discoveredEnvs
}
