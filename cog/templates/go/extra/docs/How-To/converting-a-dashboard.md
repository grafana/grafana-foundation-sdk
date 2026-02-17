# Converting a dashboard

```go
package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/grafana/grafana-foundation-sdk/go/cog/plugins"
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

func main() {
	// Required to correctly unmarshal panels and dataqueries
	plugins.RegisterDefaultPlugins()

	dashboardJSON, err := os.ReadFile("dashboard.json")
	if err != nil {
		panic(err)
	}

	dash := dashboard.Dashboard{}
	if err = json.Unmarshal(dashboardJSON, &dash); err != nil {
		panic(err)
	}

	converted := dashboard.DashboardConverter(dash)
	fmt.Println(converted)
}
```
