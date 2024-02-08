// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	"encoding/json"
	"errors"
	"fmt"

	cogvariants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
)

func VariantConfig() cogvariants.DataqueryConfig {
	return cogvariants.DataqueryConfig{
		Identifier: "azuremonitor",
		DataqueryUnmarshaler: func(raw []byte) (cogvariants.Dataquery, error) {
			dataquery := AzureMonitorQuery{}

			if err := json.Unmarshal(raw, &dataquery); err != nil {
				return nil, err
			}

			return dataquery, nil
		},
	}
}

func (resource AppInsightsMetricNameQueryOrAppInsightsGroupByQueryOrSubscriptionsQueryOrResourceGroupsQueryOrResourceNamesQueryOrMetricNamespaceQueryOrMetricDefinitionsQueryOrMetricNamesQueryOrWorkspacesQueryOrUnknownQuery) MarshalJSON() ([]byte, error) {
	if resource.AppInsightsMetricNameQuery != nil {
		return json.Marshal(resource.AppInsightsMetricNameQuery)
	}
	if resource.AppInsightsGroupByQuery != nil {
		return json.Marshal(resource.AppInsightsGroupByQuery)
	}
	if resource.SubscriptionsQuery != nil {
		return json.Marshal(resource.SubscriptionsQuery)
	}
	if resource.ResourceGroupsQuery != nil {
		return json.Marshal(resource.ResourceGroupsQuery)
	}
	if resource.ResourceNamesQuery != nil {
		return json.Marshal(resource.ResourceNamesQuery)
	}
	if resource.MetricNamespaceQuery != nil {
		return json.Marshal(resource.MetricNamespaceQuery)
	}
	if resource.MetricDefinitionsQuery != nil {
		return json.Marshal(resource.MetricDefinitionsQuery)
	}
	if resource.MetricNamesQuery != nil {
		return json.Marshal(resource.MetricNamesQuery)
	}
	if resource.WorkspacesQuery != nil {
		return json.Marshal(resource.WorkspacesQuery)
	}
	if resource.UnknownQuery != nil {
		return json.Marshal(resource.UnknownQuery)
	}

	return nil, nil
}

func (resource *AppInsightsMetricNameQueryOrAppInsightsGroupByQueryOrSubscriptionsQueryOrResourceGroupsQueryOrResourceNamesQueryOrMetricNamespaceQueryOrMetricDefinitionsQueryOrMetricNamesQueryOrWorkspacesQueryOrUnknownQuery) UnmarshalJSON(raw []byte) error {
	if raw == nil {
		return nil
	}

	// FIXME: this is wasteful, we need to find a more efficient way to unmarshal this.
	parsedAsMap := make(map[string]any)
	if err := json.Unmarshal(raw, &parsedAsMap); err != nil {
		return err
	}

	discriminator, found := parsedAsMap["kind"]
	if !found {
		return errors.New("discriminator field 'kind' not found in payload")
	}

	switch discriminator {
	case "AppInsightsGroupByQuery":
		var appInsightsGroupByQuery AppInsightsGroupByQuery
		if err := json.Unmarshal(raw, &appInsightsGroupByQuery); err != nil {
			return err
		}

		resource.AppInsightsGroupByQuery = &appInsightsGroupByQuery
		return nil
	case "AppInsightsMetricNameQuery":
		var appInsightsMetricNameQuery AppInsightsMetricNameQuery
		if err := json.Unmarshal(raw, &appInsightsMetricNameQuery); err != nil {
			return err
		}

		resource.AppInsightsMetricNameQuery = &appInsightsMetricNameQuery
		return nil
	case "MetricDefinitionsQuery":
		var metricDefinitionsQuery MetricDefinitionsQuery
		if err := json.Unmarshal(raw, &metricDefinitionsQuery); err != nil {
			return err
		}

		resource.MetricDefinitionsQuery = &metricDefinitionsQuery
		return nil
	case "MetricNamesQuery":
		var metricNamesQuery MetricNamesQuery
		if err := json.Unmarshal(raw, &metricNamesQuery); err != nil {
			return err
		}

		resource.MetricNamesQuery = &metricNamesQuery
		return nil
	case "MetricNamespaceQuery":
		var metricNamespaceQuery MetricNamespaceQuery
		if err := json.Unmarshal(raw, &metricNamespaceQuery); err != nil {
			return err
		}

		resource.MetricNamespaceQuery = &metricNamespaceQuery
		return nil
	case "ResourceGroupsQuery":
		var resourceGroupsQuery ResourceGroupsQuery
		if err := json.Unmarshal(raw, &resourceGroupsQuery); err != nil {
			return err
		}

		resource.ResourceGroupsQuery = &resourceGroupsQuery
		return nil
	case "ResourceNamesQuery":
		var resourceNamesQuery ResourceNamesQuery
		if err := json.Unmarshal(raw, &resourceNamesQuery); err != nil {
			return err
		}

		resource.ResourceNamesQuery = &resourceNamesQuery
		return nil
	case "SubscriptionsQuery":
		var subscriptionsQuery SubscriptionsQuery
		if err := json.Unmarshal(raw, &subscriptionsQuery); err != nil {
			return err
		}

		resource.SubscriptionsQuery = &subscriptionsQuery
		return nil
	case "UnknownQuery":
		var unknownQuery UnknownQuery
		if err := json.Unmarshal(raw, &unknownQuery); err != nil {
			return err
		}

		resource.UnknownQuery = &unknownQuery
		return nil
	case "WorkspacesQuery":
		var workspacesQuery WorkspacesQuery
		if err := json.Unmarshal(raw, &workspacesQuery); err != nil {
			return err
		}

		resource.WorkspacesQuery = &workspacesQuery
		return nil
	}

	return fmt.Errorf("could not unmarshal resource with `kind = %v`", discriminator)
}
