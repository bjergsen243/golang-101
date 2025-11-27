package l19

import "strings"

type sms struct {
	id      string
	content string
	tags    []string
}

func tagMessages(messages []sms, tagger func(sms) []string) []sms {
	result := make([]sms, 0, len(messages))
	for _, msg := range messages {
		msg.tags = tagger(msg)
		result = append(result, msg)

	}
	return result
}

func tagger(msg sms) []string {
	tags := []string{}
	if strings.Contains(strings.ToLower(msg.content), "urgent") {
		tags = append(tags, "Urgent")
	}
	if strings.Contains(strings.ToLower(msg.content), "sale") || strings.Contains(strings.ToLower(msg.content), "promo") {
		tags = append(tags, "Promo")
	}
	return tags
}

func taggerWithoutStrings(msg sms) []string {
	tags := []string{}

	if containsIgnoreCase(msg.content, "urgent") {
		tags = append(tags, "Urgent")
	}
	if containsIgnoreCase(msg.content, "sale") || containsIgnoreCase(msg.content, "promo") {
		tags = append(tags, "Promo")
	}
	return tags
}

func containsIgnoreCase(content, substr string) bool {
	contentLen := len(content)
	substrLen := len(substr)

	if substrLen == 0 || substrLen > contentLen {
		return false
	}
	for i := 0; i <= contentLen-substrLen; i++ {
		match := true
		for j := 0; j < substrLen; j++ {
			c1 := content[i+j]
			c2 := substr[j]
			if c1 >= 'A' && c1 <= 'Z' {
				c1 += 'a' - 'A'
			}
			if c2 >= 'A' && c2 <= 'Z' {
				c2 += 'a' - 'A'
			}
			if c1 != c2 {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}
	return false
}
