using System;

namespace RealWorldCode
{
    internal class IncomeVisitor : IVisitor
    {
        public void Visit(Element element)
        {
            Employee employee = element as Employee;

            employee.Income *= 1.10;
            Console.WriteLine($"{employee.GetType().Name} {employee.Name}'s new income: {String.Format("{0:C}", employee.Income)}");
        }
    }
}