# Nerdctl
> https://github.com/containerd/nerdctl

```bash
# rke2(k3s)로 설치된 containerd 사용하도록 설정
mkdir -p /etc/nerdctl
vi /etc/nerdctl/nerdctl.toml
---
# This is an example of /etc/nerdctl/nerdctl.toml .
# Unrelated to the daemon's /etc/containerd/config.toml .

debug          = false
debug_full     = false
address        = "unix:///run/k3s/containerd/containerd.sock"
namespace      = "k8s.io"
snapshotter    = "stargz"
cgroup_manager = "cgroupfs"
hosts_dir      = ["/etc/containerd/certs.d", "/etc/docker/certs.d"]
experimental   = true
---

nerdctl images
```
