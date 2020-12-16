package main  
  
import "fmt"  
import b64 "encoding/base64"  
func main() {  
    data := "JavaTpoint@12345!@#$%^&*()"  
    strEncode :=b64.StdEncoding.EncodeToString([]byte(data))  
    fmt.Println("value to be encode  "+data)  
    fmt.Println("Encoden value:  "+strEncode)  
  
    fmt.Println()  
  
  
    fmt.Print("Value to be decode  "+strEncode)  
    strDecode, _ := b64.StdEncoding.DecodeString(strEncode)  
    fmt.Println("Decoded value  "+string( strDecode))  
    fmt.Println()  
  
    url := "https://golang.org/ref/spec"  
  
    fmt.Println("url to be encode  "+url)  
    urlEncode := b64.URLEncoding.EncodeToString([]byte(url))  
    fmt.Println("Encoded url  "+urlEncode)  
  
    fmt.Println("value to be decode  "+urlEncode)  
    strDecode2,_ := b64.URLEncoding.DecodeString(urlEncode)  
  
    fmt.Println("Decoded value  "+string(strDecode2))  
}  