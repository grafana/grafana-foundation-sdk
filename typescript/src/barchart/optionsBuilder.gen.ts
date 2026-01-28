// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as barchart from '../barchart';
import * as common from '../common';

export class OptionsBuilder implements cog.Builder<barchart.Options> {
    protected readonly internal: barchart.Options;

    constructor() {
        this.internal = barchart.defaultOptions();
    }

    /**
     * Builds the object.
     */
    build(): barchart.Options {
        return this.internal;
    }

    // Manually select which field from the dataset to represent the x field.
    xField(xField: string): this {
        this.internal.xField = xField;
        return this;
    }

    // Use the color value for a sibling field to color each bar value.
    colorByField(colorByField: string): this {
        this.internal.colorByField = colorByField;
        return this;
    }

    // Controls the orientation of the bar chart, either vertical or horizontal.
    orientation(orientation: common.VizOrientation): this {
        this.internal.orientation = orientation;
        return this;
    }

    // Controls the radius of each bar.
    barRadius(barRadius: number): this {
        if (!(barRadius >= 0)) {
            throw new Error("barRadius must be >= 0");
        }
        if (!(barRadius <= 0.5)) {
            throw new Error("barRadius must be <= 0.5");
        }
        this.internal.barRadius = barRadius;
        return this;
    }

    // Controls the rotation of the x axis labels.
    xTickLabelRotation(xTickLabelRotation: number): this {
        if (!(xTickLabelRotation >= -90)) {
            throw new Error("xTickLabelRotation must be >= -90");
        }
        if (!(xTickLabelRotation <= 90)) {
            throw new Error("xTickLabelRotation must be <= 90");
        }
        this.internal.xTickLabelRotation = xTickLabelRotation;
        return this;
    }

    // Sets the max length that a label can have before it is truncated.
    xTickLabelMaxLength(xTickLabelMaxLength: number): this {
        if (!(xTickLabelMaxLength >= 0)) {
            throw new Error("xTickLabelMaxLength must be >= 0");
        }
        this.internal.xTickLabelMaxLength = xTickLabelMaxLength;
        return this;
    }

    // Controls the spacing between x axis labels.
    // negative values indicate backwards skipping behavior
    xTickLabelSpacing(xTickLabelSpacing: number): this {
        this.internal.xTickLabelSpacing = xTickLabelSpacing;
        return this;
    }

    // Controls whether bars are stacked or not, either normally or in percent mode.
    stacking(stacking: common.StackingMode): this {
        this.internal.stacking = stacking;
        return this;
    }

    // This controls whether values are shown on top or to the left of bars.
    showValue(showValue: common.VisibilityMode): this {
        this.internal.showValue = showValue;
        return this;
    }

    // Controls the width of bars. 1 = Max width, 0 = Min width.
    barWidth(barWidth: number): this {
        if (!(barWidth >= 0)) {
            throw new Error("barWidth must be >= 0");
        }
        if (!(barWidth <= 1)) {
            throw new Error("barWidth must be <= 1");
        }
        this.internal.barWidth = barWidth;
        return this;
    }

    // Controls the width of groups. 1 = max with, 0 = min width.
    groupWidth(groupWidth: number): this {
        if (!(groupWidth >= 0)) {
            throw new Error("groupWidth must be >= 0");
        }
        if (!(groupWidth <= 1)) {
            throw new Error("groupWidth must be <= 1");
        }
        this.internal.groupWidth = groupWidth;
        return this;
    }

    legend(legend: cog.Builder<common.VizLegendOptions>): this {
        const legendResource = legend.build();
        this.internal.legend = legendResource;
        return this;
    }

    tooltip(tooltip: cog.Builder<common.VizTooltipOptions>): this {
        const tooltipResource = tooltip.build();
        this.internal.tooltip = tooltipResource;
        return this;
    }

    text(text: cog.Builder<common.VizTextDisplayOptions>): this {
        const textResource = text.build();
        this.internal.text = textResource;
        return this;
    }

    // Enables mode which highlights the entire bar area and shows tooltip when cursor
    // hovers over highlighted area
    fullHighlight(fullHighlight: boolean): this {
        this.internal.fullHighlight = fullHighlight;
        return this;
    }
}

