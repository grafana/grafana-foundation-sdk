// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.text;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class CodeOptions {
    // The language passed to monaco code editor
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("language")
    public CodeLanguage language;
    @JsonProperty("showLineNumbers")
    public Boolean showLineNumbers;
    @JsonProperty("showMiniMap")
    public Boolean showMiniMap;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
