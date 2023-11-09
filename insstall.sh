# 关闭防火墙
systemctl stop firewalld
systemctl disable firewalld

# 关闭selinux
sed -i 's/enforcing/disabled/' /etc/selinux/config  # 永久
setenforce 0  # 临时

# 关闭swap
swapoff -a  # 临时
sed -ri 's/.*swap.*/#&/' /etc/fstab    # 永久

# 关闭完swap后，一定要重启一下虚拟机！！！
# 根据规划设置主机名
hostnamectl set-hostname k8s-master

# 在master添加hosts
cat >> /etc/hosts << EOF
172.17.250.217 k8s-master
172.17.250.216 k8s-node1
172.17.250.215 k8s-node2
EOF


# 将桥接的IPv4流量传递到iptables的链
cat > /etc/sysctl.d/k8s.conf << EOF
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
EOF

sysctl --system  # 生效


# 时间同步
yum install ntpdate -y
ntpdate time.windows.com


kubeadm init \
      --apiserver-advertise-address=172.17.250.217 \
      --image-repository registry.aliyuncs.com/google_containers \
      --kubernetes-version v1.23.6 \
      --service-cidr=10.96.0.0/12 \
      --pod-network-cidr=10.244.0.0/16

kubeadm join 192.168.113.120:6443 --token xwyg2e.rm67h7tx0qx87s9p --discovery-token-ca-cert-hash sha256:a89d7c3d69087e221386a8f595171d0b6bc1f3c977264031fe6cb9d71c2057ac