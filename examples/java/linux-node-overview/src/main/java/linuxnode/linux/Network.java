package linuxnode.linux;

import com.grafana.foundation.timeseries.PanelBuilder;
import com.grafana.foundation.units.Constants;

public class Network {

    public static PanelBuilder networkReceivedTimeseries() {
        return Common.defaultTimeSeries().
                title("Network Received").
                description("Network received (bits/s)").
                min(0.0).
                unit(Constants.BitsPerSecondSI).
                fillOpacity(0.0).
                withTarget(
                        Common.basicPrometheusQuery("rate(node_network_receive_bytes_total{job=\"integrations/raspberrypi-node\", instance=\"$instance\", device!=\"lo\"}[$__rate_interval]) * 8", "{{ device }}")
                );
    }
    public static PanelBuilder networkTransmittedTimeseries() {
        return Common.defaultTimeSeries().
                title("Network Transmitted").
                description("Network transmitted (bits/s)").
                min(0.0).
                unit(Constants.BitsPerSecondSI).
                fillOpacity(0.0).
                withTarget(
                        Common.basicPrometheusQuery("rate(node_network_transmit_bytes_total{job=\"integrations/raspberrypi-node\", instance=\"$instance\", device!=\"lo\"}[$__rate_interval]) * 8", "{{ device }}")
                );
    }
}
