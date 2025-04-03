<?php

namespace Grafana\Foundation\Expr;

final class ExprTypeClassicConditionsConditionsConverter
{
    public static function convert(\Grafana\Foundation\Expr\ExprTypeClassicConditionsConditions $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Expr\ExprTypeClassicConditionsConditionsBuilder())',
        ];
            
    
        {
    $buffer = 'evaluator(';
        $arg0 = \Grafana\Foundation\Expr\ExprTypeClassicConditionsConditionsEvaluatorConverter::convert($input->evaluator);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'operator(';
        $arg0 = \Grafana\Foundation\Expr\ExprTypeClassicConditionsConditionsOperatorConverter::convert($input->operator);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'query(';
        $arg0 = \Grafana\Foundation\Expr\ExprTypeClassicConditionsConditionsQueryConverter::convert($input->query);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'reducer(';
        $arg0 = \Grafana\Foundation\Expr\ExprTypeClassicConditionsConditionsReducerConverter::convert($input->reducer);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    

        return \implode("\n\t->", $calls);
    }
}

