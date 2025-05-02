package customquery;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.grafana.foundation.cog.variants.Dataquery;

public class CustomQuery implements Dataquery {
    @JsonProperty("refId")
    public String refId;
    @JsonProperty("hide")
    public Boolean hide;
    @JsonProperty("query")
    public String query;

    @Override
    public String dataqueryName() {
        return "custom-query";
    }
}
