package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/solo-io/solo-kit/pkg/utils/contextutils"
	"github.com/solo-io/solo-kit/projects/vcs/pkg/file"
	"github.com/solo-io/solo-kit/projects/vcs/pkg/syncer"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("%v", err)
	}
}

func run() error {
	action := flag.String("a", "k2f", "action to take")
	flag.Parse()
	printUsage()
	fmt.Printf("running, action: %v\n", *action)
	ctx := contextutils.WithLogger(context.Background(), "vcs")
	dc, err := file.NewDualClient("kube")
	if err != nil {
		return err
	}

	syncer.ApplyVcsToDeployment(ctx, dc)

	// switch *action {
	// case "k2f":
	// 	file.GenerateFilesystem(ctx, "gloo-system", dc)
	// case "f2k":
	// 	file.UpdateKube(ctx, "gloo-system", dc)
	// default:
	// 	fmt.Printf("Action not recognized: %v\n", *action)
	// }
	return nil
}

func printUsage() {
	fmt.Printf(`Usage of this demo script
go run main -a k2f - writes kubernetes data to a filesystem
go run main -a f2k - writes filesystem data to kubernetes
----------------------------
`)
}
