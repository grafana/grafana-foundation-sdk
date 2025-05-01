---
title: <span class="badge object-type-class"></span> SQLExpression
---
# <span class="badge object-type-class"></span> SQLExpression

## Definition

```java
public class SQLExpression {
  public List<QueryEditorFunctionExpression> columns;
  public String from;
  public String whereString;
  public List<QueryEditorGroupByExpression> groupBy;
  public QueryEditorPropertyExpression orderBy;
  public OrderByDirection orderByDirection;
  public Long limit;
  public Long offset;
}
```
## See also

 * <span class="badge builder"></span> [SQLExpressionBuilder](./builder-SQLExpressionBuilder.md)
