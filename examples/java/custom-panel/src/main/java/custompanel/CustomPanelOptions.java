package custompanel;

import com.fasterxml.jackson.annotation.JsonProperty;

public class CustomPanelOptions {
    @JsonProperty("makeItBeautiful")
    private Boolean makeItBeautiful;

    public void setMakeItBeautiful(Boolean makeItBeautiful) {
        this.makeItBeautiful = makeItBeautiful;
    }
}
