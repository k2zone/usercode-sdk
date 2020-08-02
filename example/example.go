package main

import (
	"fmt"
	"github.com/k2zone/usercode-sdk/codesdk"
)

func main() {
	cli := codesdk.New(&codesdk.Config{
		Timeout:    20,
		DockerId:   "9846a1921f93",
		MemorySwap: "128M",
		Memory:     "64M",
	})
	res, err := cli.Run(&codesdk.Params{
		Compiler: codesdk.SWIFT,
		Stdin:    "a",
		Script:   `print("Hello World Swift")`,
	})
	fmt.Printf("res:%+v\n", res)
	fmt.Printf("err:%+v", err)
}

/*
switch ($command) {
            case 'python':
            case 'python2':
            case 'python3':
                $commandStr .= 'file.py&&/bin/bash ' . $config['scriptName'] .' '. $command .' ' . $config['dirName'] .'file.py ' ;
                break;
            case 'php5.6':
            case 'php7.1':
                $cmd = " /usr/bin/{$command} "; //注意空格
                $commandStr .= 'file.php&&/bin/bash ' . $config['scriptName'] . $cmd . $config['dirName'] .'file.php';
                break;
            case 'gcc':
            case 'gcc64':
                $commandStr .= 'file.c&&/bin/bash ' . $config['scriptName'] .' gcc '. $config['dirName'] .'file.c ./a.out';
                break;
            case 'gcc32':
               // $commandStr = $commandStr2;
                $commandStr .= 'file.c&&/bin/bash ' .$config['scriptName'] .' "gcc -m32" '. $config['dirName'] .'file.c ./a.out';
                //dump($commandStr);
                break;
            case 'g++':
            case 'g++64':
                $commandStr .= 'file.cpp&&/bin/bash ' .  $config['scriptName'] .' g++  '. $config['dirName'] .'file.cpp  ./a.out';
                break;
            case 'g++32':
               // $commandStr = $commandStr2;
                $commandStr .= 'file.cpp&&/bin/bash ' . $config['scriptName'] .' "g++ -m32" '. $config['dirName'] .'file.cpp ./a.out';
                break;
            case 'mcs':
                $commandStr .= 'file.cs&&/bin/bash ' . $config['scriptName'] .' mcs '. $config['dirName'] .'file.cs "mono '. $config['dirName'] .'file.exe "';
                break;
            case 'java':
                //$mainClass = getJavaMainClass(base64_decode($script));
                //$mainClass = ($mainClass === false) ? 'test' : $mainClass;
                $mainClass = 'Main';
                $commandStr .= 'Main.java&&/bin/bash ' . $config['scriptName'] .' "java -jar /opt/ecj-4.2.2.jar -1.5 -encoding utf8 " '. $config['dirName'] .'Main.java "java -Dfile.encoding=UTF-8 -classpath /usercode '.$mainClass.'"';
                break;
            case 'go':
                $commandStr .= 'file.go&&/bin/bash ' .$config['scriptName'] .' go ' . '"run ' . $config['dirName'] .'file.go"';
                break;
            case 'lua':
                $commandStr .= 'file.lua&&/bin/bash ' . $config['scriptName'] .' lua '. $config['dirName'] .'file.lua';
                break;
            case 'Rscript':
                $commandStr .= 'file.R&&/bin/bash ' . $config['scriptName'] .' Rscript '. $config['dirName'] .'file.R';
                break;
            case 'shell':
                $commandStr .= 'file.sh&& ' .  $config['scriptName'] .' bash '. $config['dirName'] .'file.sh';
                break;
            case 'perl':
                $commandStr .= 'file.pl&&/bin/bash ' . $config['scriptName'] .' perl '. $config['dirName'] .'file.pl';
                break;
            case 'node':
                $commandStr .= 'file.js&&/bin/bash ' . $config['scriptName'] .' node '. $config['dirName'] .'file.js';
                break;
            case 'swift4.2':
                //$commandStr = $commandStr2;
                $commandStr .= 'file.swift&&/bin/bash ' . $config['scriptName'] .' /opt/swift-4.2.1-RELEASE-ubuntu18.04/usr/bin/swift '. $config['dirName'] .'file.swift';
                break;
            default:
                $commandStr = '';
                break;
        }
*/
