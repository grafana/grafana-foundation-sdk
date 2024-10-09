<?php

namespace Grafana\Foundation\Expr;

final class ExprTypeClassicConditionsConditionsEvaluatorConverter
{
    public static function convert(\Grafana\Foundation\Expr\ExprTypeClassicConditionsConditionsEvaluator $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Expr\ExprTypeClassicConditionsConditionsEvaluatorBuilder())',
        ];
            if (count($input->params) >= 1) {
    
        
    $buffer = 'params(';
        $tmparg0 = [];
        foreach ($input->params as $arg1) {
        $tmpparamsarg1 =\var_export($arg1, true);
        $tmparg0[] = $tmpparamsarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->type !== "") {
    
        
    $buffer = 'type(';
        $arg0 =\var_export($input->type, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

