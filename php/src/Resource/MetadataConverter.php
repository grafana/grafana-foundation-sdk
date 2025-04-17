<?php

namespace Grafana\Foundation\Resource;

final class MetadataConverter
{
    public static function convert(\Grafana\Foundation\Resource\Metadata $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Resource\MetadataBuilder())',
        ];
            if ($input->name !== "") {
    
        
    $buffer = 'name(';
        $arg0 =\var_export($input->name, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->namespace !== null && $input->namespace !== "") {
    
        
    $buffer = 'namespace(';
        $arg0 =\var_export($input->namespace, true);
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
            if ($input->uid !== null && $input->uid !== "") {
    
        
    $buffer = 'uid(';
        $arg0 =\var_export($input->uid, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->resourceVersion !== null && $input->resourceVersion !== "") {
    
        
    $buffer = 'resourceVersion(';
        $arg0 =\var_export($input->resourceVersion, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->generation !== null) {
    
        
    $buffer = 'generation(';
        $arg0 =\var_export($input->generation, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->creationTimestamp !== null && $input->creationTimestamp !== "") {
    
        
    $buffer = 'creationTimestamp(';
        $arg0 =\var_export($input->creationTimestamp, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->updateTimestamp !== null && $input->updateTimestamp !== "") {
    
        
    $buffer = 'updateTimestamp(';
        $arg0 =\var_export($input->updateTimestamp, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->deletionTimestamp !== null && $input->deletionTimestamp !== "") {
    
        
    $buffer = 'deletionTimestamp(';
        $arg0 =\var_export($input->deletionTimestamp, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

