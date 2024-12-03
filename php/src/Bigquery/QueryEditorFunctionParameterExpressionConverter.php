<?php

namespace Grafana\Foundation\Bigquery;

final class QueryEditorFunctionParameterExpressionConverter
{
    public static function convert(\Grafana\Foundation\Bigquery\QueryEditorFunctionParameterExpression $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Bigquery\QueryEditorFunctionParameterExpressionBuilder())',
        ];
            if ($input->name !== null && $input->name !== "") {
    
        
    $buffer = 'name(';
        $arg0 =\var_export($input->name, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

