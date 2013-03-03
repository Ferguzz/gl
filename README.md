gl | OpenGL Bindings for golang
===============================

You will need [GLEW](http://glew.sourceforge.net/) at least version 1.5.8.

Once GLEW is installed, you can install this package with `go get`:

    go get github.com/go-gl/gl

Contact / discussion mailing list
---------------------------------

The discussion is hosted on [google groups](https://groups.google.com/forum/#!forum/go-gl).
We can also be found periodically on IRC, [#go-gl on freenode](http://webchat.freenode.net/?randomnick=1&channels=go-gl&prompt=1).

Documentation and examples
--------------------------

Autogenerated documentation can be found on
[pkgdoc.org](http://go.pkgdoc.org/github.com/go-gl/gl), and the more thorough
[official GL SDK documentation](http://www.opengl.org/sdk/docs/man/xhtml/) applies.

Examples can be found in our [examples repository](https://github.com/go-gl/examples).

# More libraries: Easy windowing, meshes, text rendering, etc:

* [GLFW bindings](https://github.com/go-gl/glfw) for easy windowing, input etc.
* [gltext](https://github.com/go-gl/gltext) a native go library for glyph packing and text rendering
* [glu](https://github.com/go-gl/glu) GLU bindings
* [glchart](https://github.com/go-gl/glchart) a [go chart](https://www.github.com/vdobler/chart) OpenGL backend
* [gldebug](https://github.com/go-gl/gldebug) graphical timing and memory debugging utilities

# Problems and contributing

If you encounter any problems please [file an issue](https://github.com/go-gl/gl/issues/new).
Pull requests are welcome. We're looking for contributors!

# Setting your GOPATH

For the above `go get` to work without without requiring admin privileges,
please refer to [go documentation](http://golang.org/doc/code.html) or this
[five minute go screencast](http://www.youtube.com/watch?v=XCsL89YtqCs).

Operating System Specific advice
--------------------------------

The `go-gl/gl` authors primarily develop on Linux-based machines. The package
should work on other operating systems, if you encounter problems after following
the instructions below, please file an issue.

## Linux / OSX

This package uses [cgo](http://golang.org/cmd/cgo/) and therefore you will need
a C compiler such as GCC. On Debian, you can install `build-essential` to get this.

You need to install the GLEW development libraries the appropriate package manager
for your Linux distribution or by downloading the sources from the
[GLEW website](http://glew.sourceforge.net/).

For non-standard GLEW install locations, after `go get` cd to
`$GOPATH/src/github.com/go-gl/gl` and run:

    CGO_CFLAGS="-I/path/to/includes" CGO_LDFLAGS="-L/path/to/libs" go install

You may be able to determine the appropriate flags using:

    pkg-config glew --libs --cflags

Note that we import from `GL/glew.h`, so if `pkg-config --cflags` reports
`/opt/include/GL`, you may need to remove `/GL` so that `CGO_CFLAGS=-I/opt/include`.

## OSX

The following instructions are currently a best guess. If you develop on OSX
please [get in contact](http://go-gl.github.com) and let us know if this works,
or [not](https://github.com/go-gl/gl/issues/new)!

### GLEW from Homebrew / Macports

These can be used to install GLEW. Beware that you may need to set `CGO_CFLAGS`
and `CGO_LDFLAGS` - the Linux instructions above are relevant.

## Windows

Windows support is currently unknown. `gl` uses `cgo`, and therefore you will
need a C compiler. If you install the GLEW to your compiler's include and lib
directories, things should work. If they do not, please
[contact us](http://go-gl.github.com).

Forward compatibility
---------------------

It is the intent of the `go-gl` authors to keep `gl` and related packages
forward-compatible at both the API/ABI levels _per go release_ so that you can
continue to import from this github repository. Development requiring breakage
will occur on the `dev` branch, and will be merged with master when a new go
version is released.

