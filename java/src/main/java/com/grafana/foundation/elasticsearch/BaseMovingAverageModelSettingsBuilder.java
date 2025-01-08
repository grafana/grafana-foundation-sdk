// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;


public class BaseMovingAverageModelSettingsBuilder implements com.grafana.foundation.cog.Builder<BaseMovingAverageModelSettings> {
    protected final BaseMovingAverageModelSettings internal;
    
    public BaseMovingAverageModelSettingsBuilder() {
        this.internal = new BaseMovingAverageModelSettings();
    }
    public BaseMovingAverageModelSettingsBuilder model(MovingAverageModel model) {
    this.internal.model = model;
        return this;
    }
    
    public BaseMovingAverageModelSettingsBuilder window(String window) {
    this.internal.window = window;
        return this;
    }
    
    public BaseMovingAverageModelSettingsBuilder predict(String predict) {
    this.internal.predict = predict;
        return this;
    }
    public BaseMovingAverageModelSettings build() {
        return this.internal;
    }
}
