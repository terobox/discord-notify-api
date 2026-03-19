package util

import (
	"fmt"
	"strings"
)

// gpt edit v2
func BuildDiscordMessage(title, content, source string) string {
	var b strings.Builder

	title = strings.TrimSpace(title)
	content = strings.TrimSpace(content)
	source = strings.TrimSpace(source)

	// title
	if title != "" {
		b.WriteString(fmt.Sprintf("**%s**", title))
	}

	// content as blockquote
	if content != "" {
		if b.Len() > 0 {
			b.WriteString("\n\n")
		}

		lines := strings.Split(content, "\n")
		for _, line := range lines {
			line = strings.TrimRight(line, " \t")

			if strings.TrimSpace(line) == "" {
				b.WriteString(">\n")
				continue
			}

			b.WriteString("> ")
			b.WriteString(line)
			b.WriteString("\n")
		}
	}

	if source != "" {
		b.WriteString("\n-# `")
		b.WriteString(source)
		b.WriteString("` · discord-notify-api")
	} else {
		b.WriteString("\n-# discord-notify-api")
	}

	return strings.TrimRight(b.String(), "\n")
}
