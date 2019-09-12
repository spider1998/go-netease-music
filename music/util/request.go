package util

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"math/rand"
	"music/api"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var client = &http.Client{
	Timeout: time.Second * 10,
}

var userAgentList = []string{
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36",
	"Mozilla/5.0 (iPhone; CPU iPhone OS 9_1 like Mac OS X) AppleWebKit/601.1.46 (KHTML, like Gecko) Version/9.0 Mobile/13B143 Safari/601.1",
	"Mozilla/5.0 (iPhone; CPU iPhone OS 9_1 like Mac OS X) AppleWebKit/601.1.46 (KHTML, like Gecko) Version/9.0 Mobile/13B143 Safari/601.1",
	"Mozilla/5.0 (Linux; Android 5.0; SM-G900P Build/LRX21T) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Mobile Safari/537.36",
	"Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Mobile Safari/537.36",
	"Mozilla/5.0 (Linux; Android 5.1.1; Nexus 6 Build/LYZ28E) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Mobile Safari/537.36",
	"Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_2 like Mac OS X) AppleWebKit/603.2.4 (KHTML, like Gecko) Mobile/14F89;GameHelper",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/603.2.4 (KHTML, like Gecko) Version/10.1.1 Safari/603.2.4",
	"Mozilla/5.0 (iPhone; CPU iPhone OS 10_0 like Mac OS X) AppleWebKit/602.1.38 (KHTML, like Gecko) Version/10.0 Mobile/14A300 Safari/602.1",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.12; rv:46.0) Gecko/20100101 Firefox/46.0",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:46.0) Gecko/20100101 Firefox/46.0",
	"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.0)",
	"Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.0; Trident/4.0)",
	"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0)",
	"Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; Win64; x64; Trident/6.0)",
	"Mozilla/5.0 (Windows NT 6.3; Win64, x64; Trident/7.0; rv:11.0) like Gecko",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36 Edge/13.10586",
	"Mozilla/5.0 (iPad; CPU OS 10_0 like Mac OS X) AppleWebKit/602.1.38 (KHTML, like Gecko) Version/10.0 Mobile/14A300 Safari/602.1",
}

func randomUserAgent() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return userAgentList[r.Intn(19)]
}

func do(req *http.Request, cookies []*http.Cookie) (*http.Response, error) {
	basecookie := GenerateBaseCookie()
	cookies = append(cookies, basecookie...)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.8,gl;q=0.6,zh-TW;q=0.4")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Referer", "http://music.163.com")
	req.Header.Set("Host", "music.163.com")
	req.Header.Set("Cookie", "appver=2.0.2")
	req.Header.Set("User-Agent", randomUserAgent())
	for _, cookie := range cookies {
		req.Header.Add("Cookie", cookie.String())
	}

	return client.Do(req)
}

/*func GetMethond(args map[string]string, urls string, cookies []*http.Cookie) (res interface{}, err error) {
	csrf := GetCookieValueByName(cookies, "__csrf")

	URL, err := url.Parse(urls)
	if err != nil {
		return
	}
	query := URL.Query()
	for key, val := range args {
		query.Add(key, val)
	}
	URL.RawQuery = query.Encode()
	req, err := http.NewRequest(http.MethodGet, URL.String(), nil)
	if err != nil {
		return
	}
	resp, err := do(req)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	if resp.StatusCode != 200 {
		logs.Error("response error", string(b))
		var result api.APIError
		err = json.Unmarshal(b, &result)
		if err != nil {
			err = errors.WithStack(err)
			return
		}
		err = result.WithStatus(resp.StatusCode)
		return
	}
	err = json.Unmarshal(b, &res)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}*/

func CloudRequest(URL string, params map[string]interface{}, cookies []*http.Cookie) (res []byte, respCookies []*http.Cookie, err error) {
	if params == nil {
		params = make(map[string]interface{})
	}
	params["csrf_token"] = GetCookieValueByName(cookies, "__csrf")

	b, err := json.Marshal(params)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	crypto := Crypto{}
	encText, encSecKey, err := crypto.Encrypt(string(b))
	if err != nil {
		return
	}

	form := url.Values{}
	form.Set("params", encText)
	form.Set("encSecKey", encSecKey)
	body := strings.NewReader(form.Encode())

	req, err := http.NewRequest(http.MethodPost, URL, body)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	resp, err := do(req, cookies)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	defer resp.Body.Close()

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	if resp.StatusCode != 200 {
		fmt.Println(resp.StatusCode)
		var result api.APIError
		err = json.Unmarshal(b, &result)
		if err != nil {
			err = errors.WithStack(err)
			return
		}
		err = result.WithStatus(resp.StatusCode)
		return
	}
	respCookies = resp.Cookies()
	/*err = json.Unmarshal(b, &res)
	if err != nil {
		err = errors.WithStack(err)
		return
	}*/
	res = b
	return
}
