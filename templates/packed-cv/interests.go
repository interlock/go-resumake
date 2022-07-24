package packedcv

const interests = `
{{ define "Interests" }}

%==== interests ====%

\noindent\begin{minipage}{\textwidth}
\header{Interests}
\begin{tabular}{ l p{.8\textwidth} }
{{ range .JSON.Interests -}}
	{{ printf "%s:& %s\\\\\n" (escape .Name) (escape (join .Keywords ", ")) }}
{{- end -}}
\end{tabular}
\end{minipage}
\vspace{2mm}

{{- end -}}
`
