// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as common from '../common';

// TODO docs
export class LineConfigBuilder implements cog.Builder<common.LineConfig> {
    protected readonly internal: common.LineConfig;

    constructor() {
        this.internal = common.defaultLineConfig();
    }

    /**
     * Builds the object.
     */
    build(): common.LineConfig {
        return this.internal;
    }

    lineColor(lineColor: string): this {
        this.internal.lineColor = lineColor;
        return this;
    }

    lineWidth(lineWidth: number): this {
        this.internal.lineWidth = lineWidth;
        return this;
    }

    lineInterpolation(lineInterpolation: common.LineInterpolation): this {
        this.internal.lineInterpolation = lineInterpolation;
        return this;
    }

    lineStyle(lineStyle: cog.Builder<common.LineStyle>): this {
        const lineStyleResource = lineStyle.build();
        this.internal.lineStyle = lineStyleResource;
        return this;
    }

    // Indicate if null values should be treated as gaps or connected.
    // When the value is a number, it represents the maximum delta in the
    // X axis that should be considered connected.  For timeseries, this is milliseconds
    spanNulls(spanNulls: boolean | number): this {
        this.internal.spanNulls = spanNulls;
        return this;
    }
}
