<?php
defined('BASEPATH') OR exit('No direct script access allowed');

class C99 {
    private $CI;
    private $key;

    public function __construct($apikey = '')
    {
        $this->CI =& get_instance();
        $this->key = $apikey['apikey'];
    }

    private function curl_get($url)
    {
        $ch = curl_init();
        curl_setopt($ch, CURLOPT_URL, $url);
        curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1);
        curl_setopt($ch, CURLOPT_SSL_VERIFYPEER, false);
        curl_setopt($ch, CURLOPT_SSL_VERIFYHOST, false);
        $output = curl_exec($ch);
        curl_close($ch);
        return $output;
    }

    private function make_request($endpoint, $params)
    {
        $params['key'] = $this->key;
        $params['json'] = 'true';
        
        $url = "https://api.c99.nl/{$endpoint}?" . http_build_query($params);
        
        $response = $this->curl_get($url);

        return json_decode($response, true);
    }

    public function get_sub_domains($subdomain)
    {
        return $this->make_request('subdomainfinder', ['domain' => $subdomain]);
    }

    public function get_phone_info($number)
    {
        return $this->make_request('phonelookup', ['number' => $number]);
    }

    public function get_skype_user_info($username)
    {
        return $this->make_request('skyperesolver', ['username' => $username]);
    }

    public function get_skype_ip_info($ip)
    {
        return $this->make_request('ip2skype', ['ip' => $ip]);
    }

    public function firewall_resolver($domain)
    {
        return $this->make_request('firewalldetector', ['url' => $domain]);
    }

    public function port_scanner($ip)
    {
        return $this->make_request('portscanner', ['host' => $ip]);
    }

    public function check_port($host, $port)
    {
        return $this->make_request('portscanner', ['host' => $host, 'port' => $port]);
    }

    public function ping($ip)
    {
        return $this->make_request('ping', ['host' => $ip]);
    }

    public function hostname_resolver($ip)
    {
        return $this->make_request('gethostname', ['host' => $ip]);
    }

    public function dnschecker($domain)
    {
        return $this->make_request('dnschecker', ['url' => $domain]);
    }

    public function host_to_ip($host)
    {
        return $this->make_request('dnsresolver', ['host' => $host, 'server' => 'US']);
    }

    public function ip_to_domains($ip)
    {
        return $this->make_request('ip2domains', ['ip' => $ip]);
    }

    public function alexa_rank($url)
    {
        return $this->make_request('alexarank', ['url' => $url]);
    }

    public function whois_checker($domain)
    {
        return $this->make_request('whois', ['domain' => $domain]);
    }

    public function screenshot_tool($url)
    {
        return $this->make_request('createscreenshot', ['url' => $url]);
    }

    public function geo_ip($host)
    {
        return $this->make_request('geoip', ['host' => $host]);
    }

    public function website_up_or_down_checker($host)
    {
        return $this->make_request('upordown', ['host' => $host]);
    }

    public function site_reputation_checker($url)
    {
        return $this->make_request('reputationchecker', ['url' => $url]);
    }

    public function get_website_headers($host)
    {
        return $this->make_request('getheaders', ['host' => $host]);
    }

    public function link_backup($url)
    {
        return $this->make_request('linkbackup', ['url' => $url]);
    }

    public function url_shortener($url)
    {
        return $this->make_request('urlshortener', ['url' => $url]);
    }

    public function random_string_picker($textfile)
    {
        return $this->make_request('randomstringpicker', ['textfile' => $textfile]);
    }

    public function dictionary($word)
    {
        return $this->make_request('dictionary', ['word' => $word]);
    }

    public function image_reverse($url)
    {
        return $this->make_request('definepicture', ['url' => $url]);
    }

    public function synonym_finder($word)
    {
        return $this->make_request('synonym', ['word' => $word]);
    }

    public function email_validator($email)
    {
        return $this->make_request('emailvalidator', ['email' => $email]);
    }

    public function disposable_mail_check($email)
    {
        return $this->make_request('disposablemailchecker', ['email' => $email]);
    }

    public function ip_validator($ip)
    {
        return $this->make_request('ipvalidator', ['ip' => $ip]);
    }

    public function tor_checker($ip)
    {
        return $this->make_request('torchecker', ['ip' => $ip]);
    }

    public function translator($text, $tolanguage)
    {
        return $this->make_request('translate', ['text' => $text, 'tolanguage' => $tolanguage]);
    }

    public function random_info_generator($gender = 'all')
    {
        return $this->make_request('randomperson', ['gender' => $gender]);
    }

    public function youtube_video_details($videoid)
    {
        return $this->make_request('youtubedetails', ['videoid' => $videoid]);
    }

    public function youtube_to_mp3($videoid)
    {
        return $this->make_request('youtubemp3', ['videoid' => $videoid]);
    }

    public function ip_logger($action = 'viewloggers')
    {
        return $this->make_request('iplogger', ['action' => $action]);
    }

    public function bitcoin_balance($address)
    {
        return $this->make_request('bitcoinbalance', ['address' => $address]);
    }

    public function ethereum_balance($address)
    {
        return $this->make_request('ethereumbalance', ['address' => $address]);
    }

    public function currency_converter($amount, $from, $to)
    {
        return $this->make_request('currency', ['amount' => $amount, 'from' => $from, 'to' => $to]);
    }

    public function currency_rates($source)
    {
        return $this->make_request('currencyrates', ['source' => $source]);
    }

    public function weather_checker($location, $unit = 'C')
    {
        return $this->make_request('weather', ['location' => $location, 'unit' => $unit]);
    }

    public function qr_code_generator($string, $size = 150)
    {
        return $this->make_request('qrgenerator', ['string' => $string, 'size' => $size]);
    }

    public function text_parser($url)
    {
        return $this->make_request('textparser', ['url' => $url]);
    }

    public function proxy_detector($ip)
    {
        return $this->make_request('proxydetector', ['ip' => $ip]);
    }

    public function password_generator($length, $include, $customlist)
    {
        return $this->make_request('passwordgenerator', [
            'length' => $length,
            'include' => $include,
            'customlist' => $customlist
        ]);
    }

    public function random_number_generator($length = null, $between = null)
    {
        $params = [];
        if ($length !== null) $params['length'] = $length;
        if ($between !== null) $params['between'] = $between;
        return $this->make_request('randomnumber', $params);
    }

    public function license_key_generator($template, $amount = 1)
    {
        return $this->make_request('licensekeygenerator', [
            'template' => $template,
            'amount' => $amount
        ]);
    }

    public function either_or()
    {
        return $this->make_request('eitheror', []);
    }

    public function gif_finder($keyword)
    {
        return $this->make_request('gif', ['keyword' => $keyword]);
    }
}
