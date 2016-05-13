package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

var modes = []string{"bulk", "single"}

var pingstr string = "http://%v:%v/SimplePing/%v/%v"

func main() {
	t := flag.Int("t", 1, "The number of times to ping")
	m := flag.String("m", "bulk", "The mode ")
	p := flag.Int("p", 8080, "Port of the remote server ")

	flag.Parse()

	targets := flag.Args()
	if len(targets) == 2 {

		switch *m {
		case modes[0]:
			bulk(targets[0], targets[1], *t,*p)
		case modes[1]:
			single(targets[0], targets[1], *t,*p)
		default:
			fmt.Println("this mode does not exist.")
		}

	}
}

func bulk(remote string, target string, nbr int,port int) {
	url:= fmt.Sprintf(pingstr,remote,port,target,nbr)
	fmt.Printf("try to get %v\n",url)

	r, err := http.Get(url)
	if err != nil{
		fmt.Println("whoops - sth went wrong")
		return;
	}
	response, _ := ioutil.ReadAll(r.Body)
	fmt.Print(string(response))
}

func single(remote string, target string, nbr int,port int) {

	url:= fmt.Sprintf(pingstr,remote,port,target,1)
	fmt.Printf("try to get %v\n",url)

	for i:=0;i<nbr;i++{
		r, err := http.Get(url)
		if err != nil{
			fmt.Println("whoops - sth went wrong")
			return;
		}
		response, _ := ioutil.ReadAll(r.Body)
		fmt.Print(string(response))
	}
}
