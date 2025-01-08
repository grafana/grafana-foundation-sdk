// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.testdata;


public class PulseWaveQueryBuilder implements com.grafana.foundation.cog.Builder<PulseWaveQuery> {
    protected final PulseWaveQuery internal;
    
    public PulseWaveQueryBuilder() {
        this.internal = new PulseWaveQuery();
    }
    public PulseWaveQueryBuilder offCount(Long offCount) {
    this.internal.offCount = offCount;
        return this;
    }
    
    public PulseWaveQueryBuilder offValue(Double offValue) {
    this.internal.offValue = offValue;
        return this;
    }
    
    public PulseWaveQueryBuilder onCount(Long onCount) {
    this.internal.onCount = onCount;
        return this;
    }
    
    public PulseWaveQueryBuilder onValue(Double onValue) {
    this.internal.onValue = onValue;
        return this;
    }
    
    public PulseWaveQueryBuilder timeStep(Long timeStep) {
    this.internal.timeStep = timeStep;
        return this;
    }
    public PulseWaveQuery build() {
        return this.internal;
    }
}
