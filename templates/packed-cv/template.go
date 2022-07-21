package packedcv

import (
	"io"
	"text/template"

	"github.com/interlock/go-resumake/jsonresume"
)

const document = `
\documentclass[a4paper]{article}
\usepackage{fullpage}
\usepackage{amsmath}
\usepackage{amssymb}
\usepackage{textcomp}
\usepackage[utf8]{inputenc}
\usepackage[T1]{fontenc}
\usepackage[hidelinks]{hyperref}
\usepackage{geometry}
\geometry{a4paper, portrait, margin=5em}
\textheight=10in
\pagestyle{empty}

{{template "resumeDefinitions"}}

\begin{document}
\vspace*{-40pt}
{{ $data := . }}
{{ range .Sections }}
	{{- if eq . "Basics" -}}
	{{- template "Basics" $data -}}
	{{- end -}}

	{{- if eq . "Education" -}}
	{{- template "Education" $data -}}
	{{- end -}}

	{{- if eq . "Work" }}
	{{- template "Work" $data -}}
	{{ end -}}

	{{- if eq . "Skills" }}
	{{- template "Skills" $data -}}
	{{ end -}}

	{{- if eq . "Volunteer" }}
	{{- template "Volunteer" $data -}}
	{{ end -}}

	{{- if eq . "Interests" }}
	{{- template "Interests" $data -}}
	{{ end -}}
{{ end }}

%% WHITESPACE

\end{document}
`

const resumeDefinitions = `
{{ define "resumeDefinitions" }}
%\renewcommand{\encodingdefault}{cg}
%\renewcommand{\rmdefault}{lgrcmr}

\def\bull{\vrule height 0.8ex width .7ex depth -.1ex }

% DEFINITIONS FOR RESUME %%%%%%%%%%%%%%%%%%%%%%%

\newcommand{\area} [2] {
	\vspace*{-9pt}
	\begin{verse}
		\textbf{#1}   #2
	\end{verse}
}

\newcommand{\lineunder} {
	\vspace*{-8pt} \\\\
	\hspace*{-18pt} \hrulefill \\\\
}

\newcommand{\header} [1] {
	{\hspace*{-18pt} \textsc{#1}}
	\vspace*{-1em} \lineunder
}

\newcommand{\employer} [3] {
	{ \textbf{#1} (#2)\\\\ \underline{\textbf{\emph{#3}}}\\\\  }
}

\newcommand{\contact} [3] {
	\vspace*{-10pt}
	\begin{center}
		{\Huge \scshape {#1}}\\\\
		#2 \\\\ #3
	\end{center}
	\vspace*{-8pt}
}

\newenvironment{achievements}{
	\begin{list}
		{$\bullet$}{\topsep 0pt \itemsep -2pt}}{\vspace*{4pt}
	\end{list}
}

\newcommand{\schoolwithcourses} [4] {
	\textbf{#1} #2 $\bullet$ #3\\\\
	#4 \\\\
	\vspace*{5pt}
}

\newcommand{\school} [4] {
	\textbf{#1} #2 $\bullet$ #3\\\\
	#4 \\\\
}
% END RESUME DEFINITIONS %%%%%%%%%%%%%%%%%%%%%%%
{{- end -}}
`

func Render(tmpl *template.Template, resume *jsonresume.JSONResume, output io.Writer) error {
	documentTemplate := template.Must(tmpl.Parse(document))

	template.Must(documentTemplate.Parse(resumeDefinitions))
	template.Must(documentTemplate.Parse(basics))
	template.Must(documentTemplate.Parse(education))
	template.Must(documentTemplate.Parse(work))
	template.Must(documentTemplate.Parse(skills))
	template.Must(documentTemplate.Parse(volunteer))
	template.Must(documentTemplate.Parse(interests))

	sections := []string{"Basics", "Education", "Work", "Skills", "Volunteer", "Interests"}
	data := map[string]interface{}{
		"JSON":     resume,
		"Sections": sections,
	}
	err := documentTemplate.Execute(output, data)
	if err != nil {
		return err
	}

	return nil
}
