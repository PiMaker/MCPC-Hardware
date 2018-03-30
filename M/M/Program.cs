// File: Program.cs
// Created: 23.03.2018
// 
// See <summary> tags for more information.

using System;
using System.Diagnostics;
using System.IO;
using Antlr4.Runtime;
using CommandLine;
using M.Grammar;
using Parser = CommandLine.Parser;

namespace M
{
    internal class Program
    {
        private static void Main(string[] args)
        {
            CmdOptions options = null;

            // Parse arguments
            Parser.Default.ParseArguments<CmdOptions>(args)
                  .WithParsed(opts => options = opts)
                  .WithNotParsed(errs =>
                  {
                      Console.ReadKey(true);
                      if (Debugger.IsAttached)
                      {
                          Console.ReadKey(true);
                      }

                      Environment.Exit(1);
                  });

            if (!File.Exists(options.InputFile))
            {
                Console.WriteLine("ERROR: Input file does not exist.");
                if (Debugger.IsAttached)
                {
                    Console.ReadKey(true);
                }

                Environment.Exit(1);
            }

            // Read
            var code = File.ReadAllText(options.InputFile);

            var inputStream = new AntlrInputStream(code);
            var lexer = new MLexer(inputStream);
            var commonTokenStream = new CommonTokenStream(lexer);
            var parser = new MParser(commonTokenStream);

            var visitor = new AsmVisitor();
            visitor.Visit(parser.compilationUnit());
            
            // Wait for user input if in debug mode
            if (Debugger.IsAttached)
            {
                Console.ReadKey(true);
            }
        }
    }
}