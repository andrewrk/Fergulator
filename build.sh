#!/bin/bash
WORK=/tmp/go-build145183779
rm -rf $WORK
mkdir -p $WORK/_/home/andy/dev/Fergulator/_obj/
cd /home/andy/dev/Fergulator
/usr/lib/go/pkg/tool/linux_amd64/cgo -objdir $WORK/_/home/andy/dev/Fergulator/_obj/ -- -I $WORK/_/home/andy/dev/Fergulator/_obj/ jamulator.go
/usr/lib/go/pkg/tool/linux_amd64/6c -FVw -I $WORK/_/home/andy/dev/Fergulator/_obj/ -I /usr/lib/go/pkg/linux_amd64 -o $WORK/_/home/andy/dev/Fergulator/_obj/_cgo_defun.6 -DGOOS_linux -DGOARCH_amd64 $WORK/_/home/andy/dev/Fergulator/_obj/_cgo_defun.c
gcc -I . -g -O2 -fPIC -m64 -pthread -I $WORK/_/home/andy/dev/Fergulator/_obj/ -o $WORK/_/home/andy/dev/Fergulator/_obj/_cgo_main.o -c $WORK/_/home/andy/dev/Fergulator/_obj/_cgo_main.c
gcc -I . -g -O2 -fPIC -m64 -pthread -I $WORK/_/home/andy/dev/Fergulator/_obj/ -o $WORK/_/home/andy/dev/Fergulator/_obj/_cgo_export.o -c $WORK/_/home/andy/dev/Fergulator/_obj/_cgo_export.c
gcc -I . -g -O2 -fPIC -m64 -pthread -I $WORK/_/home/andy/dev/Fergulator/_obj/ -o $WORK/_/home/andy/dev/Fergulator/_obj/jamulator.cgo2.o -c $WORK/_/home/andy/dev/Fergulator/_obj/jamulator.cgo2.c
gcc -I . -g -O2 -fPIC -m64 -pthread -o $WORK/_/home/andy/dev/Fergulator/_obj/_cgo_.o $WORK/_/home/andy/dev/Fergulator/_obj/_cgo_main.o $WORK/_/home/andy/dev/Fergulator/_obj/_cgo_export.o $WORK/_/home/andy/dev/Fergulator/_obj/jamulator.cgo2.o zelda.o
/usr/lib/go/pkg/tool/linux_amd64/cgo -objdir $WORK/_/home/andy/dev/Fergulator/_obj/ -dynimport $WORK/_/home/andy/dev/Fergulator/_obj/_cgo_.o -dynout $WORK/_/home/andy/dev/Fergulator/_obj/_cgo_import.c
/usr/lib/go/pkg/tool/linux_amd64/6c -FVw -I $WORK/_/home/andy/dev/Fergulator/_obj/ -I /usr/lib/go/pkg/linux_amd64 -o $WORK/_/home/andy/dev/Fergulator/_obj/_cgo_import.6 -DGOOS_linux -DGOARCH_amd64 $WORK/_/home/andy/dev/Fergulator/_obj/_cgo_import.c
/usr/lib/go/pkg/tool/linux_amd64/6g -o $WORK/_/home/andy/dev/Fergulator/_obj/_go_.6 -p _/home/andy/dev/Fergulator -D _/home/andy/dev/Fergulator -I $WORK -I /home/andy/golang/pkg/linux_amd64 ./controller.go ./machine.go ./nametable.go ./palette.go ./ppu.go ./rom.go ./video.go $WORK/_/home/andy/dev/Fergulator/_obj/_cgo_gotypes.go $WORK/_/home/andy/dev/Fergulator/_obj/jamulator.cgo1.go
/usr/lib/go/pkg/tool/linux_amd64/pack grc $WORK/_/home/andy/dev/Fergulator.a $WORK/_/home/andy/dev/Fergulator/_obj/_go_.6 $WORK/_/home/andy/dev/Fergulator/_obj/_cgo_import.6 $WORK/_/home/andy/dev/Fergulator/_obj/_cgo_defun.6 $WORK/_/home/andy/dev/Fergulator/_obj/_cgo_export.o $WORK/_/home/andy/dev/Fergulator/_obj/jamulator.cgo2.o
go tool pack rc $WORK/_/home/andy/dev/Fergulator.a zelda.o 
cd .
/usr/lib/go/pkg/tool/linux_amd64/6l -o Fergulator -L $WORK -L /home/andy/golang/pkg/linux_amd64 $WORK/_/home/andy/dev/Fergulator.a
