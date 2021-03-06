package mideploy

import (
	"fmt"
	"strings"
)

func formatSlackMessage(author string) (*Message, error) {
	environmentTitle := strings.ToUpper(environment)
	attach := attachment{Color: "#3367d6", Text: fmt.Sprintf("Deploying *%s* to %s :among-us-sabotage:", branchName, environmentTitle)}

	message := &Message{
		ResponseType: "in_channel",
		Text:         fmt.Sprintf("%s is using *%s*", author, environmentTitle),
		Attachments:  []attachment{attach},
	}
	return message, nil
}
