package hub

import (
	"fmt"
	"net/url"
	"strings"
)

var AllowedDomains = []string{
	"channelnewsasia.com",
	"lemonde.fr",
	"techcrunch.com",
	"theverge.com",
	"openai.com",
	"finnhub.io",
	"coingecko.com",
	"opensky-network.org",
	"aisstream.io",
	"wigle.net",
	"unwiredlabs.com",
	"shodan.io",
	"opencellid.org",
}

func IsURLAllowed(targetURL string) (bool, error) {
	parsed, err := url.Parse(targetURL)
	if err != nil {
		return false, err
	}

	hostname := parsed.Hostname()
	if hostname == "" {
		return false, fmt.Errorf("invalid hostname")
	}

	for _, allowed := range AllowedDomains {
		if hostname == allowed || strings.HasSuffix(hostname, "."+allowed) {
			return true, nil
		}
	}

	return false, nil
}
