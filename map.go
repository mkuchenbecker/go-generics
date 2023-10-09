package main

// import (
// 	"bytes"
// 	"fmt"
// )

// type NodeStyle struct {
// 	Color     string
// 	Fillcolor string
// 	Style     string
// 	Shape     string
// }

// func (n NodeStyle) WithColor(color string) NodeStyle {
// 	n.Color = color
// 	return n
// }

// func (n NodeStyle) WithFillColor(fillcolor string) NodeStyle {
// 	n.Fillcolor = fillcolor
// 	return n
// }

// func (n NodeStyle) WithStyle(style string) NodeStyle {
// 	n.Style = style
// 	return n
// }

// func (n NodeStyle) WithShape(shape string) NodeStyle {
// 	n.Shape = shape
// 	return n
// }

// type GraphStyle struct {
// 	Splines string
// 	RankDir string
// }

// type LineStyle struct {
// 	Color            string
// 	Style            string
// 	IgnoreConstraint bool
// }

// type Labeled interface {
// 	GetLabel() Label
// }

// type ID interface {
// 	GetID() string
// }

// type Component interface {
// 	Labeled
// 	ID
// 	GetStyle() NodeStyle
// 	GetOutputs() Slice[Label]
// 	GetInputs() Slice[Label]
// 	NewConnection(to Component) ConnectionBuilder
// }

// type Connection interface {
// 	Labeled
// 	GetStyle() LineStyle
// 	GetFromComponent() Component
// 	GetFromLabel() Label
// 	GetToComponent() Component
// 	GetToLabel() Label
// }

// type ConnectionBuilder interface {
// 	From(from Component) ConnectionBuilder
// 	FromLabel(label Label) ConnectionBuilder
// 	ToLabel(label Label) ConnectionBuilder
// 	To(to Component) ConnectionBuilder
// 	WithLabel(label Label) ConnectionBuilder
// 	WithStyle(style LineStyle) ConnectionBuilder
// 	Build() Connection
// }

// type connectionBuilder struct {
// 	from      Component
// 	fromlabel Label
// 	to        Component
// 	tolabel   Label
// 	label     Label
// 	style     LineStyle
// }

// func (c *connectionBuilder) From(from Component) ConnectionBuilder {
// 	c.from = from
// 	return c
// }

// func (c *connectionBuilder) To(to Component) ConnectionBuilder {
// 	c.to = to
// 	return c
// }

// func (c *connectionBuilder) WithLabel(label Label) ConnectionBuilder {
// 	c.label = label
// 	return c
// }

// func (c *connectionBuilder) WithStyle(style LineStyle) ConnectionBuilder {
// 	c.style = style
// 	return c
// }

// func (c *connectionBuilder) FromLabel(label Label) ConnectionBuilder {
// 	c.fromlabel = label
// 	return c
// }

// func (c *connectionBuilder) ToLabel(label Label) ConnectionBuilder {
// 	c.tolabel = label
// 	return c
// }

// func (c *connectionBuilder) Build() Connection {
// 	return &connection{
// 		Style:     c.style,
// 		Label:     c.label,
// 		From:      c.from,
// 		To:        c.to,
// 		FromLabel: c.fromlabel,
// 		ToLabel:   c.tolabel,
// 	}
// }

// type connection struct {
// 	Style     LineStyle
// 	Label     Label
// 	From      Component
// 	FromLabel Label
// 	To        Component
// 	ToLabel   Label
// }

// func (c *connection) GetStyle() LineStyle {
// 	return c.Style
// }

// func (c *connection) GetLabel() Label {
// 	return c.Label
// }

// func (c *connection) GetFromComponent() Component {
// 	return c.From
// }

// func (c *connection) GetToComponent() Component {
// 	return c.To
// }

// func (c *connection) GetFromLabel() Label {
// 	return c.FromLabel
// }

// func (c *connection) GetToLabel() Label {
// 	return c.ToLabel
// }

// type Label struct {
// 	Tag   string
// 	Label string
// }

// type component struct {
// 	Style   NodeStyle
// 	Label   Label
// 	Outputs Slice[Label]
// 	Inputs  Slice[Label]
// 	ID      string
// }

// func (c component) GetInputs() Slice[Label] {
// 	return c.Inputs
// }

// func (c component) GetOutputs() Slice[Label] {
// 	return c.Outputs
// }

// func (c component) GetStyle() NodeStyle {
// 	return c.Style
// }

// func (c component) GetID() string {
// 	return c.ID
// }

