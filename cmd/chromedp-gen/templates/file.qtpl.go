// This file is automatically generated by qtc from "file.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line templates/file.qtpl:1
package templates

//line templates/file.qtpl:1
import (
	"sort"

	"github.com/rjeczalik/chromedp/cmd/chromedp-gen/internal"
)

// FileHeader is the file header template.

//line templates/file.qtpl:8
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line templates/file.qtpl:8
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line templates/file.qtpl:8
func StreamFileHeader(qw422016 *qt422016.Writer, pkgName string, d *internal.Domain) {
	//line templates/file.qtpl:8
	qw422016.N().S(`
`)
	//line templates/file.qtpl:9
	if d != nil {
		//line templates/file.qtpl:9
		qw422016.N().S(`// Package `)
		//line templates/file.qtpl:9
		qw422016.N().S(d.PackageName())
		//line templates/file.qtpl:9
		qw422016.N().S(` provides the Chrome Debugging Protocol
// commands, types, and events for the `)
		//line templates/file.qtpl:10
		qw422016.N().S(d.String())
		//line templates/file.qtpl:10
		qw422016.N().S(` domain.
// `)
		//line templates/file.qtpl:11
		if desc := d.Description; desc != "" {
			//line templates/file.qtpl:11
			qw422016.N().S(`
`)
			//line templates/file.qtpl:12
			qw422016.N().S(formatComment(desc, "", ""))
			//line templates/file.qtpl:12
			qw422016.N().S(`
//`)
			//line templates/file.qtpl:13
		}
		//line templates/file.qtpl:13
		qw422016.N().S(`
// Generated by the chromedp-gen command.`)
		//line templates/file.qtpl:14
	}
	//line templates/file.qtpl:14
	qw422016.N().S(`
package `)
	//line templates/file.qtpl:15
	qw422016.N().S(pkgName)
	//line templates/file.qtpl:15
	qw422016.N().S(`

`)
	//line templates/file.qtpl:17
	qw422016.N().S("// Code generated by chromedp-gen. DO NOT EDIT.")
	//line templates/file.qtpl:17
	qw422016.N().S(`
`)
//line templates/file.qtpl:18
}

//line templates/file.qtpl:18
func WriteFileHeader(qq422016 qtio422016.Writer, pkgName string, d *internal.Domain) {
	//line templates/file.qtpl:18
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line templates/file.qtpl:18
	StreamFileHeader(qw422016, pkgName, d)
	//line templates/file.qtpl:18
	qt422016.ReleaseWriter(qw422016)
//line templates/file.qtpl:18
}

//line templates/file.qtpl:18
func FileHeader(pkgName string, d *internal.Domain) string {
	//line templates/file.qtpl:18
	qb422016 := qt422016.AcquireByteBuffer()
	//line templates/file.qtpl:18
	WriteFileHeader(qb422016, pkgName, d)
	//line templates/file.qtpl:18
	qs422016 := string(qb422016.B)
	//line templates/file.qtpl:18
	qt422016.ReleaseByteBuffer(qb422016)
	//line templates/file.qtpl:18
	return qs422016
//line templates/file.qtpl:18
}

// FileImportTemplate is a general import template.

//line templates/file.qtpl:21
func StreamFileImportTemplate(qw422016 *qt422016.Writer, m map[string]string) {
	//line templates/file.qtpl:22
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	//line templates/file.qtpl:27
	qw422016.N().S(`
import (`)
	//line templates/file.qtpl:28
	for _, k := range keys {
		//line templates/file.qtpl:29
		v := m[k]

		//line templates/file.qtpl:30
		qw422016.N().S(`
	`)
		//line templates/file.qtpl:31
		if k != v {
			//line templates/file.qtpl:31
			qw422016.N().S(v)
			//line templates/file.qtpl:31
			qw422016.N().S(` `)
			//line templates/file.qtpl:31
		}
		//line templates/file.qtpl:31
		qw422016.N().S(`"`)
		//line templates/file.qtpl:31
		qw422016.N().S(k)
		//line templates/file.qtpl:31
		qw422016.N().S(`"`)
		//line templates/file.qtpl:31
	}
	//line templates/file.qtpl:31
	qw422016.N().S(`
)
`)
//line templates/file.qtpl:33
}

//line templates/file.qtpl:33
func WriteFileImportTemplate(qq422016 qtio422016.Writer, m map[string]string) {
	//line templates/file.qtpl:33
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line templates/file.qtpl:33
	StreamFileImportTemplate(qw422016, m)
	//line templates/file.qtpl:33
	qt422016.ReleaseWriter(qw422016)
//line templates/file.qtpl:33
}

//line templates/file.qtpl:33
func FileImportTemplate(m map[string]string) string {
	//line templates/file.qtpl:33
	qb422016 := qt422016.AcquireByteBuffer()
	//line templates/file.qtpl:33
	WriteFileImportTemplate(qb422016, m)
	//line templates/file.qtpl:33
	qs422016 := string(qb422016.B)
	//line templates/file.qtpl:33
	qt422016.ReleaseByteBuffer(qb422016)
	//line templates/file.qtpl:33
	return qs422016
//line templates/file.qtpl:33
}

// FileVarTemplate is a template for a single variable declaration.

//line templates/file.qtpl:36
func StreamFileVarTemplate(qw422016 *qt422016.Writer, name, value, desc string) {
	//line templates/file.qtpl:36
	qw422016.N().S(`
`)
	//line templates/file.qtpl:37
	qw422016.N().S(formatComment(desc, "", name+" "))
	//line templates/file.qtpl:37
	qw422016.N().S(`
var `)
	//line templates/file.qtpl:38
	qw422016.N().S(name)
	//line templates/file.qtpl:38
	qw422016.N().S(` = `)
	//line templates/file.qtpl:38
	qw422016.N().S(value)
	//line templates/file.qtpl:38
	qw422016.N().S(`
`)
//line templates/file.qtpl:39
}

//line templates/file.qtpl:39
func WriteFileVarTemplate(qq422016 qtio422016.Writer, name, value, desc string) {
	//line templates/file.qtpl:39
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line templates/file.qtpl:39
	StreamFileVarTemplate(qw422016, name, value, desc)
	//line templates/file.qtpl:39
	qt422016.ReleaseWriter(qw422016)
//line templates/file.qtpl:39
}

//line templates/file.qtpl:39
func FileVarTemplate(name, value, desc string) string {
	//line templates/file.qtpl:39
	qb422016 := qt422016.AcquireByteBuffer()
	//line templates/file.qtpl:39
	WriteFileVarTemplate(qb422016, name, value, desc)
	//line templates/file.qtpl:39
	qs422016 := string(qb422016.B)
	//line templates/file.qtpl:39
	qt422016.ReleaseByteBuffer(qb422016)
	//line templates/file.qtpl:39
	return qs422016
//line templates/file.qtpl:39
}
