package network

import (
	"fmt"
	"os"

	"github.com/mojo-zd/go-library/http"
	"github.com/toolkits/net"
)

var (
	RANCHER_VERSION   = "http://rancher-metadata/latest"
	RANCHER_MANAGE_IP = fmt.Sprintf("%s/%s", RANCHER_VERSION, "self/container/primary_ip")
	RANCHER_HOST_IP   = fmt.Sprintf("%s/%s", RANCHER_VERSION, "self/host/agent_ip")
)

func GetRegistryIp() (ip string) {
	ip = internalIp()

	//if len(os.Getenv("KUBERNETES_PORT")) == 0 {
	//	ip = getRancherManageIP()
	//}
	return
}

func internalIp() string {
	ips, err := net.IntranetIP()
	if err != nil {
		fmt.Errorf("get inner ip failed, error info is %s", err.Error())
		return ""
	}
	return ips[0]
}

func getRancherManageIP() (ip string) {
	URL := RANCHER_MANAGE_IP
	if len(os.Getenv("PROFILE")) > 0 {
		URL = RANCHER_HOST_IP
	}

	response := http.NewHttpClient().BuildRequestInfo(&http.RequestInfo{URL: URL}).Get()
	if response.Error != nil {
		fmt.Printf("get rancher manager ip failed, error info is %s", response.Error.Error())
		return
	}

	ip = string(response.Result)
	fmt.Println(fmt.Sprintf("rancher manager ip is %s ", ip))
	return
}
