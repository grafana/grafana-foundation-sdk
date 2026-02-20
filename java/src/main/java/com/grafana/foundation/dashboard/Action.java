// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.List;

// Dashboard action
public class Action {
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("type")
    public ActionType type;
    @JsonProperty("title")
    public String title;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("fetch")
    public FetchOptions fetch;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("infinity")
    public InfinityOptions infinity;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("confirmation")
    public String confirmation;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("oneClick")
    public Boolean oneClick;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("variables")
    public List<ActionVariable> variables;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("style")
    public DashboardActionStyle style;
    public Action() {
        this.type = ActionType.FETCH;
        this.title = "";
    }
    public Action(ActionType type,String title,FetchOptions fetch,InfinityOptions infinity,String confirmation,Boolean oneClick,List<ActionVariable> variables,DashboardActionStyle style) {
        this.type = type;
        this.title = title;
        this.fetch = fetch;
        this.infinity = infinity;
        this.confirmation = confirmation;
        this.oneClick = oneClick;
        this.variables = variables;
        this.style = style;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
