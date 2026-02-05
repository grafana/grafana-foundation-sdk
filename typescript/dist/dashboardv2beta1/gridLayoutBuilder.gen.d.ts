import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
export declare class GridLayoutBuilder implements cog.Builder<dashboardv2beta1.GridLayoutKind> {
    protected readonly internal: dashboardv2beta1.GridLayoutKind;
    constructor();
    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.GridLayoutKind;
    items(items: cog.Builder<dashboardv2beta1.GridLayoutItemKind>[]): this;
    item(item: cog.Builder<dashboardv2beta1.GridLayoutItemKind>): this;
}
