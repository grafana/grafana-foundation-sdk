// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.barchart;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;
import com.grafana.foundation.common.VizOrientation;
import com.grafana.foundation.common.StackingMode;
import com.grafana.foundation.common.VisibilityMode;
import com.grafana.foundation.common.VizLegendOptions;
import com.grafana.foundation.common.VizTooltipOptions;
import com.grafana.foundation.common.VizTextDisplayOptions;

public class Options {
    // Manually select which field from the dataset to represent the x field.
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("xField")
    public String xField;
    // Use the color value for a sibling field to color each bar value.
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("colorByField")
    public String colorByField;
    // Controls the orientation of the bar chart, either vertical or horizontal.
    @JsonProperty("orientation")
    public VizOrientation orientation;
    // Controls the radius of each bar.
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("barRadius")
    public Double barRadius;
    // Controls the rotation of the x axis labels.
    @JsonProperty("xTickLabelRotation")
    public Integer xTickLabelRotation;
    // Sets the max length that a label can have before it is truncated.
    @JsonProperty("xTickLabelMaxLength")
    public Integer xTickLabelMaxLength;
    // Controls the spacing between x axis labels.
    // negative values indicate backwards skipping behavior
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("xTickLabelSpacing")
    public Integer xTickLabelSpacing;
    // Controls whether bars are stacked or not, either normally or in percent mode.
    @JsonProperty("stacking")
    public StackingMode stacking;
    // This controls whether values are shown on top or to the left of bars.
    @JsonProperty("showValue")
    public VisibilityMode showValue;
    // Controls the width of bars. 1 = Max width, 0 = Min width.
    @JsonProperty("barWidth")
    public Double barWidth;
    // Controls the width of groups. 1 = max with, 0 = min width.
    @JsonProperty("groupWidth")
    public Double groupWidth;
    @JsonProperty("legend")
    public VizLegendOptions legend;
    @JsonProperty("tooltip")
    public VizTooltipOptions tooltip;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("text")
    public VizTextDisplayOptions text;
    // Enables mode which highlights the entire bar area and shows tooltip when cursor
    // hovers over highlighted area
    @JsonProperty("fullHighlight")
    public Boolean fullHighlight;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
