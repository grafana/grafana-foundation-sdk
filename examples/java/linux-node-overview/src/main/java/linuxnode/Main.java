package linuxnode;

import com.fasterxml.jackson.core.JsonProcessingException;
import linuxnode.linux.NodeExporter;

public class Main {
    public static void main(String[] args) {
        try {
            System.out.println(NodeExporter.Build().toJSON());
        } catch (JsonProcessingException e) {
            throw new RuntimeException(e);
        }
    }
}
