# C99 API Helper

This project provides helper libraries and command-line interfaces (CLIs) for interacting with the C99 API in PHP, Go, and Python. These implementations make it easy to integrate C99 API functionality into your projects or use them directly from the command line.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
  - [PHP](#php)
  - [Go](#go)
  - [Python](#python)
- [Available Methods](#available-methods)
- [Contributing](#contributing)
- [License](#license)

## Features

- Easy-to-use wrappers for C99 API endpoints
- Support for multiple programming languages: PHP, Go, and Python
- Command-line interfaces for quick access to API functionality
- Comprehensive set of methods covering various C99 API features

## Installation

### PHP

1. Include the `C99.php` file in your project.
2. Instantiate the `C99` class with your API key.

### Go

1. Clone the repository or copy the `c99_api.go` file into your project.
2. Build the CLI tool:
   ```
   go build c99_api.go
   ```

### Python

1. Ensure you have Python 3.x installed.
2. Install the required dependencies:
   ```
   pip install requests
   ```
3. Copy the `c99_api.py` file into your project directory.

## Usage

### PHP

```php
require_once 'C99.php';

$apiKey = 'your_api_key_here';
$c99 = new C99(['apikey' => $apiKey]);

$result = $c99->get_sub_domains('example.com');
print_r($result);
```

### Go

To use the Go CLI:

```
./c99_api <apikey> <method> [args...]
```

For example:

```
./c99_api your_api_key_here GetSubDomains example.com
```

To list all available methods:

```
./c99_api your_api_key_here list
```

### Python

To use the Python CLI:

```
python c99_api.py --apikey your_api_key_here --method method_name --args arg1 arg2
```

For example:

```
python c99_api.py --apikey your_api_key_here --method get_sub_domains --args example.com
```

To list all available methods:

```
python c99_api.py --apikey your_api_key_here --list
```

## Available Methods and Usage

Below is a comprehensive list of all available methods in the C99 API Helper, along with examples of how to use them in PHP, Go, and Python.

### 1. Get Subdomains

Finds subdomains of a given domain.

PHP:
```php
$result = $c99->get_sub_domains('example.com');
```

Go:
```
./c99_api your_api_key GetSubDomains example.com
```

Python:
```
python c99_api.py --apikey your_api_key --method get_sub_domains --args example.com
```

### 2. Get Phone Info

Get information about a phone number.

PHP:
```php
$result = $c99->get_phone_info('+1234567890');
```

Go:
```
./c99_api your_api_key GetPhoneInfo +1234567890
```

Python:
```
python c99_api.py --apikey your_api_key --method get_phone_info --args +1234567890
```

### 3. Get Skype User Info

Get information about a Skype user.

PHP:
```php
$result = $c99->get_skype_user_info('username');
```

Go:
```
./c99_api your_api_key GetSkypeUserInfo username
```

Python:
```
python c99_api.py --apikey your_api_key --method get_skype_user_info --args username
```

### 4. Get Skype IP Info

Get Skype information associated with an IP address.

PHP:
```php
$result = $c99->get_skype_ip_info('192.168.1.1');
```

Go:
```
./c99_api your_api_key GetSkypeIPInfo 192.168.1.1
```

Python:
```
python c99_api.py --apikey your_api_key --method get_skype_ip_info --args 192.168.1.1
```

### 5. Firewall Resolver

Detect firewalls on a given domain.

PHP:
```php
$result = $c99->firewall_resolver('example.com');
```

Go:
```
./c99_api your_api_key FirewallResolver example.com
```

Python:
```
python c99_api.py --apikey your_api_key --method firewall_resolver --args example.com
```

### 6. Port Scanner

Scan ports on a given IP address.

PHP:
```php
$result = $c99->port_scanner('192.168.1.1');
```

Go:
```
./c99_api your_api_key PortScanner 192.168.1.1
```

Python:
```
python c99_api.py --apikey your_api_key --method port_scanner --args 192.168.1.1
```

### 7. Check Port

Check if a specific port is open on a given host.

PHP:
```php
$result = $c99->check_port('example.com', '80');
```

Go:
```
./c99_api your_api_key CheckPort example.com 80
```

Python:
```
python c99_api.py --apikey your_api_key --method check_port --args example.com 80
```

### 8. Ping

Ping a given IP address.

PHP:
```php
$result = $c99->ping('192.168.1.1');
```

Go:
```
./c99_api your_api_key Ping 192.168.1.1
```

Python:
```
python c99_api.py --apikey your_api_key --method ping --args 192.168.1.1
```

### 9. Hostname Resolver

Resolve hostname for a given IP address.

PHP:
```php
$result = $c99->hostname_resolver('192.168.1.1');
```

Go:
```
./c99_api your_api_key HostnameResolver 192.168.1.1
```

Python:
```
python c99_api.py --apikey your_api_key --method hostname_resolver --args 192.168.1.1
```

### 10. DNS Checker

Check DNS records for a given domain.

PHP:
```php
$result = $c99->dnschecker('example.com');
```

Go:
```
./c99_api your_api_key DNSChecker example.com
```

Python:
```
python c99_api.py --apikey your_api_key --method dnschecker --args example.com
```


## Note

Remember to replace `your_api_key` with your actual C99 API key in all examples.


## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is open-source and available under the [MIT License](https://opensource.org/licenses/MIT).
