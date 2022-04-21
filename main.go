package healthcheckmodule

import (
	"fmt"
)

func example() {
	envs := DiscoverEnv("MRT_")
	fmt.Printf("Discovered: %+v\n", envs)

	reason, err := DoHealthCheck(envs)
	if err != nil {
		fmt.Printf("%s\n", reason)
		Exit(127)
	}
	fmt.Println("Exiting normally")
	Exit(0)
}
