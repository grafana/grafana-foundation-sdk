---
title: <span class="badge object-type-class"></span> Dashboard
---
# <span class="badge object-type-class"></span> Dashboard

## Definition

```java
public class Dashboard {
  public Long id;
  public String uid;
  public String title;
  public String description;
  public Long revision;
  public String gnetId;
  public List<String> tags;
  public DashboardStyle style;
  public String timezone;
  public Boolean editable;
  public DashboardCursorSync graphTooltip;
  public DashboardDashboardTime time;
  public TimePicker timepicker;
  public Integer fiscalYearStartMonth;
  public Boolean liveNow;
  public String weekStart;
  public StringOrBool refresh;
  public Short schemaVersion;
  public Integer version;
  public List<PanelOrRowPanel> panels;
  public DashboardDashboardTemplating templating;
  public AnnotationContainer annotations;
  public List<DashboardLink> links;
  public Snapshot snapshot;
}
```
## Examples

### Unmarshalling from JSON

```java
public class Main {
    public static void main(String[] args) {
        ObjectMapper mapper = new ObjectMapper();
        try {
            InputStream json = Main.class.getResourceAsStream("/dashboard.json");
            Dashboard dashboard = mapper.readValue(json, Dashboard.class);
            System.out.println(dashboard.toJSON());
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}
```
## See also

 * <span class="badge builder"></span> [DashboardBuilder](./builder-DashboardBuilder.md)
