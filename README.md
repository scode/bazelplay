# bazelplay

Bazel playground.

# Things to try

```
bazel build //...
bazel build //src/java/org/scode/bazelplay:hello_deploy.jar &&
  java -jar bazel-bin/src/java/org/scode/bazelplay/hello_deploy.jar
bazel run //src/java/org/scode/bazelplay:hello
bazel run //src/python/scode:hello
bazel build --build_python_zip //src/python/scode:hello &&
  bazel-bin/src/python/scode/hello
bazel test //...
```

# Warts?

* I am currently unsuccessful at setting the necessary Python version on Python targets using
  [python_version](https://docs.bazel.build/versions/master/be/python.html). Not sure why that is,
  but `bazel test //...` fails at the time of this writing due to not accepting the Python 3 syntax.
