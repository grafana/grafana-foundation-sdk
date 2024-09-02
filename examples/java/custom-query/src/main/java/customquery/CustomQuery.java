package customquery;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.grafana.foundation.cog.Builder;
import com.grafana.foundation.cog.variants.Dataquery;
import com.grafana.foundation.common.DataSourceRef;

public class CustomQuery implements Dataquery {

    @JsonProperty("refId")
    private String refId;
    @JsonProperty("hide")
    private Boolean hide;
    @JsonProperty("query")
    private String query;
    @JsonProperty("datasource")
    private DataSourceRef datasource;

    public static class CustomQueryBuilder implements Builder<Dataquery> {
        private final CustomQuery internal;

        public CustomQueryBuilder(String query) {
            this.internal = new CustomQuery();
            this.internal.query = query;
        }

        public CustomQueryBuilder refId(String refId) {
            this.internal.refId = refId;
            return this;
        }

        public CustomQueryBuilder hide(Boolean hide) {
            this.internal.hide = hide;
            return this;
        }

        public CustomQueryBuilder datasource(Builder<DataSourceRef> datasource) {
            this.internal.datasource = datasource.build();
            return this;
        }

        public CustomQuery build() {
            return this.internal;
        }

    }
}
