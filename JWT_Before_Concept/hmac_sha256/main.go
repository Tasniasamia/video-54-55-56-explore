package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

func main(){
	secret :=[]byte("secret-key");
	message :=[]byte("Hello Hi");

// 	generateHash,err:=hmac.New(sha256.New,secret).Write(message);


// 	if(err != nil){
// 		fmt.Println(err);

// 	}
//    fmt.Println(generateHash); //output: 8



		h:=hmac.New(sha256.New,secret);
		h.Write(message);
		hash:=h.Sum(nil);
		fmt.Println(hash);
        fmt.Println(string(hash));



}