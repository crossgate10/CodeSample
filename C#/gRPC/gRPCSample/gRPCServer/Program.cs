using Grpc.Core;
using GRPCDemo;
using System;
using System.Threading.Tasks;

namespace gRPCServer
{
    class gRPCImpl : gRPC.gRPCBase
    {

        public override Task<HelloReply> SayHello(HelloRequest request, ServerCallContext context)
        {
            return Task.FromResult(new HelloReply { Message = "Hello " + request.Name });
        }
    }

    class Program
    {
        const int Port = 9007;

        public static void Main(string[] args)
        {
            Server server = new Server
            {
                Services = { gRPC.BindService(new gRPCImpl()) },
                Ports = { new ServerPort("localhost", Port, ServerCredentials.Insecure) }
            };
            server.Start();

            Console.WriteLine("gRPC server listening on port " + Port);
            Console.WriteLine("任意鍵退出...");
            Console.ReadKey();

            server.ShutdownAsync().Wait();
        }
    }
}
