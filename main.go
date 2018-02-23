package main

import (
	"fmt"
	"poker/service"
)

func main() {
	// todo 从JSON文件中获取牌组到数组中
	if pokerFile, err := service.ReadPokerFile("./match.json"); err != nil {
		panic(err.Error())
	} else {
		// todo 循环数组比较各组牌大小
		for i := range pokerFile.File["matches"]{
			pOne := pokerFile.File["matches"][i]["alice"].(string)
			pTwo := pokerFile.File["matches"][i]["bob"].(string)
			go fmt.Printf("%s, %s ,%d\n",pOne , pTwo, service.Compare(pOne, pTwo))
		}
	}
}
