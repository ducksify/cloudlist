package inventory

import (
	"fmt"

	"github.com/ducksify/cloudlist/pkg/providers/alibaba"
	"github.com/ducksify/cloudlist/pkg/providers/arvancloud"
	"github.com/ducksify/cloudlist/pkg/providers/aws"
	"github.com/ducksify/cloudlist/pkg/providers/azure"
	"github.com/ducksify/cloudlist/pkg/providers/cloudflare"
	"github.com/ducksify/cloudlist/pkg/providers/consul"
	"github.com/ducksify/cloudlist/pkg/providers/custom"
	"github.com/ducksify/cloudlist/pkg/providers/digitalocean"
	"github.com/ducksify/cloudlist/pkg/providers/dnssimple"
	"github.com/ducksify/cloudlist/pkg/providers/fastly"
	"github.com/ducksify/cloudlist/pkg/providers/gcp"
	"github.com/ducksify/cloudlist/pkg/providers/heroku"
	"github.com/ducksify/cloudlist/pkg/providers/hetzner"
	"github.com/ducksify/cloudlist/pkg/providers/k8s"
	"github.com/ducksify/cloudlist/pkg/providers/linode"
	"github.com/ducksify/cloudlist/pkg/providers/namecheap"
	"github.com/ducksify/cloudlist/pkg/providers/nomad"
	"github.com/ducksify/cloudlist/pkg/providers/openstack"
	"github.com/ducksify/cloudlist/pkg/providers/scaleway"
	"github.com/ducksify/cloudlist/pkg/providers/terraform"
	"github.com/ducksify/cloudlist/pkg/schema"
	mapsutil "github.com/projectdiscovery/utils/maps"
)

// Inventory is an inventory of providers
type Inventory struct {
	Providers []schema.Provider
}

// New creates a new inventory of providers
func New(optionBlocks schema.Options) (*Inventory, error) {
	inventory := &Inventory{}

	for _, block := range optionBlocks {
		value, ok := block.GetMetadata("provider")
		if !ok {
			continue
		}
		provider, err := nameToProvider(value, block)
		if err != nil {
			return nil, fmt.Errorf("could not create provider %s: %s", value, err)
		}
		inventory.Providers = append(inventory.Providers, provider)
	}
	return inventory, nil
}

var Providers = map[string][]string{
	"r1c":          arvancloud.Services,
	"arvancloud":   arvancloud.Services,
	"aws":          aws.Services,
	"do":           digitalocean.Services,
	"digitalocean": digitalocean.Services,
	"gcp":          gcp.Services,
	"scw":          scaleway.Services,
	"azure":        azure.Services,
	"cloudflare":   cloudflare.Services,
	"heroku":       heroku.Services,
	"linode":       linode.Services,
	"fastly":       fastly.Services,
	"alibaba":      alibaba.Services,
	"namecheap":    namecheap.Services,
	"terraform":    terraform.Services,
	"consul":       consul.Services,
	"nomad":        nomad.Services,
	"hetzner":      hetzner.Services,
	"openstack":    openstack.Services,
	"kubernetes":   k8s.Services,
	"custom":       custom.Services,
}

func GetProviders() []string {
	return mapsutil.GetKeys(Providers)
}

func GetServices() []string {
	services := make(map[string]struct{})
	for _, s := range Providers {
		for _, service := range s {
			services[service] = struct{}{}
		}
	}
	return mapsutil.GetKeys(services)
}

// nameToProvider returns the provider for a name
func nameToProvider(value string, block schema.OptionBlock) (schema.Provider, error) {
	switch value {
	case "r1c", "arvancloud":
		return arvancloud.New(block)
	case "aws":
		return aws.New(block)
	case "do", "digitalocean":
		return digitalocean.New(block)
	case "gcp":
		return gcp.New(block)
	case "scw":
		return scaleway.New(block)
	case "azure":
		return azure.New(block)
	case "cloudflare":
		return cloudflare.New(block)
	case "heroku":
		return heroku.New(block)
	case "linode":
		return linode.New(block)
	case "fastly":
		return fastly.New(block)
	case "alibaba":
		return alibaba.New(block)
	case "namecheap":
		return namecheap.New(block)
	case "terraform":
		return terraform.New(block)
	case "consul":
		return consul.New(block)
	case "nomad":
		return nomad.New(block)
	case "hetzner":
		return hetzner.New(block)
	case "openstack":
		return openstack.New(block)
	case "kubernetes":
		return k8s.New(block)
	case "custom":
		return custom.New(block)
	case "dnssimple":
		return dnssimple.New(block)
	default:
		return nil, fmt.Errorf("invalid provider name found: %s", value)
	}
}
