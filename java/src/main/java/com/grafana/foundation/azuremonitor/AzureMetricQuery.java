// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.List;

public class AzureMetricQuery {
    // Array of resource URIs to be queried.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("resources")
    public List<AzureMonitorResource> resources;
    // metricNamespace is used as the resource type (or resource namespace).
    // It's usually equal to the target metric namespace. e.g. microsoft.storage/storageaccounts
    // Kept the name of the variable as metricNamespace to avoid backward incompatibility issues.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("metricNamespace")
    public String metricNamespace;
    // Used as the value for the metricNamespace property when it's different from the resource namespace.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("customNamespace")
    public String customNamespace;
    // The metric to query data for within the specified metricNamespace. e.g. UsedCapacity
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("metricName")
    public String metricName;
    // The Azure region containing the resource(s).
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("region")
    public String region;
    // The granularity of data points to be queried. Defaults to auto.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("timeGrain")
    public String timeGrain;
    // The aggregation to be used within the query. Defaults to the primaryAggregationType defined by the metric.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("aggregation")
    public String aggregation;
    // Filters to reduce the set of data returned. Dimensions that can be filtered on are defined by the metric.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("dimensionFilters")
    public List<AzureMetricDimension> dimensionFilters;
    // Maximum number of records to return. Defaults to 10.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("top")
    public String top;
    // Time grains that are supported by the metric.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("allowedTimeGrainsMs")
    public List<Long> allowedTimeGrainsMs;
    // Aliases can be set to modify the legend labels. e.g. {{ resourceGroup }}. See docs for more detail.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("alias")
    public String alias;
    // @deprecated
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("timeGrainUnit")
    public String timeGrainUnit;
    // @deprecated This property was migrated to dimensionFilters and should only be accessed in the migration
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("dimension")
    public String dimension;
    // @deprecated This property was migrated to dimensionFilters and should only be accessed in the migration
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("dimensionFilter")
    public String dimensionFilter;
    // @deprecated Use metricNamespace instead
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("metricDefinition")
    public String metricDefinition;
    // @deprecated Use resourceGroup, resourceName and metricNamespace instead
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("resourceUri")
    public String resourceUri;
    // @deprecated Use resources instead
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("resourceGroup")
    public String resourceGroup;
    // @deprecated Use resources instead
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("resourceName")
    public String resourceName;
    public AzureMetricQuery() {
    }
    public AzureMetricQuery(List<AzureMonitorResource> resources,String metricNamespace,String customNamespace,String metricName,String region,String timeGrain,String aggregation,List<AzureMetricDimension> dimensionFilters,String top,List<Long> allowedTimeGrainsMs,String alias,String timeGrainUnit,String dimension,String dimensionFilter,String metricDefinition,String resourceUri,String resourceGroup,String resourceName) {
        this.resources = resources;
        this.metricNamespace = metricNamespace;
        this.customNamespace = customNamespace;
        this.metricName = metricName;
        this.region = region;
        this.timeGrain = timeGrain;
        this.aggregation = aggregation;
        this.dimensionFilters = dimensionFilters;
        this.top = top;
        this.allowedTimeGrainsMs = allowedTimeGrainsMs;
        this.alias = alias;
        this.timeGrainUnit = timeGrainUnit;
        this.dimension = dimension;
        this.dimensionFilter = dimensionFilter;
        this.metricDefinition = metricDefinition;
        this.resourceUri = resourceUri;
        this.resourceGroup = resourceGroup;
        this.resourceName = resourceName;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
