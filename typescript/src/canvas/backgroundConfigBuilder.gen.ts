// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as canvas from '../canvas';
import * as common from '../common';

export class BackgroundConfigBuilder implements cog.Builder<canvas.BackgroundConfig> {
    protected readonly internal: canvas.BackgroundConfig;

    constructor() {
        this.internal = canvas.defaultBackgroundConfig();
    }

    /**
     * Builds the object.
     */
    build(): canvas.BackgroundConfig {
        return this.internal;
    }

    color(color: cog.Builder<common.ColorDimensionConfig>): this {
        const colorResource = color.build();
        this.internal.color = colorResource;
        return this;
    }

    image(image: cog.Builder<common.ResourceDimensionConfig>): this {
        const imageResource = image.build();
        this.internal.image = imageResource;
        return this;
    }

    size(size: canvas.BackgroundImageSize): this {
        this.internal.size = size;
        return this;
    }
}
