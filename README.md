# kernel_deb_builder
利用 GitHub Actions 自动编译 Linux 内核为 deb 包。

# 如何使用
如果您想要利用我的这个自动化脚本根据自己的需求编译内核，请参考如下步骤：

#### 1. Fork 仓库
访问 [https://github.com/debuggerx01/kernel_deb_builder](https://github.com/debuggerx01/kernel_deb_builder) ，点击右上角的 `Fork` 按钮，并 clone 到本地

#### 2. 更新 config 文件
在本地将您获取的 config 文件替换根目录下的 `config`，可以从您系统的 `/boot/config*` 文件复制，或者手动编辑

#### 3. 编写自定义修改脚本
当前 `/patch.d/` 目录下的修改脚本是只针对我自己的需求编写的，建议您先将其删掉，然后编写自己的脚本放在这个目录下，在脚本执行过程中会自动应用该目录下的所有脚本

#### 4. 推送修改
推送后，action 自动触发，可以在您的仓库页面的 `Actions` 选项卡查看进度详情。

Enjoy～ :grin:

了解更多请看：

[利用 GitHub Actions 自动构建 Linux 内核为 deb 包](https://www.debuggerx.com/2021/08/17/Using-github-actions-to-build-kernel-deb-packages/)

[修改 Linux 内核使系统启动时间缩短约 30 秒](https://www.debuggerx.com/2021/07/07/Modify-the-linux-kernel-to-reduce-the-boot-speed-by-about-30-seconds/)
