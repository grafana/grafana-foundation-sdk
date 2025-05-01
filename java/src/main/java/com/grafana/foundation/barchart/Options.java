// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.barchart;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.grafana.foundation.common.VizOrientation;
import com.grafana.foundation.common.StackingMode;
import com.grafana.foundation.common.VisibilityMode;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.grafana.foundation.common.VizLegendOptions;
import com.grafana.foundation.common.VizTooltipOptions;
import com.grafana.foundation.common.VizTextDisplayOptions;

public class Options {
    // Manually select which field from the dataset to represent the x field.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("xField")
    public String xField;
    // Use the color value for a sibling field to color each bar value.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("colorByField")
    public String colorByField;
    // Controls the orientation of the bar chart, either vertical or horizontal.
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("orientation")
    public VizOrientation orientation;
    // Controls the radius of each bar.
    @JsonInclude(JsonInclude.Include.NON_NULL)
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
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("xTickLabelSpacing")
    public Integer xTickLabelSpacing;
    // Controls whether bars are stacked or not, either normally or in percent mode.
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("stacking")
    public StackingMode stacking;
    // This controls whether values are shown on top or to the left of bars.
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("showValue")
    public VisibilityMode showValue;
    // Controls the width of bars. 1 = Max width, 0 = Min width.
    @JsonProperty("barWidth")
    public Double barWidth;
    // Controls the width of groups. 1 = max with, 0 = min width.
    @JsonProperty("groupWidth")
    public Double groupWidth;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("legend")
    public VizLegendOptions legend;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("tooltip")
    public VizTooltipOptions tooltip;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("text")
    public VizTextDisplayOptions text;
    // Enables mode which highlights the entire bar area and shows tooltip when cursor
    // hovers over highlighted area
    @JsonProperty("fullHighlight")
    public Boolean fullHighlight;
    public Options() {
        this.orientation = VizOrientation.AUTO;
        this.barRadius = 0.0;
        this.xTickLabelRotation = 0;
        this.xTickLabelMaxLength = 0;
        this.xTickLabelSpacing = 0;
        this.stacking = StackingMode.NONE;
        this.showValue = VisibilityMode.AUTO;
        this.barWidth = 1.0;
        this.groupWidth = 0.7;
        this.legend = new com.grafana.foundation.common.VizLegendOptionsBuilder().build();
        this.tooltip = new com.grafana.foundation.common.VizTooltipOptionsBuilder().build();
        this.fullHighlight = false;
    }
    public Options(String xField,String colorByField,VizOrientation orientation,Double barRadius,Integer xTickLabelRotation,Integer xTickLabelMaxLength,Integer xTickLabelSpacing,StackingMode stacking,VisibilityMode showValue,Double barWidth,Double groupWidth,VizLegendOptions legend,VizTooltipOptions tooltip,VizTextDisplayOptions text,Boolean fullHighlight) {
        this.xField = xField;
        this.colorByField = colorByField;
        this.orientation = orientation;
        this.barRadius = barRadius;
        this.xTickLabelRotation = xTickLabelRotation;
        this.xTickLabelMaxLength = xTickLabelMaxLength;
        this.xTickLabelSpacing = xTickLabelSpacing;
        this.stacking = stacking;
        this.showValue = showValue;
        this.barWidth = barWidth;
        this.groupWidth = groupWidth;
        this.legend = legend;
        this.tooltip = tooltip;
        this.text = text;
        this.fullHighlight = fullHighlight;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
