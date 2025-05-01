// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.testdata;

import java.util.Map;

public class SimulationQueryBuilder implements com.grafana.foundation.cog.Builder<SimulationQuery> {
    protected final SimulationQuery internal;
    
    public SimulationQueryBuilder() {
        this.internal = new SimulationQuery();
    }
    public SimulationQueryBuilder key(com.grafana.foundation.cog.Builder<Key> key) {
    Key keyResource = key.build();
        this.internal.key = keyResource;
        return this;
    }
    
    public SimulationQueryBuilder config(Map<String, Object> config) {
        this.internal.config = config;
        return this;
    }
    
    public SimulationQueryBuilder stream(Boolean stream) {
        this.internal.stream = stream;
        return this;
    }
    
    public SimulationQueryBuilder last(Boolean last) {
        this.internal.last = last;
        return this;
    }
    public SimulationQuery build() {
        return this.internal;
    }
}
