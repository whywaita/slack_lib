package slack_lib

import (
	"github.com/pkg/errors"

	"github.com/nlopes/slack"
	"github.com/nlopes/slack/slackutilsx"
)

const (
	ErrMethodNotSupportedForChannelType = "method_not_supported_for_channel_type"
)

func ConvertDisplayChannelName(api *slack.Client, ev *slack.MessageEvent) (fromType, name string, err error) {
	// identify channel or group (as known as private channel) or DM

	channelType := slackutilsx.DetectChannelType(ev.Channel)
	switch channelType {
	case slackutilsx.CTypeChannel:
		fromType = "channel"
		info, err := api.GetChannelInfo(ev.Channel)
		if err != nil {
			if err.Error() == ErrMethodNotSupportedForChannelType {
				// This error occurred by the private channels only converted from the public channel.
				// So, this is private channel if this error.
				info, err := api.GetGroupInfo(ev.Channel)
				if err != nil {
					return "", "", err
				}
				return "group", info.Name, nil
			} else {
				return "", "", err
			}
		}

		return fromType, info.Name, nil

	case slackutilsx.CTypeGroup:
		fromType = "group"
		info, err := api.GetGroupInfo(ev.Channel)
		if err != nil {
			return "", "", err
		}

		return fromType, info.Name, nil

	case slackutilsx.CTypeDM:
		if ev.Msg.SubType != "" {
			// SubType is not define user
		} else {
			fromType = "DM"
			info, err := api.GetUserInfo(ev.Msg.User)
			if err != nil {
				return "", "", err
			}

			return fromType, info.Name, nil
		}
	default:
		fromType = ""
		name = ""
	}

	return "", "", errors.New("channel not found")
}

func ConvertDisplayUserName(api *slack.Client, ev *slack.MessageEvent, id string) (username, usertype string, err error) {
	// user id to display name

	if id != "" {
		// specific id (maybe user)
		info, err := api.GetUserInfo(id)
		if err != nil {
			return "", "", err
		}

		return info.Name, "user", nil
	}

	// return self id
	if ev.Msg.BotID == "B01" {
		// this is slackbot
		return "Slack bot", "bot", nil
	} else if ev.Msg.BotID != "" {
		// this is bot
		byInfo, err := api.GetBotInfo(ev.Msg.BotID)
		if err != nil {
			return "", "", err
		}

		return byInfo.Name, "bot", nil
	} else if ev.Msg.SubType != "" {
		// SubType is not define user
		return ev.Msg.SubType, "status", nil
	} else {
		// user
		byInfo, err := api.GetUserInfo(ev.Msg.User)
		if err != nil {
			return "", "", err
		}

		return byInfo.Name, "user", nil
	}

	return "", "", errors.New("user not found")
}