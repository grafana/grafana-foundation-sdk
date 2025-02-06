// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import java.util.Map;

public class MovingAverageBuilder implements com.grafana.foundation.cog.Builder<MovingAverage> {
    protected final MovingAverage internal;
    
    public MovingAverageBuilder() {
        this.internal = new MovingAverage();
        this.internal.type = "moving_avg";
    }
    public MovingAverageBuilder pipelineAgg(String pipelineAgg) {
        this.internal.pipelineAgg = pipelineAgg;
        return this;
    }
    
    public MovingAverageBuilder field(String field) {
        this.internal.field = field;
        return this;
    }
    
    public MovingAverageBuilder id(String id) {
        this.internal.id = id;
        return this;
    }
    
    public MovingAverageBuilder settings(Map<String, Object> settings) {
        this.internal.settings = settings;
        return this;
    }
    
    public MovingAverageBuilder hide(Boolean hide) {
        this.internal.hide = hide;
        return this;
    }
    public MovingAverage build() {
        return this.internal;
    }
}
