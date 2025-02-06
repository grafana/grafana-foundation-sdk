// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;


public class MovingAverageHoltModelSettingsBuilder implements com.grafana.foundation.cog.Builder<MovingAverageHoltModelSettings> {
    protected final MovingAverageHoltModelSettings internal;
    
    public MovingAverageHoltModelSettingsBuilder() {
        this.internal = new MovingAverageHoltModelSettings();
        this.internal.model = "holt";
    }
    public MovingAverageHoltModelSettingsBuilder settings(com.grafana.foundation.cog.Builder<ElasticsearchMovingAverageHoltModelSettingsSettings> settings) {
        this.internal.settings = settings.build();
        return this;
    }
    
    public MovingAverageHoltModelSettingsBuilder window(String window) {
        this.internal.window = window;
        return this;
    }
    
    public MovingAverageHoltModelSettingsBuilder minimize(Boolean minimize) {
        this.internal.minimize = minimize;
        return this;
    }
    
    public MovingAverageHoltModelSettingsBuilder predict(String predict) {
        this.internal.predict = predict;
        return this;
    }
    public MovingAverageHoltModelSettings build() {
        return this.internal;
    }
}
