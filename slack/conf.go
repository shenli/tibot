package slack

type Config struct {
	Title string
	Slack slackInfo
}

type slackInfo struct {
	Token string `toml:"token"`
}
