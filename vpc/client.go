package vpc

import (
	"os"

	"github.com/hdksky/aliyungo/common"
)

type Client struct {
	common.Client
}

const (
	// VPCDefaultEndpoint is the default API endpoint of VPC services
	VPCDefaultEndpoint = "https://vpc.aliyuncs.com"
	VPCAPIVersion      = "2016-04-28"

	DefaultTimeout         = 60
	DefaultWaitForInterval = 5
)

// NewClient creates a new instance of ECS client
func NewClient(accessKeyId, accessKeySecret string) *Client {
	endpoint := os.Getenv("VPC_ENDPOINT")
	if endpoint == "" {
		endpoint = VPCDefaultEndpoint
	}
	return NewClientWithEndpoint(endpoint, accessKeyId, accessKeySecret)
}

func NewClientWithEndpoint(endpoint string, accessKeyId, accessKeySecret string) *Client {
	client := &Client{}
	client.Init(endpoint, VPCAPIVersion, accessKeyId, accessKeySecret)
	return client
}
