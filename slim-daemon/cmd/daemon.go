package cmd

import (
	"github.com/cilium/cilium/slim-daemon/k8s"
)

type Daemon struct {
	watcher *k8sWatcher
}

func NewDaemon() *Daemon {
	d := &Daemon{
		watcher: NewK8sWatcher(),
	}

	return d
}

/*
func (d *Daemon) UpdateIdentities(added, deleted cache.IdentityCache) {}

func (d *Daemon) GetNodeSuffix() string { return "" }
*/

func (d *Daemon) Run() {
	client := k8s.WatcherClient()
	d.watcher.initPodWatcher(client)
}
