!/usr/bin/bash
if [ ! -e "go.mod" ]:
then
	make initmod
fi
make
