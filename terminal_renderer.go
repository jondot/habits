package main

import (
  "bytes"
  "github.com/russross/blackfriday"
)

const (
  TermEscapePrefix = "\033["
  TermEscapeSuffix = "m"

  Reset = TermEscapePrefix + "0" + TermEscapeSuffix

  Bold = TermEscapePrefix + "1" + TermEscapeSuffix

  Green = TermEscapePrefix + "32" + TermEscapeSuffix
  Cyan  = TermEscapePrefix + "36" + TermEscapeSuffix
  White = TermEscapePrefix + "37" + TermEscapeSuffix

  OnBlue = TermEscapePrefix + "44" + TermEscapeSuffix
  OnGrey = TermEscapePrefix + "40" + TermEscapeSuffix
)

type Terminal struct {
  flags    int
}


func NewTerminalRenderer(flags int) blackfriday.Renderer {
  return &Terminal{
    flags:      flags,
  }
}

func (options *Terminal) GetFlags() int {
  return options.flags
}

func (options *Terminal) TitleBlock(out *bytes.Buffer, text []byte) {
  text = bytes.TrimPrefix(text, []byte("% "))
  text = bytes.Replace(text, []byte("\n% "), []byte("\n"), -1)
  out.WriteString("[")
  out.Write(text)
  out.WriteString("]")
}

func (options *Terminal) Header(out *bytes.Buffer, text func() bool, level int, id string) {
  marker := out.Len()

  defer out.WriteString(Reset)
  out.WriteString(OnBlue)
  if !text() {
    out.Truncate(marker)
    return
  }

  return
}

func (options *Terminal) BlockHtml(out *bytes.Buffer, text []byte) {
  return
}

func (options *Terminal) HRule(out *bytes.Buffer) {
  out.WriteString("----------")
}

func (options *Terminal) BlockCode(out *bytes.Buffer, text []byte, lang string) {
  out.WriteString(string(text))
}

func (options *Terminal) BlockQuote(out *bytes.Buffer, text []byte) {
  out.WriteString("<blockquote>\n")
  out.Write(text)
  out.WriteString("</blockquote>\n")
}

func (options *Terminal) Table(out *bytes.Buffer, header []byte, body []byte, columnData []int) {
  out.Write(header)
  out.WriteString("\n")
  out.Write(body)
}

func (options *Terminal) TableRow(out *bytes.Buffer, text []byte) {
  out.Write(text)
}

func (options *Terminal) TableHeaderCell(out *bytes.Buffer, text []byte, align int) {
  out.Write(text)
}

func (options *Terminal) TableCell(out *bytes.Buffer, text []byte, align int) {
  out.Write(text)
}

func (options *Terminal) Footnotes(out *bytes.Buffer, text func() bool) {

}

func (options *Terminal) FootnoteItem(out *bytes.Buffer, name, text []byte, flags int) {
  return
}

func (options *Terminal) List(out *bytes.Buffer, text func() bool, flags int) {
  return
}

func (options *Terminal) ListItem(out *bytes.Buffer, text []byte, flags int) {
  out.Write(text)
}

func (options *Terminal) Paragraph(out *bytes.Buffer, text func() bool) {
  marker := out.Len()

  if !text() {
    out.Truncate(marker)
    return
  }
}

func (options *Terminal) AutoLink(out *bytes.Buffer, link []byte, kind int) {
  out.Write(link)
}

func (options *Terminal) CodeSpan(out *bytes.Buffer, text []byte) {
  out.WriteString("<")
  out.Write(text)
  out.WriteString(">")
}

func (options *Terminal) DoubleEmphasis(out *bytes.Buffer, text []byte) {
  out.WriteString(Bold)
  out.Write(text)
  out.WriteString(Reset)
}

func (options *Terminal) Emphasis(out *bytes.Buffer, text []byte) {
  if len(text) == 0 {
    return
  }
  out.WriteString(Bold)
  out.Write(text)
  out.WriteString(Reset)
}


func (options *Terminal) Image(out *bytes.Buffer, link []byte, title []byte, alt []byte) {
  return
}

func (options *Terminal) LineBreak(out *bytes.Buffer) {
  out.WriteString("\n")
}

func (options *Terminal) Link(out *bytes.Buffer, link []byte, title []byte, content []byte) {
  out.Write(content)
  return
}

func (options *Terminal) RawHtmlTag(out *bytes.Buffer, text []byte) {
  out.Write(text)
}

func (options *Terminal) TripleEmphasis(out *bytes.Buffer, text []byte) {
  out.WriteString("<strong><em>")
  out.Write(text)
  out.WriteString("</em></strong>")
}

func (options *Terminal) StrikeThrough(out *bytes.Buffer, text []byte) {
  out.WriteString("<del>")
  out.Write(text)
  out.WriteString("</del>")
}

func (options *Terminal) FootnoteRef(out *bytes.Buffer, ref []byte, id int) {

}

func (options *Terminal) Entity(out *bytes.Buffer, entity []byte) {
  out.Write(entity)
}

func (options *Terminal) NormalText(out *bytes.Buffer, text []byte) {
  out.Write(text)
}

func (options *Terminal) Smartypants(out *bytes.Buffer, text []byte) {
  out.Write(text)
}

func (options *Terminal) DocumentHeader(out *bytes.Buffer) {

}

func (options *Terminal) DocumentFooter(out *bytes.Buffer) {

}

func (options *Terminal) TocHeader(text []byte, level int) {

}

func (options *Terminal) TocFinalize() {

}
