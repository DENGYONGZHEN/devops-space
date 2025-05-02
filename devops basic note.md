# devops basic note

### 🖥️linux

1.修改hostname

```bash
sudo hostnamectl set-hostname [new-hostname]
```

2.修改IP地址

```bash
sudo vi /etc/netplan/50-cloud-init.yaml

network:
  version: 2
  renderer: networkd
  ethernets:
    ens33:
      dhcp4: no
      addresses:
        - 192.168.93.240/24
      routes:
        - to: default
          via: 192.168.93.2
      nameservers:
        addresses:
          - 114.114.114.114
          - 8.8.8.8
          - 119.29.29.29

```

```
docker run -p 5173:5173 -v /app/node_modules -v ${pwd}:/app ee44ef8e97c1 npm run dev -- --host 0.0.0.0
```



```sh
# 1. 添加远程仓库别名
git remote add origin git@github.com:DENGYONGZHEN/devops-app.git

# 2. 拉取远程分支
git fetch origin

# 3. 合并远程分支（允许无关历史）
git merge origin/master --allow-unrelated-histories

# 4. 设置上游跟踪分支
git branch --set-upstream-to=origin/master master

# 5. 推送本地代码到远程（可选）
git push
```



