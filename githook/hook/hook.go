package hook

import (
	"encoding/json"
	"fmt"

	"github.com/google/go-github/github"
	"github.com/juju/errors"
	"github.com/ngaut/log"
	"github.com/shenli/tibot/slack"
)

type Hook struct {
	slackClient *slack.Slack
}

func (h *Hook) SetSlackClient(s *slack.Slack) {
	h.slackClient = s
}

func (h *Hook) HandleEvent(event string, requestBody []byte) error {
	switch event {
	case "issues":
		return h.handleIssue(requestBody)
	case "issue_comment":
		return h.handleIssueComment(requestBody)
	default:
		fmt.Println("Unnown event: ", event)
	}
	return nil
}

func (h *Hook) handleIssue(req []byte) (err error) {
	issue := &github.IssuesEvent{}

	err = json.Unmarshal(req, issue)
	if err != nil {
		return errors.Trace(err)
	}
	if *issue.Action != "opened" {
		return
	}
	preText, text := buildIssueMsg(issue)
	fmt.Println("Parse issue: ", preText, text)
	if h == nil {
		fmt.Println(" h is nul")
	}
	err = h.slackClient.SendMsg(preText, text)
	return
}

func buildIssueMsg(event *github.IssuesEvent) (preText, text string) {
	issue := event.Issue
	preText = fmt.Sprintf("New Issue <%s|%s> by user: <%s|%s>", *issue.HTMLURL, *issue.Title, *issue.User.HTMLURL, *issue.User.Login)
	text = *issue.Body
	return
}

func (h *Hook) handleIssueComment(req []byte) (err error) {
	comment := &github.IssueCommentEvent{}
	err = json.Unmarshal(req, comment)
	if err != nil {
		return errors.Trace(err)
	}
	log.Info("Get Comment:")
	log.Infof("%s", *comment)
	if comment.Issue != nil && comment.Issue.PullRequestLinks != nil {
		return nil
	}
	preText, text := buildIssueCommentMsg(comment)
	err = h.slackClient.SendMsg(preText, text)
	return
}

func buildIssueCommentMsg(event *github.IssueCommentEvent) (preText, text string) {
	issue := event.Issue
	comment := event.Comment
	preText = fmt.Sprintf("New comment for <%s|%s> by user: <%s|%s>", *comment.HTMLURL, *issue.Title, *comment.User.HTMLURL, *comment.User.Login)
	text = *comment.Body
	return
}
