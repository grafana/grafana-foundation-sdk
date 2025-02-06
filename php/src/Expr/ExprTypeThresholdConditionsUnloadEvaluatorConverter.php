<?php

namespace Grafana\Foundation\Expr;

final class ExprTypeThresholdConditionsUnloadEvaluatorConverter
{
    public static function convert(\Grafana\Foundation\Expr\ExprTypeThresholdConditionsUnloadEvaluator $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Expr\ExprTypeThresholdConditionsUnloadEvaluatorBuilder())',
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
            
    
        {
    $buffer = 'type(';
        $arg0 ='\Grafana\Foundation\Expr\ExprTypeThresholdConditionsUnloadEvaluatorType::fromValue("'.$input->type.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    

        return \implode("\n\t->", $calls);
    }
}

