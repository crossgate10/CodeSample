using System;
using System.Collections.Generic;

namespace RealWorldCode
{
    internal class Employees
    {
        private List<Employee> employees = new List<Employee>();

        internal void Attach(Employee employee)
        {
            this.employees.Add(employee);
        }

        public void Detach(Employee employee)
        {
            this.employees.Remove(employee);
        }

        internal void Accept(IVisitor visitor)
        {
            foreach (Employee employee in this.employees)
            {
                employee.Accept(visitor);
            }
            Console.WriteLine();
        }
    }
}