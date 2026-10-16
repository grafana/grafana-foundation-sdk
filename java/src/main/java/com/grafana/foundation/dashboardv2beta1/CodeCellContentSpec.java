// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.List;

public class CodeCellContentSpec {
    @JsonProperty("language")
    public String language;
    @JsonProperty("code")
    public String code;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("highlight")
    public List<Long> highlight;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("annotation")
    public String annotation;
    public CodeCellContentSpec() {
        this.language = "";
        this.code = "";
    }
    public CodeCellContentSpec(String language,String code,List<Long> highlight,String annotation) {
        this.language = language;
        this.code = code;
        this.highlight = highlight;
        this.annotation = annotation;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
