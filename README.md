# goSCF
Session Cookie Finder - It helps you to find the main session cookie/s (upto 4) from the bunch of cookies, which is responsible for the user authentication.

## Install
```
go get -u github.com/darklotuskdb/goSCF
```
## Screenshot
![ss](https://user-images.githubusercontent.com/29382875/135152542-d1995524-a055-4a39-9319-7ce4df290536.png)

## Usage
```
./goSCF -u 'https://target.com/myprofile' -c 'JSESSIONID=C0E10C52075E1E5; AltoroAccounts=ODuZ34jxNzIuNDR8' -i 'Hello Admin User'
```
```
-c string
        Provide Cookies from Burp Suite
  -i string
        Identifier like username, email, Response Header, etc
  -m int
        0 : For All
        1 : Single Cookie Check (Default)
        2 : Double Cookie Check
        3 : Triple Cookie Check
        4 : Quadruple Cookie Check (default 1)
  -u string
        Authenticated Web Page URL
```        

## Is It Helpful To You?
[BuyMeACoffee](https://www.buymeacoffee.com/darklotus)

## About Me

* **DarkLotus** - *Cyber Security Researcher* - [DarkLotusKDB](https://darklotuskdb.github.io/KDBhati/)

### Social Media Handles
* [Twitter](https://twitter.com/darklotuskdb)
* [Medium](https://darklotus.medium.com/)
* [Linkedin](https://www.linkedin.com/in/kamaldeepbhati/)
* [Instagram](https://www.instagram.com/kamaldeepbhati/)
