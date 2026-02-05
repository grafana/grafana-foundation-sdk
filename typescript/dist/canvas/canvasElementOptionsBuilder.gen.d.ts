import * as cog from '../cog';
import * as canvas from '../canvas';
export declare class CanvasElementOptionsBuilder implements cog.Builder<canvas.CanvasElementOptions> {
    protected readonly internal: canvas.CanvasElementOptions;
    constructor();
    /**
     * Builds the object.
     */
    build(): canvas.CanvasElementOptions;
    name(name: string): this;
    type(type: string): this;
    config(config: any): this;
    constraint(constraint: cog.Builder<canvas.Constraint>): this;
    placement(placement: cog.Builder<canvas.Placement>): this;
    background(background: cog.Builder<canvas.BackgroundConfig>): this;
    border(border: cog.Builder<canvas.LineConfig>): this;
    connections(connections: cog.Builder<canvas.CanvasConnection>[]): this;
}
