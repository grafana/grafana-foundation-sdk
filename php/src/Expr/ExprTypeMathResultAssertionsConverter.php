<?php

namespace Grafana\Foundation\Expr;

final class ExprTypeMathResultAssertionsConverter
{
    public static function convert(\Grafana\Foundation\Expr\ExprTypeMathResultAssertions $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Expr\ExprTypeMathResultAssertionsBuilder())',
        ];
            if ($input->maxFrames !== null) {
    
        
    $buffer = 'maxFrames(';
        $arg0 =\var_export($input->maxFrames, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->type !== null) {
    
        
    $buffer = 'type(';
        $arg0 ='\Grafana\Foundation\Expr\TypeMathType::fromValue("'.$input->type.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if (count($input->typeVersion) >= 1) {
    
        
    $buffer = 'typeVersion(';
        $tmparg0 = [];
        foreach ($input->typeVersion as $arg1) {
        $tmptypeVersionarg1 =\var_export($arg1, true);
        $tmparg0[] = $tmptypeVersionarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}
