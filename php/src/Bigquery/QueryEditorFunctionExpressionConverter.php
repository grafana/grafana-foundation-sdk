<?php

namespace Grafana\Foundation\Bigquery;

final class QueryEditorFunctionExpressionConverter
{
    public static function convert(\Grafana\Foundation\Bigquery\QueryEditorFunctionExpression $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Bigquery\QueryEditorFunctionExpressionBuilder())',
        ];
            if ($input->name !== null && $input->name !== "") {
    
        
    $buffer = 'name(';
        $arg0 =\var_export($input->name, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->parameters !== null && count($input->parameters) >= 1) {
    
        
    $buffer = 'parameters(';
        $tmparg0 = [];
        foreach ($input->parameters as $arg1) {
        $tmpparametersarg1 = \Grafana\Foundation\Bigquery\QueryEditorFunctionParameterExpressionConverter::convert($arg1);
        $tmparg0[] = $tmpparametersarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

