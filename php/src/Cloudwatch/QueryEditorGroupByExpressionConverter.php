<?php

namespace Grafana\Foundation\Cloudwatch;

final class QueryEditorGroupByExpressionConverter
{
    public static function convert(\Grafana\Foundation\Cloudwatch\QueryEditorGroupByExpression $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Cloudwatch\QueryEditorGroupByExpressionBuilder())',
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

