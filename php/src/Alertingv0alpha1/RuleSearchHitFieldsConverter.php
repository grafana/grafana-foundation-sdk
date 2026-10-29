<?php

namespace Grafana\Foundation\Alertingv0alpha1;

final class RuleSearchHitFieldsConverter
{
    public static function convert(\Grafana\Foundation\Alertingv0alpha1\RuleSearchHitFields $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Alertingv0alpha1\RuleSearchHitFieldsBuilder())',
        ];
            if ($input->title !== null && $input->title !== "") {
    
        
    $buffer = 'title(';
        $arg0 =\var_export($input->title, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->folder !== null && $input->folder !== "") {
    
        
    $buffer = 'folder(';
        $arg0 =\var_export($input->folder, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->type !== null && $input->type !== "") {
    
        
    $buffer = 'type(';
        $arg0 =\var_export($input->type, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->interval !== null && $input->interval !== "") {
    
        
    $buffer = 'interval(';
        $arg0 =\var_export($input->interval, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->paused !== null) {
    
        
    $buffer = 'paused(';
        $arg0 =\var_export($input->paused, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->labels !== null) {
    
        
    $buffer = 'labels(';
        $arg0 = "[";
        foreach ($input->labels as $key => $arg1) {
            $tmplabelsarg1 =\var_export($arg1, true);
            $arg0 .= "\t".var_export($key, true)." => $tmplabelsarg1,";
        }
        $arg0 .= "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->datasourceUIDs !== null && count($input->datasourceUIDs) >= 1) {
    
        
    $buffer = 'datasourceUIDs(';
        $tmparg0 = [];
        foreach ($input->datasourceUIDs as $arg1) {
        $tmpdatasourceUIDsarg1 =\var_export($arg1, true);
        $tmparg0[] = $tmpdatasourceUIDsarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->annotations !== null) {
    
        
    $buffer = 'annotations(';
        $arg0 = "[";
        foreach ($input->annotations as $key => $arg1) {
            $tmpannotationsarg1 =\var_export($arg1, true);
            $arg0 .= "\t".var_export($key, true)." => $tmpannotationsarg1,";
        }
        $arg0 .= "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->for !== null && $input->for !== "") {
    
        
    $buffer = 'for(';
        $arg0 =\var_export($input->for, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->keepFiringFor !== null && $input->keepFiringFor !== "") {
    
        
    $buffer = 'keepFiringFor(';
        $arg0 =\var_export($input->keepFiringFor, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->dashboardUID !== null && $input->dashboardUID !== "") {
    
        
    $buffer = 'dashboardUID(';
        $arg0 =\var_export($input->dashboardUID, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->panelID !== null) {
    
        
    $buffer = 'panelID(';
        $arg0 =\var_export($input->panelID, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->receiver !== null && $input->receiver !== "") {
    
        
    $buffer = 'receiver(';
        $arg0 =\var_export($input->receiver, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->notificationType !== null && $input->notificationType !== "") {
    
        
    $buffer = 'notificationType(';
        $arg0 =\var_export($input->notificationType, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->routingTree !== null && $input->routingTree !== "") {
    
        
    $buffer = 'routingTree(';
        $arg0 =\var_export($input->routingTree, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->metric !== null && $input->metric !== "") {
    
        
    $buffer = 'metric(';
        $arg0 =\var_export($input->metric, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->targetDatasourceUID !== null && $input->targetDatasourceUID !== "") {
    
        
    $buffer = 'targetDatasourceUID(';
        $arg0 =\var_export($input->targetDatasourceUID, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

