package slack

import (
	"github.com/BurntSushi/toml"
	"github.com/juju/errors"
	"github.com/ngaut/log"
	"github.com/nlopes/slack"
)

type Slack struct {
	conf *Config
	api  *slack.Client
}

func NewSlack(config string) (*Slack, error) {

	s := &Slack{}
	s.conf = &Config{}
	if _, err := toml.DecodeFile(config, s.conf); err != nil {
		log.Errorf("Read slack config file meet error: %v", err)
		return nil, errors.Trace(err)
	}
	s.init()
	return s, nil
}

func (s *Slack) init() {
	if s.api == nil {
		s.api = slack.New(s.conf.Slack.Token)
		log.Infof("New slack with token %s", s.conf.Slack.Token)
	}
}

func (s *Slack) SendMsg(preText, text string) error {

	log.Infof("SenMsg")

	params := slack.PostMessageParameters{}
	attachment := slack.Attachment{
		Pretext: preText,
		Text:    text,
		// Uncomment the following part to send a field too
		/*
			Fields: []slack.AttachmentField{
				slack.AttachmentField{
					Title: "a",
					Value: "no",
				},
			},
		*/
	}
	params.Attachments = []slack.Attachment{attachment}

	notify := "@" + s.conf.Slack.At
	toChan := s.conf.Slack.TargetChan
	log.Infof("SenMsg1")
	channelID, timestamp, err := s.api.PostMessage(s.conf.Slack.TargetChan, notify, params)
	if err != nil {
		log.Errorf("Send message to %s failed with error: %v", toChan, err)
		return errors.Trace(err)
	}
	log.Debug("Send message to %s succ ChanID: %s, Time: %s", toChan, channelID, timestamp)
	return nil
}
