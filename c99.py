import requests
import argparse
import json
import sys

class C99:
    def __init__(self, apikey):
        self.key = apikey
        self.base_url = "https://api.c99.nl/"

    def make_request(self, endpoint, params):
        params['key'] = self.key
        params['json'] = 'true'
        
        url = f"{self.base_url}{endpoint}"
        response = requests.get(url, params=params)
        return response.json()

    def get_sub_domains(self, subdomain):
        """Find subdomains of a given domain."""
        return self.make_request('subdomainfinder', {'domain': subdomain})

    def get_phone_info(self, number):
        """Get information about a phone number."""
        return self.make_request('phonelookup', {'number': number})

    def get_skype_user_info(self, username):
        """Get information about a Skype user."""
        return self.make_request('skyperesolver', {'username': username})

    def get_skype_ip_info(self, ip):
        """Get Skype information associated with an IP address."""
        return self.make_request('ip2skype', {'ip': ip})

    def firewall_resolver(self, domain):
        """Detect firewalls on a given domain."""
        return self.make_request('firewalldetector', {'url': domain})

    def port_scanner(self, ip):
        """Scan ports on a given IP address."""
        return self.make_request('portscanner', {'host': ip})

    def check_port(self, host, port):
        """Check if a specific port is open on a given host."""
        return self.make_request('portscanner', {'host': host, 'port': port})

    def ping(self, ip):
        """Ping a given IP address."""
        return self.make_request('ping', {'host': ip})

    def hostname_resolver(self, ip):
        """Resolve hostname for a given IP address."""
        return self.make_request('gethostname', {'host': ip})

    def dnschecker(self, domain):
        """Check DNS records for a given domain."""
        return self.make_request('dnschecker', {'url': domain})

    def host_to_ip(self, host):
        """Convert a hostname to an IP address."""
        return self.make_request('dnsresolver', {'host': host, 'server': 'US'})

    def ip_to_domains(self, ip):
        """Find domains associated with a given IP address."""
        return self.make_request('ip2domains', {'ip': ip})

    def alexa_rank(self, url):
        """Get the Alexa rank for a given URL."""
        return self.make_request('alexarank', {'url': url})

    def whois_checker(self, domain):
        """Perform a WHOIS lookup for a given domain."""
        return self.make_request('whois', {'domain': domain})

    def screenshot_tool(self, url):
        """Take a screenshot of a given URL."""
        return self.make_request('createscreenshot', {'url': url})

    def geo_ip(self, host):
        """Get geolocation information for a given IP address."""
        return self.make_request('geoip', {'host': host})

    def website_up_or_down_checker(self, host):
        """Check if a website is up or down."""
        return self.make_request('upordown', {'host': host})

    def site_reputation_checker(self, url):
        """Check the reputation of a given URL."""
        return self.make_request('reputationchecker', {'url': url})

    def get_website_headers(self, host):
        """Get HTTP headers for a given website."""
        return self.make_request('getheaders', {'host': host})

    def link_backup(self, url):
        """Create a backup of a given URL."""
        return self.make_request('linkbackup', {'url': url})

    def url_shortener(self, url):
        """Shorten a given URL."""
        return self.make_request('urlshortener', {'url': url})

    def random_string_picker(self, textfile):
        """Pick a random string from a given text file."""
        return self.make_request('randomstringpicker', {'textfile': textfile})

    def dictionary(self, word):
        """Look up the definition of a word."""
        return self.make_request('dictionary', {'word': word})

    def image_reverse(self, url):
        """Perform a reverse image search."""
        return self.make_request('definepicture', {'url': url})

    def synonym_finder(self, word):
        """Find synonyms for a given word."""
        return self.make_request('synonym', {'word': word})

    def email_validator(self, email):
        """Validate an email address."""
        return self.make_request('emailvalidator', {'email': email})

    def disposable_mail_check(self, email):
        """Check if an email is from a disposable email service."""
        return self.make_request('disposablemailchecker', {'email': email})

    def ip_validator(self, ip):
        """Validate an IP address."""
        return self.make_request('ipvalidator', {'ip': ip})

    def tor_checker(self, ip):
        """Check if an IP address is a Tor exit node."""
        return self.make_request('torchecker', {'ip': ip})

    def translator(self, text, tolanguage):
        """Translate text to a specified language."""
        return self.make_request('translate', {'text': text, 'tolanguage': tolanguage})

    def random_info_generator(self, gender='all'):
        """Generate random person information."""
        return self.make_request('randomperson', {'gender': gender})

    def youtube_video_details(self, videoid):
        """Get details about a YouTube video."""
        return self.make_request('youtubedetails', {'videoid': videoid})

    def youtube_to_mp3(self, videoid):
        """Convert a YouTube video to MP3."""
        return self.make_request('youtubemp3', {'videoid': videoid})

    def ip_logger(self, action='viewloggers'):
        """Log IP addresses."""
        return self.make_request('iplogger', {'action': action})

    def bitcoin_balance(self, address):
        """Check the balance of a Bitcoin address."""
        return self.make_request('bitcoinbalance', {'address': address})

    def ethereum_balance(self, address):
        """Check the balance of an Ethereum address."""
        return self.make_request('ethereumbalance', {'address': address})

    def currency_converter(self, amount, from_currency, to_currency):
        """Convert between currencies."""
        return self.make_request('currency', {'amount': amount, 'from': from_currency, 'to': to_currency})

    def currency_rates(self, source):
        """Get current currency exchange rates."""
        return self.make_request('currencyrates', {'source': source})

    def weather_checker(self, location, unit='C'):
        """Check the weather for a given location."""
        return self.make_request('weather', {'location': location, 'unit': unit})

    def qr_code_generator(self, string, size=150):
        """Generate a QR code."""
        return self.make_request('qrgenerator', {'string': string, 'size': size})

    def text_parser(self, url):
        """Parse text from a given URL."""
        return self.make_request('textparser', {'url': url})

    def proxy_detector(self, ip):
        """Detect if an IP address is a proxy."""
        return self.make_request('proxydetector', {'ip': ip})

    def password_generator(self, length, include, customlist):
        """Generate a random password."""
        return self.make_request('passwordgenerator', {
            'length': length,
            'include': include,
            'customlist': customlist
        })

    def random_number_generator(self, length=None, between=None):
        """Generate a random number."""
        params = {}
        if length is not None:
            params['length'] = length
        if between is not None:
            params['between'] = between
        return self.make_request('randomnumber', params)

    def license_key_generator(self, template, amount=1):
        """Generate license keys."""
        return self.make_request('licensekeygenerator', {
            'template': template,
            'amount': amount
        })

    def either_or(self):
        """Get a random 'either/or' question."""
        return self.make_request('eitheror', {})

    def gif_finder(self, keyword):
        """Find a GIF based on a keyword."""
        return self.make_request('gif', {'keyword': keyword})

def main():
    parser = argparse.ArgumentParser(description='C99 API CLI - A command-line interface for the C99 API')
    parser.add_argument('--apikey', required=True, help='Your C99 API key')
    parser.add_argument('--method', required=True, help='The method to call')
    parser.add_argument('--args', nargs='*', help='Arguments for the method')
    parser.add_argument('--list', action='store_true', help='List all available methods')

    args = parser.parse_args()

    c99 = C99(args.apikey)

    if args.list:
        print("Available methods:")
        for method_name in dir(C99):
            if not method_name.startswith('_') and callable(getattr(C99, method_name)):
                method = getattr(C99, method_name)
                print(f"  {method_name}: {method.__doc__}")
        sys.exit(0)

    if hasattr(c99, args.method):
        method = getattr(c99, args.method)
        print(f"Executing: {args.method}")
        print(f"Description: {method.__doc__}")
        result = method(*args.args) if args.args else method()
        print(json.dumps(result, indent=2))
    else:
        print(f"Error: Method '{args.method}' not found")
        print("Use --list to see all available methods")
        sys.exit(1)

if __name__ == "__main__":
    main()
