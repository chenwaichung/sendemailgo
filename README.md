# A simple email client based on gomail

Introduction
------------
Most options are compatible with sendEmail[https://github.com/mogaal/sendemail]


Build and Usage
----------------------
To build it, run:
    
```bash
    go build
```
Display usage, run:

```bash
./sendemailgo --help

Usage of ./sendemailgo:
  -f string
        from (sender) email address.
  -h    display the usage and exit.
  -help
        display the usage and exit.
  -m string
        message body.
  -p int
        the mail server port, default 25. (default 25)
  -s string
        the mail server address.
  -t string
        to email address(receiver).
  -tls
        enable the tls option.
  -u string
        message subject.
  -v    display the version and exit.
  -version
        display the version and exit.
  -xp string
        password for SMTP authentication.
  -xu string
        username for SMTP authentication.
```

Example
-------

```bash
    ./sendemailgo -f example@example.com -s smtp.example.com -xu example@example.com -xp password -t receiver@example.com　-u "Test" -m "Test mail" 
```