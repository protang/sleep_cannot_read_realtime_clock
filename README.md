# sleep: cannot read realtime clock: Invalid argument

Currently Window Store by default provides Ubuntu 20.04 if users install Ubuntu without specifying the version.  
After you install Ubuntu 20.04 in WSL, you may get the following error: "sleep: cannot read realtime clock: Invalid argument"  
This code is designed to solve this problem.  
Related link: <https://github.com/microsoft/WSL/issues/4898>  

当前，如果用户在未指定版本的情况下安装Ubuntu，则默认情况下Window Store将提供Ubuntu 20.04。  
在WSL中安装Ubuntu 20.04后，可能会出现这样的错误: "sleep: cannot read realtime clock: Invalid argument"  
本代码就是为了解决这个问题。  
相关链接： <https://github.com/microsoft/WSL/issues/4898>  

按下面的方法编译代码后替换(/bin/sleep):   
Compile the code as follows and replace (/bin/sleep):  
```
go build -o sleep
mv /bin/sleep /bin/sleep.old
cp sleep /bin/sleep
chmod +x /bin/sleep
```

如果你没有Go的环境或不知如何编译，可以直接用我编译好的 sleep 文件替换也可以.   
If you don’t have a Go environment or don’t know how to compile, you can directly replace it with the `sleep` file I compiled.   


