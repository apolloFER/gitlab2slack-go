package main

import (
	"text/template"
	"bytes"
	"net/http"
	"encoding/json"
)

type TemplateExec struct {
	Message SlackMessage
}

type SlackCommitMessage struct {
	Author string
	Message string
	Commit string
	CommitUrl string
}

type SlackMessage struct {
	ProjectName string
	ProjectUrl string
	ProjectBranch string
	Author string
	Commits []SlackCommitMessage
	CompareUrl string
	BranchUrl string
}

type Slack struct {
	HookUrl string
	Channels []string
}

type SlackPayload struct {
	Channel string `json:"channel"`
	Text string `json:"text"`
}

const messageTemplate = `
{{.Message.Author}} pushed to branch <{{.Message.BranchUrl}}|{{.Message.ProjectBranch}}> of <{{.Message.ProjectUrl}}|{{.Message.ProjectName}}> (<{{.Message.CompareUrl}}|Compare changes>)
{{range $commit := .Message.Commits}}
> <{{$commit.CommitUrl}}|{{$commit.Commit}}>: {{$commit.Message}} - {{$commit.Author}}{{end}}
`


func SendSlackMessage(message SlackMessage, slack Slack) error {

	t := template.Must(template.New("message").Parse(messageTemplate))
	var buf bytes.Buffer
	if err := t.Execute(&buf, TemplateExec{Message: message}); err != nil {
		return err
	}

	for _, channel := range slack.Channels {
		payload := new(SlackPayload)
		payload.Channel = channel
		payload.Text = buf.String()

		requestData, err := json.Marshal(payload)

		if err != nil {
			return err
		}

		toSend := bytes.NewBuffer(requestData)
		_, err = http.Post(slack.HookUrl, "application/json", toSend)

		if err != nil {
			return err
		}

	}

	return nil
}
