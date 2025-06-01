#!/bin/sh

ARGS="/runtime/example/$2"
OUTPUT="/runtime/output"

cmd_dot() {
  ./symtable "$ARGS" >"$OUTPUT/symtable.gv"
  dot -Tsvg "$OUTPUT/symtable.gv" -o "$OUTPUT/symtable.svg"
  dot -Tpng "$OUTPUT/symtable.gv" -o "$OUTPUT/symtable.png"
  dot -Tpdf "$OUTPUT/symtable.gv" -o "$OUTPUT/symtable.pdf"
}

# 定义其他任务
cmd_symtable() { ./symtable "$ARGS"; }
cmd_parser() { ./parser "$ARGS"; }
cmd_ir() { ./ir "$ARGS"; }

case "$1" in
dot)
  cmd_dot
  ;;
symtable)
  cmd_symtable
  ;;
parser)
  cmd_parser
  ;;
ir)
  cmd_ir
  ;;
*)
  exit 1
  ;;
esac
