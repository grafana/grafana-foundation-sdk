// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.candlestick;

import com.grafana.foundation.common.VizLegendOptions;

public class OptionsBuilder implements com.grafana.foundation.cog.Builder<Options> {
    protected final Options internal;
    
    public OptionsBuilder() {
        this.internal = new Options();
    }
    public OptionsBuilder mode(VizDisplayMode mode) {
        this.internal.mode = mode;
        return this;
    }
    
    public OptionsBuilder candleStyle(CandleStyle candleStyle) {
        this.internal.candleStyle = candleStyle;
        return this;
    }
    
    public OptionsBuilder colorStrategy(ColorStrategy colorStrategy) {
        this.internal.colorStrategy = colorStrategy;
        return this;
    }
    
    public OptionsBuilder fields(com.grafana.foundation.cog.Builder<CandlestickFieldMap> fields) {
    CandlestickFieldMap fieldsResource = fields.build();
        this.internal.fields = fieldsResource;
        return this;
    }
    
    public OptionsBuilder colors(com.grafana.foundation.cog.Builder<CandlestickColors> colors) {
    CandlestickColors colorsResource = colors.build();
        this.internal.colors = colorsResource;
        return this;
    }
    
    public OptionsBuilder legend(com.grafana.foundation.cog.Builder<VizLegendOptions> legend) {
    VizLegendOptions legendResource = legend.build();
        this.internal.legend = legendResource;
        return this;
    }
    
    public OptionsBuilder includeAllFields(Boolean includeAllFields) {
        this.internal.includeAllFields = includeAllFields;
        return this;
    }
    public Options build() {
        return this.internal;
    }
}
