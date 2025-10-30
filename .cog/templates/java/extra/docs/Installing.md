# Installing

=== "Gradle"
    ```kotlin
    implementation("com.grafana:grafana-foundation-sdk:{{ if ne .Extra.ReleaseTag ""}}{{ .Extra.ReleaseTag }}{{ else }}{{ .Extra.GrafanaVersion|registryToSemver }}-{{ .Extra.BuildTimestamp }}{{ end }}")
    ```
=== "Maven"
    ```xml
    <dependency>
        <groupId>com.grafana</groupId>
        <artifactId>grafana-foundation-sdk</artifactId>
        <version>{{ .Extra.GrafanaVersion|registryToSemver }}-{{ .Extra.BuildTimestamp }}</version>
    </dependency>
    ```
