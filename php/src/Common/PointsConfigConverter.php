<?php

namespace Grafana\Foundation\Common;

final class PointsConfigConverter
{
    public static function convert(\Grafana\Foundation\Common\PointsConfig $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Common\PointsConfigBuilder())',
        ];
            if ($input->showPoints !== null) {
    
        
    $buffer = 'showPoints(';
        $arg0 ='\Grafana\Foundation\Common\VisibilityMode::fromValue("'.$input->showPoints.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->pointSize !== null) {
    
        
    $buffer = 'pointSize(';
        $arg0 =\var_export($input->pointSize, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->pointColor !== null && $input->pointColor !== "") {
    
        
    $buffer = 'pointColor(';
        $arg0 =\var_export($input->pointColor, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->pointSymbol !== null && $input->pointSymbol !== "") {
    
        
    $buffer = 'pointSymbol(';
        $arg0 =\var_export($input->pointSymbol, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

