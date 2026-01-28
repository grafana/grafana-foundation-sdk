<?php

namespace Grafana\Foundation\Bargauge;

final class OptionsConverter
{
    public static function convert(\Grafana\Foundation\Bargauge\Options $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Bargauge\OptionsBuilder())',
        ];
            
    
        {
    $buffer = 'displayMode(';
        $arg0 ='\Grafana\Foundation\Common\BarGaugeDisplayMode::fromValue("'.$input->displayMode.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'valueMode(';
        $arg0 ='\Grafana\Foundation\Common\BarGaugeValueMode::fromValue("'.$input->valueMode.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->showUnfilled !== true) {
    
        
    $buffer = 'showUnfilled(';
        $arg0 =\var_export($input->showUnfilled, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->minVizWidth !== 0) {
    
        
    $buffer = 'minVizWidth(';
        $arg0 =\var_export($input->minVizWidth, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'reduceOptions(';
        $arg0 = \Grafana\Foundation\Common\ReduceDataOptionsConverter::convert($input->reduceOptions);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->text !== null) {
    
        
    $buffer = 'text(';
        $arg0 = \Grafana\Foundation\Common\VizTextDisplayOptionsConverter::convert($input->text);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->minVizHeight !== 10) {
    
        
    $buffer = 'minVizHeight(';
        $arg0 =\var_export($input->minVizHeight, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'orientation(';
        $arg0 ='\Grafana\Foundation\Common\VizOrientation::fromValue("'.$input->orientation.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    

        return \implode("\n\t->", $calls);
    }
}

