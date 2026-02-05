import * as cog from '../cog';
import * as canvas from '../canvas';
import * as common from '../common';
export declare class CanvasConnectionBuilder implements cog.Builder<canvas.CanvasConnection> {
    protected readonly internal: canvas.CanvasConnection;
    constructor();
    /**
     * Builds the object.
     */
    build(): canvas.CanvasConnection;
    source(source: cog.Builder<canvas.ConnectionCoordinates>): this;
    target(target: cog.Builder<canvas.ConnectionCoordinates>): this;
    targetName(targetName: string): this;
    color(color: cog.Builder<common.ColorDimensionConfig>): this;
    size(size: cog.Builder<common.ScaleDimensionConfig>): this;
    vertices(vertices: cog.Builder<canvas.ConnectionCoordinates>[]): this;
    sourceOriginal(sourceOriginal: cog.Builder<canvas.ConnectionCoordinates>): this;
    targetOriginal(targetOriginal: cog.Builder<canvas.ConnectionCoordinates>): this;
}
