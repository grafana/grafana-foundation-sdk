// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;


public class OptionsWithTooltipBuilder implements com.grafana.foundation.cog.Builder<OptionsWithTooltip> {
    protected final OptionsWithTooltip internal;
    
    public OptionsWithTooltipBuilder() {
        this.internal = new OptionsWithTooltip();
    }
    public OptionsWithTooltipBuilder tooltip(com.grafana.foundation.cog.Builder<VizTooltipOptions> tooltip) {
        this.internal.tooltip = tooltip.build();
        return this;
    }
    public OptionsWithTooltip build() {
        return this.internal;
    }
}
