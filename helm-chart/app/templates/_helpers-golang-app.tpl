{{- define "golangapp.selectorLabels" -}}
app: {{ template "app.shortname" . }}-golangapp
release: {{ .Release.Name }}
component: golangapp
date: "{{ now | unixEpoch }}"
{{- end }}

{{- define "golangapp.labels" -}}
chart: {{ include "app.chart" . }}
{{ include "golangapp.selectorLabels" . }}
heritage: {{ .Release.Service }}
{{- end }}