import { Builder, Dataquery } from '@grafana/grafana-foundation-sdk/cog';

export interface CustomQuery {
  // refId and hide are expected on all queries
  refId?: string;
  hide?: boolean;


  // query is specific to the CustomQuery type
  query: string;

  // Let cog know that CustomQuery is a Dataquery variant
  _implementsDataqueryVariant(): void;
}

export const defaultCustomQuery = (): CustomQuery => ({
  query: "",
  _implementsDataqueryVariant() { },
});

export class CustomQueryBuilder implements Builder<Dataquery> {
  private readonly internal: CustomQuery;

  constructor(query: string) {
    this.internal = defaultCustomQuery();
    this.internal.query = query;
  }

  build(): CustomQuery {
    return this.internal;
  }

  refId(refId: string): this {
    this.internal.refId = refId;
    return this;
  }

  hide(hide: boolean): this {
    this.internal.hide = hide;
    return this;
  }
}
