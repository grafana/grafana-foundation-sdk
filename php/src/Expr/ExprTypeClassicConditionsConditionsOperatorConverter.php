<?php

namespace Grafana\Foundation\Expr;

final class ExprTypeClassicConditionsConditionsOperatorConverter
{
    public static function convert(\Grafana\Foundation\Expr\ExprTypeClassicConditionsConditionsOperator $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Expr\ExprTypeClassicConditionsConditionsOperatorBuilder())',
        ];
            
    
        {
    $buffer = 'type(';
        $arg0 ='\Grafana\Foundation\Expr\TypeClassicConditionsType::fromValue("'.$input->type.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    

        return \implode("\n\t->", $calls);
    }
}

