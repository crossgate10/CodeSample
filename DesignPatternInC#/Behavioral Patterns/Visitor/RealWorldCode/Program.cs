using System;

namespace RealWorldCode
{
    class Program
    {
        static void Main(string[] args)
        {
            Employees employees = new Employees();
            employees.Attach(new Clerk());
            employees.Attach(new Director());
            employees.Attach(new President());

            employees.Accept(new IncomeVisitor());
            employees.Accept(new VacationVisitor());

            Console.ReadKey();
        }
    }
}
