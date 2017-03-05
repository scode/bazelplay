# bazelplay

Bazel playground.

# Things to try

```
bazel build //...
bazel build //src/java/org/scode/bazelplay:hello_deploy.jar &&
  java -jar bazel-bin/src/java/org/scode/bazelplay/hello_deploy.jar
```
