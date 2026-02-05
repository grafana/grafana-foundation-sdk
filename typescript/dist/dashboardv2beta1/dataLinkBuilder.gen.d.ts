import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
export declare class DataLinkBuilder implements cog.Builder<dashboardv2beta1.DataLink> {
    protected readonly internal: dashboardv2beta1.DataLink;
    constructor();
    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.DataLink;
    title(title: string): this;
    url(url: string): this;
    targetBlank(targetBlank: boolean): this;
}
