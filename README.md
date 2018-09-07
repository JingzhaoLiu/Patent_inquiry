# Patent_inquiry
通过读取excel表格中的公司名称，在表格中写入公司的专利个数和种类


## 版本v1.0

可以读取excel表格中的公司名称，在表格中写入公司的专利个数和种类



		

### 使用
1. ex.xlsx 中公司名是第一列，也就是 A1、 A2、 A3依次排列公司名称


    青岛金大地工艺品有限公司 | header 2| header 2| header 2
    ---|---|---|---
    青岛船舱工艺品有限公司 | row 1 col 2| row 1 col 2| row 1 col 2
    青岛大唐工艺品有限公司 | row 2 col 2| row 1 col 2| row 1 col 2

2. excopy.xlsx是ex.xlsx的复制文件

3. 把文件 ex.xlsx和excopy.xlsx 放到slipe.go同一个文件夹下

4. 写入时会写入到D列，D1、 D2、 D3依次排列


    青岛金大地工艺品有限公司 | header 2| header 2| 外观专利3
    ---|---|---|---
    青岛船舱工艺品有限公司 | row 1 col 2| row 1 col 2| 外观专利3
    青岛大唐工艺品有限公司 | row 2 col 2| row 1 col 2| 外观专利5

- 问题
1. 公司的专利获取不全，获取不到超过5项之后的数据
