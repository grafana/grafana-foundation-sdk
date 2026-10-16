// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import java.util.LinkedList;
import java.util.HashMap;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonSetter;
import com.fasterxml.jackson.annotation.Nulls;
import java.util.List;
import java.util.Map;

/**
 * A notebook spec is a dashboard spec with narrative content. It has a title, optional description, tags, time
 * settings, and a map of elements (panels and cells) referenced by the layout.
 */
public class NotebookSpec {
    @JsonProperty("title")
    public String title;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("description")
    public String description;
    @JsonSetter(nulls = Nulls.AS_EMPTY)
    @JsonProperty("tags")
    public List<String> tags;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("timeSettings")
    public TimeSettingsSpec timeSettings;
    @JsonSetter(nulls = Nulls.AS_EMPTY)
    @JsonProperty("elements")
    public Map<String, NotebookElement> elements;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("layout")
    public NotebookLayoutKind layout;
    public NotebookSpec() {
        this.title = "";
        this.tags = new LinkedList<>();
        this.timeSettings = new com.grafana.foundation.dashboardv2beta1.TimeSettingsSpec();
        this.elements = new HashMap<>();
        this.layout = new com.grafana.foundation.dashboardv2beta1.NotebookLayoutKind();
    }
    public NotebookSpec(String title,String description,List<String> tags,TimeSettingsSpec timeSettings,Map<String, NotebookElement> elements,NotebookLayoutKind layout) {
        this.title = title;
        this.description = description;
        this.tags = tags;
        this.timeSettings = timeSettings;
        this.elements = elements;
        this.layout = layout;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
