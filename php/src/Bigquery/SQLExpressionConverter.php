<?php

namespace Grafana\Foundation\Bigquery;

final class SQLExpressionConverter
{
    public static function convert(\Grafana\Foundation\Bigquery\SQLExpression $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Bigquery\SQLExpressionBuilder())',
        ];
            if ($input->columns !== null && count($input->columns) >= 1) {
    
        
    $buffer = 'columns(';
        $tmparg0 = [];
        foreach ($input->columns as $arg1) {
        $tmpcolumnsarg1 = \Grafana\Foundation\Bigquery\QueryEditorFunctionExpressionConverter::convert($arg1);
        $tmparg0[] = $tmpcolumnsarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->from !== null && $input->from !== "") {
    
        
    $buffer = 'from(';
        $arg0 =\var_export($input->from, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->whereString !== null && $input->whereString !== "") {
    
        
    $buffer = 'whereString(';
        $arg0 =\var_export($input->whereString, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->groupBy !== null && count($input->groupBy) >= 1) {
    
        
    $buffer = 'groupBy(';
        $tmparg0 = [];
        foreach ($input->groupBy as $arg1) {
        $tmpgroupByarg1 = \Grafana\Foundation\Bigquery\QueryEditorGroupByExpressionConverter::convert($arg1);
        $tmparg0[] = $tmpgroupByarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->orderBy !== null) {
    
        
    $buffer = 'orderBy(';
        $arg0 = \Grafana\Foundation\Bigquery\QueryEditorPropertyExpressionConverter::convert($input->orderBy);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->orderByDirection !== null) {
    
        
    $buffer = 'orderByDirection(';
        $arg0 ='\Grafana\Foundation\Bigquery\OrderByDirection::fromValue("'.$input->orderByDirection.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->limit !== null) {
    
        
    $buffer = 'limit(';
        $arg0 =\var_export($input->limit, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->offset !== null) {
    
        
    $buffer = 'offset(';
        $arg0 =\var_export($input->offset, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

