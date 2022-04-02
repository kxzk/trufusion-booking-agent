package main

import (
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strings"
)

// TODO: fix this, hacky AF
const (
	USERNAME = ""
	PASSWORD = ""
)

var (
	session  = "https://cart.mindbodyonline.com/sites/14486/session/new"
	login    = "https://cart.mindbodyonline.com/sites/14486/session/"
	checkout = "https://cart.mindbodyonline.com/sites/14486/cart/proceed_to_checkout"
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

func main() {
	client := httpClient()

	// TODO: create function for all these requests + responses
	req, _ := http.NewRequest("GET", session, nil)
	req.Header.Add("Connection", "keep-alive")

	resp, _ := client.Do(req)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	strBody := string(body)

	// TODO: extract out into function
	// TODO: improve regex to be more exact
	//
	// only need `utf8` and `authenticity_token`
	re := regexp.MustCompile(`<input.*>`)

	matches := re.FindAllStringSubmatch(strBody, -1)

	splitMatches := strings.Split(matches[0][0], " ")

	var parts []string
	// parts[0] -> utf8
	// parts[1] -> authenticity_token

	for _, s := range splitMatches {
		if strings.HasPrefix(s, "value=") {
			s = strings.Replace(s, "value=", "", -1)
			s = strings.Replace(s, "\"", "", -1)
			parts = append(parts, s)
		}
	}

	// TODO: extract out into function
	user := url.Values{}
	user.Add("mb_client_session[username]", USERNAME)
	user.Add("mb_client_session[password]", PASSWORD)
	user.Add("utf8", parts[0])
	user.Add("authenticity_token", parts[1])

	// login to mindbodyonline with account
	req2, _ := http.NewRequest("POST", login, strings.NewReader(user.Encode()))
	req2.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp2, _ := client.Do(req2)
	defer resp2.Body.Close()

	nextWeekDate := getNextWeekDate()
	_, isoWeek := nextWeekDate.ISOWeek()
	dayOfWeek := nextWeekDate.Weekday().String()

	// TODO: add check to make sure only runs on Mon-Wed +  Fri

	class := classSchedule[dayOfWeek][classPreference[dayOfWeek]]

	classDate := getFormattedDate(nextWeekDate, class.time)
	classID := class.id + getClassIDOffset(isoWeek)
	classURL := classURLFormat[class.name]

	classLink := getClassLink(classDate, classID, classURL)

	// adds desired class to our cart
	req3, _ := http.NewRequest("GET", classLink, nil)

	resp3, _ := client.Do(req3)
	defer resp3.Body.Close()

	// this request will proceed with checking out the class in our cart
	req4, _ := http.NewRequest("GET", checkout, nil)

	resp4, _ := client.Do(req4)
	defer resp4.Body.Close()
}
