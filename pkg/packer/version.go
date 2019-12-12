package packer

import (
	"github.com/vdice/porter-packer/pkg"
	"get.porter.sh/porter/pkg/mixin"
	"get.porter.sh/porter/pkg/porter/version"
)

func (m *Mixin) PrintVersion(opts version.Options) error {
	metadata := mixin.Metadata{
		Name: "packer",
		VersionInfo: mixin.VersionInfo{
			Version: pkg.Version,
			Commit:  pkg.Commit,
			Author:  "Vaughn Dice",
		},
	}
	return version.PrintVersion(m.Context, opts, metadata)
}
