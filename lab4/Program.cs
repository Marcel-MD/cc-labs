using System.Diagnostics;
using lab4;

var n = 14;

Graph g = new Graph(n);

g.AddEdge(1, 2);
g.AddEdge(1, 3);
g.AddEdge(2, 4);
g.AddEdge(3, 4);
g.AddEdge(3, 5);
g.AddEdge(4, 6);
g.AddEdge(5, 7);
g.AddEdge(5, 8);
g.AddEdge(6, 9);
g.AddEdge(7, 10);
g.AddEdge(7, 11);
g.AddEdge(10, 12);
g.AddEdge(11, 13);

// Measure Execution Time
Stopwatch stopwatch = new Stopwatch();

stopwatch.Start();
g.Bfs(1, 2);
stopwatch.Stop();

stopwatch.Reset();
stopwatch.Start();
g.Dfs(1, 2);
stopwatch.Stop();

stopwatch.Start();
g.Bfs(1, 13);
stopwatch.Stop();

stopwatch.Reset();
stopwatch.Start();
g.Dfs(1, 13);
stopwatch.Stop();

for (int i = 2; i < 14; i++)
{
    Console.WriteLine($"\n------------ {i} -------------");
    stopwatch.Start();
    g.Bfs(1, i);
    stopwatch.Stop();
    Console.WriteLine("Bfs elapsed time (ms): {0}", stopwatch.Elapsed.TotalMilliseconds);

    stopwatch.Reset();
    stopwatch.Start();
    g.Dfs(1, i);
    stopwatch.Stop();
    Console.WriteLine("Dfs elapsed time (ms): {0}", stopwatch.Elapsed.TotalMilliseconds);
} 

