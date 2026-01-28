<?php

namespace Grafana\Foundation\Barchart;

final class OptionsConverter
{
    public static function convert(\Grafana\Foundation\Barchart\Options $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Barchart\OptionsBuilder())',
        ];
            if ($input->xField !== null && $input->xField !== "") {
    
        
    $buffer = 'xField(';
        $arg0 =\var_export($input->xField, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->colorByField !== null && $input->colorByField !== "") {
    
        
    $buffer = 'colorByField(';
        $arg0 =\var_export($input->colorByField, true);
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
    
    
            if ($input->barRadius !== null && $input->barRadius !== (float) 0) {
    
        
    $buffer = 'barRadius(';
        $arg0 =\var_export($input->barRadius, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->xTickLabelRotation !== 0) {
    
        
    $buffer = 'xTickLabelRotation(';
        $arg0 =\var_export($input->xTickLabelRotation, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'xTickLabelMaxLength(';
        $arg0 =\var_export($input->xTickLabelMaxLength, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->xTickLabelSpacing !== null && $input->xTickLabelSpacing !== 0) {
    
        
    $buffer = 'xTickLabelSpacing(';
        $arg0 =\var_export($input->xTickLabelSpacing, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'stacking(';
        $arg0 ='\Grafana\Foundation\Common\StackingMode::fromValue("'.$input->stacking.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'showValue(';
        $arg0 ='\Grafana\Foundation\Common\VisibilityMode::fromValue("'.$input->showValue.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->barWidth !== (float) 0.97) {
    
        
    $buffer = 'barWidth(';
        $arg0 =\var_export($input->barWidth, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->groupWidth !== (float) 0.7) {
    
        
    $buffer = 'groupWidth(';
        $arg0 =\var_export($input->groupWidth, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'legend(';
        $arg0 = \Grafana\Foundation\Common\VizLegendOptionsConverter::convert($input->legend);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'tooltip(';
        $arg0 = \Grafana\Foundation\Common\VizTooltipOptionsConverter::convert($input->tooltip);
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
            if ($input->fullHighlight !== false) {
    
        
    $buffer = 'fullHighlight(';
        $arg0 =\var_export($input->fullHighlight, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

