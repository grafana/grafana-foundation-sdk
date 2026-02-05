import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
export declare class TabsLayoutBuilder implements cog.Builder<dashboardv2beta1.TabsLayoutKind> {
    protected readonly internal: dashboardv2beta1.TabsLayoutKind;
    constructor();
    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.TabsLayoutKind;
    tabs(tabs: cog.Builder<dashboardv2beta1.TabsLayoutTabKind>[]): this;
    tab(tab: cog.Builder<dashboardv2beta1.TabsLayoutTabKind>): this;
}
