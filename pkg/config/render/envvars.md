# Available environment variable bindings

{{range . -}}
## {{ .Name }}
{{ if .Description }}
{{ .Description }}
{{ end -}}

{{ range .Options}}
  {{- if not .Undocumented }}
  - `{{ .Name }}`: {{ if .EnvvarDesc }}{{ .EnvvarDesc }}{{ else }}{{ .Description }}{{ end }}
  {{- end -}}
{{- end -}}
{{ end }}
