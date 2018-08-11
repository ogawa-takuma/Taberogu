package main

import(
	"fmt"
	"os"
	"io/ioutil"
	"encoding/json"
)

type TabelogInfo struct {
	Field string `json:"Field"`
	Omit string `json:"Omit"`
	OmitEmpty string `json:"OmitEmpty"`
	Num int `json:"num"`
}

func makeJSONFile(fileName string) {
	file, err := os.Create("./tabelog_info/" + fileName)
	if err != nil {
		fmt.Println("ファイルが作成できませんでした。")
	}else {
		fmt.Println("ファイルを作成しました。\n%v\n", file)
	}
}

func writeTabeLogInfo(info []byte, fileName string) {
	ioutil.WriteFile(fileName, info, os.ModePerm)
}

func main(){

	var fileName = "TabelogInfo.json"
	
	tabeLogInfo := TabelogInfo{
		Field: "field",
		Omit: "omit",
		OmitEmpty: "omit",
		Num: 28,
	}

	bytes, _ := json.Marshal(&tabeLogInfo)
	fmt.Println(string(bytes))

	makeJSONFile(fileName)
	writeTabeLogInfo(bytes, "./tabelog_info/" + fileName)
}

