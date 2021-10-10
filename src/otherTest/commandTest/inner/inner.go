package inner

import (
	"fmt"
	"github.com/urfave/cli/v2"
)


var SubCommand = &cli.Command{
	Name:        "go",
	Description: fmt.Sprintf("%s","go参数用于输出"),
	Usage: "print the first",
    Subcommands: []*cli.Command{
		StudyCmd,
	},
	//Flags: []cli.Flag{
	//	&cli.StringFlag{
	//		Name:  "say",
	//		Usage: "",
	//	},
	//},
}

var StudyCmd = &cli.Command{
	Name:        "study",
	Usage: "show study which",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "say",
			Usage: "say hello",
			Value: "hello",
		},
	},
	Action: func(cctx *cli.Context) error {
		fmt.Println("hello")
		return nil
	},

}
