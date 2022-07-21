package packedcv

const basics = `
{{ define "Basics" }}
{{- if .JSON.Basics -}}
{{- with .JSON.Basics -}}

	{{- $address := "" -}}
	{{- if .Location -}}
	{{- $address = escape .Location.Address -}}
	{{- end -}}
	  
	{{- $line1 := "" -}} 
  {{- if .Name -}}
	{{- $line1 = printf "{\\Huge \\scshape {%s}}\\\\" (escape .Name) -}}
	{{- end -}}
	  
	{{- $line2 := newSlice -}}
	{{- if $address -}}
	{{- $line2 = appendSlice $line2 $address -}}
	{{- end -}}

	{{- if .Email -}}
	{{- $line2 = appendSlice $line2 (printf "\\href{mailto:%s}{%s}" (escape .Email) (escape .Email)) -}}
	{{- end -}}

	{{- if .Phone -}}
	{{- $line2 = appendSlice $line2 (printf "\\href{tel:%s}{%s}" (escape .Phone) (escape .Phone)) -}}
	{{- end -}}

	{{- if .Website -}}
	{{- $line2 = appendSlice $line2 (printf "\\href{%s}{%s}" (escape .Website) (escape .Website)) -}}
	{{- end -}}

	{{- $line2 = join $line2 " $\\cdot$ " -}}
	

%==== Profile ====%

\vspace*{-10pt}
\begin{center}
	{{ $line1 }}
	{{ $line2 }}
\end{center}
{{- end -}}
{{- else -}}
%% Basic skip
{{- end -}}
{{ end }}
`
