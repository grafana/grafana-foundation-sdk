// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.xychart;


public class XychartFieldConfigPointSizeBuilder implements com.grafana.foundation.cog.Builder<XychartFieldConfigPointSize> {
    protected final XychartFieldConfigPointSize internal;
    
    public XychartFieldConfigPointSizeBuilder() {
        this.internal = new XychartFieldConfigPointSize();
    }
    public XychartFieldConfigPointSizeBuilder fixed(Integer fixed) {
        if (!(fixed >= 0)) {
            throw new IllegalArgumentException("fixed must be >= 0");
        }
    this.internal.fixed = fixed;
        return this;
    }
    
    public XychartFieldConfigPointSizeBuilder min(Integer min) {
        if (!(min >= 0)) {
            throw new IllegalArgumentException("min must be >= 0");
        }
    this.internal.min = min;
        return this;
    }
    
    public XychartFieldConfigPointSizeBuilder max(Integer max) {
        if (!(max >= 0)) {
            throw new IllegalArgumentException("max must be >= 0");
        }
    this.internal.max = max;
        return this;
    }
    public XychartFieldConfigPointSize build() {
        return this.internal;
    }
}
