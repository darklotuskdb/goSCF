package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)
func main() {
	Crst := "\033[0m"
	Cy := "\033[33m"
	Cg:= "\033[32m"

	var Url, ck, idf string
	var mth int
	flag.StringVar(&Url, "u", "", "Authenticated Web Page URL")
	flag.StringVar(&ck, "c", "", "Provide Cookies from Burp Suite")
	flag.StringVar(&idf, "i", "", "Identifier like username, email, Response Header, etc")
	flag.IntVar(&mth, "m", 1, "0 : For All \n1 : Single Cookie Check (Default) \n2 : Double Cookie Check \n3 : Tripe Cookie Check \n4 : Quadruple Cookie Check")
	flag.Parse()

	req, err := http.NewRequest("GET", Url, nil)
	if err != nil {
		log.Fatal(err)
	}

	tc, _ := strconv.Atoi(fmt.Sprint(strings.Count(ck, ";"))) //converting string to integer
if mth == 1 || mth == 0 {
	fmt.Println(string(Cy), "\n[+] ================ Checking For Single Session Cookie =================", string(Crst))
	for i := 0; i <= tc; i++ {
		w := fmt.Sprint(strings.Split(ck, ";")[i])
		fw := fmt.Sprint(strings.Replace(w, " ", "", -1))

		req.Header.Set("Cookie", fw) //passing the cookies one by one
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Fatal(err)
		}

		defer resp.Body.Close()              //will execute after the next function or code
		body, _ := ioutil.ReadAll(resp.Body) //after this defer will execute

		val , _ := regexp.MatchString("(?i)" + idf, string(body))
		if val == true {
			fmt.Println(string(Cg), val, " : ", fw, string(Crst))
		}
	}
}
	if mth == 2 || mth == 0 {
	fmt.Println(string(Cy), "\n[+] ================ Checking For Double Session Cookies ================", string(Crst))
	for i := 0; i <= tc; i++ {
		for j := 0; j <= tc; j++ {
			if i == j {
				break
			} else {
				x := fmt.Sprint(strings.Split(ck, ";")[i])
				y := fmt.Sprint(strings.Split(ck, ";")[j])

				fx := fmt.Sprint(strings.Replace(x, " ", "", -1))
				fy := fmt.Sprint(strings.Replace(y, " ", "", -1))
				fxy := fmt.Sprint(string(fx), "; ", string(fy))

				req.Header.Set("Cookie", fxy)

				resp, err := http.DefaultClient.Do(req)
				if err != nil {
					log.Fatal(err)
				}

				defer resp.Body.Close()
				body, _ := ioutil.ReadAll(resp.Body)

				val , _ := regexp.MatchString("(?i)" + idf, string(body))
				if val == true {
					fmt.Println(string(Cg), val, " : ", fxy, string(Crst))
				}
			}
		}
	}
}
	if mth == 3 || mth == 0 {
	fmt.Println(string(Cy), "\n[+] ================ Checking For Tripe Session Cookies =================", string(Crst))
	for i := 0; i <= tc; i++ {
		for j := 0; j <= tc; j++ {
			for k := 0; k <= tc; k++ {
				if i == j || i == k || j == k {
					break
				} else {
					x := fmt.Sprint(strings.Split(ck, ";")[i])
					y := fmt.Sprint(strings.Split(ck, ";")[j])
					z := fmt.Sprint(strings.Split(ck, ";")[k])

					fx := fmt.Sprint(strings.Replace(x, " ", "", -1))
					fy := fmt.Sprint(strings.Replace(y, " ", "", -1))
					fz := fmt.Sprint(strings.Replace(z, " ", "", -1))
					fxyz := fmt.Sprint(string(fx), "; ", string(fy), "; ", string(fz))

					req.Header.Set("Cookie", fxyz)

					resp, err := http.DefaultClient.Do(req)
					if err != nil {
						log.Fatal(err)
					}

					defer resp.Body.Close()
					body, _ := ioutil.ReadAll(resp.Body)

					val , _ := regexp.MatchString("(?i)" + idf, string(body))
					if val == true {
						fmt.Println(string(Cg), val, " : ", fxyz, string(Crst))
					}
				}
			}
		}
	}
}
	if mth == 4 || mth == 0 {
	fmt.Println(string(Cy), "\n[+] =============== Checking For Quadruple Session Cookies ==============", string(Crst))
	for i := 0; i <= tc; i++ {
		for j := 0; j <= tc; j++ {
			for k := 0; k <= tc; k++ {
				for l := 0; l <= tc; l++ {
					if i == j || i == k || i == l || j == k || j == l || k == l {
						break
					} else {
						x := fmt.Sprint(strings.Split(ck, ";")[i])
						y := fmt.Sprint(strings.Split(ck, ";")[j])
						z := fmt.Sprint(strings.Split(ck, ";")[k])
						w := fmt.Sprint(strings.Split(ck, ";")[l])

						fx := fmt.Sprint(strings.Replace(x, " ", "", -1))
						fy := fmt.Sprint(strings.Replace(y, " ", "", -1))
						fz := fmt.Sprint(strings.Replace(z, " ", "", -1))
						fw := fmt.Sprint(strings.Replace(w, " ", "", -1))
						fxyzw := fmt.Sprint(string(fx), "; ", string(fy), "; ", string(fz), "; ", string(fw))
						req.Header.Set("Cookie", fxyzw)

						resp, err := http.DefaultClient.Do(req)
						if err != nil {
							log.Fatal(err)
						}

						defer resp.Body.Close()
						body, _ := ioutil.ReadAll(resp.Body)

						val , _ := regexp.MatchString("(?i)" + idf, string(body))
						if val == true {
							fmt.Println(string(Cg), val, " : ", fxyzw, string(Crst))
						}
					}
				}
			}
		}
	}
}
		fmt.Println()
}
