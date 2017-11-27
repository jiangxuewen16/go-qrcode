# go-qrcode
goLang 二维码导出

go build qrcode   #编译代码

执行
qrcode -i /home/测试用例.txt -o /home/outQrcode -s 1024

-i 要转换成二维码的文件名称（包含地址）
-o 二维码输出的目录
-s 二维码的像素(px)