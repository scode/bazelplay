# bazelplay

Bazel playground.

# Things to try

```
bazel build //...
bazel build //src/java/org/scode/bazelplay:hello_deploy.jar &&
  java -jar bazel-bin/src/java/org/scode/bazelplay/hello_deploy.jar
bazel run //src/java/org/scode/bazelplay:hello
bazel run --python_top=//:pyruntime //src/python/scode:hello
bazel build --python_top=//:pyruntime --build_python_zip //src/python/scode:hello &&
  bazel-bin/src/python/scode/hello
bazel test --python_top=//:pyruntime //...
```
