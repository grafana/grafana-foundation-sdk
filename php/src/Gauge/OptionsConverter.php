<?php

namespace Grafana\Foundation\Gauge;

final class OptionsConverter
{
    public static function convert(\Grafana\Foundation\Gauge\Options $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Gauge\OptionsBuilder())',
        ];
            if ($input->showThresholdLabels !== false) {
    
        
    $buffer = 'showThresholdLabels(';
        $arg0 =\var_export($input->showThresholdLabels, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->showThresholdMarkers !== true) {
    
        
    $buffer = 'showThresholdMarkers(';
        $arg0 =\var_export($input->showThresholdMarkers, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'sizing(';
        $arg0 ='\Grafana\Foundation\Common\BarGaugeSizing::fromValue("'.$input->sizing.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->minVizWidth !== 75) {
    
        
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
            if ($input->minVizHeight !== 75) {
    
        
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

