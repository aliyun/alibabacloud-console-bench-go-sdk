<p align="center">
<a href=" https://www.alibabacloud.com"><img src="https://aliyunsdk-pages.alicdn.com/icons/Aliyun.svg"></a>
</p>

<h1 align="center">企业工作台 golang sdk使用</h1>

## 基本原理

在官网 SDK 的基础上，对 Client进行重写，满足企业工作台的调用逻辑，同时完全兼容官网 SDK，这样就形成了 企业工作台定制 Client + 官网 SDK 提供 APIMETA 的模式。

## 环境要求

- 找阿里云企业工作台团队，提供 OpenAPI 访问凭证(consoleKey、consoleSecret)

## 安装

使用 `go get` 下载安装 SDK:

```go
go get -u github.com/aliyun/alibabacloud-console-bench-go-sdk
```


## 快速使用

```go
package main
import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
)
// 可以放在单独文件中
func CreateClient() (CommonClient, error) {
	open := true
	if open {
		client, err := NewClientWithAccessKey(${regionId}, ${consoleKey}, ${consoleSecret})
		client.Domain = "console-bench.aliyuncs.com"
		client.Scheme = "HTTP"
		client.PathPattern = "/api/acs/openapi"
		client.Method = "GET"
		return client, err
	} else {
		client, err := sdk.NewClientWithAccessKey("REGION_ID", "ACCESS_KEY_ID", "ACCESS_KEY_SECRET")
		return client, err
	}
}

// 在具体业务中
func main() {
	client, err := CreateClient()
	if err != nil {
		panic(err)
	}

	req := requests.NewCommonRequest()
	req.Product = "Ecs"
	req.Version = "2014-05-26"
	req.ApiName = "DescribeInstances"

	req.QueryParams["RegionId"] = "cn-hangzhou"
	req.QueryParams["IdToken"] = "idToken"

	resp, err := client.ProcessCommonRequest(req)
	fmt.Println(err)
	fmt.Println(resp)
}
```

说明

- endpoint: 测试环境下需要 host 绑定 114.55.202.134 console-work.aliyuncs.com

## 许可证

[Apache-2.0](http://www.apache.org/licenses/LICENSE-2.0)

Copyright (c) 2009-present, Alibaba Cloud All rights reserved.
