package main

import (
	"fmt"
	"os"

	"github.com/paketo-community/pip/pip"
	"github.com/paketo-community/pip/python_packages"

	"github.com/cloudfoundry/libcfbuildpack/build"
	"github.com/cloudfoundry/libcfbuildpack/buildpackplan"
)

func main() {
	context, err := build.DefaultBuild()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to create default build context: %s", err)
		os.Exit(100)
	}

	code, err := runBuild(context)
	if err != nil {
		context.Logger.Info(err.Error())
	}

	os.Exit(code)
}

func runBuild(context build.Build) (int, error) {
	context.Logger.Title(context.Buildpack)

	pipPackageManager := pip.PIP{Logger: context.Logger}

	packagesContributor, willContribute, err := python_packages.NewContributor(context, pipPackageManager)
	if err != nil {
		return context.Failure(102), err
	}

	if willContribute {
		if err := packagesContributor.Contribute(); err != nil {
			return context.Failure(103), err
		}
	}

	return context.Success(buildpackplan.Plan{})
}
