<?php

namespace Grafana\Foundation\Canvas;

final class ConstraintConverter
{
    public static function convert(\Grafana\Foundation\Canvas\Constraint $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Canvas\ConstraintBuilder())',
        ];
            if ($input->horizontal !== null) {
    
        
    $buffer = 'horizontal(';
        $arg0 ='\Grafana\Foundation\Canvas\HorizontalConstraint::fromValue("'.$input->horizontal.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->vertical !== null) {
    
        
    $buffer = 'vertical(';
        $arg0 ='\Grafana\Foundation\Canvas\VerticalConstraint::fromValue("'.$input->vertical.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

