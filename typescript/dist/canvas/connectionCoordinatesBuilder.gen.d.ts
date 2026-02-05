import * as cog from '../cog';
import * as canvas from '../canvas';
export declare class ConnectionCoordinatesBuilder implements cog.Builder<canvas.ConnectionCoordinates> {
    protected readonly internal: canvas.ConnectionCoordinates;
    constructor();
    /**
     * Builds the object.
     */
    build(): canvas.ConnectionCoordinates;
    x(x: number): this;
    y(y: number): this;
}
