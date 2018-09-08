# Patent_inquiry
通过读取excel表格中的公司名称，在表格中写入公司的专利个数和种类



## 版本v3.0

功能已经完成，有时间会优化还有并发版

## go交叉编译


> Mac 下编译 Linux 和 Windows 64位可执行程序

``` shell
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go

CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
```

> Linux 下编译 Mac 和 Windows 64位可执行程序


``` shell
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build main.go

CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
```

> Windows 下编译 Mac 和 Linux 64位可执行程序
``` shell
SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build main.go


SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build main.go

```

> 修改
- 增加了windows下的版本 `bin/slipe.exe`，已经在联想电脑测试通过 

> 修复
- 解决专利准确问题，之前通过专利申请号判断专利，后来发现申请号不一致，开始有的是年，有的是月。已经修改为`申请公布号`判断。

-  解决`freq包`数组去重不能返回多条的问题，之前没有字符串拼接给重复赋值掉了（没注意）




## 版本v2.0 

专利申请号前4位为专利申请的年份。第五位如果是1，则代表是发明专利，如果是2，则代表实用新型专利，如果是3，则代表外观专利。如某专利号为：201510xxxxxx.x则代表这是2015年申请的一件发明专利。如果是201530xxxxxx.x，则代表这是一件2015年申请的外观专利。

专利申请号不准确，已经发现存在  CN`2006`20084377.0和CN`01`268819.3 等年份月份两种。

v3.0改为`申请公布号CN2511690Y`,通过第三位判断。现为`2` 实用新型专利。

> 修改

- 修改为使用公司的专利页面获取信息
  https://www.tianyancha.com/search/t402?key=

> 修复

- 解决公司的专利获取不全问题

## 版本v1.0

可以读取excel表格中的公司名称，在表格中写入公司的专利个数和种类



		

### 使用
1. ex.xlsx 中公司名是第一列，也就是 A1、 A2、 A3依次排列公司名称


    青岛金大地工艺品有限公司 | header 2| header 2| header 2
    ---|---|---|---
    青岛船舱工艺品有限公司 | row 1 col 2| row 1 col 2| row 1 col 2
    青岛大唐工艺品有限公司 | row 2 col 2| row 1 col 2| row 1 col 2

2. excopy.xlsx是ex.xlsx的复制文件

3. 把文件 ex.xlsx和excopy.xlsx 放到slipe.go同一个文件夹下,运行 `go run slipe.go` 。MAC下可以把文件 ex.xlsx和excopy.xlsx 放到bin中，点击编译好的文件 `bin/slipe`。  

4. 写入时会写入到D列，D1、 D2、 D3依次排列


    青岛金大地工艺品有限公司 | header 2| header 2| 外观专利3
    ---|---|---|---
    青岛船舱工艺品有限公司 | row 1 col 2| row 1 col 2| 外观专利3
    青岛大唐工艺品有限公司 | row 2 col 2| row 1 col 2| 外观专利5

> 问题

1. 公司的专利获取不全，获取不到超过5项之后的数据
