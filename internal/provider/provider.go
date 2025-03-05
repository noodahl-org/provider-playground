package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ provider.Provider = &playgroundProvider{}
)

type playgroundProviderModel struct {
	Region types.String `tfsdk:"region"`
}

type playgroundProvider struct {
	version string
}

// New is a helper function to simplify provider server and testing implementation.
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &playgroundProvider{
			version: version,
		}
	}
}

// Metadata returns the provider type name.
func (p *playgroundProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "provider-playground"
	resp.Version = p.version
}

// Schema defines the provider-level schema for configuration data.
func (p *playgroundProvider) Schema(_ context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"region": schema.StringAttribute{
				Optional: true,
			},
			"os": schema.StringAttribute{
				Required: true,
			},
			"aws_access_key": schema.StringAttribute{
				Optional: true,
			},
			"secret_key": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
		},
	}
}

func (p *playgroundProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var config playgroundProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	

}

func (p *playgroundProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return nil
}

func (p *playgroundProvider) Resources(_ context.Context) []func() resource.Resource {
	return nil
}
