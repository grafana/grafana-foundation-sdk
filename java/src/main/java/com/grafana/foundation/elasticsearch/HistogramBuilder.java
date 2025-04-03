// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;


public class HistogramBuilder implements com.grafana.foundation.cog.Builder<Histogram> {
    protected final Histogram internal;
    
    public HistogramBuilder() {
        this.internal = new Histogram();
    }
    public HistogramBuilder field(String field) {
        this.internal.field = field;
        return this;
    }
    
    public HistogramBuilder id(String id) {
        this.internal.id = id;
        return this;
    }
    
    public HistogramBuilder settings(com.grafana.foundation.cog.Builder<ElasticsearchHistogramSettings> settings) {
        this.internal.settings = settings.build();
        return this;
    }
    public Histogram build() {
        return this.internal;
    }
}
