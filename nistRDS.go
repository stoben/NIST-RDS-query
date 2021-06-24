package main


import (
	"crypto/md5"
	"encoding/hex"
	"fmt"	
	"net/http" 
	"os"
	"io"
	"io/ioutil"
)


func main() {
	md5hash, err := md5FromFile(os.Args[0])
	if(err != nil) {
		fmt.Println(err)
		return
	}
	
	baseUrl := fmt.Sprintf("https://hashlookup.circl.lu/lookup/md5/%s", md5hash)
	
	client := &http.Client { }

	req, _ := http.NewRequest("GET", string(baseUrl), nil)
	req.Header.Set("Accept", "application/json")

	resp, _ := client.Do(req)

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println(string(body))
}


func md5FromFile(filepath string) (string, error) {
	var ret string

	file, err:= os.Open(filepath)
	if(err != nil) {
		return ret, err
	}

	defer file.Close()

	hash := md5.New()

	if _, err:= io.Copy(hash, file); err != nil {
		return ret, err
	}

	hashbytes := hash.Sum(nil)[:16]

	ret = hex.EncodeToString(hashbytes)

	return ret, nil

}
