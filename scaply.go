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
	Url         url.URL

	HostingService Hosting

	Project string
	Repo    string
	IssueId string
}

func (s *Scalp) GithubIssue() *octokit.Issue {
	client := octokit.NewClient(nil)
	issue_service := client.Issues(&s.Url)
	issue, _ := issue_service.One()

	return issue
}

func ScalpUrl(url_str string) (scalp *Scalp) {
	scalp = &Scalp{}
	scalp.OriginalUrl = url_str

	url, err := url.Parse(url_str)
	if err != nil {
		log.Printf("Can't parse url %v\n", err)
		return nil
	}

	if strings.Contains(url.Host, "github.com") {
		scalp.HostingService = Github
	} else {
		log.Printf("Invalid service: %s\n", url.Host)
		return nil
	}

	re := regexp.MustCompile("^/(.*)/(.*)/issues/([0-9]+)")
	matches := re.FindStringSubmatch(url.Path)

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
