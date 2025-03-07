package provider

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/noodahl-org/provider-playground/internal/clients"
)

type postgresResourceModel struct {
	ID      types.String `tfsdk:"id"`
	Version types.String `tfsdk:"version"`
	Status  types.String `tfsdk:"status"`
}
type postgresResource struct {
	cmd clients.CmdClient
}

func NewPostgresResource() resource.Resource {
	return &postgresResource{}
}

func (r *postgresResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_postgres"
}

func (r *postgresResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"status": schema.StringAttribute{
				Computed: true,
			},
			"version": schema.StringAttribute{
				Required: true,
			},
		},
	}
}

func (p *postgresResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan postgresResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	//todo fix version param
	_, err := p.cmd.Command(ctx, "brew", []string{"install", fmt.Sprintf("postgresql@%s", plan.Version.ValueString())})
	if err != nil {
		resp.Diagnostics.AddError(
			"Error running brew install",
			err.Error(),
		)
		return
	}

	serviceResp, err := p.cmd.Command(ctx, "brew", []string{"services", "start", fmt.Sprintf("postgresql@%s", plan.Version.ValueString())})
	if err != nil {
		resp.Diagnostics.AddError(
			"Error running brew install",
			err.Error(),
		)
		return
	}

	status, err := p.cmd.Command(ctx, "sh", []string{"-c", "brew services list | grep postgresql | awk '{print $2}'"})
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to determine postgres status",
			err.Error(),
		)
		return
	}

	hash := sha256.Sum256([]byte(serviceResp))
	plan.ID = types.StringValue(base64.StdEncoding.EncodeToString(hash[:]))
	plan.Status = types.StringValue(status)

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

}

func (r *postgresResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan postgresResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.cmd.Command(ctx, "brew", []string{"services", "stop", "postgresql@" + plan.Version.ValueString()})
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to stop postgresql service via homebrew",
			err.Error(),
		)
	}

	_, err = r.cmd.Command(ctx, "brew", []string{"reinstall", "postgresql@" + plan.Version.ValueString()})
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to stop postgresql service via homebrew",
			err.Error(),
		)
	}

	_, err = r.cmd.Command(ctx, "brew", []string{"services", "start", "postgresql@" + plan.Version.ValueString()})
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to stop postgresql service via homebrew",
			err.Error(),
		)
	}
}

func (r *postgresResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state postgresResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.cmd.Command(ctx, "brew", []string{"services", "stop", "postgresql@" + state.Version.ValueString()})
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to stop postgresql service via homebrew",
			err.Error(),
		)
	}

	_, err = r.cmd.Command(ctx, "brew", []string{"uninstall", "postgresql@" + state.Version.ValueString()})
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to stop postgresql service via homebrew",
			err.Error(),
		)
	}

	result, err := r.cmd.Command(ctx, "sh", []string{"-c", "brew services list | grep postgresql | awk '{print $2}'"})
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to determine postgres status",
			err.Error(),
		)
		return
	}

	state.Status = types.StringValue(result)
}

func (r *postgresResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state postgresResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	result, err := r.cmd.Command(ctx, "sh", []string{"-c", "brew services list | grep postgresql | awk '{print $2}'"})
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to determine postgres status",
			err.Error(),
		)
		return
	}

	state.Status = types.StringValue(result)
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
