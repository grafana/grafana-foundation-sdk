import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
export declare class FetchOptionsBuilder implements cog.Builder<dashboardv2beta1.FetchOptions> {
    protected readonly internal: dashboardv2beta1.FetchOptions;
    constructor();
    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.FetchOptions;
    method(method: dashboardv2beta1.HttpRequestMethod): this;
    url(url: string): this;
    body(body: string): this;
    queryParams(queryParams: string[][]): this;
    headers(headers: string[][]): this;
}
