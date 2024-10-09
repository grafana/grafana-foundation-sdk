<?php

namespace Grafana\Foundation\Cloudwatch;

final class QueryEditorFunctionExpressionConverter
{
    public static function convert(\Grafana\Foundation\Cloudwatch\QueryEditorFunctionExpression $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Cloudwatch\QueryEditorFunctionExpressionBuilder())',
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
        $tmpparametersarg1 = \Grafana\Foundation\Cloudwatch\QueryEditorFunctionParameterExpressionConverter::convert($arg1);
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

