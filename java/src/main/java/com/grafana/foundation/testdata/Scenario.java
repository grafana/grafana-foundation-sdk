// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.testdata;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

// TODO: Should this live here given it's not used in the dataquery?
public class Scenario {
    @JsonProperty("id")
    public String id;
    @JsonProperty("name")
    public String name;
    @JsonProperty("stringInput")
    public String stringInput;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("description")
    public String description;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("hideAliasField")
    public Boolean hideAliasField;
    public Scenario() {
    }
    
    public Scenario(String id,String name,String stringInput,String description,Boolean hideAliasField) {
        this.id = id;
        this.name = name;
        this.stringInput = stringInput;
        this.description = description;
        this.hideAliasField = hideAliasField;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
