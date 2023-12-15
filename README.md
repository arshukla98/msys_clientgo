
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

## Downloading dependencies

- Make sure these are the environment variables set.

```
controlplane msys_clientgo on  main via 🐹 v1.19 ➜  echo $GOPATH, $GO111MODULE
/root/go,

controlplane msys_clientgo on  main via 🐹 v1.19 ➜  export GO111MODULE=off

controlplane msys_clientgo on  main via 🐹 v1.19 ➜  echo $GOPATH, $GO111MODULE
/root/go, off
```
- Here is the status of source folder.

```
controlplane msys_clientgo on  main via 🐹 v1.19 ➜  ls -la $GOPATH/src
total 12
drwxr-xr-x 3 root root 4096 Dec 15 02:26 .
drwxr-xr-x 5 root root 4096 Dec 15 02:26 ..
drwxr-xr-x 3 root root 4096 Dec 15 02:26 github.com
```
- Execute the commands shown below to download dependencies.

```
controlplane msys_clientgo on  main via 🐹 v1.19 ➜  go get -d ./...
package cmp: unrecognized import path "cmp": import path does not begin with hostname

controlplane msys_clientgo on  main via 🐹 v1.19 ✖ go get k8s.io/code-generator k8s.io/gengo/...
```
- Now you will notice that the dependencies are stored in source folder.

```
controlplane msys_clientgo on  main via 🐹 v1.19 ➜  ls -la $GOPATH/src
total 28
drwxr-xr-x 7 root root 4096 Dec 15 02:29 .
drwxr-xr-x 5 root root 4096 Dec 15 02:26 ..
drwxr-xr-x 7 root root 4096 Dec 15 02:29 github.com
drwxr-xr-x 3 root root 4096 Dec 15 02:29 golang.org
drwxr-xr-x 3 root root 4096 Dec 15 02:29 gopkg.in
drwxr-xr-x 7 root root 4096 Dec 15 02:29 k8s.io
drwxr-xr-x 5 root root 4096 Dec 15 02:29 sigs.k8s.io

controlplane msys_clientgo on  main via 🐹 v1.19 ➜  ls $GOPATH/src/k8s.io
apimachinery  code-generator  gengo  klog  utils

controlplane msys_clientgo on  main via 🐹 v1.19 ➜
```

## Generate Clientset, Listers, Informer and Deepcopy funcs

- Navigate to $GOPATH/src. Note the exact time.

```
controlplane msys_clientgo on  main [⇡] via 🐹 v1.19 ➜  cd $GOPATH/src

controlplane ~/go/src ➜  date
Fri 15 Dec 2023 02:48:38 AM EST
```
- Note the status of pkg directory.

```
controlplane ~/go/src ➜  ls -la /root/go/src/github.com/arshukla98/msys_clientgo/pkg
total 12
drwxr-xr-x 3 root root 4096 Dec 15 02:26 .
drwxr-xr-x 4 root root 4096 Dec 15 02:35 ..
drwxr-xr-x 3 root root 4096 Dec 15 02:26 apis
```

- Execute the following command.

```
controlplane ~/go/src ➜  bash k8s.io/code-generator/generate-groups.sh all \
> github.com/arshukla98/msys_clientgo/pkg/generated \
> github.com/arshukla98/msys_clientgo/pkg/apis monitoring:v1 \
> --go-header-file k8s.io/code-generator/examples/hack/boilerplate.go.txt
```

The First lines execute the bash file with all option that is an alias for 
"applyconfiguration,client,deepcopy,informer,lister".

The Paths in second line defines the folder for generated files
with respect to $GOPATH/src.

The Paths in Third line are relative with respect to $GOPATH/src. 
"monitoring:v1" represents to look for monitoring/v1 folder within 
github.com/arshukla98/msys_clientgo/pkg/apis.

The Fourth line depicts the location of header file with respect to $GOPATH/src.


- Here is the output.

```
WARNING: generate-groups.sh is deprecated.
WARNING: Please use k8s.io/code-generator/kube_codegen.sh instead.

WARNING: Specifying "all" as a generator is deprecated.
WARNING: Please list the specific generators needed.
WARNING: "all" is now an alias for "applyconfiguration,client,deepcopy,informer,lister"; new code generators WILL NOT be added to this set

WARNING: generate-internal-groups.sh is deprecated.
WARNING: Please use k8s.io/code-generator/kube_codegen.sh instead.

no required module provides package github.com/arshukla98/msys_clientgo/pkg/apis/monitoring/v1: go.mod file not found in current directory or any parent directory; see 'go help modules'
Generating deepcopy funcs
Generating apply configuration for monitoring:v1 at github.com/arshukla98/msys_clientgo/pkg/generated/applyconfiguration
Generating clientset for monitoring:v1 at github.com/arshukla98/msys_clientgo/pkg/generated/clientset
Generating listers for monitoring:v1 at github.com/arshukla98/msys_clientgo/pkg/generated/listers
Generating informers for monitoring:v1 at github.com/arshukla98/msys_clientgo/pkg/generated/informers
```

- Here comes the generated directory few mins ago. Inspect it.

```
controlplane ~/go/src ➜  ls -la /root/go/src/github.com/arshukla98/msys_clientgo/pkg
total 16
drwxr-xr-x 4 root root 4096 Dec 15 02:50 .
drwxr-xr-x 4 root root 4096 Dec 15 02:35 ..
drwxr-xr-x 3 root root 4096 Dec 15 02:26 apis
drwxr-xr-x 6 root root 4096 Dec 15 02:50 generated

controlplane ~/go/src ➜  ls -la /root/go/src/github.com/arshukla98/msys_clientgo/pkg/generated
total 24
drwxr-xr-x 6 root root 4096 Dec 15 02:50 .
drwxr-xr-x 4 root root 4096 Dec 15 02:50 ..
drwxr-xr-x 4 root root 4096 Dec 15 02:50 applyconfiguration
drwxr-xr-x 3 root root 4096 Dec 15 02:50 clientset
drwxr-xr-x 3 root root 4096 Dec 15 02:50 informers
drwxr-xr-x 3 root root 4096 Dec 15 02:50 listers

controlplane ~/go/src ➜  date
Fri 15 Dec 2023 02:50:29 AM EST
```

- Finally, we generated all the required files. You can see the files in git commit with title 
"Generating Clientset, Informers and Listers".


