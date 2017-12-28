package slack_lib

import (
	"errors"

	"github.com/nlopes/slack"
)

func ConvertDisplayChannelName(api *slack.Client, ev *slack.MessageEvent) (fromType, name string, err error) {
	// identify channel or group (as known as private channel) or DM

	// Channel prefix : C
	// Group prefix : G
	// Direct message prefix : D
	for _, c := range ev.Channel {
		if string(c) == "C" {
			fromType = "channel"
			info, err := api.GetChannelInfo(ev.Channel)

			if err != nil {
				return "", "", err
			}

			name = info.Name

			return fromType, name, nil
		} else if string(c) == "G" {
			fromType = "group"
			info, err := api.GetGroupInfo(ev.Channel)

			if err != nil {
				return "", "", err
			}

			name = info.Name

			return fromType, name, nil
		} else if string(c) == "D" {
			if ev.Msg.SubType != "" {
				// SubType is not define user
			} else {
				fromType = "DM"
				info, err := api.GetUserInfo(ev.Msg.User)

				if err != nil {
					return "", "", err
				}

				name = info.Name

				return fromType, name, nil
			}
		} else {
			fromType = ""
			name = ""
		}
	}

	return "", "", errors.New("channel not found")
}

func ConvertDisplayName(api *slack.Client, id string) (string, error) {
	// user id to display name
	info, err := api.GetUserInfo(id)
	if err != nil {
		return "", err
	}

	return info.Name, nil
}
