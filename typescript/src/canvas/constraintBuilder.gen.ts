// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as canvas from '../canvas';

export class ConstraintBuilder implements cog.Builder<canvas.Constraint> {
    protected readonly internal: canvas.Constraint;

    constructor() {
        this.internal = canvas.defaultConstraint();
    }

    /**
     * Builds the object.
     */
    build(): canvas.Constraint {
        return this.internal;
    }

    horizontal(horizontal: canvas.HorizontalConstraint): this {
        this.internal.horizontal = horizontal;
        return this;
    }

    vertical(vertical: canvas.VerticalConstraint): this {
        this.internal.vertical = vertical;
        return this;
    }
}
