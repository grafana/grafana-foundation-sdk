// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.Map;

public class AnnotationQuerySpec {
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("query")
    public DataQueryKind query;
    @JsonProperty("enable")
    public Boolean enable;
    @JsonProperty("hide")
    public Boolean hide;
    @JsonProperty("iconColor")
    public String iconColor;
    @JsonProperty("name")
    public String name;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("builtIn")
    public Boolean builtIn;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("filter")
    public AnnotationPanelFilter filter;
    // Placement can be used to display the annotation query somewhere else on the dashboard other than the default location.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("placement")
    public String placement;
    // Mappings define how to convert data frame fields to annotation event fields.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("mappings")
    public Map<String, AnnotationEventFieldMapping> mappings;
    // Catch-all field for datasource-specific properties. Should not be available in as code tooling.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("legacyOptions")
    public Map<String, Object> legacyOptions;
    public AnnotationQuerySpec() {
        this.query = new com.grafana.foundation.dashboardv2beta1.DataQueryKind();
        this.enable = false;
        this.hide = false;
        this.iconColor = "";
        this.name = "";
        this.builtIn = false;
        this.placement = Constants.AnnotationQueryPlacement;
    }
    public AnnotationQuerySpec(DataQueryKind query,Boolean enable,Boolean hide,String iconColor,String name,Boolean builtIn,AnnotationPanelFilter filter,Map<String, AnnotationEventFieldMapping> mappings,Map<String, Object> legacyOptions) {
        this.query = query;
        this.enable = enable;
        this.hide = hide;
        this.iconColor = iconColor;
        this.name = name;
        this.builtIn = builtIn;
        this.filter = filter;
        this.placement = Constants.AnnotationQueryPlacement;
        this.mappings = mappings;
        this.legacyOptions = legacyOptions;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
