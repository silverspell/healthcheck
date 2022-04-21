package healthcheckmodule

import (
	"fmt"
	"os"
	"strings"

	"github.com/silverspell/commonlibs"
)

func DiscoverEnv(prefix string) []string {
	discoveredEnvs := commonlibs.Filter(os.Environ(), func(item string) bool { return strings.HasPrefix(item, prefix) })
	discoveredEnvs = commonlibs.Map(discoveredEnvs, func(item string) string { return strings.ReplaceAll(item, prefix, "") })
	discoveredEnvs = commonlibs.Map(discoveredEnvs, func(item string) string { return strings.Split(item, "=")[0] })
	fmt.Printf("Discovered envs %+v\n", discoveredEnvs)
	return discoveredEnvs
}

func DoHealthCheck(types []string) (string, error) {
	var err error
	for _, val := range types {
		fmt.Println(val)
		switch val {
		case "CASSANDRA":
			fmt.Println("Trying cassandra")
			cassandra := new(CassandraConnection)
			err = cassandra.Connect()
		case "REDIS":
			fmt.Println("Trying redis")
			redis := new(RedisConnection)
			err = redis.Connect()
		default:
			fmt.Println("unknown")
			err = nil
		}
	}

	if err != nil {
		return err.Error(), err
	}
	return "", nil
}

func Exit(code int) {
	os.Exit(code)
}
