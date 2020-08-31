我是光年实验室高级招聘经理。
我在github上访问了你的开源项目，你的代码超赞。你最近有没有在看工作机会，我们在招软件开发工程师，拉钩和BOSS等招聘网站也发布了相关岗位，有公司和职位的详细信息。
我们公司在杭州，业务主要做流量增长，是很多大型互联网公司的流量顾问。公司弹性工作制，福利齐全，发展潜力大，良好的办公环境和学习氛围。
公司官网是http://www.gnlab.com,公司地址是杭州市西湖区古墩路紫金广场B座，若你感兴趣，欢迎与我联系，
电话是0571-88839161，手机号：18668131388，微信号：echo 'bGhsaGxoMTEyNAo='|base64 -D ,静待佳音。如有打扰，还请见谅，祝生活愉快工作顺利。

## Ark [![Build Status](https://api.travis-ci.org/ark-lang/ark.svg?branch=master)][1] [![license](http://img.shields.io/badge/license-MIT-brightgreen.svg)](https://raw.githubusercontent.com/ark-lang/ark/master/LICENSE)
[1]: https://travis-ci.org/ark-lang/ark "Build Status"

<img width="50%" align="right" style="display: block; margin:40px auto;" src="https://raw.githubusercontent.com/ark-lang/ark-gl/master/example.gif">

[Ark](//www.ark-lang.org) is a systems programming language focused on being practical and pragmatic. We want a language that is simple, easy to write, yet powerful enough to create readable, performant and concise code for any problem

On the right is a gif of an example program [ark-gl](//github.com/ark-lang/ark-gl) written in Ark using OpenGL and GLFW.

## Index
* [Getting Involved](#getting-involved)
* [Example](#example)
* [Installing](#installing)
    * [Dependencies](#dependencies)
    * [Building](#building)
    * [Compiling Ark code](#compiling-ark-code)

## <a name="getting-involed"></a> Getting Involved
Check out the [contributing guide](/CONTRIBUTING.md), there's a lot of information there to give you ideas of how you can help out.

## <a name="example"></a> Example
Ark is still a work in progress, this code sample reflects what Ark can
do *currently*, though the way you write the following will likely change
in the near future.

More examples can be found [here](/examples).

```rust
// binding to printf
[c] func printf(fmt: ^u8, ...);

pub func main(argc: int, argv: ^^u8) -> int {
    // accessed via the C module
    C::printf(c"Running %s\n", ^argv);

    // mutable i, type inferred
    mut i := 0;

    for i < 5 {
        C::printf(c"%d\n", i);

        i += 1;
    }
    return 0;
}
```

## <a name="installing"></a> Installing
Installing Ark is simple, you'll need a few dependencies
before you get started:

### <a name="dependencies"></a> Dependencies
* Go installed and `$GOPATH` setup - [Instructions on setting up GOPATH](//golang.org/doc/code.html#GOPATH)
* For building LLVM bindings:
  * Subversion
  * A C++ Compiler
  * CMake installed
  * `libedit-dev` installed

### <a name="building"></a> Building
Once you have your dependencies setup, building ark from scratch is done by
running the following commands:

```bash
$ git clone https://github.com/ark-lang/go-llvm.git $GOPATH/src/github.com/ark-lang/go-llvm
$ cd $GOPATH/src/github.com/ark-lang/go-llvm
$ ./build.sh
$ go get github.com/ark-lang/ark/...
```

The `ark` binary will be built in `$GOPATH/bin`. To use the compiler,
make sure `$GOPATH/bin` is in your `$PATH`.

### <a name="compiling-ark-code"></a> Compiling Ark code
Currently the module system Ark uses is a work in progress. As of writing this,
each ark file represents a module. A module has a child-module "C" which
contains all of the C functions and other bindings you may write.

Given the following project structure:

    src/
      - entities/
        - entity.ark
        - player.ark
      - main.ark

To compile this, you would pass through the file which contains the main
entry point (main function) to your program, which is conventionally named "main.ark".

Since our main file is in another folder, we need to set the src folder as
an include directory so that the compiler doesn't think it's a module. We use
the `-I` flag for this:

    ark build -I src src/main.ark --loglevel=debug

This should compile your code, and produce an executable called "main", which
you can then run.

For more information on the module system and how it works,
refer to the ["Modules and Dependencies"](http://book.ark-lang.org/modules.html)
section in the Ark reference.

For more information on program flags, refer to the
["Program Input"](http://book.ark-lang.org/source.html), section in the Ark
reference.
