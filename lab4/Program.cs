using System.Diagnostics;
using lab4;

Console.Write("Number of vertices: ");
var n = int.Parse(Console.ReadLine());

Graph g = new Graph(n);

// Generate Random Graph
var rnd = new Random();
var vertices = Enumerable.Range(0, n).ToList();
vertices = vertices.OrderBy(item => rnd.Next()).ToList();

for (int i = 0; i < vertices.Count-1; i++)
{
    g.AddEdge(vertices[i],vertices[i+1]);    
}

foreach (var i in vertices)
{
    var j = rnd.Next(vertices.Count - 1);
    var k = rnd.Next(vertices.Count - 1);
    
    g.AddEdge(vertices[j],vertices[k]);
}

// Measure Execution Time
Stopwatch stopwatch = new Stopwatch();

stopwatch.Start();
g.Bfs(0);
stopwatch.Stop();
Console.WriteLine("\nBfs elapsed time (ms): {0}", stopwatch.Elapsed.TotalMilliseconds);

stopwatch.Reset();
stopwatch.Start();
g.Dfs(0);
stopwatch.Stop();
Console.WriteLine("\nDfs elapsed time (ms): {0}", stopwatch.Elapsed.TotalMilliseconds);