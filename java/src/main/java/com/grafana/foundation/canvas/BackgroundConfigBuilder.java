// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.canvas;

import com.grafana.foundation.common.ColorDimensionConfig;
import com.grafana.foundation.common.ResourceDimensionConfig;

public class BackgroundConfigBuilder implements com.grafana.foundation.cog.Builder<BackgroundConfig> {
    protected final BackgroundConfig internal;
    
    public BackgroundConfigBuilder() {
        this.internal = new BackgroundConfig();
    }
    public BackgroundConfigBuilder color(com.grafana.foundation.cog.Builder<ColorDimensionConfig> color) {
    ColorDimensionConfig colorResource = color.build();
        this.internal.color = colorResource;
        return this;
    }
    
    public BackgroundConfigBuilder image(com.grafana.foundation.cog.Builder<ResourceDimensionConfig> image) {
    ResourceDimensionConfig imageResource = image.build();
        this.internal.image = imageResource;
        return this;
    }
    
    public BackgroundConfigBuilder size(BackgroundImageSize size) {
        this.internal.size = size;
        return this;
    }
    public BackgroundConfig build() {
        return this.internal;
    }
}
