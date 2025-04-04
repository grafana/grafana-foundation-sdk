# Building a dashboard

```java
public class Main {
    public static void main(String[] args) {
        Dashboard dashboard = new Dashboard.Builder("Sample Dashboard").
                Uid("generated-from-java").
                Tags(List.of("generated", "from", "java")).
                Refresh("1m").Time(new DashboardDashboardTime.Builder().
                        From("now-30m").
                        To("now")
                ).
                Timezone(Constants.TimeZoneBrowser).
                WithRow(new RowPanel.Builder("Overview")).
                WithPanel(new PanelBuilder().
                        Title("Network Received").
                        Unit("bps").
                        Min(0.0).
                        WithTarget(new Dataquery.Builder().
                                Expr("rate(node_network_receive_bytes_total{job=\"integrations/raspberrypi-node\", device!=\"lo\"}[$__rate_interval]) * 8").
                                LegendFormat({{ `"{{ device }}"` }})
                        )
                ).build();

        try {
            System.out.println(dashboard.toJSON());
        } catch (JsonProcessingException e) {
            e.printStackTrace();
        }
    }
}
```
