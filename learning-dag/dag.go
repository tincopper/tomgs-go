package main

//图结构
type DAG struct {
	Vertexs []*Vertex
}

//顶点
type Vertex struct {
	Key      string
	Value    interface{}
	Parents  []*Vertex
	Children []*Vertex
}

//添加顶点
func (dag *DAG) AddVertex(v *Vertex) {
	dag.Vertexs = append(dag.Vertexs, v)
}

//添加边
func (dag *DAG) AddEdge(from, to *Vertex) {
	from.Children = append(from.Children, to)
	to.Parents = append(to.Parents, from)
}

//生成图，返回dag和其根顶点
func NewDAG1() (*DAG, *Vertex) {
	var dag = &DAG{}
	va := &Vertex{Key: "a", Value: "1"}
	/*
		 a
	   / | \
	  b  c  d
	  /  |    \
	 e   f     g
	 |  / | \  |
	 |  h i  j |     1                     a(4)
	 \  | |  | |                b(3)       c(3)    d(3)
	  \|  |  |/						       f(2)
	    \ | /                e     h		i(1)		j		g 1
	      u         0                  	   u(0)
	 */
	return dag, va
}
func NewDAG() (*DAG, *Vertex) {
	var dag = &DAG{}
	va := &Vertex{Key: "a", Value: "1"}
	vb := &Vertex{Key: "b", Value: "2"}
	vc := &Vertex{Key: "c", Value: "3"}
	vd := &Vertex{Key: "d", Value: "4"}
	ve := &Vertex{Key: "e", Value: "5"}
	
	/*
	   a
	  / \
	 b   c
	  \  /
	   d
	   |
	   e
	 */
	 
	dag.AddEdge(va, vb)
	dag.AddEdge(va, vc)
	
	dag.AddEdge(vb, vd)
	dag.AddEdge(vc, vd)
	
	dag.AddEdge(vd, ve)
	
	return dag, va
}

/*func NewDAG() (*DAG, *Vertex) {
	var dag = &DAG{}
	
	va := &Vertex{Key: "a", Value: "1"}
	vb := &Vertex{Key: "b", Value: "2"}
	vc := &Vertex{Key: "c", Value: "3"}
	vd := &Vertex{Key: "d", Value: "4"}
	ve := &Vertex{Key: "e", Value: "5"}
	vf := &Vertex{Key: "f", Value: "6"}
	vg := &Vertex{Key: "g", Value: "7"}
	vh := &Vertex{Key: "h", Value: "8"}
	vi := &Vertex{Key: "i", Value: "9"}
	vx := &Vertex{Key: "x", Value: "10"}
	vy := &Vertex{Key: "y", Value: "11"}
	
	dag.AddEdge(va, vb)
	dag.AddEdge(va, vc)
	dag.AddEdge(va, vd)
	dag.AddEdge(vb, ve)
	dag.AddEdge(vb, vh)
	dag.AddEdge(vb, vf)
	dag.AddEdge(vc, vf)
	dag.AddEdge(vc, vg)
	dag.AddEdge(vd, vg)
	dag.AddEdge(vh, vi)
	dag.AddEdge(ve, vi)
	dag.AddEdge(vf, vi)
	dag.AddEdge(vg, vi)
	dag.AddEdge(vx, vd)
	dag.AddEdge(vy, vi)
	return dag, va
}*/
