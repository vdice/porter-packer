package packer

import "fmt"

const packerVersion = "1.4.5"
const dockerfileLines = `ENV PACKER_VERSION=%s
RUN apt-get update && apt-get install -y gpg curl unzip && \
	curl -Os https://releases.hashicorp.com/packer/${PACKER_VERSION}/packer_${PACKER_VERSION}_linux_amd64.zip && \
	unzip packer_${PACKER_VERSION}_linux_amd64.zip -d /usr/bin
`

// Build will generate the necessary Dockerfile lines
// for an invocation image using this mixin
func (m *Mixin) Build() error {
	fmt.Fprintf(m.Out, dockerfileLines, packerVersion)
	return nil
}
