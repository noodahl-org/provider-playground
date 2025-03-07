package provider

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/noodahl-org/provider-playground/internal/clients"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ provider.Provider = &playgroundProvider{}
)

type playgroundProviderModel struct {
	OS types.String `tfsdk:"os"`
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
			"os": schema.StringAttribute{
				Required: true,
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

	if config.OS.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("os"),
			"Unknown OS",
			"The provider cannot setup on this system",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	//todo replace
	opsys := os.Getenv("OS")

	if !config.OS.IsNull() {
		opsys = config.OS.ValueString()
	}

	if opsys == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("os"),
			"Missing OS",
			"The provider cannot determine which operating system to deploy on",
		)
	}

	client := clients.NewCmdClient()
	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *playgroundProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewPostgresDataSource,
	}
}

func (p *playgroundProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewPostgresResource,
	}
}
