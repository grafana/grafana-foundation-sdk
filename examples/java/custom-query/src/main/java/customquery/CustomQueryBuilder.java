package customquery;

import com.grafana.foundation.cog.Builder;
import com.grafana.foundation.cog.variants.Dataquery;

public class CustomQueryBuilder implements Builder<Dataquery> {
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

    @Override
    public CustomQuery build() {
        return this.internal;
    }
}
