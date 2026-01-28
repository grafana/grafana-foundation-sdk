<?php

namespace Grafana\Foundation\Logs;

final class OptionsConverter
{
    public static function convert(\Grafana\Foundation\Logs\Options $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Logs\OptionsBuilder())',
        ];
            
    
        {
    $buffer = 'showLabels(';
        $arg0 =\var_export($input->showLabels, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'showCommonLabels(';
        $arg0 =\var_export($input->showCommonLabels, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'showTime(';
        $arg0 =\var_export($input->showTime, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'showLogContextToggle(';
        $arg0 =\var_export($input->showLogContextToggle, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'wrapLogMessage(';
        $arg0 =\var_export($input->wrapLogMessage, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'prettifyLogMessage(';
        $arg0 =\var_export($input->prettifyLogMessage, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'enableLogDetails(';
        $arg0 =\var_export($input->enableLogDetails, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'sortOrder(';
        $arg0 ='\Grafana\Foundation\Common\LogsSortOrder::fromValue("'.$input->sortOrder.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'dedupStrategy(';
        $arg0 ='\Grafana\Foundation\Common\LogsDedupStrategy::fromValue("'.$input->dedupStrategy.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->onClickFilterLabel !== null) {
    
        
    $buffer = 'onClickFilterLabel(';
        $arg0 =\var_export($input->onClickFilterLabel, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->onClickFilterOutLabel !== null) {
    
        
    $buffer = 'onClickFilterOutLabel(';
        $arg0 =\var_export($input->onClickFilterOutLabel, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->isFilterLabelActive !== null) {
    
        
    $buffer = 'isFilterLabelActive(';
        $arg0 =\var_export($input->isFilterLabelActive, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->onClickFilterString !== null) {
    
        
    $buffer = 'onClickFilterString(';
        $arg0 =\var_export($input->onClickFilterString, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->onClickFilterOutString !== null) {
    
        
    $buffer = 'onClickFilterOutString(';
        $arg0 =\var_export($input->onClickFilterOutString, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->onClickShowField !== null) {
    
        
    $buffer = 'onClickShowField(';
        $arg0 =\var_export($input->onClickShowField, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->onClickHideField !== null) {
    
        
    $buffer = 'onClickHideField(';
        $arg0 =\var_export($input->onClickHideField, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->displayedFields !== null && count($input->displayedFields) >= 1) {
    
        
    $buffer = 'displayedFields(';
        $tmparg0 = [];
        foreach ($input->displayedFields as $arg1) {
        $tmpdisplayedFieldsarg1 =\var_export($arg1, true);
        $tmparg0[] = $tmpdisplayedFieldsarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

