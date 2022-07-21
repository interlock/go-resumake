package packedcv

const work = `
{{ define "Work" }}
{{ if .JSON.Work }}
%==== Experience ====%

\header{Experience}
\vspace{1mm}
{{- range .JSON.Work -}}
	{{- $line1 := "" -}}
	{{- $line2 := "" -}}
	{{- $highlightLines := newSlice -}}
	
	{{- if .Company -}}
		{{- $line1 = printf "%s\\textbf{%s}" $line1 (escape .Company) -}}
  	{{- end -}}

	{{- if .Position -}}
		{{- $line1 = printf "%s $\\cdot$ \\textit{%s}" $line1 (escape .Position) -}}
	{{- end -}}

	{{- if and .StartDate .EndDate -}}
		{{- $line1 = printf "%s \\hfill %s $\\cdot$ %s" $line1 (dateFormat .StartDate "2006-Jan") (dateFormat .EndDate "2006-Jan") -}}
	{{- else -}}
		{{- if .StartDate -}}
			{{- $line1 = printf "%s \\hfill %s $\\cdot$ Present" $line1 (dateFormat .StartDate "2006-Jan") -}}
		{{- end -}}
		{{- if .EndDate -}}
			{{- $line1 = printf "%s \\hfill %s" $line1 (dateFormat .EndDate "2006-Jan") -}}
		{{- end -}}
	{{- end -}}

	{{- if $line1 -}}{{- $line1 = printf "%s" $line1 -}}{{- end -}}
	{{- if $line2 -}}{{- $line2 = printf "%s\\\\" $line2 -}}{{- end -}}
	
	{{- if .Highlights -}}
		{{- $highlightLines = appendSlice $highlightLines "\\vspace{-1mm}" -}}
		{{- $highlightLines = appendSlice $highlightLines "\\begin{itemize} \\itemsep 1pt" -}}
		{{- range .Highlights -}}
			{{- $highlightLines = appendSlice $highlightLines (printf "\\item %s" (escape .)) -}}
		{{- end -}}
		{{- $highlightLines = appendSlice $highlightLines "\\end{itemize}" -}}
	{{- end -}}
	
	{{ $line1 }}
	{{ $line2 }}
	{{ join $highlightLines "\n" }}
	{{ printf "\n" }}
{{- end -}}
{{- end -}}
{{- end -}}
`
