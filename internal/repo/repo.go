package repo

import (
	"fmt"
	"net/url"
	"os/exec"
)

const command string = "/usr/bin/dnf"

const repoArgValue string = "--repofrompath=test,%s"

var args = []string{"", "repoquery", "rpm"}

func QueryRepo(repoUrl *url.URL, packageName string) (*string, error) {
	args[0] = fmt.Sprintf(repoArgValue, repoUrl.String())
	args := append(args, packageName)
	out, err := exec.Command(command, args...).Output()
	if err != nil {
		return nil, err
	}

	outStr := string(out)

	return &outStr, nil
}
