<?php

namespace Grafana\Foundation\Cloudwatch;

final class QueryEditorOperatorExpressionConverter
{
    public static function convert(\Grafana\Foundation\Cloudwatch\QueryEditorOperatorExpression $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Cloudwatch\QueryEditorOperatorExpressionBuilder())',
        ];
            
    
        {
    $buffer = 'property(';
        $arg0 = \Grafana\Foundation\Cloudwatch\QueryEditorPropertyConverter::convert($input->property);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'operator(';
        $arg0 = \Grafana\Foundation\Cloudwatch\QueryEditorOperatorConverter::convert($input->operator);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    

        return \implode("\n\t->", $calls);
    }
}

