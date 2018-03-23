// File: CmdOptions.cs
// Created: 23.03.2018
// 
// See <summary> tags for more information.

using System.Collections.Generic;
using CommandLine;

namespace M
{
    internal class CmdOptions
    {
        [Option('i', "input", Required = true, HelpText = "Input file to be compiled.")]
        public string InputFile { get; set; }

        [Option('o', "output", Required = true, HelpText = "Output file to be created.")]
        public string OutputFile { get; set; }

        //[Option(Default = false, HelpText = "Prints verbose messages to standard output.")]
        //public bool Verbose { get; set; }
    }
}
