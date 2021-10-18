package fetcher

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	_ "golang.org/x/net/html"
	"golang.org/x/net/html/charset"
	_ "golang.org/x/text"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

func Fetch(url string) ([]byte, error) {

	//resp, err := http.Get(url)
	//if err != nil {
	//	return nil, err
	//}
	newUrl := strings.Replace(url, "http://", "https://", 1)
	request, _ := http.NewRequest(http.MethodGet, newUrl, nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")

	cookie := "FSSBBIl1UgzbN7NO=5qbtGqqsLPmVd37otBOwTth3yVkHcGwcDzXMcvGWhzk1X1Fd3VMcAIV0Svn7fGa9LmNomMMmMBt47iI1qW6aTPq; sid=08778025-f037-4d3d-8eeb-fb9bc7e7b798; ec=cnkm23xF-1629867206527-1fa049a63b0731404074495; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1629867206,1629960243; _exid=nxtjFEhYW7GEqAtZJ%2BvzystEO%2F21q1QAxEjeAzkainjD4VWb1fxqiRl76MzAPkSFYR5bJZyRNBLcvyFYlCdW3Q%3D%3D; _efmdata=zkl6jKkh7G5OZPWum7vmqCO5Ex0f21kpy91aDqEFUoZXtd0ca4rab4yj9pvSsUS4S6ZQocwYSuEzmpWeb3gXbhNISlYjh%2BQl6tvkiYq4CCI%3D; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1629967265; FSSBBIl1UgzbN7NP=531wsNKlvV_qqqqm4GR5lKaOlFEbcFetOIq6gOs4RHZmfgoGPVSy5nQVEknGLgd3FhH8zShGQrDgqNDVxfoWwyGekSCe5Ma8p3fvwkE6EHKLilIp_iIm4vrdKaAtvI2xY5tb8584uoGiGGc.7fajK8M0lXT_5W_sRrSUx16AsoXRbHi5Hm6wGNE4GbKgn6O4ygO_5cbqamnl7Cq3StctPRiQ0w.B9xJqxbJMsKR1peBoEVJPZpT74hZre8lQOVRwg8AhM0.ynEXXSdvAdBH1aJL"

	request.Header.Add("cookie", cookie)
	resp, _ := http.DefaultClient.Do(request)

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error:status code:%d", resp.StatusCode)
	}

	//如果页面传来的不是utf8，我们需要转为utf8格式
	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())

	return ioutil.ReadAll(utf8Reader)

}
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Ftcher error:%v", err)
		return unicode.UTF8
	}

	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
