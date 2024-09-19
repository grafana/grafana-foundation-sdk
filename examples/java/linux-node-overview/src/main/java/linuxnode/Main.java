package linuxnode;

import com.fasterxml.jackson.core.JsonProcessingException;
import linuxnode.linux.DashboardBuilder;

public class Main {
    public static void main(String[] args) {
        try {
            System.out.println(DashboardBuilder.Build().toJSON());
        } catch (JsonProcessingException e) {
            throw new RuntimeException(e);
        }
    }
}
