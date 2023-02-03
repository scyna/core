package scyna_utils

import (
	"fmt"
	"regexp"
	"strings"
)

var pathrgxp = regexp.MustCompile(`:[A-z,0-9,$,-,_,.,+,!,*,',(,),\\,]{1,}`)

func PublishURL(urlPath string) string {
	ret := strings.Replace(urlPath, "/", ".", -1)
	ret = fmt.Sprintf("API%s", ret)
	return ret
}

func SubscriberURL(urlPath string) string {
	subURL := pathrgxp.ReplaceAllString(urlPath, "*")
	subURL = strings.Replace(subURL, "/", ".", -1)
	subURL = fmt.Sprintf("API%s", subURL)
	return subURL
}
