package report

import (
	"fmt"
	"regexp"
	"testing"
)

func TestGetFileCovered(t *testing.T) {
	rePrefix := regexp.MustCompile("<span class=\"fc\" id=\"L-?[0-9]\\d*\">")
	reSuffix := regexp.MustCompile("</span>")
	line := "<span class=\"fc\" id=\"L31\">public class LogGrpcInterceptor implements ServerInterceptor {</span>"
	line = rePrefix.ReplaceAllString(line,"")
	line = reSuffix.ReplaceAllString(line,"")
	fmt.Print(line)
}
