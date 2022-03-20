namespace lab4;

public class Graph
{
    // No. of vertices    
    private int _V;
 
    //Adjacency Lists
    List<int>[] _adj;
 
    public Graph(int V)
    {
        _adj = new List<int>[V];
        for(int i = 0; i < _adj.Length; i++)
        {
            _adj[i] = new List<int>();
        }
        _V = V;
    }
 
    // Function to add an edge into the graph
    public void AddEdge(int v, int w)
    {        
        _adj[v].Add(w);
    }
    
    public void Bfs(int s)
    {
        // Mark all the vertices as not
        // visited(By default set as false)
        bool[] visited = new bool[_V];
        for(int i = 0; i < _V; i++)
            visited[i] = false;
     
        // Create a queue for BFS
        LinkedList<int> queue = new LinkedList<int>();
     
        // Mark the current node as
        // visited and enqueue it
        visited[s] = true;
        queue.AddLast(s);        
 
        while(queue.Any())
        {
         
            // Dequeue a vertex from queue
            // and print it
            s = queue.First();
            // Console.Write(s + " " );
            queue.RemoveFirst();
         
            // Get all adjacent vertices of the
            // dequeued vertex s. If a adjacent
            // has not been visited, then mark it
            // visited and enqueue it
            List<int> list = _adj[s];
 
            foreach (var val in list)            
            {
                if (!visited[val])
                {
                    visited[val] = true;
                    queue.AddLast(val);
                }
            }
        }
    }
    
    private void DfsUtil(int v, bool[] visited)
    {
        // Mark the current node as visited
        // and print it
        visited[v] = true;
        // Console.Write(v + " ");
 
        // Recur for all the vertices
        // adjacent to this vertex
        List<int> vList = _adj[v];
        foreach(var n in vList)
        {
            if (!visited[n])
                DfsUtil(n, visited);
        }
    }
 
    // The function to do DFS traversal.
    // It uses recursive DFSUtil()
    public void Dfs(int v)
    {
        // Mark all the vertices as not visited
        // (set as false by default in c#)
        bool[] visited = new bool[_V];
 
        // Call the recursive helper function
        // to print DFS traversal
        DfsUtil(v, visited);
    }
}