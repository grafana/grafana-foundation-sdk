// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;


public class OptionsWithTextFormattingBuilder implements com.grafana.foundation.cog.Builder<OptionsWithTextFormatting> {
    protected final OptionsWithTextFormatting internal;
    
    public OptionsWithTextFormattingBuilder() {
        this.internal = new OptionsWithTextFormatting();
    }
    public OptionsWithTextFormattingBuilder text(com.grafana.foundation.cog.Builder<VizTextDisplayOptions> text) {
    this.internal.text = text.build();
        return this;
    }
    public OptionsWithTextFormatting build() {
        return this.internal;
    }
}
