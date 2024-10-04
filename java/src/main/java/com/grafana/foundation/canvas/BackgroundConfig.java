// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.canvas;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.grafana.foundation.common.ColorDimensionConfig;
import com.grafana.foundation.common.ResourceDimensionConfig;

public class BackgroundConfig {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("color")
    public ColorDimensionConfig color;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("image")
    public ResourceDimensionConfig image;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("size")
    public BackgroundImageSize size;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<BackgroundConfig> {
        protected final BackgroundConfig internal;
        
        public Builder() {
            this.internal = new BackgroundConfig();
        }
    public Builder color(com.grafana.foundation.cog.Builder<ColorDimensionConfig> color) {
    this.internal.color = color.build();
        return this;
    }
    
    public Builder image(com.grafana.foundation.cog.Builder<ResourceDimensionConfig> image) {
    this.internal.image = image.build();
        return this;
    }
    
    public Builder size(BackgroundImageSize size) {
    this.internal.size = size;
        return this;
    }
    public BackgroundConfig build() {
            return this.internal;
        }
    }
}
