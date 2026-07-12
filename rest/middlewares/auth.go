package middleware

import (
	"log"

	"net/http"
	"strings"
	"mains/util"
)

func AuthMiddleware(mux http.Handler) http.Handler {
	handler := func(w http.ResponseWriter, r *http.Request) {
     
     header :=r.Header.Get("Authorization");

	 if(header == "") {
		 http.Error(w, "Unauthorized", http.StatusUnauthorized)
		 return
	 }

	 headerArr :=strings.Split(header, " ");
	 log.Println("headerArr",headerArr);

	 if(len(headerArr) != 2) {
		 http.Error(w, "Invalid Token", http.StatusUnauthorized)
		 return
	 }

	 accessToken :=headerArr[1];
	 log.Println("accessToken",accessToken);

	 tokenParts :=strings.Split(accessToken, ".");
	 log.Println("tokenParts",tokenParts);

	 jwtHeader :=tokenParts[0];
	 jwtPayload :=tokenParts[1];
	 jwtSignature :=tokenParts[2];

	 byMessase :=jwtHeader+"."+jwtPayload;
	 
	generateHashSignature := util.ConvertHMAC_SHA256([]byte(byMessase), []byte("your_secret_key"));
	generateBase64HashSignature :=util.ConvertBase64(generateHashSignature);

	if(generateBase64HashSignature != jwtSignature) {
		http.Error(w, "Invalid Token", http.StatusUnauthorized)
		return
	}

	mux.ServeHTTP(w,r);

    }
	return http.HandlerFunc(handler);
}
