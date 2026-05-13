package doppler

import (
	"fmt"
	"path/filepath"
	"unicode"

	// Allow embedding bridge-metadata.json in the provider.
	_ "embed"

	doppler "github.com/DopplerHQ/terraform-provider-doppler/doppler" // Import the upstream provider

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	tfbridgetokens "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfgen"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"

	"github.com/nellisauction/pulumi-doppler/provider/pkg/version"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	mainPkg = "doppler"
	// modules:
	mainMod = "index" // the doppler module
)

// makeMember manufactures a type token for the package and the given module and type.
func makeMember(mod string, mem string) tokens.ModuleMember {
	return tokens.ModuleMember(mainPkg + ":" + mod + ":" + mem)
}

// makeType manufactures a type token for the package and the given module and type.
func makeType(mod string, typ string) tokens.Type {
	return tokens.Type(makeMember(mod, typ))
}

// makeResource manufactures a standard resource token given a module and resource name.  It
// automatically uses the main package and names the file by simply lower casing the resource's
// first character.
func makeResource(mod string, res string) tokens.Type {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return makeType(mod+"/"+fn, res)
}

//go:embed cmd/pulumi-resource-doppler/bridge-metadata.json
var metadata []byte

// Provider returns additional overlaid schema and metadata associated with the provider.
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := shimv2.NewProvider(doppler.Provider())

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:                 p,
		Name:              "doppler",
		Version:           version.Version,
		DisplayName:       "Doppler",
		Publisher:         "NellisAuction",
		LogoURL:           "",
		PluginDownloadURL: "github://api.github.com/nellisauction/pulumi-doppler",
		Description:       "A Pulumi package for managing Doppler secrets and configuration.",
		Keywords:          []string{"pulumi", "doppler", "category/utility"},
		License:           "Apache-2.0",
		Homepage:          "https://github.com/nellisauction/pulumi-doppler",
		Repository:        "https://github.com/nellisauction/pulumi-doppler",
		GitHubOrg:         "DopplerHQ",
		DocRules:          &tfbridge.DocRuleInfo{EditRules: docEditRules},
		Config: map[string]*tfbridge.SchemaInfo{
			"doppler_token": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"DOPPLER_TOKEN"},
				},
			},
			"host": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"DOPPLER_API_HOST"},
				},
			},
			"verify_tls": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"DOPPLER_VERIFY_TLS"},
				},
			},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			PackageName:          "@nellisauction/pulumi-doppler",
			RespectSchemaVersion: true,
		},
		Python: (func() *tfbridge.PythonInfo {
			i := &tfbridge.PythonInfo{
				RespectSchemaVersion: true,
			}
			i.PyProject.Enabled = true
			return i
		})(),
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/nellisauction/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
			RespectSchemaVersion:           true,
		},
		CSharp: &tfbridge.CSharpInfo{
			RespectSchemaVersion: true,
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
			Namespaces: map[string]string{
				mainPkg: "Doppler",
			},
		},
		MetadataInfo:                   tfbridge.NewProviderMetadata(metadata),
		EnableZeroDefaultSchemaVersion: true,
		EnableAccurateBridgePreview:    true,
	}

	prov.MustComputeTokens(tfbridgetokens.SingleModule("doppler_", mainMod,
		tfbridgetokens.MakeStandard(mainPkg)))

	prov.MustApplyAutoAliases()
	prov.SetAutonaming(255, "-")

	return prov
}

func docEditRules(defaults []tfbridge.DocsEdit) []tfbridge.DocsEdit {
	return append(
		defaults,
		skipInstallationSections...,
	)
}

var skipInstallationSections = []tfbridge.DocsEdit{
	// TF Variable do not apply to Pulumi
	{
		Path: "index.html.markdown",
		Edit: func(_ string, content []byte) ([]byte, error) {
			return tfgen.SkipSectionByHeaderContent(content, func(headerText string) bool {
				return headerText == "Terraform Variables"
			})
		},
	},
}
