package red;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.grafana.foundation.dashboard.Dashboard;

import java.util.List;

import static red.Red.red;

public class Main {
    public static void main(String[] args) {
        Dashboard dashboard = red("RED method", List.of("sample-service", "payments", "front-gateway")).build();

        try {
            System.out.println(dashboard.toJSON());
        } catch (JsonProcessingException e) {
            throw new RuntimeException(e);
        }
    }
}
