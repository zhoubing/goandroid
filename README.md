#### 设置Android NDK编译环境
source env.sh


### Windows下使用Idea配置WSL中的Go环境要点
正常创建Windows项目，然后在Edit Configurations中选择Go build，在Run on中选择或者新建配置好的WSL环境(WSL环境中已经配置好Go环境)，即可使用此环境构建工程


#### 交叉编译
https://www.cnblogs.com/jing332/p/16671425.html
https://www.cnblogs.com/develon/p/16464371.html

Settings -> Go -> GoModules -> Environment
1. CC=C:\Users\zhoubing\AppData\Local\Android\Sdk\ndk\21.0.6113669\toolchains\llvm\prebuilt\windows-x86_64\bin\armv7a-linux-androideabi29-clang.cmd;
2. GOPROXY=https://goproxy.cn,direct;
3. GOOS=android;
4. GOARCH=arm;
5. GO111MODULE=on;
6. CGO_ENABLED=1

#### 七牛云
https://goproxy.cn/

#### goland配合wsl2直接调用wsl2里go环境的方法
https://zhuanlan.zhihu.com/p/378437571


#### 串口
https://www.jianshu.com/p/bfda4c8fb09d

#### TCP
https://blog.csdn.net/qq_39445165/article/details/124393872
https://gitee.com/Aceld/zinx


#### SysCall
https://blog.csdn.net/YU20211221/article/details/126422159

#### mmap
https://geektutu.com/post/quick-go-mmap.html
https://studygolang.com/articles/14021
https://studygolang.com/articles/14022


http://androidxref.com/9.0.0_r3/xref/frameworks/native/cmds/servicemanager/service_manager.c
http://androidxref.com/9.0.0_r3/xref/frameworks/native/cmds/servicemanager/binder.c


run java on android
https://raccoon.onyxbits.de/blog/run-java-app-android/
https://stackoverflow.com/questions/36539308/is-it-possible-to-install-the-jdk-on-an-android-device


### Java Socket
https://www.jianshu.com/p/c1d7182b12eb
https://blog.csdn.net/qq23001186/article/details/124669844

### Go Socket
https://www.cnblogs.com/cndeveloper/p/14392374.html

### Binder
http://aospxref.com/android-9.0.0_r61/xref/frameworks/native/libs/binder/IServiceManager.cpp
http://aospxref.com/android-9.0.0_r61/xref/frameworks/native/cmds/servicemanager/binder.c

### GPS
http://aospxref.com/android-11.0.0_r21/xref/hardware/qcom/gps/msm8909/loc_api/libloc_api_50001/gps.c
http://aospxref.com/android-11.0.0_r21/xref/hardware/qcom/gps/msm8960/loc_api/libloc_api_50001/loc.cpp
https://www.cnblogs.com/bcfx/articles/2914687.html
https://blog.csdn.net/weixin_39600704/article/details/117501474
https://blog.csdn.net/weixin_39845825/article/details/117493152
https://www.cnblogs.com/ljf181275034/articles/3319420.html


#### Read Write File
https://www.jianshu.com/p/247b69e7c0dd




https://stackoverflow.com/questions/14750459/change-the-permissions-and-ownership-of-tty-port-in-android-programmatically

#### goroutine pool
https://github.com/panjf2000/ants/blob/master/README_ZH.md
