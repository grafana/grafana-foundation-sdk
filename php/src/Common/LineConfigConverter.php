<?php

namespace Grafana\Foundation\Common;

final class LineConfigConverter
{
    public static function convert(\Grafana\Foundation\Common\LineConfig $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Common\LineConfigBuilder())',
        ];
            if ($input->lineColor !== null && $input->lineColor !== "") {
    
        
    $buffer = 'lineColor(';
        $arg0 =\var_export($input->lineColor, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->lineWidth !== null) {
    
        
    $buffer = 'lineWidth(';
        $arg0 =\var_export($input->lineWidth, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->lineInterpolation !== null) {
    
        
    $buffer = 'lineInterpolation(';
        $arg0 ='\Grafana\Foundation\Common\LineInterpolation::fromValue("'.$input->lineInterpolation.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->lineStyle !== null) {
    
        
    $buffer = 'lineStyle(';
        $arg0 = \Grafana\Foundation\Common\LineStyleConverter::convert($input->lineStyle);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spanNulls !== null) {
    
        
    $buffer = 'spanNulls(';
        switch (true) {
            case is_bool($input->spanNulls):
                $disjunctionspanNulls =\var_export($input->spanNulls, true);
                $arg0 = $disjunctionspanNulls;
                break;
            case is_float($input->spanNulls):
                $disjunctionspanNulls =\var_export($input->spanNulls, true);
                $arg0 = $disjunctionspanNulls;
                break;
            default:
                throw new \ValueError('disjunction branch not handled');
        }
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

