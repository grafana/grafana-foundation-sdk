// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;


public class MovingAverageLinearModelSettingsBuilder implements com.grafana.foundation.cog.Builder<MovingAverageLinearModelSettings> {
    protected final MovingAverageLinearModelSettings internal;
    
    public MovingAverageLinearModelSettingsBuilder() {
        this.internal = new MovingAverageLinearModelSettings();
    this.internal.model = "linear";
    }
    public MovingAverageLinearModelSettingsBuilder window(String window) {
    this.internal.window = window;
        return this;
    }
    
    public MovingAverageLinearModelSettingsBuilder predict(String predict) {
    this.internal.predict = predict;
        return this;
    }
    public MovingAverageLinearModelSettings build() {
        return this.internal;
    }
}
