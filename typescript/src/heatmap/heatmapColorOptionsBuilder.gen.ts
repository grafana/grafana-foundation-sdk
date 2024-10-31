// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as heatmap from '../heatmap';

// Controls various color options
export class HeatmapColorOptionsBuilder implements cog.Builder<heatmap.HeatmapColorOptions> {
    protected readonly internal: heatmap.HeatmapColorOptions;

    constructor() {
        this.internal = heatmap.defaultHeatmapColorOptions();
    }

    /**
     * Builds the object.
     */
    build(): heatmap.HeatmapColorOptions {
        return this.internal;
    }

    // Sets the color mode
    mode(mode: heatmap.HeatmapColorMode): this {
        this.internal.mode = mode;
        return this;
    }

    // Controls the color scheme used
    scheme(scheme: string): this {
        this.internal.scheme = scheme;
        return this;
    }

    // Controls the color fill when in opacity mode
    fill(fill: string): this {
        this.internal.fill = fill;
        return this;
    }

    // Controls the color scale
    scale(scale: heatmap.HeatmapColorScale): this {
        this.internal.scale = scale;
        return this;
    }

    // Controls the exponent when scale is set to exponential
    exponent(exponent: number): this {
        this.internal.exponent = exponent;
        return this;
    }

    // Controls the number of color steps
    steps(steps: number): this {
        if (!(steps >= 2)) {
            throw new Error("steps must be >= 2");
        }
        if (!(steps <= 128)) {
            throw new Error("steps must be <= 128");
        }
        this.internal.steps = steps;
        return this;
    }

    // Reverses the color scheme
    reverse(reverse: boolean): this {
        this.internal.reverse = reverse;
        return this;
    }

    // Sets the minimum value for the color scale
    min(min: number): this {
        this.internal.min = min;
        return this;
    }

    // Sets the maximum value for the color scale
    max(max: number): this {
        this.internal.max = max;
        return this;
    }
}
