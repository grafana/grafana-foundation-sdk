// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

// TODO docs
// FROM: AnnotationQuery in grafana-data/src/types/annotations.ts
public class AnnotationQuery {
    // Name of annotation.
    @JsonProperty("name")
    public String name;
    // Datasource where the annotations data is
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("datasource")
    public DataSourceRef datasource;
    // When enabled the annotation query is issued with every dashboard refresh
    @JsonProperty("enable")
    public Boolean enable;
    // Annotation queries can be toggled on or off at the top of the dashboard.
    // When hide is true, the toggle is not shown in the dashboard.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("hide")
    public Boolean hide;
    // Color to use for the annotation event markers
    @JsonProperty("iconColor")
    public String iconColor;
    // Filters to apply when fetching annotations
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("filter")
    public AnnotationPanelFilter filter;
    // TODO.. this should just be a normal query target
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("target")
    public AnnotationTarget target;
    // TODO -- this should not exist here, it is based on the --grafana-- datasource
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("type")
    public String type;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<AnnotationQuery> {
        private final AnnotationQuery internal;
        
        public Builder() {
            this.internal = new AnnotationQuery();
        this.enable(true);
        this.hide(false);
        }
    public Builder name(String name) {
    this.internal.name = name;
        return this;
    }
    
    public Builder datasource(DataSourceRef datasource) {
    this.internal.datasource = datasource;
        return this;
    }
    
    public Builder enable(Boolean enable) {
    this.internal.enable = enable;
        return this;
    }
    
    public Builder hide(Boolean hide) {
    this.internal.hide = hide;
        return this;
    }
    
    public Builder iconColor(String iconColor) {
    this.internal.iconColor = iconColor;
        return this;
    }
    
    public Builder filter(com.grafana.foundation.cog.Builder<AnnotationPanelFilter> filter) {
    this.internal.filter = filter.build();
        return this;
    }
    
    public Builder target(com.grafana.foundation.cog.Builder<AnnotationTarget> target) {
    this.internal.target = target.build();
        return this;
    }
    
    public Builder type(String type) {
    this.internal.type = type;
        return this;
    }
    public AnnotationQuery build() {
            return this.internal;
        }
    }
}
