using System;

namespace Visitor
{
    class Program
    {
        static void Main(string[] args)
        {
            ObjectStructure structure = new ObjectStructure();
            structure.Attach(new ConcreteElementA());
            structure.Attach(new ConcreteElementB());

            ConcreteVisitor1 visitor1 = new ConcreteVisitor1();
            ConcreteVisitor2 visitor2 = new ConcreteVisitor2();

            structure.Accept(visitor1);
            structure.Accept(visitor2);

            Console.ReadKey();
        }
    }
}
