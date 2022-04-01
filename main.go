package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strings"
)

// TODO: reduce all the duplicative code
// TODO: maybe add .yaml file and read in for configuration
// TODO: better strucutre/organization
// TODO: better way to structure all http.NewRequest's
// TODO: better extraction of utf8 and auth_token
// TODO: add tests

// TODO: fix this, hacky AF
const (
	USERNAME = ""
	PASSWORD = ""
)

func main() {

	jar, _ := cookiejar.New(nil)
	c := &http.Client{Jar: jar}

	// TODO: extract out core domain name, more efficient with creating links
	req, _ := http.NewRequest("GET", "https://cart.mindbodyonline.com/sites/14486/session/new", nil)
	req.Header.Add("Connection", "keep-alive")

	resp, _ := c.Do(req)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	strBody := string(body)

	// TODO: improve this match to be more specific, reduce amount of work
	// I have to do afterwards
	re := regexp.MustCompile(`<input.*>`)

	matches := re.FindAllStringSubmatch(strBody, -1)

	splitMatches := strings.Split(matches[0][0], " ")

	var parts []string

	for _, s := range splitMatches {
		if strings.HasPrefix(s, "value=") {
			s = strings.Replace(s, "value=", "", -1)
			s = strings.Replace(s, "\"", "", -1)
			parts = append(parts, s)
		}
	}

	// parts[0] -> utf8
	// parts[1] -> authenticity_token

	user := url.Values{}
	user.Add("mb_client_session[username]", USERNAME)
	user.Add("mb_client_session[password]", PASSWORD)
	user.Add("utf8", parts[0])
	user.Add("authenticity_token", parts[1])

	signIn := "https://cart.mindbodyonline.com/sites/14486/session/"

	req2, _ := http.NewRequest("POST", signIn, strings.NewReader(user.Encode()))
	req2.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp2, _ := c.Do(req2)
	defer resp2.Body.Close()

	body2, _ := ioutil.ReadAll(resp2.Body)

	// TODO: figure out what to do with these since I don't need them
	_ = body2

	// TODO: figure out how to programatically generate these for each of Katrina's classes
	// theoretically, will  need to run Mon-Thurs and it will just be current_date() + 7 days
	// teacher  name and class will be the same
	//
	// need to figure out if ID is increasing, random or able to be pre-populated or going
	// to have to scrape to get it -> item_mbo_id
	req3, _ := http.NewRequest("GET", "https://cart.mindbodyonline.com/sites/14486/cart/add_booking?item%5Binfo%5D=Mon.+Apr++4%2C+2022++9%3A00+pm&item%5Bmbo_id%5D=114627&item%5Bmbo_location_id%5D=1&item%5Bname%5D=60+Min.+Hot+Power+Vinyasa+%28Level+1%2F2%29+&item%5Btype%5D=Class&source=schedule_v1&ga_client_id=undefined", nil)

	resp3, _ := c.Do(req3)
	defer resp3.Body.Close()

	body3, _ := ioutil.ReadAll(resp3.Body)

	fmt.Println(string(body3))

	req4, _ := http.NewRequest("GET", "https://cart.mindbodyonline.com/sites/14486/cart/proceed_to_checkout", nil)

	resp4, _ := c.Do(req4)
	defer resp4.Body.Close()

	body4, _ := ioutil.ReadAll(resp4.Body)

	fmt.Println(string(body4))

}
