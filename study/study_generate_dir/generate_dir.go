package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

var (
	stRootDir   string
	stSeparator string
	iRootNode   Node
)

const (
	stJsonFileName = "file.json"
)

type Node struct {
	Text     string `json:"text"`
	Children []Node `json:"children"`
}

func loadJsonFile() {
	stWorkDir, _ := os.Getwd()
	stSeparator = string(os.PathSeparator)
	stRootDir = stWorkDir[:strings.LastIndex(stWorkDir, stSeparator)]
	stJsonfileByte, err := os.ReadFile(stWorkDir + stSeparator + stJsonFileName)
	if err != nil {
		panic("Read Json File Error :" + err.Error())
	}

	err = json.Unmarshal(stJsonfileByte, &iRootNode)
	if err != nil {
		panic("Load Json Date Error :" + err.Error())
	}

}

func parseNode(iNode Node, parentDir string) {
	if iNode.Text != "" {
		creatPath := stRootDir + stSeparator
		if parentDir != "" {
			creatPath = creatPath + parentDir + stSeparator
		}
		creatDir(creatPath + iNode.Text)
	}
	if iNode.Children != nil {
		for _, n := range iNode.Children {
			parentPath := ""
			if parentDir != "" {
				parentPath = parentDir + stSeparator + iNode.Text
			} else {
				parentPath = iNode.Text
			}
			parseNode(n, parentPath)
		}
	}
}

func creatDir(path string) {
	if path == "" {
		return
	}
	fmt.Println(path)
}

func main() {
	loadJsonFile()
	parseNode(iRootNode, "")
}
