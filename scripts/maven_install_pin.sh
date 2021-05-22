#!/bin/bash

set -eux

bazelisk run @unpinned_maven//:pin
