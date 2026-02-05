---
title: <span class="badge object-type-class"></span> SQLExpression
---
# <span class="badge object-type-class"></span> SQLExpression

## Definition

```java
public class SQLExpression {
  public QueryEditorFunctionExpression select;
  public QueryEditorPropertyExpressionOrQueryEditorFunctionExpression from;
  public QueryEditorArrayExpression where;
  public QueryEditorArrayExpression groupBy;
  public QueryEditorFunctionExpression orderBy;
  public String orderByDirection;
  public Long limit;
}
```
## See also

 * <span class="badge builder"></span> [SQLExpressionBuilder](./builder-SQLExpressionBuilder.md)
