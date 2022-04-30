package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	//name := "olisa michael"
	name1 := os.Args[1]

	fmt.Println(os.Args[0])

	//	tpl := `
	//	<!DOCTYPE html>
	//	<html lang="en">
	//	<head>
	//    <meta charset="UTF-8">
	//    <meta http-equiv="X-UA-Compatible" content="IE=edge">
	//    <meta name="viewport" content="width=device-width, initial-scale=1.0">
	//    <title>Document</title>
	//	</head>
	//	<body>
	//    <h2> ` + name + `<h2>
	//</body>
	//</html>
	//`
	//	str := fmt.Sprintf(tpl)

	tpl2 := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
	</head>
	<body>
    <h2> ` + name1 + `<h2>
</body>
</html>
`

	//fmt.Printf("%v", str)

	newFile, err := os.Create("index4.html")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	_, err1 := io.Copy(newFile, strings.NewReader(tpl2))
	if err1 != nil {
		log.Fatal(err1)
	}

}
