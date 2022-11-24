package gfPatterns

func xss(){
    //File creation
    file, err := os.Create("gf_Pattern/xss.txt")
    if err != nil {
		log.Fatalln(err)
 	}
     defer file.Close()

    //gf-tool execution
    out, err := exec.Command("cat waybackurls.txt | gf xss").Output()
    if err != nil {
        fmt.Printf("%s", err)
    }
    output := string(out[:])
	_, err2 := file.WriteString(output)
    if err2 != nil {
        log.Fatal(err2)
    }
}

// func ssti(){

// }

// func idor(){
    
// }

// func lfi(){

// }

// func rce(){
    
// }

// func sqli(){

// }

// func ssrf(){
    
// }

// func jsvar(){

// }


func gfPatterns(){
	//Checking for Linux system
	if runtime.GOOS == "windows" {
        fmt.Println("Can't Execute this on a windows machine")
		os.Exit(0)
    }
    //Creating gf directory
    out, err := exec.Command("mkdir gf_Pattern")
    if err != nil {
		log.Fatalln(err)
 	}

    //Calling Pattern Functions
    xss()
    // ssti()
    // idor()
    // lfi()
    // rce()
    // sqli()
    // ssrf()
    // jsvar()
}