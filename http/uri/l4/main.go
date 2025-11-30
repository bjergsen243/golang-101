package l4

import (
	"net/url"
)

func newParsedURL(urlString string) ParsedURL {
	parsedUrl, err := url.Parse(urlString)
	if err != nil {
		return ParsedURL{}
	}

	return ParsedURL{
		protocol: parsedUrl.Scheme,
		username: parsedUrl.User.Username(),
		password: func() string {
			password, _ := parsedUrl.User.Password()
			return password
		}(),
		hostname: parsedUrl.Hostname(),
		port:     parsedUrl.Port(),
		pathname: parsedUrl.Path,
		search: func() string {
			if len(parsedUrl.RawQuery) > 0 {
				return parsedUrl.RawQuery
			}
			return ""
		}(),
		hash: func() string {
			if len(parsedUrl.Fragment) > 0 {
				return parsedUrl.Fragment
			}
			return ""
		}(),
	}

}
