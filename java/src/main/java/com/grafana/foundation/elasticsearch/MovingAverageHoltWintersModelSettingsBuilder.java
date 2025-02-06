// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;


public class MovingAverageHoltWintersModelSettingsBuilder implements com.grafana.foundation.cog.Builder<MovingAverageHoltWintersModelSettings> {
    protected final MovingAverageHoltWintersModelSettings internal;
    
    public MovingAverageHoltWintersModelSettingsBuilder() {
        this.internal = new MovingAverageHoltWintersModelSettings();
        this.internal.model = "holt_winters";
    }
    public MovingAverageHoltWintersModelSettingsBuilder settings(com.grafana.foundation.cog.Builder<ElasticsearchMovingAverageHoltWintersModelSettingsSettings> settings) {
        this.internal.settings = settings.build();
        return this;
    }
    
    public MovingAverageHoltWintersModelSettingsBuilder window(String window) {
        this.internal.window = window;
        return this;
    }
    
    public MovingAverageHoltWintersModelSettingsBuilder minimize(Boolean minimize) {
        this.internal.minimize = minimize;
        return this;
    }
    
    public MovingAverageHoltWintersModelSettingsBuilder predict(String predict) {
        this.internal.predict = predict;
        return this;
    }
    public MovingAverageHoltWintersModelSettings build() {
        return this.internal;
    }
}
