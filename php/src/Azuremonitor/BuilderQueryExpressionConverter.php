<?php

namespace Grafana\Foundation\Azuremonitor;

final class BuilderQueryExpressionConverter
{
    public static function convert(\Grafana\Foundation\Azuremonitor\BuilderQueryExpression $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Azuremonitor\BuilderQueryExpressionBuilder())',
        ];
            if ($input->from !== null) {
    
        
    $buffer = 'from(';
        $arg0 = \Grafana\Foundation\Azuremonitor\BuilderQueryEditorPropertyExpressionConverter::convert($input->from);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->columns !== null) {
    
        
    $buffer = 'columns(';
        $arg0 = \Grafana\Foundation\Azuremonitor\BuilderQueryEditorColumnsExpressionConverter::convert($input->columns);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->where !== null) {
    
        
    $buffer = 'where(';
        $arg0 = \Grafana\Foundation\Azuremonitor\BuilderQueryEditorWhereExpressionArrayConverter::convert($input->where);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->reduce !== null) {
    
        
    $buffer = 'reduce(';
        $arg0 = \Grafana\Foundation\Azuremonitor\BuilderQueryEditorReduceExpressionArrayConverter::convert($input->reduce);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->groupBy !== null) {
    
        
    $buffer = 'groupBy(';
        $arg0 = \Grafana\Foundation\Azuremonitor\BuilderQueryEditorGroupByExpressionArrayConverter::convert($input->groupBy);
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
            if ($input->orderBy !== null) {
    
        
    $buffer = 'orderBy(';
        $arg0 = \Grafana\Foundation\Azuremonitor\BuilderQueryEditorOrderByExpressionArrayConverter::convert($input->orderBy);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fuzzySearch !== null) {
    
        
    $buffer = 'fuzzySearch(';
        $arg0 = \Grafana\Foundation\Azuremonitor\BuilderQueryEditorWhereExpressionArrayConverter::convert($input->fuzzySearch);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->timeFilter !== null) {
    
        
    $buffer = 'timeFilter(';
        $arg0 = \Grafana\Foundation\Azuremonitor\BuilderQueryEditorWhereExpressionArrayConverter::convert($input->timeFilter);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

