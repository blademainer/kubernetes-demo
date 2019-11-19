module github.com/blademainer/kubernetes-demo

go 1.13

require (
	github.com/bazelbuild/bazel-gazelle v0.19.1 // indirect
	github.com/bazelbuild/buildtools v0.0.0-20191118122812-e1746d65f6dd // indirect
	github.com/cespare/prettybench v0.0.0-20150116022406-03b8cfe5406c // indirect
	github.com/go-bindata/go-bindata v3.1.2+incompatible // indirect
	github.com/gogo/protobuf v1.3.1
	golang.org/x/tools v0.0.0-20191118222007-07fc4c7f2b98 // indirect
	gonum.org/v1/gonum v0.6.0 // indirect
	gotest.tools/gotestsum v0.4.0 // indirect
	k8s.io/apimachinery v0.0.0-20191115015347-3c7067801da2
	k8s.io/client-go v0.0.0-20190918160344-1fbdaa4c8d90
	k8s.io/code-generator v0.0.0-20191114215150-2a85f169f05f
	k8s.io/gengo v0.0.0-20191108084044-e500ee069b5c // indirect
	sigs.k8s.io/controller-runtime v0.4.0 // indirect
)

replace (
	golang.org/x/sys => golang.org/x/sys v0.0.0-20190813064441-fde4db37ae7a // pinned to release-branch.go1.13
	golang.org/x/tools => golang.org/x/tools v0.0.0-20190821162956-65e3620a7ae7 // pinned to release-branch.go1.13
	k8s.io/api => k8s.io/api v0.0.0-20191115135540-bbc9463b57e5
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20191115015347-3c7067801da2
	k8s.io/client-go => k8s.io/client-go v0.0.0-20191115215802-0a8a1d7b7fae
	k8s.io/code-generator => k8s.io/code-generator v0.0.0-20191114215150-2a85f169f05f
)
