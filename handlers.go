package main

import (
	"fmt"
	"strings"
	"net/http"
	"encoding/json"
)

type GitlabPayloadRepo struct {
	Name string `json:"name"`
	Homepage string `json:"homepage"`
}

type GitlabPayloadCommitAuthor struct {
	Name string `json:"name"`
}

type GitlabPayloadCommit struct {
	Message string `json:"message"`
	Url string `json:"url"`
	Id string `json:"id"`
	Author GitlabPayloadCommitAuthor `json:"author"`
}

type GitlabPayload struct {
	Before string `json:"before"`
	After string `json:"after"`
	User string `json:"user_name"`
	Ref string `json:"ref"`
	Repository GitlabPayloadRepo `json:"repository"`
	Commits []GitlabPayloadCommit `json:"commits"`
}

func gitlabHandler(w http.ResponseWriter, r *http.Request) {

	var decodedPayload GitlabPayload
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&decodedPayload)

	if err != nil {
		return
	}

	if decodedPayload.Commits == nil {
		return
	}

	branchSlice := strings.Split(decodedPayload.Ref, "/")
	branch := branchSlice[len(branchSlice) - 1]

	slackCommits := []SlackCommitMessage{}

	for _, commit := range decodedPayload.Commits {
		var slackCommit SlackCommitMessage

		slackCommit.Author = commit.Author.Name
		slackCommit.Commit = commit.Id[:10]
		slackCommit.CommitUrl = commit.Url
		slackCommit.Message = commit.Message

		slackCommits = append(slackCommits, slackCommit)
	}

	var slackMessage SlackMessage

	slackMessage.Commits = slackCommits
	slackMessage.Author = decodedPayload.User
	slackMessage.ProjectName = decodedPayload.Repository.Name
	slackMessage.ProjectUrl = decodedPayload.Repository.Homepage
	slackMessage.ProjectBranch = branch
	slackMessage.CompareUrl = fmt.Sprintf("%s/compare/%s...%s", decodedPayload.Repository.Homepage, decodedPayload.Before, decodedPayload.After)
	slackMessage.BranchUrl = fmt.Sprintf("%s/commits/%s", decodedPayload.Repository.Homepage, branch)

	err = SendSlackMessage(slackMessage, slack)

	if err != nil {
		fmt.Println(err)
	}
}
