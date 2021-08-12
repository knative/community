module knative.dev/community/mechanics/tools/gen-aliases

go 1.16

require (
	k8s.io/apimachinery v0.21.1
	k8s.io/test-infra v0.0.0-20210812083547-0e507b656399
	sigs.k8s.io/yaml v1.2.0
)

// k8s.io/test-infra requires old tektoncd/pipeline that'd pull old knative.dev/pkg
replace github.com/tektoncd/pipeline => github.com/tektoncd/pipeline v0.26.1-0.20210811222006-76cb481f504c
