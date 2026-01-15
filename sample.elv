#!/usr/bin/env elvish

echo "Running script:" (src)[name]
if (> (count $args) 1) {
  echo "Arguments received:"
  var c = (num 0)
  for arg $args[0..] {
    echo "  at index" $c":" $arg
    set c = (+ $c 1)
  }
} else {
  echo "No additional arguments received."
}
