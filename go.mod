module github.com/vdice/porter-packer/mixin/packer

go 1.13

require (
	get.porter.sh/porter v0.22.2-beta.1
	github.com/ghodss/yaml v1.0.0
	github.com/gobuffalo/packr/v2 v2.7.1
	github.com/spf13/cobra v0.0.5
	github.com/stretchr/testify v1.4.0
	github.com/vdice/porter-packer v0.0.0-20191220213930-1aad6dcec64f
	github.com/xeipuuv/gojsonschema v1.2.0
	gopkg.in/yaml.v2 v2.2.4
)

replace github.com/hashicorp/go-plugin => github.com/carolynvs/go-plugin v1.0.1-acceptstdin
