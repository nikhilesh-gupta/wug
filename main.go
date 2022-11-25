package main

import (
   "io/ioutil"
   "log"
   "net/http"
   "fmt"
   "os"
   "os/exec"
   "runtime"
   "flag"
   "github.com/nikhilesh-gupta/wug/gf/gfPatterns"
)

func waybackurls(&domain string){
   var link string
   link = "https://web.archive.org/cdx/search/cdx?url=*." + domain + "/*&output=text&fl=original&collapse=urlkey"
   resp, err := http.Get(link)
   if err != nil {
      log.Fatalln(err)
   }
   body, err := ioutil.ReadAll(resp.Body)
   if err != nil {
      log.Fatalln(err)
   }
   waybackurls_ := string(body)
//    log.Printf(waybackurls_)

   //File creation
   file, err := os.Create("waybackurls.txt")
   if err != nil {
		log.Fatalln(err)
 	}
	defer file.Close()
	_, err2 := file.WriteString(waybackurls_)

    if err2 != nil {
        log.Fatal(err2)
    }
	fmt.Println("done")
}


func main() {
   //Command line Flags
   file := flag.String("f", ,"to enter data from a file")
   domain := flag.String("d", ,"to enter domain")
   flag.Parse()

   //Calling functions
   waybackurls(domain)
   gfPatterns()
}
