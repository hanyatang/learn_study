package main

import (
    "fmt"
    "math/rand"
    "time"
    "github.com/360EntSecGroup-Skylar/excelize"
)
func main() {
    var (
        first int
        two int
        three int
        four int
        five int
        six int
        seven int
        eight int
        nine int
        ten int
        eleven int
        twelev int
        thireen int
        fourteen int
        fifteen int
        firstjieguo int

    )
     //step1: 设置种子数
     rand.Seed(time.Now().UnixNano())
     //step2：获取随机数
     first = rand.Intn(100000) //[0,100)
     fmt.Println("随机生成的数字为",first)
     two = rand.Intn(100000) //[0,100)
     fmt.Println("随机生成的数字为",two)
     three = rand.Intn(100000) //[0,100)
     fmt.Println("随机生成的数字为",three)
     four = rand.Intn(100000) //[0,100)
     fmt.Println("随机生成的数字为",four)
     five = rand.Intn(100000) //[0,100)
     fmt.Println("随机生成的数字为",five)
     six = rand.Intn(100000) //[0,100)
     fmt.Println("随机生成的数字为",six)
     seven = rand.Intn(100000) //[0,100)
     fmt.Println("随机生成的数字为",seven)
     eight = rand.Intn(100000) //[0,100)
     fmt.Println("随机生成的数字为",eight)
     nine = rand.Intn(100000) //[0,100)
     fmt.Println("随机生成的数字为",nine)
     ten = rand.Intn(100000) //[0,100)
     fmt.Println("随机生成的数字为",ten)
     eleven = rand.Intn(100000) //[0,100)
     fmt.Println("随机生成的数字为",eleven)
     twelev = rand.Intn(100000) //[0,100)
     fmt.Println("随机生成的数字为",twelev)
     thireen = rand.Intn(100000) //[0,100)
     fmt.Println("随机生成的数字为",thireen)
     fourteen = rand.Intn(100000) //[0,100)
     fmt.Println("随机生成的数字为",fourteen)
     fifteen = rand.Intn(100000) //[0,100)
     fmt.Println("随机生成的数字为",fifteen)
     firstjieguo = first+two+three+four+five+six+seven+eight+nine+ten+eleven+twelev+thireen+fourteen+fifteen
  
    f := excelize.NewFile()
    // 创建一个工作表
    index := f.NewSheet("Sheet1")
    // 设置单元格的值
    f.SetCellValue("Sheet1", "A1", first)
    f.SetCellValue("Sheet1", "A2", two)
    f.SetCellValue("Sheet1", "A3",three)
    f.SetCellValue("Sheet1", "A4",four)
    f.SetCellValue("Sheet1", "A5",five)
    f.SetCellValue("Sheet1", "A6",six)
    f.SetCellValue("Sheet1", "A7",seven)
    f.SetCellValue("Sheet1", "A8",eight)
    f.SetCellValue("Sheet1", "A9",nine)
    f.SetCellValue("Sheet1", "A10",ten)
    f.SetCellValue("Sheet1", "A11",eleven)
    f.SetCellValue("Sheet1", "A12",twelev)
    f.SetCellValue("Sheet1", "A13",thireen)
    f.SetCellValue("Sheet1", "A14",fourteen)
    f.SetCellValue("Sheet1", "A15",fifteen)
    f.SetCellValue("Sheet1", "A16",firstjieguo)
    
    // 设置工作簿的默认工作表
    f.SetActiveSheet(index)
    // 根据指定路径保存文件
    err := f.SaveAs("珠心算.xlsx")
    if err != nil {
        fmt.Println(err)
    }
}