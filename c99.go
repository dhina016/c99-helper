package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strings"
)

type C99 struct {
	Key     string
	BaseURL string
}

type MethodInfo struct {
	Name        string
	Description string
	Params      []ParamInfo
}

type ParamInfo struct {
	Name     string
	Type     string
	Required bool
}

func NewC99(apikey string) *C99 {
	return &C99{
		Key:     apikey,
		BaseURL: "https://api.c99.nl/",
	}
}

func (c *C99) makeRequest(endpoint string, params map[string]string) (map[string]interface{}, error) {
	params["key"] = c.Key
	params["json"] = "true"

	u, err := url.Parse(c.BaseURL + endpoint)
	if err != nil {
		return nil, err
	}

	q := u.Query()
	for k, v := range params {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetSubDomains finds subdomains of a given domain.
func (c *C99) GetSubDomains(subdomain string) (map[string]interface{}, error) {
	return c.makeRequest("subdomainfinder", map[string]string{"domain": subdomain})
}

// GetPhoneInfo gets information about a phone number.
func (c *C99) GetPhoneInfo(number string) (map[string]interface{}, error) {
	return c.makeRequest("phonelookup", map[string]string{"number": number})
}

// GetSkypeUserInfo gets information about a Skype user.
func (c *C99) GetSkypeUserInfo(username string) (map[string]interface{}, error) {
	return c.makeRequest("skyperesolver", map[string]string{"username": username})
}

// GetSkypeIPInfo gets Skype information associated with an IP address.
func (c *C99) GetSkypeIPInfo(ip string) (map[string]interface{}, error) {
	return c.makeRequest("ip2skype", map[string]string{"ip": ip})
}

// FirewallResolver detects firewalls on a given domain.
func (c *C99) FirewallResolver(domain string) (map[string]interface{}, error) {
	return c.makeRequest("firewalldetector", map[string]string{"url": domain})
}

// PortScanner scans ports on a given IP address.
func (c *C99) PortScanner(ip string) (map[string]interface{}, error) {
	return c.makeRequest("portscanner", map[string]string{"host": ip})
}

// CheckPort checks if a specific port is open on a given host.
func (c *C99) CheckPort(host, port string) (map[string]interface{}, error) {
	return c.makeRequest("portscanner", map[string]string{"host": host, "port": port})
}

// Ping pings a given IP address.
func (c *C99) Ping(ip string) (map[string]interface{}, error) {
	return c.makeRequest("ping", map[string]string{"host": ip})
}

// HostnameResolver resolves hostname for a given IP address.
func (c *C99) HostnameResolver(ip string) (map[string]interface{}, error) {
	return c.makeRequest("gethostname", map[string]string{"host": ip})
}

// DNSChecker checks DNS records for a given domain.
func (c *C99) DNSChecker(domain string) (map[string]interface{}, error) {
	return c.makeRequest("dnschecker", map[string]string{"url": domain})
}

// HostToIP converts a hostname to an IP address.
func (c *C99) HostToIP(host string) (map[string]interface{}, error) {
	return c.makeRequest("dnsresolver", map[string]string{"host": host, "server": "US"})
}

// IPToDomains finds domains associated with a given IP address.
func (c *C99) IPToDomains(ip string) (map[string]interface{}, error) {
	return c.makeRequest("ip2domains", map[string]string{"ip": ip})
}

// AlexaRank gets the Alexa rank for a given URL.
func (c *C99) AlexaRank(url string) (map[string]interface{}, error) {
	return c.makeRequest("alexarank", map[string]string{"url": url})
}

// WhoisChecker performs a WHOIS lookup for a given domain.
func (c *C99) WhoisChecker(domain string) (map[string]interface{}, error) {
	return c.makeRequest("whois", map[string]string{"domain": domain})
}

// ScreenshotTool takes a screenshot of a given URL.
func (c *C99) ScreenshotTool(url string) (map[string]interface{}, error) {
	return c.makeRequest("createscreenshot", map[string]string{"url": url})
}

// GeoIP gets geolocation information for a given IP address.
func (c *C99) GeoIP(host string) (map[string]interface{}, error) {
	return c.makeRequest("geoip", map[string]string{"host": host})
}

// WebsiteUpOrDownChecker checks if a website is up or down.
func (c *C99) WebsiteUpOrDownChecker(host string) (map[string]interface{}, error) {
	return c.makeRequest("upordown", map[string]string{"host": host})
}

// SiteReputationChecker checks the reputation of a given URL.
func (c *C99) SiteReputationChecker(url string) (map[string]interface{}, error) {
	return c.makeRequest("reputationchecker", map[string]string{"url": url})
}

// GetWebsiteHeaders gets HTTP headers for a given website.
func (c *C99) GetWebsiteHeaders(host string) (map[string]interface{}, error) {
	return c.makeRequest("getheaders", map[string]string{"host": host})
}

// LinkBackup creates a backup of a given URL.
func (c *C99) LinkBackup(url string) (map[string]interface{}, error) {
	return c.makeRequest("linkbackup", map[string]string{"url": url})
}

// URLShortener shortens a given URL.
func (c *C99) URLShortener(url string) (map[string]interface{}, error) {
	return c.makeRequest("urlshortener", map[string]string{"url": url})
}

// RandomStringPicker picks a random string from a given text file.
func (c *C99) RandomStringPicker(textfile string) (map[string]interface{}, error) {
	return c.makeRequest("randomstringpicker", map[string]string{"textfile": textfile})
}

// Dictionary looks up the definition of a word.
func (c *C99) Dictionary(word string) (map[string]interface{}, error) {
	return c.makeRequest("dictionary", map[string]string{"word": word})
}

// ImageReverse performs a reverse image search.
func (c *C99) ImageReverse(url string) (map[string]interface{}, error) {
	return c.makeRequest("definepicture", map[string]string{"url": url})
}

// SynonymFinder finds synonyms for a given word.
func (c *C99) SynonymFinder(word string) (map[string]interface{}, error) {
	return c.makeRequest("synonym", map[string]string{"word": word})
}

// EmailValidator validates an email address.
func (c *C99) EmailValidator(email string) (map[string]interface{}, error) {
	return c.makeRequest("emailvalidator", map[string]string{"email": email})
}

// DisposableMailCheck checks if an email is from a disposable email service.
func (c *C99) DisposableMailCheck(email string) (map[string]interface{}, error) {
	return c.makeRequest("disposablemailchecker", map[string]string{"email": email})
}

// IPValidator validates an IP address.
func (c *C99) IPValidator(ip string) (map[string]interface{}, error) {
	return c.makeRequest("ipvalidator", map[string]string{"ip": ip})
}

// TorChecker checks if an IP address is a Tor exit node.
func (c *C99) TorChecker(ip string) (map[string]interface{}, error) {
	return c.makeRequest("torchecker", map[string]string{"ip": ip})
}

// Translator translates text to a specified language.
func (c *C99) Translator(text, tolanguage string) (map[string]interface{}, error) {
	return c.makeRequest("translate", map[string]string{"text": text, "tolanguage": tolanguage})
}

// RandomInfoGenerator generates random person information.
func (c *C99) RandomInfoGenerator(gender string) (map[string]interface{}, error) {
	return c.makeRequest("randomperson", map[string]string{"gender": gender})
}

// YouTubeVideoDetails gets details about a YouTube video.
func (c *C99) YouTubeVideoDetails(videoid string) (map[string]interface{}, error) {
	return c.makeRequest("youtubedetails", map[string]string{"videoid": videoid})
}

// YouTubeToMP3 converts a YouTube video to MP3.
func (c *C99) YouTubeToMP3(videoid string) (map[string]interface{}, error) {
	return c.makeRequest("youtubemp3", map[string]string{"videoid": videoid})
}

// IPLogger logs IP addresses.
func (c *C99) IPLogger(action string) (map[string]interface{}, error) {
	return c.makeRequest("iplogger", map[string]string{"action": action})
}

// BitcoinBalance checks the balance of a Bitcoin address.
func (c *C99) BitcoinBalance(address string) (map[string]interface{}, error) {
	return c.makeRequest("bitcoinbalance", map[string]string{"address": address})
}

// EthereumBalance checks the balance of an Ethereum address.
func (c *C99) EthereumBalance(address string) (map[string]interface{}, error) {
	return c.makeRequest("ethereumbalance", map[string]string{"address": address})
}

// CurrencyConverter converts between currencies.
func (c *C99) CurrencyConverter(amount, fromCurrency, toCurrency string) (map[string]interface{}, error) {
	return c.makeRequest("currency", map[string]string{"amount": amount, "from": fromCurrency, "to": toCurrency})
}

// CurrencyRates gets current currency exchange rates.
func (c *C99) CurrencyRates(source string) (map[string]interface{}, error) {
	return c.makeRequest("currencyrates", map[string]string{"source": source})
}

// WeatherChecker checks the weather for a given location.
func (c *C99) WeatherChecker(location, unit string) (map[string]interface{}, error) {
	return c.makeRequest("weather", map[string]string{"location": location, "unit": unit})
}

// QRCodeGenerator generates a QR code.
func (c *C99) QRCodeGenerator(str string, size string) (map[string]interface{}, error) {
	return c.makeRequest("qrgenerator", map[string]string{"string": str, "size": size})
}

// TextParser parses text from a given URL.
func (c *C99) TextParser(url string) (map[string]interface{}, error) {
	return c.makeRequest("textparser", map[string]string{"url": url})
}

// ProxyDetector detects if an IP address is a proxy.
func (c *C99) ProxyDetector(ip string) (map[string]interface{}, error) {
	return c.makeRequest("proxydetector", map[string]string{"ip": ip})
}

// PasswordGenerator generates a random password.
func (c *C99) PasswordGenerator(length, include, customlist string) (map[string]interface{}, error) {
	return c.makeRequest("passwordgenerator", map[string]string{
		"length":     length,
		"include":    include,
		"customlist": customlist,
	})
}

// RandomNumberGenerator generates a random number.
func (c *C99) RandomNumberGenerator(length, between string) (map[string]interface{}, error) {
	params := make(map[string]string)
	if length != "" {
		params["length"] = length
	}
	if between != "" {
		params["between"] = between
	}
	return c.makeRequest("randomnumber", params)
}

// LicenseKeyGenerator generates license keys.
func (c *C99) LicenseKeyGenerator(template string, amount string) (map[string]interface{}, error) {
	return c.makeRequest("licensekeygenerator", map[string]string{
		"template": template,
		"amount":   amount,
	})
}

// EitherOr gets a random 'either/or' question.
func (c *C99) EitherOr() (map[string]interface{}, error) {
	return c.makeRequest("eitheror", map[string]string{})
}

// GIFFinder finds a GIF based on a keyword.
func (c *C99) GIFFinder(keyword string) (map[string]interface{}, error) {
	return c.makeRequest("gif", map[string]string{"keyword": keyword})
}

// Define method information using struct tags
var methodInfos = []MethodInfo{
	{
		Name:        "GetSubDomains",
		Description: "Find subdomains of a given domain.",
		Params: []ParamInfo{
			{Name: "subdomain", Type: "string", Required: true},
		},
	},
	{
		Name:        "GetPhoneInfo",
		Description: "Get information about a phone number.",
		Params: []ParamInfo{
			{Name: "number", Type: "string", Required: true},
		},
	},
	{
		Name:        "GetSkypeUserInfo",
		Description: "Get information about a Skype user.",
		Params: []ParamInfo{
			{Name: "username", Type: "string", Required: true},
		},
	},
	{
		Name:        "GetSkypeIPInfo",
		Description: "Get Skype information associated with an IP address.",
		Params: []ParamInfo{
			{Name: "ip", Type: "string", Required: true},
		},
	},
	{
		Name:        "FirewallResolver",
		Description: "Detect firewalls on a given domain.",
		Params: []ParamInfo{
			{Name: "domain", Type: "string", Required: true},
		},
	},
	{
		Name:        "PortScanner",
		Description: "Scan ports on a given IP address.",
		Params: []ParamInfo{
			{Name: "ip", Type: "string", Required: true},
		},
	},
	{
		Name:        "CheckPort",
		Description: "Check if a specific port is open on a given host.",
		Params: []ParamInfo{
			{Name: "host", Type: "string", Required: true},
			{Name: "port", Type: "string", Required: true},
		},
	},
	{
		Name:        "Ping",
		Description: "Ping a given IP address.",
		Params: []ParamInfo{
			{Name: "ip", Type: "string", Required: true},
		},
	},
	{
		Name:        "HostnameResolver",
		Description: "Resolve hostname for a given IP address.",
		Params: []ParamInfo{
			{Name: "ip", Type: "string", Required: true},
		},
	},
	{
		Name:        "DNSChecker",
		Description: "Check DNS records for a given domain.",
		Params: []ParamInfo{
			{Name: "domain", Type: "string", Required: true},
		},
	},
	{
		Name:        "HostToIP",
		Description: "Convert a hostname to an IP address.",
		Params: []ParamInfo{
			{Name: "host", Type: "string", Required: true},
		},
	},
	{
		Name:        "IPToDomains",
		Description: "Find domains associated with a given IP address.",
		Params: []ParamInfo{
			{Name: "ip", Type: "string", Required: true},
		},
	},
	{
		Name:        "AlexaRank",
		Description: "Get the Alexa rank for a given URL.",
		Params: []ParamInfo{
			{Name: "url", Type: "string", Required: true},
		},
	},
	{
		Name:        "WhoisChecker",
		Description: "Perform a WHOIS lookup for a given domain.",
		Params: []ParamInfo{
			{Name: "domain", Type: "string", Required: true},
		},
	},
	{
		Name:        "ScreenshotTool",
		Description: "Take a screenshot of a given URL.",
		Params: []ParamInfo{
			{Name: "url", Type: "string", Required: true},
		},
	},
	{
		Name:        "GeoIP",
		Description: "Get geolocation information for a given IP address.",
		Params: []ParamInfo{
			{Name: "host", Type: "string", Required: true},
		},
	},
	{
		Name:        "WebsiteUpOrDownChecker",
		Description: "Check if a website is up or down.",
		Params: []ParamInfo{
			{Name: "host", Type: "string", Required: true},
		},
	},
	{
		Name:        "SiteReputationChecker",
		Description: "Check the reputation of a given URL.",
		Params: []ParamInfo{
			{Name: "url", Type: "string", Required: true},
		},
	},
	{
		Name:        "GetWebsiteHeaders",
		Description: "Get HTTP headers for a given website.",
		Params: []ParamInfo{
			{Name: "host", Type: "string", Required: true},
		},
	},
	{
		Name:        "LinkBackup",
		Description: "Create a backup of a given URL.",
		Params: []ParamInfo{
			{Name: "url", Type: "string", Required: true},
		},
	},
	{
		Name:        "URLShortener",
		Description: "Shorten a given URL.",
		Params: []ParamInfo{
			{Name: "url", Type: "string", Required: true},
		},
	},
	{
		Name:        "RandomStringPicker",
		Description: "Pick a random string from a given text file.",
		Params: []ParamInfo{
			{Name: "textfile", Type: "string", Required: true},
		},
	},
	{
		Name:        "Dictionary",
		Description: "Look up the definition of a word.",
		Params: []ParamInfo{
			{Name: "word", Type: "string", Required: true},
		},
	},
	{
		Name:        "ImageReverse",
		Description: "Perform a reverse image search.",
		Params: []ParamInfo{
			{Name: "url", Type: "string", Required: true},
		},
	},
	{
		Name:        "SynonymFinder",
		Description: "Find synonyms for a given word.",
		Params: []ParamInfo{
			{Name: "word", Type: "string", Required: true},
		},
	},
	{
		Name:        "EmailValidator",
		Description: "Validate an email address.",
		Params: []ParamInfo{
			{Name: "email", Type: "string", Required: true},
		},
	},
	{
		Name:        "DisposableMailCheck",
		Description: "Check if an email is from a disposable email service.",
		Params: []ParamInfo{
			{Name: "email", Type: "string", Required: true},
		},
	},
	{
		Name:        "IPValidator",
		Description: "Validate an IP address.",
		Params: []ParamInfo{
			{Name: "ip", Type: "string", Required: true},
		},
	},
	{
		Name:        "TorChecker",
		Description: "Check if an IP address is a Tor exit node.",
		Params: []ParamInfo{
			{Name: "ip", Type: "string", Required: true},
		},
	},
	{
		Name:        "Translator",
		Description: "Translate text to a specified language.",
		Params: []ParamInfo{
			{Name: "text", Type: "string", Required: true},
			{Name: "tolanguage", Type: "string", Required: true},
		},
	},
	{
		Name:        "RandomInfoGenerator",
		Description: "Generate random person information.",
		Params: []ParamInfo{
			{Name: "gender", Type: "string", Required: true},
		},
	},
	{
		Name:        "YouTubeVideoDetails",
		Description: "Get details about a YouTube video.",
		Params: []ParamInfo{
			{Name: "videoid", Type: "string", Required: true},
		},
	},
	{
		Name:        "YouTubeToMP3",
		Description: "Convert a YouTube video to MP3.",
		Params: []ParamInfo{
			{Name: "videoid", Type: "string", Required: true},
		},
	},
	{
		Name:        "IPLogger",
		Description: "Log IP addresses.",
		Params: []ParamInfo{
			{Name: "action", Type: "string", Required: true},
		},
	},
	{
		Name:        "BitcoinBalance",
		Description: "Check the balance of a Bitcoin address.",
		Params: []ParamInfo{
			{Name: "address", Type: "string", Required: true},
		},
	},
	{
		Name:        "EthereumBalance",
		Description: "Check the balance of an Ethereum address.",
		Params: []ParamInfo{
			{Name: "address", Type: "string", Required: true},
		},
	},
	{
		Name:        "CurrencyConverter",
		Description: "Convert between currencies.",
		Params: []ParamInfo{
			{Name: "amount", Type: "string", Required: true},
			{Name: "fromCurrency", Type: "string", Required: true},
			{Name: "toCurrency", Type: "string", Required: true},
		},
	},
	{
		Name:        "CurrencyRates",
		Description: "Get current currency exchange rates.",
		Params: []ParamInfo{
			{Name: "source", Type: "string", Required: true},
		},
	},
	{
		Name:        "WeatherChecker",
		Description: "Check the weather for a given location.",
		Params: []ParamInfo{
			{Name: "location", Type: "string", Required: true},
			{Name: "unit", Type: "string", Required: true},
		},
	},
	{
		Name:        "QRCodeGenerator",
		Description: "Generate a QR code.",
		Params: []ParamInfo{
			{Name: "str", Type: "string", Required: true},
			{Name: "size", Type: "string", Required: true},
		},
	},
	{
		Name:        "TextParser",
		Description: "Parse text from a given URL.",
		Params: []ParamInfo{
			{Name: "url", Type: "string", Required: true},
		},
	},
	{
		Name:        "ProxyDetector",
		Description: "Detect if an IP address is a proxy.",
		Params: []ParamInfo{
			{Name: "ip", Type: "string", Required: true},
		},
	},
	{
		Name:        "PasswordGenerator",
		Description: "Generate a random password.",
		Params: []ParamInfo{
			{Name: "length", Type: "string", Required: true},
			{Name: "include", Type: "string", Required: true},
			{Name: "customlist", Type: "string", Required: true},
		},
	},
	{
		Name:        "RandomNumberGenerator",
		Description: "Generate a random number.",
		Params: []ParamInfo{
			{Name: "length", Type: "string", Required: false},
			{Name: "between", Type: "string", Required: false},
		},
	},
	{
		Name:        "LicenseKeyGenerator",
		Description: "Generate license keys.",
		Params: []ParamInfo{
			{Name: "template", Type: "string", Required: true},
			{Name: "amount", Type: "string", Required: true},
		},
	},
	{
		Name:        "EitherOr",
		Description: "Get a random 'either/or' question.",
		Params:      []ParamInfo{},
	},
	{
		Name:        "GIFFinder",
		Description: "Find a GIF based on a keyword.",
		Params: []ParamInfo{
			{Name: "keyword", Type: "string", Required: true},
		},
	},
}

func getMethodInfo(name string) *MethodInfo {
	for _, info := range methodInfos {
		if info.Name == name {
			return &info
		}
	}
	return nil
}

func printMethodUsage(info *MethodInfo) {
	fmt.Printf("Usage: %s ", info.Name)
	for _, param := range info.Params {
		if param.Required {
			fmt.Printf("<%s> ", param.Name)
		} else {
			fmt.Printf("[%s] ", param.Name)
		}
	}
	fmt.Printf("\nDescription: %s\n", info.Description)
	fmt.Println("Parameters:")
	for _, param := range info.Params {
		req := "optional"
		if param.Required {
			req = "required"
		}
		fmt.Printf("  %s (%s): %s\n", param.Name, param.Type, req)
	}
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run c99_api.go <apikey> <method> [args...]")
		fmt.Println("Use 'list' as the method to see all available methods")
		os.Exit(1)
	}

	apiKey := os.Args[1]
	method := os.Args[2]
	args := os.Args[3:]

	c99 := NewC99(apiKey)

	if method == "list" {
		fmt.Println("Available methods:")
		for _, info := range methodInfos {
			fmt.Printf("  %s: %s\n", info.Name, info.Description)
		}
		os.Exit(0)
	}

	methodInfo := getMethodInfo(method)
	if methodInfo == nil {
		fmt.Printf("Error: Method '%s' not found\n", method)
		fmt.Println("Use 'list' to see all available methods")
		os.Exit(1)
	}

	if len(args) < len(methodInfo.Params) {
		fmt.Printf("Error: Not enough arguments for method '%s'\n", method)
		printMethodUsage(methodInfo)
		os.Exit(1)
	}

	m := reflect.ValueOf(c99).MethodByName(method)
	if !m.IsValid() {
		fmt.Printf("Error: Method '%s' not found\n", method)
		fmt.Println("Use 'list' to see all available methods")
		os.Exit(1)
	}

	params := make([]reflect.Value, len(args))
	for i, arg := range args {
		params[i] = reflect.ValueOf(arg)
	}

	result := m.Call(params)
	if len(result) == 2 && !result[1].IsNil() {
		fmt.Printf("Error: %v\n", result[1].Interface())
		os.Exit(1)
	}

	jsonResult, err := json.MarshalIndent(result[0].Interface(), "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(jsonResult))
}
