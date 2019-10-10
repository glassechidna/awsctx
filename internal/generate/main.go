package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type visitor struct {
	sb      io.Writer
	pkgName string
	fset    *token.FileSet
}

type method struct {
	Name       string
	Args       string
	ReturnType string
}

type api struct {
	Name          string
	SourcePackage string
	Methods       []method
}

func (v visitor) Visit(node ast.Node) (w ast.Visitor) {
	ts, ok := node.(*ast.TypeSpec)
	if ok {
		name := ts.Name.Name
		if !strings.HasSuffix(name, "API") {
			return v
		}

		myapi := &api{
			Name:          strings.TrimSuffix(name, "API"),
			SourcePackage: v.pkgName,
			Methods:       nil,
		}

		it := ts.Type.(*ast.InterfaceType)
		for _, meth := range it.Methods.List {
			name := meth.Names[0].Name
			if !strings.HasSuffix(name, "WithContext") {
				continue
			}

			if strings.HasSuffix(name, "PagesWithContext") {
				continue
			}

			if strings.HasPrefix(name, "WaitUntil") {
				continue
			}

			ft := meth.Type.(*ast.FuncType)

			inputType := &strings.Builder{}
			printer.Fprint(inputType, v.fset, ft.Params.List[1].Type)
			args := fmt.Sprintf("ctx context.Context, input %s, opts ...request.Option", inputType.String())

			returnType := &strings.Builder{}
			printer.Fprint(returnType, v.fset, ft.Results.List[0].Type)

			myapi.Methods = append(myapi.Methods, method{
				Name:       name,
				Args:       args,
				ReturnType: returnType.String(),
			})
		}

		err := tmpl.Execute(v.sb, myapi)
		if err != nil {
			panic(err)
		}
	}

	return v
}

var tmpltxt = `
// Code generated by internal/generate/main.go. DO NOT EDIT.

package {{ .SourcePackage }}ctx

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/{{ .SourcePackage }}"
	"github.com/aws/aws-sdk-go/service/{{ .SourcePackage }}/{{ .SourcePackage }}iface"
	"github.com/glassechidna/awsctx"
)

type {{ .Name }} interface {
{{- range .Methods }}
	{{ .Name }}({{ .Args }}) ({{ .ReturnType }}, error)
{{- end }}
}

type Client struct {
	{{ .SourcePackage }}iface.{{ .Name }}API
	Contexter awsctx.Contexter
}

func New(base {{ .SourcePackage }}iface.{{ .Name }}API, ctxer awsctx.Contexter) {{ .Name }} {
	return &Client{
		{{ .Name }}API: base,
		Contexter: ctxer,
	}
}

var _ {{ .Name }} = (*{{ .SourcePackage }}.{{ .Name }})(nil)
var _ {{ .Name }} = (*Client)(nil)
{{ range .Methods }}
func (c *Client) {{ .Name }}({{ .Args }}) ({{ .ReturnType }}, error) {
	req := &awsctx.AwsRequest{
		Service: "{{ $.SourcePackage }}",
		Action:  "{{ .Name }}",
		Input:   input,
		Output:  ({{ .ReturnType }})(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.{{ $.Name }}API.{{ .Name }}(ctx, input, opts...)
	})

	return req.Output.({{ .ReturnType }}), req.Error
}
{{ end }}
`
var tmpl = template.Must(template.New("").Parse(tmpltxt[1 : len(tmpltxt)-1]))

func main() {
	if len(os.Args) == 1 {
		fmt.Println("usage: generator path-to-generated-dir [sdk-version]")
		os.Exit(1)
	}

	genpath := os.Args[1]
	sdkVersion := ""
	if len(os.Args) == 3 {
		sdkVersion = os.Args[2]
	}

	sdkPath := awsSdkPath(sdkVersion)
	fmt.Printf("module path: %s\n", sdkPath)
	glob := sdkPath + "/service/**/*iface/interface.go"
	paths, err := filepath.Glob(glob)
	if err != nil {
		panic(err)
	}

	fset := token.NewFileSet()

	for _, path := range paths {
		f, err := parser.ParseFile(fset, path, nil, 0)
		if err != nil {
			panic(err)
		}
		dirname := filepath.Base(filepath.Dir(path))
		pkgName := strings.TrimSuffix(dirname, "iface")
		dirToMake := genpath + "/" + pkgName + "ctx"
		fmt.Println(dirToMake)
		os.MkdirAll(dirToMake, 0777)
		outf, err := os.Create(dirToMake + "/interface.go")
		if err != nil {
			panic(err)
		}
		defer outf.Close()
		v := &visitor{fset: fset, sb: outf, pkgName: pkgName}
		ast.Walk(v, f)
	}
}
