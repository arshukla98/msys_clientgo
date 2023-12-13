
## Steps to Generate Clientsets, Informers and Listers

- This is the config file to install go. You may have to export path again.

```
controlplane ~ ➜  cat >> config.sh
sudo apt update && sudo apt install -y make build-essential

echo 'Installing Go 1.19'

sleep 2

export VERSION=1.19

curl  -L https://golang.org/dl/go${VERSION}.linux-amd64.tar.gz -o go${VERSION}.linux-amd64.tar.gz

tar -xzf go${VERSION}.linux-amd64.tar.gz -C /usr/local

export PATH=$PATH:/usr/local/go/bin

go version

^C
```

- Run config bash file.

```
controlplane ~ ➜ bash config.sh

```

- You will notice these lines confirming the fact that go is installed. Export the path and find version.

```
Installing Go 1.19
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    78  100    78    0     0   2689      0 --:--:-- --:--:-- --:--:--  2689
100    73  100    73    0     0    281      0 --:--:-- --:--:-- --:--:--   281
100  141M  100  141M    0     0   145M      0 --:--:-- --:--:-- --:--:--  276M
go version go1.19 linux/amd64

controlplane ~ ➜  export PATH=$PATH:/usr/local/go/bin

controlplane ~ ➜  go version
go version go1.19 linux/amd64
```

- Export the gopath variable and prepare the structure.

```
controlplane ~ ➜  export GOPATH=~/go

controlplane ~ ➜  mkdir -p go/bin go/pkg go/src

controlplane ~ ➜  cd go/src
```

- Create the directory as shown below. Format=github.com/GH_Username/Repo_Name

```
controlplane ~/go/src ➜  mkdir -p github.com/arshukla98/msys_clientgo

controlplane ~/go/src ➜  cd github.com/arshukla98/msys_clientgo
```

- Intialize this directory as go module.

```
controlplane github.com/arshukla98/msys_clientgo ➜  go mod init github.com/arshukla98/msys_clientgo
go: creating new go.mod: module github.com/arshukla98/msys_clientgo
```

- Create directory in the format -> pkg/apis/<group_name>/<version>

```
controlplane msys_clientgo on  master [?] via 🐹 v1.19 ➜  mkdir -p pkg/apis/monitoring/v1

controlplane msys_clientgo on  master [?] via 🐹 v1.19 ➜  ls -la
total 20
drwxr-xr-x 4 root root 4096 Dec 13 01:28 .
drwxr-xr-x 3 root root 4096 Dec 13 01:26 ..
drwxr-xr-x 7 root root 4096 Dec 13 01:28 .git
-rw-r--r-- 1 root root   52 Dec 13 01:28 go.mod
drwxr-xr-x 3 root root 4096 Dec 13 01:28 pkg
```

- Create the file register.go in this location "pkg/apis/monitoring".

```
controlplane msys_clientgo on  master [?] via 🐹 v1.19 ➜  cat > pkg/apis/monitoring/register.go
package monitoring

const (
        GroupName = "monitoring.core.example.com"
)
^C
```
This means as per go "monitoring" is the name of package "pkg/apis/monitoring".

- Now enter in v1 directory and create 3 files doc.go, types.go and register.go.

```
controlplane msys_clientgo on  master [?] via 🐹 v1.19 ➜  cat > pkg/apis/monitoring/v1/doc.go
// +k8s:deepcopy-gen=package,register

// Package v1 is the v1 version of the API.
// +groupName=monitoring.core.example.com
package v1
^C
```

This is the same groupName i am using in "pkg/apis/monitoring/register.go".
As per go "v1" is the name of package "pkg/apis/monitoring/v1".

- Please look for types.go in the following link.

types.go -> https://github.com/arshukla98/PrometheusCRD/blob/createCRD/api/v1/promcr_types.go

- In types.go, You will notice certain comments with '+'. That are actually ANNOTATIONS and VERY IMP.

- Sample of register.go is shown in the following link.

register.go -> https://github.com/kubernetes/apiextensions-apiserver/blob/v0.28.4/examples/client-go/pkg/apis/cr/v1/register.go

  You need to change at 3 points. No need to worry as you can see the changes in git version history.

