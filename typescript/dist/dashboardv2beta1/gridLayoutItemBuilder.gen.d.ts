import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
export declare class GridLayoutItemBuilder implements cog.Builder<dashboardv2beta1.GridLayoutItemKind> {
    protected readonly internal: dashboardv2beta1.GridLayoutItemKind;
    constructor(name: string);
    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.GridLayoutItemKind;
    x(x: number): this;
    y(y: number): this;
    width(width: number): this;
    height(height: number): this;
    repeat(repeat: cog.Builder<dashboardv2beta1.RepeatOptions>): this;
    element(name: string): this;
}
