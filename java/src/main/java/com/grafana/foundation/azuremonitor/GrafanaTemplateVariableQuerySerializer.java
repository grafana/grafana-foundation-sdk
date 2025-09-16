// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;

import com.fasterxml.jackson.core.JsonGenerator;
import com.fasterxml.jackson.databind.JsonSerializer;
import com.fasterxml.jackson.databind.SerializerProvider;

import java.io.IOException;

public class GrafanaTemplateVariableQuerySerializer extends JsonSerializer<GrafanaTemplateVariableQuery> {

    @Override
    public void serialize(GrafanaTemplateVariableQuery value, JsonGenerator gen, SerializerProvider serializerProvider) throws IOException {
        if (value.appInsightsMetricNameQuery != null) {
            gen.writeObject(value.appInsightsMetricNameQuery);
        }
        else if (value.appInsightsGroupByQuery != null) {
            gen.writeObject(value.appInsightsGroupByQuery);
        }
        else if (value.subscriptionsQuery != null) {
            gen.writeObject(value.subscriptionsQuery);
        }
        else if (value.resourceGroupsQuery != null) {
            gen.writeObject(value.resourceGroupsQuery);
        }
        else if (value.resourceNamesQuery != null) {
            gen.writeObject(value.resourceNamesQuery);
        }
        else if (value.metricNamespaceQuery != null) {
            gen.writeObject(value.metricNamespaceQuery);
        }
        else if (value.metricDefinitionsQuery != null) {
            gen.writeObject(value.metricDefinitionsQuery);
        }
        else if (value.metricNamesQuery != null) {
            gen.writeObject(value.metricNamesQuery);
        }
        else if (value.workspacesQuery != null) {
            gen.writeObject(value.workspacesQuery);
        }
        else if (value.unknownQuery != null) {
            gen.writeObject(value.unknownQuery);
        }
    }
}
