<?php

namespace Grafana\Foundation\Azuremonitor;

final class BuilderQueryEditorGroupByExpressionConverter
{
    public static function convert(\Grafana\Foundation\Azuremonitor\BuilderQueryEditorGroupByExpression $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Azuremonitor\BuilderQueryEditorGroupByExpressionBuilder())',
        ];
            if ($input->property !== null) {
    
        
    $buffer = 'property(';
        $arg0 = \Grafana\Foundation\Azuremonitor\BuilderQueryEditorPropertyConverter::convert($input->property);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->interval !== null) {
    
        
    $buffer = 'interval(';
        $arg0 = \Grafana\Foundation\Azuremonitor\BuilderQueryEditorPropertyConverter::convert($input->interval);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->focus !== null) {
    
        
    $buffer = 'focus(';
        $arg0 =\var_export($input->focus, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->type !== null) {
    
        
    $buffer = 'type(';
        $arg0 ='\Grafana\Foundation\Azuremonitor\BuilderQueryEditorExpressionType::fromValue("'.$input->type.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

