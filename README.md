对阿里云上的常见API进行封装，方便不会编程的同学也可以方便地通过命令行管理阿里云上的实例。

参数概览：
``` bash
-configFilePath string
      -configFilePath <path> 可选，若不存在则会首次启动提示输入参数并自动生成。 (default "./config.yaml")
```
yaml文件支持的参数如下:
``` golang
// Config 保存本程序的设置
type Config struct {
	Action          string `yaml:"action,omitempty"`
	AccessKeyID     string `yaml:"accessKeyID,omitempty"`
	AccessSecret    string `yaml:"accessSecret,omitempty"`
	RegionID        string `yaml:"regionID,omitempty"`
	InstanceID      string `yaml:"instanceID,omitempty"`
	SecurityGroupID string `yaml:"securityGroupID,omitempty"`
}
```

action参数就是要进行的操作，和阿里云上的api名称一一对应。

现在支持describeInstanceStatus, describeSecurityGroups, createSecurityGroup, stopInstance，startInstance。joinSecurityGroup，执行的操作分别是，获取所有服务器信息，获取所有安全组信息，创建安全组（默认全屏蔽），停止实例，启动实例，把实例加入安全组。

有些操作如stopInstance，startInstance需要指定实例id，可以通过describeInstanceStatus获取。joinSecurityGroup除了实例id还需要安全组id，可以通过describeSecurityGroups获取或createSecurityGroup创建一个新的。

regionID是阿里云上服务器处于哪个大区的代号，比如华东2上海就是cn-shanghai，华东1杭州就是cn-hangzhou，华北2北京就是cn-beijing。accessKeyID和accessSecret是阿里云用于验证用户身份的凭据，需要在这里生成: https://ram.console.aliyun.com/users , 可以参考以下图片：

![教程](./example_1.png)
![教程](./example_2.png)
