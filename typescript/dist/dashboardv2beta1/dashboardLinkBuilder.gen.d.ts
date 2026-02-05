import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
export declare class DashboardLinkBuilder implements cog.Builder<dashboardv2beta1.DashboardLink> {
    protected readonly internal: dashboardv2beta1.DashboardLink;
    constructor();
    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.DashboardLink;
    title(title: string): this;
    type(type: dashboardv2beta1.DashboardLinkType): this;
    icon(icon: string): this;
    tooltip(tooltip: string): this;
    url(url: string): this;
    tags(tags: string[]): this;
    asDropdown(asDropdown: boolean): this;
    targetBlank(targetBlank: boolean): this;
    includeVars(includeVars: boolean): this;
    keepTime(keepTime: boolean): this;
    placement(placement: "inControlsMenu"): this;
}
