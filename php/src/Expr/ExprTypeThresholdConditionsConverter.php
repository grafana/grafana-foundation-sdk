<?php

namespace Grafana\Foundation\Expr;

final class ExprTypeThresholdConditionsConverter
{
    public static function convert(\Grafana\Foundation\Expr\ExprTypeThresholdConditions $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Expr\ExprTypeThresholdConditionsBuilder())',
        ];
            
    
        {
    $buffer = 'evaluator(';
        $arg0 = \Grafana\Foundation\Expr\ExprTypeThresholdConditionsEvaluatorConverter::convert($input->evaluator);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->loadedDimensions !== null) {
    
        
    $buffer = 'loadedDimensions(';
        $arg0 =\var_export($input->loadedDimensions, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->unloadEvaluator !== null) {
    
        
    $buffer = 'unloadEvaluator(';
        $arg0 = \Grafana\Foundation\Expr\ExprTypeThresholdConditionsUnloadEvaluatorConverter::convert($input->unloadEvaluator);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

