// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.testdata;


public class StreamingQueryBuilder implements com.grafana.foundation.cog.Builder<StreamingQuery> {
    protected final StreamingQuery internal;
    
    public StreamingQueryBuilder() {
        this.internal = new StreamingQuery();
    }
    public StreamingQueryBuilder type(StreamingQueryType type) {
    this.internal.type = type;
        return this;
    }
    
    public StreamingQueryBuilder speed(Integer speed) {
    this.internal.speed = speed;
        return this;
    }
    
    public StreamingQueryBuilder spread(Integer spread) {
    this.internal.spread = spread;
        return this;
    }
    
    public StreamingQueryBuilder noise(Integer noise) {
    this.internal.noise = noise;
        return this;
    }
    
    public StreamingQueryBuilder bands(Integer bands) {
    this.internal.bands = bands;
        return this;
    }
    
    public StreamingQueryBuilder url(String url) {
    this.internal.url = url;
        return this;
    }
    public StreamingQuery build() {
        return this.internal;
    }
}
