package telegram

const MarkdownTypeMessage = "markdown"
const HtmlTypeMessage = "html"

var typeMessageAllowed = [2]string{MarkdownTypeMessage, HtmlTypeMessage}

func (c *Client) getTypeMessage(typeMessage string) string {
	tm := c.defaultTypeMessage
	for _, a := range typeMessageAllowed {
		if a == typeMessage {
			tm = typeMessage
			break
		}
	}

	return tm
}
