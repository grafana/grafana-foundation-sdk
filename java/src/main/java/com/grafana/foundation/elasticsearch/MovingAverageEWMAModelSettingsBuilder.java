// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;


public class MovingAverageEWMAModelSettingsBuilder implements com.grafana.foundation.cog.Builder<MovingAverageEWMAModelSettings> {
    protected final MovingAverageEWMAModelSettings internal;
    
    public MovingAverageEWMAModelSettingsBuilder() {
        this.internal = new MovingAverageEWMAModelSettings();
        this.internal.model = "ewma";
    }
    public MovingAverageEWMAModelSettingsBuilder settings(com.grafana.foundation.cog.Builder<ElasticsearchMovingAverageEWMAModelSettingsSettings> settings) {
        this.internal.settings = settings.build();
        return this;
    }
    
    public MovingAverageEWMAModelSettingsBuilder window(String window) {
        this.internal.window = window;
        return this;
    }
    
    public MovingAverageEWMAModelSettingsBuilder minimize(Boolean minimize) {
        this.internal.minimize = minimize;
        return this;
    }
    
    public MovingAverageEWMAModelSettingsBuilder predict(String predict) {
        this.internal.predict = predict;
        return this;
    }
    public MovingAverageEWMAModelSettings build() {
        return this.internal;
    }
}
