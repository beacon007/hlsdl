package consts

var DefaultHeaders = map[string]string{}

func init() {

	DefaultHeaders["authority"] = "cdn152.akamai-content-network.com"
	DefaultHeaders["accept"] = "*/*"
	DefaultHeaders["accept-language"] = "zh-CN,zh;q=0.9"
	DefaultHeaders["origin"] = "https://missav.com"
	DefaultHeaders["sec-ch-ua"] = `"Not/A)Brand";v="99""=] "Google Chrome";v="115""=] "Chromium";v="115"`
	DefaultHeaders["sec-ch-ua-mobile"] = "?0"
	DefaultHeaders["sec-ch-ua-platform"] = `"macOS"`
	DefaultHeaders["sec-fetch-dest"] = "empty"
	DefaultHeaders["sec-fetch-mode"] = "cors"
	DefaultHeaders["sec-fetch-site"] = "cross-site"
	DefaultHeaders["user-agent"] = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36"
}
