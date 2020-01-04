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

# Caveats

* The build relies on a locally provided JDK. I'm not sure how to avoid that at the time of this writing,
  but if you run into [this issue](https://github.com/bazelbuild/bazel/issues/6993) you most likely need
  to set JAVA_HOME.
  * On macOS, with openjdk installed through homebrew, try
    ``JAVA_HOME=/usr/local/opt/openjdk/libexec/openjdk.jdk/Contents/Home``.
