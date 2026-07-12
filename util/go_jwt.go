package util;

import (
	"encoding/base64"
	"encoding/json"
	"crypto/hmac"
	"crypto/sha256"

)

type Header struct{
	ALG string `json:"alg"`;
	TYP string `json:"typ"`;
}


type Payload struct{
	Sub string `json:"sub"`;  //user id
	Name string `json:"name"`;
    Email string `json:"email"`;
	
}

func ConvertBase64(data []byte) string {
	enc :=base64.URLEncoding;
	enc =enc.WithPadding(base64.NoPadding);
    b64 := enc.EncodeToString(data);
	return b64;
}

func ConvertHMAC_SHA256(data []byte,secret []byte)[]byte{
	
h := hmac.New(sha256.New, secret)
h.Write(data);
return h.Sum(nil)

}

// func CreateJwtToken(secret string,payload Payload) string {
// 	header := Header{
// 		ALG: "HS256",
// 		TYP: "JWT",
// 	}

// 	headerBytes, err:= json.Marshal(header)
// 	if err != nil {
// 		panic(err)
// 	}
// 	payloadBytes, err2:= json.Marshal(payload)
// 	if err2 != nil {
// 		panic(err2)
// 	}	
// 	headerEncoded:= ConvertBase64(headerBytes)
// 	payloadEncoded:= ConvertBase64(payloadBytes)
//     signature:= ConvertHMAC_SHA256([]byte(headerEncoded + "." + payloadEncoded), []byte(secret));
// 	signatureBase64:= ConvertBase64(signature)
// 	return headerEncoded + "." + payloadEncoded + "." + signatureBase64
// }


func CreateJwtToken(secret string,payloadObj Payload) string {
	headerObj  := Header{
		ALG: "HS256",
		TYP: "JWT",
	}

	headerBytes, err:= json.Marshal(headerObj )
	if err != nil {
		panic(err)
	}
	payloadBytes, err2:= json.Marshal(payloadObj)
	if err2 != nil {
		panic(err2)
	}	
	headerEncoded:= ConvertBase64(headerBytes)
    payloadEncoded:= ConvertBase64(payloadBytes)
   
	headerPayloadBase64:=headerEncoded + "." + payloadEncoded;

	byteArr1:=[]byte(headerPayloadBase64)
	byteArr2:=[]byte(secret);

    signature:= ConvertHMAC_SHA256(byteArr1, byteArr2);
	signatureBase64:= ConvertBase64(signature)
	jwtToken:= headerEncoded + "." + payloadEncoded + "." + signatureBase64;

	return jwtToken;
}