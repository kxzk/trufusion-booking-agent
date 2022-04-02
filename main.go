package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strings"
)

// TODO: fix this, hacky AF
const (
	username = ""
	password = ""
	session  = "https://cart.mindbodyonline.com/sites/14486/session/new"
	login    = "https://cart.mindbodyonline.com/sites/14486/session/"
	checkout = "https://cart.mindbodyonline.com/sites/14486/cart/proceed_to_checkout"
	utf8     = "&#x2713"
)

var classPreference = map[string]int{
	"Monday":    1, // bootcamp
	"Tuesday":   1, // bootcamp
	"Wednesday": 0, // bootcamp
	"Friday":    0, // kettlebell
}

func httpClient() *http.Client {
	jar, _ := cookiejar.New(nil)
	client := &http.Client{Jar: jar}
	return client
}

func request(client *http.Client, method, url string, body io.Reader) *http.Response {
	req, _ := http.NewRequest(method, url, body)
	resp, _ := client.Do(req)
	return resp
}

func getAuthToken(body string) string {
	re := regexp.MustCompile(`authenticity_token\" value=\"(.*)"`)
	match := re.FindStringSubmatch(body)[1] // grab actual value

	// TODO: make test
	// if len(match) != 86 {
	// }

	return match
}

func getEncodedVals(user, pass, utf, auth string) string {
	vals := url.Values{}
	vals.Add("mb_client_session[username]", user)
	vals.Add("mb_client_session[password]", pass)
	vals.Add("utf8", utf)
	vals.Add("authenticity_token", auth)

	return vals.Encode()
}

func main() {
	client := httpClient()

	// start session and get `authenticity_token` from body
	sessReq := request(client, "GET", session, nil)
	defer sessReq.Body.Close()

	body, _ := ioutil.ReadAll(sessReq.Body)

	authToken := getAuthToken(string(body))
	encodedVals := getEncodedVals(username, password, utf8, authToken)

	// login to mindbodyonline with account
	loginReq := request(client, "POST", login, strings.NewReader(encodeVals))
	defer loginReq.Body.Close()

	nextWeekDate := getNextWeekDate()
	_, isoWeek := nextWeekDate.ISOWeek()
	dayOfWeek := nextWeekDate.Weekday().String()

	class := classSchedule[dayOfWeek][classPreference[dayOfWeek]]

	classDate := getFormattedDate(nextWeekDate, class.time)
	classID := class.id + getClassIDOffset(isoWeek)
	classURL := classURLFormat[class.name]

	classLink := getClassLink(classDate, classID, classURL)

	// adds desired class to our cart
	classLinkReq := request(client, "GET", classLink, nil)
	defer classLinkReq.Body.Close()

	// this request will proceed with checking out the class in our cart
	checkoutReq := request(client, "GET", checkout, nil)
	defer checkoutReq.Body.Close()
}
