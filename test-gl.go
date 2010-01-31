package main

import "sdl"

import "gl"

import "unsafe"

func main() {

	sdl.Init(sdl.INIT_VIDEO);

	var screen = sdl.SetVideoMode(640, 480, 32, sdl.OPENGL);

	if screen == nil {
		panic("sdl error")
	}

	if gl.Init() != 0 {
		panic("glew error")
	}

    vertex_shader := gl.CreateShader(gl.VERTEX_SHADER);

    vertex_shader.Source("  attribute vec3 aVertexPosition;\
\
  void main(void)\
  {\
    gl_Position = vec4(aVertexPosition, 1.0);\
  }");

    vertex_shader.Compile();
    print(vertex_shader.GetInfoLog());

    fragment_shader := gl.CreateShader(gl.FRAGMENT_SHADER);
    fragment_shader.Source("void main(void){gl_FragColor = vec4(1.0, 0.0, 0.0, 1.0);}");
    fragment_shader.Compile();
    print(fragment_shader.GetInfoLog());

    program := gl.CreateProgram();
    program.AttachShader(vertex_shader);
    program.AttachShader(fragment_shader);
    program.Link();
    print(program.GetInfoLog());
    program.Use();

    //buffer := gl.CreateBuffer();
    //buffer.Bind(gl.ELEMENT_ARRAY_BUFFER);
    //gl.BufferData(gl.ELEMENT_ARRAY_BUFFER,9,unsafe.Pointer(&array[0]),gl.STATIC_DRAW);

	var running = true;

	for running {

        if(gl.GetError()!=0) {
            println(gl.GetError());
        }

		e := &sdl.Event{};

		for e.Poll() {
			switch e.Type {
			case sdl.QUIT:
				running = false;
				break;
			case sdl.KEYDOWN:
				running = false;
				break;
			}
		}

        program.Use();

		vertex_loc:=program.GetAttribLocation("aVertexPosition");

		gl.EnableVertexAttribArray(vertex_loc);
		gl.VertexAttribPointer(vertex_loc, 3, gl.FLOAT, false, 0, unsafe.Pointer(&suzanne[0]));

        gl.DrawArrays(gl.TRIANGLES, 0, len(suzanne)/9);

		sdl.GL_SwapBuffers();

		sdl.Delay(25);
	}

	sdl.Quit();

}
