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
bazel test bazel test --python_top=//:pyruntime //...
```

# Warts?

* Python version selection is [currently
  buggy](https://github.com/bazelbuild/bazel/issues/4815). The
  [workaround](https://github.com/bazelbuild/bazel/issues/4815#issuecomment-460777113)
  is applied, hence the `--python_top` arguments above. Also, `BUILD`
  must be modified if `/usr/bin/python{2,3}` are not valid on your
  system.
