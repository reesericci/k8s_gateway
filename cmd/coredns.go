package main

import (
	_ "github.com/coredns/coredns/core/plugin"
	_ "github.com/ori-edge/k8s_gateway"
	_ "github.com/coredns/caddy/onevent"
	_ "github.com/coredns/coredns/plugin/acl"
	_ "github.com/coredns/coredns/plugin/any"
	_ "github.com/coredns/coredns/plugin/auto"
	_ "github.com/coredns/coredns/plugin/autopath"
	_ "github.com/coredns/coredns/plugin/azure"
	_ "github.com/coredns/coredns/plugin/bind"
	_ "github.com/coredns/coredns/plugin/bufsize"
	_ "github.com/coredns/coredns/plugin/cache"
	_ "github.com/coredns/coredns/plugin/cancel"
	_ "github.com/coredns/coredns/plugin/chaos"
	_ "github.com/coredns/coredns/plugin/clouddns"
	_ "github.com/coredns/coredns/plugin/debug"
	_ "github.com/coredns/coredns/plugin/dns64"
	_ "github.com/coredns/coredns/plugin/dnssec"
	_ "github.com/coredns/coredns/plugin/dnstap"
	_ "github.com/coredns/coredns/plugin/erratic"
	_ "github.com/coredns/coredns/plugin/errors"
	_ "github.com/coredns/coredns/plugin/etcd"
	_ "github.com/coredns/coredns/plugin/file"
	_ "github.com/coredns/coredns/plugin/forward"
	_ "github.com/coredns/coredns/plugin/grpc"
	_ "github.com/coredns/coredns/plugin/health"
	_ "github.com/coredns/coredns/plugin/hosts"
	_ "github.com/coredns/coredns/plugin/k8s_external"
	_ "github.com/coredns/coredns/plugin/kubernetes"
	_ "github.com/coredns/coredns/plugin/loadbalance"
	_ "github.com/coredns/coredns/plugin/local"
	_ "github.com/coredns/coredns/plugin/log"
	_ "github.com/coredns/coredns/plugin/loop"
	_ "github.com/coredns/coredns/plugin/metadata"
	_ "github.com/coredns/coredns/plugin/metrics"
	_ "github.com/coredns/coredns/plugin/nsid"
	_ "github.com/coredns/coredns/plugin/pprof"
	_ "github.com/coredns/coredns/plugin/ready"
	_ "github.com/coredns/coredns/plugin/reload"
	_ "github.com/coredns/coredns/plugin/rewrite"
	_ "github.com/coredns/coredns/plugin/root"
	_ "github.com/coredns/coredns/plugin/route53"
	_ "github.com/coredns/coredns/plugin/secondary"
	_ "github.com/coredns/coredns/plugin/sign"
	_ "github.com/coredns/coredns/plugin/template"
	_ "github.com/coredns/coredns/plugin/tls"
	_ "github.com/coredns/coredns/plugin/trace"
	_ "github.com/coredns/coredns/plugin/transfer"
	_ "github.com/coredns/coredns/plugin/whoami"

	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/coremain"
)

var dropPlugins = map[string]bool{
	"kubernetes":   true,
	"k8s_external": true,
}

func init() {
	var directives []string
	var alreadyAdded bool

	for _, name := range dnsserver.Directives {

		if dropPlugins[name] {
			if !alreadyAdded {
				directives = append(directives, "k8s_gateway")
				alreadyAdded = true
			}
			continue
		}
		directives = append(directives, name)
	}

	dnsserver.Directives = directives

}

func main() {
	coremain.Run()
}