// func (c component) NewConnection(to Component) ConnectionBuilder {
// 	builder := &connectionBuilder{from: &c, to: to}
// 	return builder
// }

// func CollectLabels(buf *bytes.Buffer) func(l Label) {
// 	return func(l Label) {
// 		if buf.Len() > 0 {
// 			buf.WriteString("|")
// 		}
// 		buf.WriteString(fmt.Sprintf("<%s> %s", l.Tag, l.Label))
// 	}
// }

// func (c component) GetLabel() Label {
// 	inputBuf := bytes.NewBufferString("")
// 	c.Inputs.Do(CollectLabels(inputBuf))

// 	outputBuf := bytes.NewBufferString("")
// 	c.Outputs.Do(CollectLabels(outputBuf))

// 	return Label{Label: fmt.Sprintf("{{%s}|%s|{%s}}", inputBuf.String(), c.Label.Label, outputBuf.String())}
// }

// type Graph struct {
// 	Components   Slice[Component]
// 	Connections  Slice[Connection]
// 	Subgraphs    Slice[Graph]
// 	DefaultStyle NodeStyle
// 	GraphStyle   GraphStyle
// 	Name         string
// 	Label 	  Label
// }

// func (g Graph) Contents() string {
// 	buf := bytes.NewBufferString("")
// 	if g.GraphStyle.RankDir != "" {
// 		buf.WriteString(fmt.Sprintf("rankdir=%s;\n", g.GraphStyle.RankDir))
// 	}
// 	if g.GraphStyle.Splines != "" {
// 		buf.WriteString(fmt.Sprintf("splines=%s;\n", g.GraphStyle.Splines))
// 	}

// 	stlebuf := bytes.NewBufferString("")
// 	if g.DefaultStyle.Shape != "" {
// 		stlebuf.WriteString(fmt.Sprintf("shape=%s", g.DefaultStyle.Shape))
// 	}
// 	if g.DefaultStyle.Color != "" {
// 		stlebuf.WriteString(fmt.Sprintf("color=%s", g.DefaultStyle.Color))
// 	}
// 	if g.DefaultStyle.Fillcolor != "" {
// 		stlebuf.WriteString(fmt.Sprintf("fillcolor=%s", g.DefaultStyle.Fillcolor))
// 	}
// 	if g.DefaultStyle.Style != "" {
// 		stlebuf.WriteString(fmt.Sprintf("style=%s", g.DefaultStyle.Style))
// 	}
// 	if stlebuf.Len() > 0 {
// 		buf.WriteString(fmt.Sprintf("node [%s];\n", stlebuf.String()))
// 	}

// 	if g.Label.Label != "" {
// 		buf.WriteString(fmt.Sprintf("label=\"%s\";\n", g.Label.Label))
// 	}

// 	g.Subgraphs.Do(func(s Graph) {
// 		buf.WriteString(fmt.Sprintf("subgraph cluster_%s {", s.Name))
// 		buf.WriteString(s.Contents())
// 		buf.WriteString("}")
// 	})
// 	g.Components.Do(func(c Component) {
// 		buf.WriteString(componentString(c))
// 	})
// 	g.Connections.Do(func(c Connection) {
// 		buf.WriteString(connectionString(c))
// 	})
// 	return buf.String()
// }

// func (g Graph) ToGraph() *bytes.Buffer {
// 	return bytes.NewBufferString(fmt.Sprintf("digraph %s {\n%s}", g.Name, g.Contents()))
// }

// type T any

// type O any
// type Slice[T any] []T

// func (f Slice[T]) Map(mapper func(t T) O) Slice[O] {
// 	return Map[T, O](f, mapper)
// }

// func (f Slice[T]) Filter(filter func(t T) bool) Slice[T] {
// 	return Filter[T](f, filter)
// }

// func (f Slice[T]) Do(doer func(t T)) {
// 	Do[T](f, doer)
// }

// func Map[T any, O any](ts []T, mapper func(t T) O) []O {
// 	r := make([]O, 0, len(ts))
// 	for i, t := range ts {
// 		r[i] = mapper(t)
// 	}
// 	return r
// }

// func Do[T any](ts []T, doer func(t T)) {
// 	for _, t := range ts {
// 		doer(t)
// 	}
// }

// func Filter[T any](ts []T, filter func(t T) bool) []T {
// 	r := make([]T, 0)
// 	for _, t := range ts {
// 		if filter(t) {
// 			r = append(r, t)
// 		}
// 	}
// 	return r
// }
