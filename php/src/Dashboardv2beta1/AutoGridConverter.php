<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class AutoGridConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutKind $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\AutoGridBuilder())',
        ];
            if ($input->spec->maxColumnCount !== null && $input->spec->maxColumnCount !== (float) 3) {
    
        
    $buffer = 'maxColumnCount(';
        $arg0 =\var_export($input->spec->maxColumnCount, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'columnWidthMode(';
        $arg0 ='\Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutSpecColumnWidthMode::fromValue("'.$input->spec->columnWidthMode.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->spec->columnWidth !== null) {
    
        
    $buffer = 'columnWidth(';
        $arg0 =\var_export($input->spec->columnWidth, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'rowHeightMode(';
        $arg0 ='\Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutSpecRowHeightMode::fromValue("'.$input->spec->rowHeightMode.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->spec->rowHeight !== null) {
    
        
    $buffer = 'rowHeight(';
        $arg0 =\var_export($input->spec->rowHeight, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->fillScreen !== null && $input->spec->fillScreen !== false) {
    
        
    $buffer = 'fillScreen(';
        $arg0 =\var_export($input->spec->fillScreen, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if (count($input->spec->items) >= 1) {
    
        
    $buffer = 'items(';
        $tmparg0 = [];
        foreach ($input->spec->items as $arg1) {
        $tmpitemsarg1 = \Grafana\Foundation\Dashboardv2beta1\AutoGridItemConverter::convert($arg1);
        $tmparg0[] = $tmpitemsarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if (count($input->spec->items) >= 1) {
    foreach ($input->spec->items as $item) {
        
    $buffer = 'withItem(';
        $arg0 =\var_export($item->kind, true);
        $buffer .= $arg0;
        $buffer .= ', ';
        $arg1 ='(new \Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutItemSpec(element: '.'(new \Grafana\Foundation\Dashboardv2beta1\ElementReference(kind: '.\var_export($item->spec->element->kind, true).',name: '.\var_export($item->spec->element->name, true).',))'.','.(($item->spec->repeat !== null) ? 'repeat: '.'(new \Grafana\Foundation\Dashboardv2beta1\AutoGridRepeatOptions(mode: '.\var_export($item->spec->repeat->mode, true).',value: '.\var_export($item->spec->repeat->value, true).',))'.', ' : '').''.(($item->spec->conditionalRendering !== null) ? 'conditionalRendering: '.'(new \Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingGroupKind(kind: '.\var_export($item->spec->conditionalRendering->kind, true).',spec: '.'(new \Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingGroupSpec(visibility: '.'\Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingGroupSpecVisibility::fromValue("'.$item->spec->conditionalRendering->spec->visibility.'")'.',condition: '.'\Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingGroupSpecCondition::fromValue("'.$item->spec->conditionalRendering->spec->condition.'")'.',items: '.\var_export($item->spec->conditionalRendering->spec->items, true).',))'.',))'.', ' : '').'))';
        $buffer .= $arg1;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    }
    }

        return \implode("\n\t->", $calls);
    }
}

