<?php

namespace Grafana\Foundation\Bigquery;

final class QueryEditorPropertyExpressionConverter
{
    public static function convert(\Grafana\Foundation\Bigquery\QueryEditorPropertyExpression $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Bigquery\QueryEditorPropertyExpressionBuilder())',
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

