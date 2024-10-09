<?php

namespace Grafana\Foundation\Cloudwatch;

final class QueryEditorPropertyExpressionConverter
{
    public static function convert(\Grafana\Foundation\Cloudwatch\QueryEditorPropertyExpression $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Cloudwatch\QueryEditorPropertyExpressionBuilder())',
        ];
            
    
        {
    $buffer = 'property(';
        $arg0 = \Grafana\Foundation\Cloudwatch\QueryEditorPropertyConverter::convert($input->property);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    

        return \implode("\n\t->", $calls);
    }
}

