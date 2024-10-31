// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as xychart from '../xychart';

export class XYDimensionConfigBuilder implements cog.Builder<xychart.XYDimensionConfig> {
    protected readonly internal: xychart.XYDimensionConfig;

    constructor() {
        this.internal = xychart.defaultXYDimensionConfig();
    }

    /**
     * Builds the object.
     */
    build(): xychart.XYDimensionConfig {
        return this.internal;
    }

    frame(frame: number): this {
        if (!(frame >= 0)) {
            throw new Error("frame must be >= 0");
        }
        this.internal.frame = frame;
        return this;
    }

    x(x: string): this {
        this.internal.x = x;
        return this;
    }

    exclude(exclude: string[]): this {
        this.internal.exclude = exclude;
        return this;
    }
}
