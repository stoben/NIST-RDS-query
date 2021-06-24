package main


import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
	"net/http" 
	"os"
	"io"
	"io/ioutil"
)


/*
examples
nistRDS.go -f file
nistRDS.go -h 8ED4B4ED952526D89899E723F3488DE4


*/
func main() {

	usage := "Usage -f for creating MD5 on the fly -h for setting md5 directly";

	if(len(os.Args) < 3) {
		fmt.Println(usage)
		return
	}

	var md5hash string

	if(os.Args[1] == "-f") {		
		var err error
		md5hash, err = md5FromFile(os.Args[2])
		if(err != nil) {
			fmt.Println(err)
			return
		}
	} else if(os.Args[1] == "-h") {
		md5hash = strings.TrimSpace(os.Args[2])
	} else {
		fmt.Println(usage) //unknown command
		return;
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
