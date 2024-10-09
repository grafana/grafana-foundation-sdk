<?php

namespace Grafana\Foundation\Canvas;

final class CanvasElementOptionsConverter
{
    public static function convert(\Grafana\Foundation\Canvas\CanvasElementOptions $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Canvas\CanvasElementOptionsBuilder())',
        ];
            if ($input->name !== "") {
    
        
    $buffer = 'name(';
        $arg0 =\var_export($input->name, true);
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
            if ($input->config !== null) {
    
        
    $buffer = 'config(';
        $arg0 =\var_export($input->config, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->constraint !== null) {
    
        
    $buffer = 'constraint(';
        $arg0 = \Grafana\Foundation\Canvas\ConstraintConverter::convert($input->constraint);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->placement !== null) {
    
        
    $buffer = 'placement(';
        $arg0 = \Grafana\Foundation\Canvas\PlacementConverter::convert($input->placement);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->background !== null) {
    
        
    $buffer = 'background(';
        $arg0 = \Grafana\Foundation\Canvas\BackgroundConfigConverter::convert($input->background);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->border !== null) {
    
        
    $buffer = 'border(';
        $arg0 = \Grafana\Foundation\Canvas\LineConfigConverter::convert($input->border);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->connections !== null && count($input->connections) >= 1) {
    
        
    $buffer = 'connections(';
        $tmparg0 = [];
        foreach ($input->connections as $arg1) {
        $tmpconnectionsarg1 = \Grafana\Foundation\Canvas\CanvasConnectionConverter::convert($arg1);
        $tmparg0[] = $tmpconnectionsarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

