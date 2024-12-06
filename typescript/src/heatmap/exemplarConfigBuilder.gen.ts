// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as heatmap from '../heatmap';

// Controls exemplar options
export class ExemplarConfigBuilder implements cog.Builder<heatmap.ExemplarConfig> {
    protected readonly internal: heatmap.ExemplarConfig;

    constructor() {
        this.internal = heatmap.defaultExemplarConfig();
    }

    /**
     * Builds the object.
     */
    build(): heatmap.ExemplarConfig {
        return this.internal;
    }

    // Sets the color of the exemplar markers
    color(color: string): this {
        this.internal.color = color;
        return this;
    }
}
