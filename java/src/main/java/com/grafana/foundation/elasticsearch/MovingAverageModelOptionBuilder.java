// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;


public class MovingAverageModelOptionBuilder implements com.grafana.foundation.cog.Builder<MovingAverageModelOption> {
    protected final MovingAverageModelOption internal;
    
    public MovingAverageModelOptionBuilder() {
        this.internal = new MovingAverageModelOption();
    }
    public MovingAverageModelOptionBuilder label(String label) {
    this.internal.label = label;
        return this;
    }
    
    public MovingAverageModelOptionBuilder value(MovingAverageModel value) {
    this.internal.value = value;
        return this;
    }
    public MovingAverageModelOption build() {
        return this.internal;
    }
}
