package main

import(
	"fmt"
	"os"
	"io/ioutil"
	"net/http"
)

type TabelogInfo struct {
	Field string `json:"Field"`
	Omit string `json:"Omit"`
	OmitEmpty string `json:"OmitEmpty"`
	Num int `json:"num"`
}

func makeTabeLogInfoFile(info []byte, fileName string){

	file, err := os.Create("./tabelog_info/" + fileName)
	if err != nil {
		fmt.Println("ファイルが作成できませんでした。")
	}else {
		fmt.Println("ファイルを作成しました。\n%v\n", file)
	}	
	
	ioutil.WriteFile("./tabelog_info/" + fileName, info, os.ModePerm)
}

func getHttpInfo(url string) (string, error) {

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	return string(b), nil	

}

func main(){

	var url = "https://tabelog.com/osaka/A2701/A270202/27040810/"
	var fileName = "TabeLogInfo.json"
	
	bodyInfo, err := getHttpInfo(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(bodyInfo)
	makeTabeLogInfoFile([]byte(bodyInfo), fileName)
}

