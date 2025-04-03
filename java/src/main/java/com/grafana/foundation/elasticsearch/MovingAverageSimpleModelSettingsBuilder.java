// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;


public class MovingAverageSimpleModelSettingsBuilder implements com.grafana.foundation.cog.Builder<MovingAverageSimpleModelSettings> {
    protected final MovingAverageSimpleModelSettings internal;
    
    public MovingAverageSimpleModelSettingsBuilder() {
        this.internal = new MovingAverageSimpleModelSettings();
    }
    public MovingAverageSimpleModelSettingsBuilder window(String window) {
        this.internal.window = window;
        return this;
    }
    
    public MovingAverageSimpleModelSettingsBuilder predict(String predict) {
        this.internal.predict = predict;
        return this;
    }
    public MovingAverageSimpleModelSettings build() {
        return this.internal;
    }
}
