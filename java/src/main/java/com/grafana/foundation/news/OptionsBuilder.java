// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.news;


public class OptionsBuilder implements com.grafana.foundation.cog.Builder<Options> {
    protected final Options internal;
    
    public OptionsBuilder() {
        this.internal = new Options();
    }
    public OptionsBuilder feedUrl(String feedUrl) {
        this.internal.feedUrl = feedUrl;
        return this;
    }
    
    public OptionsBuilder showImage(Boolean showImage) {
        this.internal.showImage = showImage;
        return this;
    }
    public Options build() {
        return this.internal;
    }
}
