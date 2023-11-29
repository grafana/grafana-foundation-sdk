// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package team

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Team] = (*TeamBuilder)(nil)

type TeamBuilder struct {
	internal *Team
	errors   map[string]cog.BuildErrors
}

func NewTeamBuilder(name string) *TeamBuilder {
	resource := &Team{}
	builder := &TeamBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()
	builder.internal.Name = name

	return builder
}

func (builder *TeamBuilder) Build() (Team, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("Team", err)...)
	}

	if len(errs) != 0 {
		return Team{}, errs
	}

	return *builder.internal, nil
}

// Name of the team.
func (builder *TeamBuilder) Name(name string) *TeamBuilder {
	builder.internal.Name = name

	return builder
}

// Email of the team.
func (builder *TeamBuilder) Email(email string) *TeamBuilder {
	builder.internal.Email = &email

	return builder
}

func (builder *TeamBuilder) applyDefaults() {
}
