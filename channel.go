package slack_lib

import "github.com/nlopes/slack"

func GetChannel(api *slack.Client, ev *slack.MessageEvent) (fromType, name string, error) {
	// 	identify channel (public and private) or or group or DM

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
			
      name := info.Name
		} else if string(c) == "G" {
			fromType = "group"
			info, err := api.GetGroupInfo(ev.Channel)

      if err != nil {
        return "", "", err
      }

      name := info.Name
		} else if string(c) == "D" {
			if ev.Msg.SubType != "" {
				// SubType is not define user
			} else {
        fromType = "DM"
				info, err := api.GetUserInfo(ev.Msg.User)

        if err != nil {
          return "","", err 
        }

			  name := info.Name
			}
		} else {
      fromType = ""
      name = ""
		}
	}

	return fromType, name, nil
}
