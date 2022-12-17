find  . -iname '*.sh' -printf '%f\n'| cut -d '.' -f 1 | sort -nr # Descending
