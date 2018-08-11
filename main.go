package main

import(
	"fmt"
	"os"
)

var fileName = "TabelogInfo.json"

func makeJSONFile() {
	file, err := os.Create("./tabelog_info/" + fileName)
	if err != nil {
		fmt.Println("ファイルが作成できませんでした。")
	}else {
		fmt.Println("ファイルを作成しました。\n%v\n", file)
	}
}

func main(){
	makeJSONFile()
}
