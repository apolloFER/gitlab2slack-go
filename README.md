## Gitlab2Slack

Standalone Golang webserver which receives webhooks from Gitlab and forwards them to more than one channel in Slack.
Current Slack service integration in Gitlab specifies only one channel.

Currently supports only commit webhooks.

### Usage

After building with ```go build``` run it as follows:

```gitlab2slack -c #channel -c @user -d domain -t token -l 0.0.0.0:8080```

#### Parameters:

 - ```-c```, ```--channel``` - Slack channel(s) (one or more) to send Gitlab messages to
 - ```-d```, ```--domain``` - Slack domain
 - ```-t```, ```--token``` - Slack incoming webhook token
 - ```-l```, ```--listen``` - IP:port pair on which to listen for requests

Incoming Gitlab webook requests should arrive on ```http://your-address:8080/gitlab```. Add a webhook with that URL for a project on your Gitlab.

### Docker

Buil the docker image as follows:

```docker build -t gitlab2slack .```

Then run it with the following options:

```docker run -p 8080:5000 gitlab2slack -c #channel -c @user -d domain -t token```
