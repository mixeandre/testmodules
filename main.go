package main

import (
	"fmt"
	"github.com/mixeandre/moduleone"
	"github.com/mixeandre/moduletwo"
)

func main()  {
	fmt.Println(moduleone.GetData())
	fmt.Println(moduletwo.GetData())
}
