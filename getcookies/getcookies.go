package getcookies

import (
	"fmt"
	"net/http"
	"regexp"
)

// curl -iL 'https://www.dmv.ca.gov/wasapp/foa/clear.do?goTo=officeVisit&localeName=en'
// 302 to
// https://www.dmv.ca.gov/wasapp/foa/officeVisit.do
// find and keep the last:
// Set-Cookie: AMWEBJCT!%2Fwasapp!JSESSIONID=0000oD6oZ3U_x3nodWfysT4CYOZ:18u4cegug; Path=/; Secure; HttpOnly

const startURL = "https://www.dmv.ca.gov/wasapp/foa/clear.do?goTo=officeVisit&localeName=en"

func getCookies(url string) ([]string, error) {
	resp, err := http.Get(url)
	res := []string{}
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	for k, v := range resp.Header {
		if k == "Set-Cookie" {
			res = append(res, v...)
		}
	}
	return res, nil
}

//GetJSESSIONID will get the latest JSESSIONID
func GetJSESSIONID() (string, error) {
	cks, err := getCookies(startURL)
	if err != nil {
		return "", err
	}

	r, _ := regexp.Compile("(JSESSIONID=(.*?);)")

	for _, s := range cks {
		fmt.Println(s)
		match := r.FindStringSubmatch(s)
		if len(match) < 3 {
			return "", err
		}
		return match[2], nil

	}

	return "", nil
}
