package slack_lib

import (
	"github.com/jinzhu/copier"
	"github.com/nlopes/slack"
)

func ConvertReadableName(api *slack.Client, ev *slack.MessageEvent) (slack.Msg, error) {
	var err error
	result := slack.Msg{}
	msg := ev.Msg

	copier.Copy(&result, &msg)

	rUser, err := api.GetUserInfo(msg.User)
	if err != nil {
		return slack.Msg{}, err
	}

	_, channelName, err := GetFromName(api, ev)
	if err != nil {
		return slack.Msg{}, err
	}

	rTeam, err := api.GetTeamInfo()
	if err != nil {
		return slack.Msg{}, err
	}

	result.User = rUser.Name
	result.Channel = channelName
	result.Team = rTeam.Name

	return result, nil
}
