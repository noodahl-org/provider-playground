package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/noodahl-org/provider-playground/internal/clients"
)

type postgresDataSourceModel struct {
	Status string `tfsdk:"status"`
}
type postgresDataSource struct {
	cmd *clients.CmdClient
}

func NewPostgresDataSource() datasource.DataSource {
	return &postgresDataSource{}
}

func (d *postgresDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_postgres"
}

func (d *postgresDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"status": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (d *postgresDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state postgresDataSourceModel
	result, err := d.cmd.Command(ctx, "sh", []string{"-c", "brew services list | grep postgresql | awk '{print $2}'"})
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to determine postgres status",
			err.Error(),
		)
		return
	}

	state.Status = string(result)
	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

}

func (d *postgresDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client, ok := req.ProviderData.(*clients.CmdClient)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *CmdClient, go: %T.", req.ProviderData),
		)
		return
	}
	d.cmd = client
}
