// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;

// Position and dimensions of a panel in the grid
public class GridPos {
    // Panel height. The height is the number of rows from the top edge of the panel.
    @JsonProperty("h")
    public Integer h;
    // Panel width. The width is the number of columns from the left edge of the panel.
    @JsonProperty("w")
    public Integer w;
    // Panel x. The x coordinate is the number of columns from the left edge of the grid
    @JsonProperty("x")
    public Integer x;
    // Panel y. The y coordinate is the number of rows from the top edge of the grid
    @JsonProperty("y")
    public Integer y;
    // Whether the panel is fixed within the grid. If true, the panel will not be affected by other panels' interactions
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("static")
    public Boolean staticArg;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
