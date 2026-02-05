import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
export declare class InfinityOptionsBuilder implements cog.Builder<dashboardv2beta1.InfinityOptions> {
    protected readonly internal: dashboardv2beta1.InfinityOptions;
    constructor();
    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.InfinityOptions;
    method(method: dashboardv2beta1.HttpRequestMethod): this;
    url(url: string): this;
    body(body: string): this;
    queryParams(queryParams: string[][]): this;
    datasourceUid(datasourceUid: string): this;
    headers(headers: string[][]): this;
}
