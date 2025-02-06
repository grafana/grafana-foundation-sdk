// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.testdata;


public class CSVWaveBuilder implements com.grafana.foundation.cog.Builder<CSVWave> {
    protected final CSVWave internal;
    
    public CSVWaveBuilder() {
        this.internal = new CSVWave();
    }
    public CSVWaveBuilder labels(String labels) {
        this.internal.labels = labels;
        return this;
    }
    
    public CSVWaveBuilder name(String name) {
        this.internal.name = name;
        return this;
    }
    
    public CSVWaveBuilder timeStep(Long timeStep) {
        this.internal.timeStep = timeStep;
        return this;
    }
    
    public CSVWaveBuilder valuesCSV(String valuesCSV) {
        this.internal.valuesCSV = valuesCSV;
        return this;
    }
    public CSVWave build() {
        return this.internal;
    }
}
