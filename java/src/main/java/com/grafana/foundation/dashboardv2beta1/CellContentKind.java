// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.annotation.JsonUnwrapped;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import com.fasterxml.jackson.annotation.JsonInclude;

/**
 * Pluggable cell content discriminated by `kind`. New content types are added
 * by extending this union with another <Name>CellContentKind member.
 */
@JsonDeserialize(using = CellContentKindDeserializer.class)
public class CellContentKind {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected MarkdownCellContentKind markdownCellContentKind;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected CodeCellContentKind codeCellContentKind;
    protected CellContentKind() {}
    public static CellContentKind createMarkdownCellContentKind(MarkdownCellContentKind markdownCellContentKind) {
        CellContentKind cellContentKind = new CellContentKind();
        cellContentKind.markdownCellContentKind = markdownCellContentKind;
        return cellContentKind;
    }
    public static CellContentKind createCodeCellContentKind(CodeCellContentKind codeCellContentKind) {
        CellContentKind cellContentKind = new CellContentKind();
        cellContentKind.codeCellContentKind = codeCellContentKind;
        return cellContentKind;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
