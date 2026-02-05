import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
export declare class TabsLayoutSpecBuilder implements cog.Builder<dashboardv2beta1.TabsLayoutSpec> {
    protected readonly internal: dashboardv2beta1.TabsLayoutSpec;
    constructor();
    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.TabsLayoutSpec;
    tabs(tabs: cog.Builder<dashboardv2beta1.TabsLayoutTabKind>[]): this;
}
