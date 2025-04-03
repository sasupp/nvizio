package url

import "strings"

type linkBuilderConfig struct {
	urls        []string
	queryParams map[string]string
}

type LinkOption func(*linkBuilderConfig)

func LinkPath(urlPath string) LinkOption {
	return func(lb *linkBuilderConfig) {
		lb.urls = append(lb.urls, urlPath)
	}
}

func LinkParam(queryParam string, value string) LinkOption {
	return func(lb *linkBuilderConfig) {
		lb.queryParams[queryParam] = value
	}
}

func BuildLink(options ...LinkOption) string {
	builder := linkBuilderConfig{urls: []string{}, queryParams: make(map[string]string)}
	for _, option := range options {
		option(&builder)
	}
	var sb strings.Builder
	for _, url := range builder.urls {
		sb.WriteString(url)
	}
	sb.WriteString("?")
	for p, v := range builder.queryParams {
		sb.WriteString(p)
		sb.WriteString("=")
		sb.WriteString(v)
	}
	return sb.String()
}
