package util

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"math"
	"math/rand"
	"music/api"
	"net/http"
	"net/url"
	"strconv"
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

func do(req *http.Request, cookies []*http.Cookie, api string, header map[string]interface{}) (*http.Response, error) {
	basecookie := GenerateBaseCookie()
	cookies = append(cookies, basecookie...)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.8,gl;q=0.6,zh-TW;q=0.4")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Referer", "http://music.163.com")
	req.Header.Set("Host", "music.163.com")
	req.Header.Set("Cookie", "appver=2.0.2")
	if api == "linux" {
		req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.90 Safari/537.36")
	} else {
		req.Header.Set("User-Agent", randomUserAgent())
	}
	for _, cookie := range cookies {
		req.Header.Add("Cookie", cookie.String())
	}
	//req.Header.Set("Cookie", "osver=undefined; deviceId=undefined; appver=6.1.1; versioncode=140; mobilename=undefined; buildver="+header["buildver"].(string)+"; resolution=1920x1080; __csrf=; os=android; channel=undefined; requestId="+header["requestId"].(string))
	return client.Do(req)
}

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
	resp, err := do(req, cookies, "weapi", nil)
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

func LAPICloudRequest(URL string, params map[string]interface{}, cookies []*http.Cookie) (res []byte, respCookies []*http.Cookie, err error) {
	if params == nil {
		params = make(map[string]interface{})
	}
	type LinuxApi struct {
		Method string                 `json:"method"`
		URL    string                 `json:"url"`
		Params map[string]interface{} `json:"params"`
	}
	var ob LinuxApi
	ob.Method = "POST"
	ob.Params = params
	ob.URL = strings.Replace(URL, "weapi", "api", -1)
	b, err := json.Marshal(ob)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	encText := AesEncryptECB(string(b), "lapi")
	form := url.Values{}
	form.Set("eparams", encText)
	body := strings.NewReader(form.Encode())
	lurl := `https://music.163.com/api/linux/forward`

	req, err := http.NewRequest(http.MethodPost, lurl, body)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	resp, err := do(req, cookies, "linux", nil)
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

func EAPICloudRequest(URL string, params map[string]interface{}, cookies []*http.Cookie, dURL string) (res []byte, respCookies []*http.Cookie, err error) {
	if params == nil {
		params = make(map[string]interface{})
	}
	cookieMap := map[string]interface{}{}
	for _, cookie := range cookies {
		cookieMap[cookie.Name] = cookie.Value
	}
	rand.Seed(time.Now().Unix())

	requestID := strconv.FormatFloat(math.Floor(float64(rand.Intn(1000))), 'f', -1, 64)
	if len(requestID) == 1 {
		requestID = "000" + requestID
	} else if len(requestID) == 2 {
		requestID = "00" + requestID
	} else if len(requestID) == 3 {
		requestID = "0" + requestID
	}
	header := map[string]interface{}{
		"osver":       cookieMap["osver"],
		"deviceId":    cookieMap["deviceId"],
		"appver":      `6.1.1`,
		"versioncode": "140",
		"mobilename":  cookieMap["mobilename"],
		"buildver":    strconv.FormatInt((time.Now().UnixNano() / 1e6), 10)[:10],
		"resolution":  "1920x1080",
		"__csrf":      cookieMap["__csrf"],
		"os":          "android",
		"channel":     cookieMap["channel"],
		"requestId":   strconv.FormatInt((time.Now().UnixNano()/1e6), 10) + `_` + requestID,
	}
	if cookieMap["MUSIC_U"] != nil {
		header["MUSIC_U"] = cookieMap["MUSIC_U"]
	} else if cookieMap["MUSIC_A"] != nil {
		header["MUSIC_A"] = cookieMap["MUSIC_A"]
	}
	params["header"] = header
	b, err := json.Marshal(params)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	lurl := strings.Replace(URL, "weapi", "eapi", -1)

	encText := EAPIAesEncryptECB(string(b), dURL)
	form := url.Values{}
	form.Set("params", encText)
	body := strings.NewReader(form.Encode())
	req, err := http.NewRequest(http.MethodPost, lurl, body)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	resp, err := do(req, cookies, "eapi", header)
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

//AB786F8DF53844553FC327689D15CB12D5736D1C59E786E1C198C85D34D798CFA64B970985C975A90515F3B94
// 3CDB31E7987F74C5A518DF48750D5920011F63CB6BEB7DA24EAB342F5717ACD61B6BAE091E0742D40A827412
// 8C085BF999F92974F7DBDB8A415C46451770B82C26C917E2DCFCD0FE3DC246A97FA29CEF9F32A45FA6E22F2FD
// 1D73DD95A24D9EF56138F0AFBEB788BAABF569873D35D203111923693B7F723D1FD4057F24CD7BDB55E38C117
// B7BECA13AF812FFCB6AB615B6360936F70BB019F9F72B44F7815238182B1337D3AF7E28FAB9A61F3AF675EDDA
// FDB5F5D87FA332040217F46FA35EBB7E2424FE102AE342B7ADCA8FA367140018186ACB355E9513AEC4F59C487
// AE929AB0ED8C39E770E77FD65AF6A77C21A742DA7217B4FEF0F7BC84E6B8766D124F8AACE8A727CF93452405DB
// 018D56B73B89B3C9EE5136D7C8330DDBE6C8FB2F092C05A5933C4745F701671A8FC6B560E6BD46C20

//AF047AB9ACC436C08101E8542E2D2378AB110DB4C7B0F4BC32B33E80214D1C93505F488051A7322EB8029486A03FC3
// 6BEAA551A43FE77B3B021D577B90D0D028DA77DDAB687FF8335AF34F1CEE72B2B74067152BC2FF7A379F1929BDE32
// 06FDE6C997E67CE2E4C5E7256D5C9B731C1149DC72EB230CD4018EE1DAF6744E4FD665BCC3C7CBDE88825429811321
// F51227A68972E12B0409D98A2D0A79394E7C48CB828842778F99FFE415A35EECFD5DECC62EABDE766EB72F92FE5DB4
// BBE9AD2255A0B4E447FBC29C4F1F2DD4AC9FFC0352204C67320CF51CD44E2EB3F5EFBBC09BBF21C9291E9A23CE4733
// 012EB3E55604BDE4BD27DC60C326AEED29F753FCEB390E4D0C3A8726504180BB763CC33BF7C6AA3B102FBE7296AB0DB9EA5C46AD12B
