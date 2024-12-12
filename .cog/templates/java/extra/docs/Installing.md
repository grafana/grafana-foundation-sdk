# Installing

=== "Gradle"
    ```kotlin
    implementation("com.grafana:grafana-foundation-sdk:{{ .Extra.GrafanaVersion|registryToSemver }}-{{ .Extra.BuildTimestamp }}")
    ```
=== "Maven"
    ```xml
    <dependency>
        <groupId>com.grafana</groupId>
        <artifactId>grafana-foundation-sdk</artifactId>
        <version>{{ .Extra.GrafanaVersion|registryToSemver }}-{{ .Extra.BuildTimestamp }}</version>
    </dependency>
    ```
