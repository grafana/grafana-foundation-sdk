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
    this.internal.color = color.build();
        return this;
    }
    
    public BackgroundConfigBuilder image(com.grafana.foundation.cog.Builder<ResourceDimensionConfig> image) {
    this.internal.image = image.build();
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
