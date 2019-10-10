package main

import (
	"bufio"
	"bytes"
	"github.com/Masterminds/semver"
	"github.com/pkg/errors"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func awsSdkPath(version string) string {
	if version == "" {
		version = latestAwsSdkVersion()
	}

	return gopath() + "/pkg/mod/github.com/aws/aws-sdk-go@v" + version
}

func gopath() string {
	cmd := exec.Command("go", "env")
	output, err := cmd.CombinedOutput()
	if err != nil {
		panic(errors.WithStack(err))
	}

	prefix := "GOPATH="
	scan := bufio.NewScanner(bytes.NewBuffer(output))

	for scan.Scan() {
		line := scan.Text()
		if strings.HasPrefix(line, prefix) {
			rest := strings.TrimPrefix(line, prefix)
			val, err := strconv.Unquote(rest)
			if err != nil {
				panic(errors.WithStack(err))
			}
			return val
		}
	}

	panic(errors.New("gopath not found"))
}

func latestAwsSdkVersion() string {
	base := gopath()
	paths, err := filepath.Glob(base + "/pkg/mod/github.com/aws/aws-sdk-go@v*")
	if err != nil {
		panic(errors.WithStack(err))
	}

	versions := semver.Collection{}
	for _, path := range paths {
		bits := strings.Split(path, "@")
		version := bits[len(bits)-1]
		v, err := semver.NewVersion(version)
		if err != nil {
			panic(errors.WithStack(err))
		}
		versions = append(versions, v)
	}

	sort.Sort(versions)
	latest := versions[len(versions)-1]
	return latest.String()
}
