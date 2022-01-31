package tasks

import (
	"fmt"

	"github.com/beego/beego/v2/adapter/toolbox"
)

func init() {
	fmt.Println("task base init")
	toolbox.AddTask("test", toolbox.NewTask("test", "0/10 * * * * *", func() error {
		fmt.Println("tttt")
		return nil
	}))
}
