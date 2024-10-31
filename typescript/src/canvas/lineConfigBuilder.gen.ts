// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as canvas from '../canvas';
import * as common from '../common';

export class LineConfigBuilder implements cog.Builder<canvas.LineConfig> {
    protected readonly internal: canvas.LineConfig;

    constructor() {
        this.internal = canvas.defaultLineConfig();
    }

    /**
     * Builds the object.
     */
    build(): canvas.LineConfig {
        return this.internal;
    }

    color(color: cog.Builder<common.ColorDimensionConfig>): this {
        const colorResource = color.build();
        this.internal.color = colorResource;
        return this;
    }

    width(width: number): this {
        this.internal.width = width;
        return this;
    }
}
