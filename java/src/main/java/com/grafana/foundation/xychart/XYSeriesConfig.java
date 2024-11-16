// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.xychart;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class XYSeriesConfig {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("name")
    public XychartXYSeriesConfigName name;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("frame")
    public XychartXYSeriesConfigFrame frame;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("x")
    public XychartXYSeriesConfigX x;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("y")
    public XychartXYSeriesConfigY y;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("color")
    public XychartXYSeriesConfigColor color;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("size")
    public XychartXYSeriesConfigSize size;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<XYSeriesConfig> {
        protected final XYSeriesConfig internal;
        
        public Builder() {
            this.internal = new XYSeriesConfig();
        }
    public Builder name(com.grafana.foundation.cog.Builder<XychartXYSeriesConfigName> name) {
    this.internal.name = name.build();
        return this;
    }
    
    public Builder frame(com.grafana.foundation.cog.Builder<XychartXYSeriesConfigFrame> frame) {
    this.internal.frame = frame.build();
        return this;
    }
    
    public Builder x(com.grafana.foundation.cog.Builder<XychartXYSeriesConfigX> x) {
    this.internal.x = x.build();
        return this;
    }
    
    public Builder y(com.grafana.foundation.cog.Builder<XychartXYSeriesConfigY> y) {
    this.internal.y = y.build();
        return this;
    }
    
    public Builder color(com.grafana.foundation.cog.Builder<XychartXYSeriesConfigColor> color) {
    this.internal.color = color.build();
        return this;
    }
    
    public Builder size(com.grafana.foundation.cog.Builder<XychartXYSeriesConfigSize> size) {
    this.internal.size = size.build();
        return this;
    }
    public XYSeriesConfig build() {
            return this.internal;
        }
    }
}
