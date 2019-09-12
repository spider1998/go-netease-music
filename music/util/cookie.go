package util

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"io/ioutil"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func GenerateBaseCookie() []*http.Cookie {
	randomStr := GenerateRandomString(`0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKMNOPQRSTUVWXYZ/+`, 176)
	timestamp := strconv.FormatInt(time.Now().UnixNano(), 10)

	jsessionid := randomStr + ":" + timestamp
	nuid := GenerateRandomString("0123456789abcdefghijklmnopqrstuvwxyz", 32)

	// cookieStr := `JSESSIONID-WYYY=` + jsessionid + `;_iuqxldmzr_=32;_ntes_nnid=` + nuid + "," + strconv.FormatInt(time.Now().UnixNano(), 10) + `;_ntes_nuid=` + nuid
	baseCookies := make([]*http.Cookie, 4)
	baseCookies[0] = &http.Cookie{Name: `JSESSIONID-WYYY`, Value: jsessionid}
	baseCookies[1] = &http.Cookie{Name: `_iuqxldmzr_`, Value: "32"}
	baseCookies[2] = &http.Cookie{Name: `_ntes_nnid`, Value: nuid + "," + strconv.FormatInt(time.Now().UnixNano(), 10)}
	baseCookies[3] = &http.Cookie{Name: `_ntes_nuid`, Value: nuid}

	return baseCookies
}

func SetupDefaultCookie() []*http.Cookie {
	cookies := make([]*http.Cookie, 4)
	cookies[0] = &http.Cookie{Name: "appver", Value: "1.5.9"}
	cookies[1] = &http.Cookie{Name: "os", Value: "osx"}
	cookies[2] = &http.Cookie{Name: "channel", Value: "netease"}
	cookies[3] = &http.Cookie{Name: "osver", Value: "%e7%89%88%e6%9c%ac+10.13.2%ef%bc%88%e7%89%88%e5%8f%b7+17C88%ef%bc%89"}
	return cookies
}

func GetCookieValueByName(cookies []*http.Cookie, name string) string {
	for _, cookie := range cookies {
		if strings.EqualFold(cookie.Name, name) {
			return cookie.Value
		}
	}
	return ""
}

func GenerateRandomString(originStr string, length int) string {
	target := ""
	for i := 0; i < length; i++ {
		r := rand.New(rand.NewSource(time.Now().Unix() * int64(i)))
		pi := uint8(math.Floor(r.Float64() * float64(len(originStr))))

		target += string(originStr[pi])
	}

	return target
}

var (
	Name       string
	Value      string
	Path       string
	Domain     string
	Expires    string
	RawExpires string
	MaxAge     string
	Secure     string
	HttpOnly   string
	SameSite   string
	Raw        string
	Unparsed   string
)

type EnvCookies struct {
	Version    string `json:"vresion"`
	Name       string `json:"name"`
	Value      string `json:"value"`
	Path       string `json:"path"`
	Domain     string `json:"domain"`
	Expires    string `json:"expires"`
	RawExpires string `json:"raw_expires"`
	MaxAge     string `json:"max_age"`
	Secure     string `json:"secure"`
	HttpOnly   string `json:"http_only"`
	SameSite   string `json:"same_site"`
	Raw        string `json:"raw"`
	Unparsed   string `json:"unparsed"`
}

func WriteCookieToEnv(cookies []*http.Cookie) {
	for _, cookie := range cookies {
		Name += "&" + cookie.Name
		Value += "&" + cookie.Value
		Path += "&" + cookie.Path
		Domain += "&" + cookie.Domain
		Expires += "&" + cookie.Expires.String()
		RawExpires += "&" + cookie.RawExpires
		MaxAge += "&" + strconv.Itoa(cookie.MaxAge)
		Secure += "&" + strconv.FormatBool(cookie.Secure)
		HttpOnly += "&" + strconv.FormatBool(cookie.HttpOnly)
		SameSite += "&" + strconv.Itoa(int(cookie.SameSite))
		Raw += "&" + cookie.Raw
		Unparsed += "&" + strings.Join(cookie.Unparsed, ",")
	}
	name := ".env"
	content := "VERSION=1.0" +
		"\nNAME=" + Name +
		"\nVALUE=" + Value +
		"\nPATH=" + Path +
		"\nDOMAIN=" + Domain +
		"\nEXPIRES=" + Expires +
		"\nRAWEXPIRES=" + RawExpires +
		"\nMAXAGE=" + MaxAge +
		"\nSECURE=" + Secure +
		"\nHTTPONLY=" + HttpOnly +
		"\nSAMESITE=" + SameSite +
		"\nRAW=" + Raw +
		"\nUNPARSED=" + Unparsed
	writeWithIoutil(name, content)
}

func writeWithIoutil(name, content string) {
	data := []byte(content)
	if ioutil.WriteFile(name, data, 0644) == nil {
		fmt.Println("写入Cookies成功:", content)
	}
}

func GetCookiesForEnv() (cookies []*http.Cookie, err error) {
	godotenv.Load()
	var envCookies EnvCookies
	err = envconfig.Process("", &envCookies)
	if err != nil {
		return
	}
	fmt.Println(envCookies.Version)
	if len(envCookies.Name) == 0 {
		return
	}
	num := len(strings.Split(envCookies.Name, "&"))
	for i := 1; i < num; i++ {
		var cookie http.Cookie
		cookie.Name = strings.Split(envCookies.Name, "&")[i]
		cookie.Value = strings.Split(envCookies.Value, "&")[i]
		cookie.Path = strings.Split(envCookies.Path, "&")[i]
		cookie.Domain = strings.Split(envCookies.Domain, "&")[i]
		cookie.Expires = parseTime(strings.Split(envCookies.Expires, "&")[i]).UTC()
		cookie.RawExpires = strings.Split(envCookies.RawExpires, "&")[i]
		cookie.MaxAge = stringToInt(strings.Split(envCookies.MaxAge, "&")[i])
		cookie.Secure = stringToBool(strings.Split(envCookies.Secure, "&")[i])
		cookie.HttpOnly = stringToBool(strings.Split(envCookies.HttpOnly, "&")[i])
		cookie.SameSite = http.SameSite(stringToInt(strings.Split(envCookies.SameSite, "&")[i]))
		cookie.Raw = strings.Split(envCookies.Raw, "&")[i]
		cookie.Unparsed = strings.Split(strings.Split(envCookies.Unparsed, "&")[i], ",")
		cookies = append(cookies, &cookie)
	}
	return
}

func parseTime(timeString string) (t time.Time) {
	parsedTime, err := time.Parse("2006-01-02 15:04:05", timeString[:19])
	if err != nil {
		panic(err)
		return
	}
	t = time.Time(parsedTime)
	return
}

func stringToInt(str string) (num int) {
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
		return
	}
	return
}

func stringToBool(str string) (b bool) {
	b, err := strconv.ParseBool(str)
	if err != nil {
		panic(err)
		return
	}
	return
}
