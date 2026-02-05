import * as cog from '../cog';
import * as canvas from '../canvas';
export declare class PlacementBuilder implements cog.Builder<canvas.Placement> {
    protected readonly internal: canvas.Placement;
    constructor();
    /**
     * Builds the object.
     */
    build(): canvas.Placement;
    top(top: number): this;
    left(left: number): this;
    right(right: number): this;
    bottom(bottom: number): this;
    width(width: number): this;
    height(height: number): this;
    rotation(rotation: number): this;
}
