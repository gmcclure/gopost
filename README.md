GoPost
======

This is my (gmcclure@gmail.com) personal, Go-based blogging engine.

Templates are located in `tmpl/`, so the blog should be started from the
project's root directory: 

    bin/main

For now, static files are stored in the toplevel `static` directory, which is
symlinked in `bin`. It's hacky, and there should be a config var somewhere or
something, but it works for now. An `ls -l bin` should produce a line that
looks like this:

    static -> ../static
