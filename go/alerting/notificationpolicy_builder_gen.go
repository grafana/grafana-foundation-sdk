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
	resource := NewNotificationPolicy()
	builder := &NotificationPolicyBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	return builder
}

func (builder *NotificationPolicyBuilder) Build() (NotificationPolicy, error) {
	if err := builder.internal.Validate(); err != nil {
		return NotificationPolicy{}, err
	}

	return *builder.internal, nil
}

func (builder *NotificationPolicyBuilder) Continue(continueArg bool) *NotificationPolicyBuilder {
	builder.internal.Continue = &continueArg

	return builder
}

func (builder *NotificationPolicyBuilder) GroupBy(groupBy []string) *NotificationPolicyBuilder {
	builder.internal.GroupBy = groupBy

	return builder
}

func (builder *NotificationPolicyBuilder) GroupInterval(groupInterval string) *NotificationPolicyBuilder {
	builder.internal.GroupInterval = &groupInterval

	return builder
}

func (builder *NotificationPolicyBuilder) GroupWait(groupWait string) *NotificationPolicyBuilder {
	builder.internal.GroupWait = &groupWait

	return builder
}

// Deprecated. Remove before v1.0 release.
func (builder *NotificationPolicyBuilder) Match(match map[string]string) *NotificationPolicyBuilder {
	builder.internal.Match = match

	return builder
}

func (builder *NotificationPolicyBuilder) MatchRe(matchRe MatchRegexps) *NotificationPolicyBuilder {
	builder.internal.MatchRe = &matchRe

	return builder
}

// Matchers is a slice of Matchers that is sortable, implements Stringer, and
// provides a Matches method to match a LabelSet against all Matchers in the
// slice. Note that some users of Matchers might require it to be sorted.
func (builder *NotificationPolicyBuilder) Matchers(matchers Matchers) *NotificationPolicyBuilder {
	builder.internal.Matchers = &matchers

	return builder
}

func (builder *NotificationPolicyBuilder) MuteTimeIntervals(muteTimeIntervals []string) *NotificationPolicyBuilder {
	builder.internal.MuteTimeIntervals = muteTimeIntervals

	return builder
}

func (builder *NotificationPolicyBuilder) ObjectMatchers(objectMatchers ObjectMatchers) *NotificationPolicyBuilder {
	builder.internal.ObjectMatchers = &objectMatchers

	return builder
}

func (builder *NotificationPolicyBuilder) Provenance(provenance Provenance) *NotificationPolicyBuilder {
	builder.internal.Provenance = &provenance

	return builder
}

func (builder *NotificationPolicyBuilder) Receiver(receiver string) *NotificationPolicyBuilder {
	builder.internal.Receiver = &receiver

	return builder
}

func (builder *NotificationPolicyBuilder) RepeatInterval(repeatInterval string) *NotificationPolicyBuilder {
	builder.internal.RepeatInterval = &repeatInterval

	return builder
}

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
