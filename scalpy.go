package scalpy

import (
	"github.com/octokit/go-octokit/octokit"
	"log"
	"net/url"
	"regexp"
	"strings"
)

type Hosting int

const (
	Github Hosting = iota
)

type Scalp struct {
	OriginalUrl string
	Url         *url.URL

	HostingService Hosting

	Project string
	Repo    string
	IssueId string
}

func (s *Scalp) GithubIssue() (issue *octokit.Issue) {
	client := octokit.NewClient(nil)
	url, _ := octokit.RepoIssuesURL.Expand(octokit.M{
		"owner":  s.Project,
		"repo":   s.Repo,
		"number": s.IssueId,
	})

	issue, _ = client.Issues(url).One()

	log.Printf("url %s %v", s.Url, issue)

	return
}

func ScalpUrl(url_str string) (scalp *Scalp) {
	scalp = &Scalp{
		OriginalUrl: url_str,
	}

	var err error
	scalp.Url, err = url.Parse(url_str)

	if err != nil {
		log.Printf("Can't parse url %v\n", err)
		return nil
	}

	if strings.Contains(scalp.Url.Host, "github.com") {
		scalp.HostingService = Github
	} else {
		log.Printf("Invalid service: %s\n", scalp.Url.Host)
		return nil
	}

	re := regexp.MustCompile("^/(.*)/(.*)/issues/([0-9]+)")
	matches := re.FindStringSubmatch(scalp.Url.Path)

	if matches == nil {
		log.Printf("Invalid path %v\n", matches)
		return nil
	} else {
		scalp.Project = matches[1]
		scalp.Repo = matches[2]
		scalp.IssueId = matches[3]
	}

	return scalp
}
