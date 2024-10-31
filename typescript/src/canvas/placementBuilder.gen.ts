// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as canvas from '../canvas';

export class PlacementBuilder implements cog.Builder<canvas.Placement> {
    protected readonly internal: canvas.Placement;

    constructor() {
        this.internal = canvas.defaultPlacement();
    }

    /**
     * Builds the object.
     */
    build(): canvas.Placement {
        return this.internal;
    }

    top(top: number): this {
        this.internal.top = top;
        return this;
    }

    left(left: number): this {
        this.internal.left = left;
        return this;
    }

    right(right: number): this {
        this.internal.right = right;
        return this;
    }

    bottom(bottom: number): this {
        this.internal.bottom = bottom;
        return this;
    }

    width(width: number): this {
        this.internal.width = width;
        return this;
    }

    height(height: number): this {
        this.internal.height = height;
        return this;
    }
}
