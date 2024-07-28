package commands

var applicationHelpTemplate string = `
{{.Name}}: {{.ShortDescription}}

Description
    {{- $wrappedText := wrap .Description .Config.DescriptionMaxWidth 4}}

    {{$wrappedText}}

Usage

    {{.Name}} <command> [arguments]

Commands
{{$longestLen := longestStringLength .Commands -}}
{{range .Commands}}
    {{$padded := padEnd .Name $longestLen -}}
    {{white $padded}}  {{.ShortDescription}}{{end}}
{{$paddedHelp := padEnd "help" $longestLen}}    {{white $paddedHelp}}  Display this help message

Run '{{.Name}} help <command>' for more information on a command.

`
