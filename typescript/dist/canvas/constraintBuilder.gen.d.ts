import * as cog from '../cog';
import * as canvas from '../canvas';
export declare class ConstraintBuilder implements cog.Builder<canvas.Constraint> {
    protected readonly internal: canvas.Constraint;
    constructor();
    /**
     * Builds the object.
     */
    build(): canvas.Constraint;
    horizontal(horizontal: canvas.HorizontalConstraint): this;
    vertical(vertical: canvas.VerticalConstraint): this;
}
