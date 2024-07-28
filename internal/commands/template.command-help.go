package commands

var commandHelpTemplate string = `
{{getCallerName}}: {{whiteCommand .Name}} — {{.ShortDescription}}

Description
    {{- $wrappedText := wrap .Description getMaxWidth 4}}

    {{$wrappedText}}

Usage

    {{white getCallerName}} {{whiteCommand .Name}} {{range .Arguments}}{{if .Optional}}[{{.Name}}]{{else}}<{{.Name}}>{{end}} {{end}}

Arguments
{{range .Arguments}}
    {{- $argDesc := wrap .Description getMaxWidth 4}}
    {{.Label}}

    {{$argDesc}}
{{end}}
Run '{{getCallerName}} help' for information on additional commands.

`
