// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alerting

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[NotificationPolicy] = (*NotificationPolicyBuilder)(nil)

// A Route is a node that contains definitions of how to handle alerts. This is modified
// from the upstream alertmanager in that it adds the ObjectMatchers property.
type NotificationPolicyBuilder struct {
	internal *NotificationPolicy
	errors   map[string]cog.BuildErrors
}

func NewNotificationPolicyBuilder() *NotificationPolicyBuilder {
	resource := &NotificationPolicy{}
	builder := &NotificationPolicyBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *NotificationPolicyBuilder) Build() (NotificationPolicy, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("NotificationPolicy", err)...)
	}

	if len(errs) != 0 {
		return NotificationPolicy{}, errs
	}

	return *builder.internal, nil
}

// A Route is a node that contains definitions of how to handle alerts. This is modified
// from the upstream alertmanager in that it adds the ObjectMatchers property.
func (builder *NotificationPolicyBuilder) Continue(continueArg bool) *NotificationPolicyBuilder {
	builder.internal.Continue = &continueArg

	return builder
}

// A Route is a node that contains definitions of how to handle alerts. This is modified
// from the upstream alertmanager in that it adds the ObjectMatchers property.
func (builder *NotificationPolicyBuilder) GroupBy(groupBy []string) *NotificationPolicyBuilder {
	builder.internal.GroupBy = groupBy

	return builder
}

// A Route is a node that contains definitions of how to handle alerts. This is modified
// from the upstream alertmanager in that it adds the ObjectMatchers property.
func (builder *NotificationPolicyBuilder) GroupInterval(groupInterval string) *NotificationPolicyBuilder {
	builder.internal.GroupInterval = &groupInterval

	return builder
}

// A Route is a node that contains definitions of how to handle alerts. This is modified
// from the upstream alertmanager in that it adds the ObjectMatchers property.
func (builder *NotificationPolicyBuilder) GroupWait(groupWait string) *NotificationPolicyBuilder {
	builder.internal.GroupWait = &groupWait

	return builder
}

// A Route is a node that contains definitions of how to handle alerts. This is modified
// from the upstream alertmanager in that it adds the ObjectMatchers property.
func (builder *NotificationPolicyBuilder) Match(match map[string]string) *NotificationPolicyBuilder {
	builder.internal.Match = match

	return builder
}

// A Route is a node that contains definitions of how to handle alerts. This is modified
// from the upstream alertmanager in that it adds the ObjectMatchers property.
func (builder *NotificationPolicyBuilder) MatchRe(matchRe MatchRegexps) *NotificationPolicyBuilder {
	builder.internal.MatchRe = &matchRe

	return builder
}

// A Route is a node that contains definitions of how to handle alerts. This is modified
// from the upstream alertmanager in that it adds the ObjectMatchers property.
func (builder *NotificationPolicyBuilder) Matchers(matchers Matchers) *NotificationPolicyBuilder {
	builder.internal.Matchers = &matchers

	return builder
}

// A Route is a node that contains definitions of how to handle alerts. This is modified
// from the upstream alertmanager in that it adds the ObjectMatchers property.
func (builder *NotificationPolicyBuilder) MuteTimeIntervals(muteTimeIntervals []string) *NotificationPolicyBuilder {
	builder.internal.MuteTimeIntervals = muteTimeIntervals

	return builder
}

// A Route is a node that contains definitions of how to handle alerts. This is modified
// from the upstream alertmanager in that it adds the ObjectMatchers property.
func (builder *NotificationPolicyBuilder) ObjectMatchers(objectMatchers ObjectMatchers) *NotificationPolicyBuilder {
	builder.internal.ObjectMatchers = &objectMatchers

	return builder
}

// A Route is a node that contains definitions of how to handle alerts. This is modified
// from the upstream alertmanager in that it adds the ObjectMatchers property.
func (builder *NotificationPolicyBuilder) Provenance(provenance Provenance) *NotificationPolicyBuilder {
	builder.internal.Provenance = &provenance

	return builder
}

// A Route is a node that contains definitions of how to handle alerts. This is modified
// from the upstream alertmanager in that it adds the ObjectMatchers property.
func (builder *NotificationPolicyBuilder) Receiver(receiver string) *NotificationPolicyBuilder {
	builder.internal.Receiver = &receiver

	return builder
}

// A Route is a node that contains definitions of how to handle alerts. This is modified
// from the upstream alertmanager in that it adds the ObjectMatchers property.
func (builder *NotificationPolicyBuilder) RepeatInterval(repeatInterval string) *NotificationPolicyBuilder {
	builder.internal.RepeatInterval = &repeatInterval

	return builder
}

// A Route is a node that contains definitions of how to handle alerts. This is modified
// from the upstream alertmanager in that it adds the ObjectMatchers property.
func (builder *NotificationPolicyBuilder) Routes(routes []cog.Builder[NotificationPolicy]) *NotificationPolicyBuilder {
	routesResources := make([]NotificationPolicy, 0, len(routes))
	for _, r1 := range routes {
		routesDepth1, err := r1.Build()
		if err != nil {
			builder.errors["routes"] = err.(cog.BuildErrors)
			return builder
		}
		routesResources = append(routesResources, routesDepth1)
	}
	builder.internal.Routes = routesResources

	return builder
}

func (builder *NotificationPolicyBuilder) applyDefaults() {
}
