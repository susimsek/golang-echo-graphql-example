{{- define "golangapp.selectorLabels" -}}
app: {{ template "app.shortname" . }}-golangapp
release: {{ .Release.Name }}
component: golangapp
rollme: {{ randAlphaNum 5 | quote }}
{{- end }}

{{- define "golangapp.labels" -}}
chart: {{ include "app.chart" . }}
{{ include "golangapp.selectorLabels" . }}
heritage: {{ .Release.Service }}
{{- end }}