// File: ASTNodeMath.cs
// Created: 23.03.2018
// 
// See <summary> tags for more information.

namespace M.AST
{
    internal class ASTNodeMath : ASTNode
    {
        public enum Operation
        {
            Add,
            Subtract,
            Multiply,
            Divide
        }

        public Operation Op { get; private set; }

        public ASTNodeMath(Operation op)
        {
            this.Op = op;
        }
    }
}