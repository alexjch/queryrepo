package args

import (
	"flag"
	"log"
	"net/url"
	"strings"
)

type Args struct {
	RepoUrl *url.URL
	Port    int
}

type URLValue struct {
	url *url.URL
}

func ParseArgs() *Args {
	repoUrlStr := flag.String("repoUrl", "", "Repo URL")
	servicePort := flag.Int("servicePort", 9001, "Service port")

	flag.Parse()

	if strings.Trim(*repoUrlStr, " ") == "" {
		log.Fatal("repoUrl can not be empty")
	}

	repoUrl, err := url.Parse(*repoUrlStr)
	if err != nil {
		log.Fatal(err)
	}

	return &Args{
		RepoUrl: repoUrl,
		Port:    *servicePort,
	}
}
