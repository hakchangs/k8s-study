# NFS

### nfs 설치

```shell
dnf install nfs-utils -y #RHEL

systemctl enable nfs-server
mkdir -p /mnt/pv
chmod 707 /mnt/pv
chown -R 65534:65534 /mnt/pv
systemctl start nfs-server

vi /etc/exports
---
# 모든 Node 에서 접근하도록 설정
/mnt/pv 192.168.122.21(rw,sync,no_root_squash)
/mnt/pv 192.168.122.22(rw,sync,no_root_squash)
---

# 재구동 및 설정확인
systemctl restart nfs-server
exportfs -v
```

### 드라이버 설치 (csi-nfs)
> https://github.com/kubernetes-csi/csi-driver-nfs \
> https://github.com/kubernetes-csi/csi-driver-nfs/blob/master/docs/driver-parameters.md

csi-nfs 설치
```shell
curl -skSL https://raw.githubusercontent.com/kubernetes-csi/csi-driver-nfs/v4.5.0/deploy/install-driver.sh | bash -s v4.5.0 --
```

storageclass 추가
```shell
cat << EOF > nfs-sc.yml
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
  name: nfs-csi
  annotations:
    storageclass.kubernetes.io/is-default-class: "true"
provisioner: nfs.csi.k8s.io
parameters:
  server: 192.168.122.11
  share: /mnt/pv
  mountPermissions: "0777"
reclaimPolicy: Retain 
volumeBindingMode: Immediate
mountOptions:
  - nfsvers=4.1
EOF

kubectl apply -f nfs-sc.yml
```

pvc 생성 테스트 - bound 되면 성공
```shell
kubectl create -f https://raw.githubusercontent.com/kubernetes-csi/csi-driver-nfs/master/deploy/example/pvc-nfs-csi-dynamic.yaml
kubectl get pvc
```

