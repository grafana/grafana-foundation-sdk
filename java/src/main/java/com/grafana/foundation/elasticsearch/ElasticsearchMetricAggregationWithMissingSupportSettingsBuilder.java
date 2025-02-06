// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;


public class ElasticsearchMetricAggregationWithMissingSupportSettingsBuilder implements com.grafana.foundation.cog.Builder<ElasticsearchMetricAggregationWithMissingSupportSettings> {
    protected final ElasticsearchMetricAggregationWithMissingSupportSettings internal;
    
    public ElasticsearchMetricAggregationWithMissingSupportSettingsBuilder() {
        this.internal = new ElasticsearchMetricAggregationWithMissingSupportSettings();
    }
    public ElasticsearchMetricAggregationWithMissingSupportSettingsBuilder missing(String missing) {
        this.internal.missing = missing;
        return this;
    }
    public ElasticsearchMetricAggregationWithMissingSupportSettings build() {
        return this.internal;
    }
}
