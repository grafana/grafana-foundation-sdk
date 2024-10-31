// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as canvas from '../canvas';

export class ConnectionCoordinatesBuilder implements cog.Builder<canvas.ConnectionCoordinates> {
    protected readonly internal: canvas.ConnectionCoordinates;

    constructor() {
        this.internal = canvas.defaultConnectionCoordinates();
    }

    /**
     * Builds the object.
     */
    build(): canvas.ConnectionCoordinates {
        return this.internal;
    }

    x(x: number): this {
        this.internal.x = x;
        return this;
    }

    y(y: number): this {
        this.internal.y = y;
        return this;
    }
}
