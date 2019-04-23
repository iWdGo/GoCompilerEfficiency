// urlbuilder evaluates any difference between returning a url using the type or a string
package urlbuilder

import "net/url"

// getAppString returns a string with the url
func getAppString(s string) string {
	u := new(url.URL)
	u.Scheme = "https"
	u.Host = "testinghello-in-the-cloud.appspot.com"
	u.Path = s
	return u.String()
}

// getAppUrl returns a pointer to the url
func getAppUrl(s string) *url.URL {
	u := new(url.URL)
	u.Scheme = "https"
	u.Host = "testinghello-in-the-cloud.appspot.com"
	u.Path = s
	return u
}
