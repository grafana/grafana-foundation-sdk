// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;


public class DateHistogramBuilder implements com.grafana.foundation.cog.Builder<DateHistogram> {
    protected final DateHistogram internal;
    
    public DateHistogramBuilder() {
        this.internal = new DateHistogram();
    }
    public DateHistogramBuilder field(String field) {
        this.internal.field = field;
        return this;
    }
    
    public DateHistogramBuilder id(String id) {
        this.internal.id = id;
        return this;
    }
    
    public DateHistogramBuilder settings(com.grafana.foundation.cog.Builder<ElasticsearchDateHistogramSettings> settings) {
        this.internal.settings = settings.build();
        return this;
    }
    public DateHistogram build() {
        return this.internal;
    }
}
