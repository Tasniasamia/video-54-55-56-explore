package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
)
func main(){

	str :="hello";
    byteArr :=[]byte(str);

	fmt.Println(byteArr);

base64Conv:=base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(byteArr);

fmt.Println(base64Conv);

decodeBase64,err:=base64.URLEncoding.WithPadding(base64.NoPadding).DecodeString(base64Conv);
if(err != nil){
	log.Fatal(err);
	os.Exit(1);
}

fmt.Println(string(decodeBase64));

	
}