{{- define "golangapp.selectorLabels" -}}
app: {{ template "app.shortname" . }}-golangapp
release: {{ .Release.Name }}
component: golangapp
helm-revision: {{ .Release.Revision }}
{{- end }}

{{- define "golangapp.labels" -}}
chart: {{ include "app.chart" . }}
{{ include "golangapp.selectorLabels" . }}
heritage: {{ .Release.Service }}
{{- end }}