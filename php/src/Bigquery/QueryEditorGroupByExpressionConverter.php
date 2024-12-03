<?php

namespace Grafana\Foundation\Bigquery;

final class QueryEditorGroupByExpressionConverter
{
    public static function convert(\Grafana\Foundation\Bigquery\QueryEditorGroupByExpression $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Bigquery\QueryEditorGroupByExpressionBuilder())',
        ];
            
    
        {
    $buffer = 'property(';
        $arg0 = \Grafana\Foundation\Bigquery\QueryEditorPropertyConverter::convert($input->property);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    

        return \implode("\n\t->", $calls);
    }
}

