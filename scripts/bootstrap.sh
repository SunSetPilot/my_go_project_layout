#!/bin/bash

BinaryName="YOUR_BINARY_NAME"

CURDIR=$(cd $(dirname $0); pwd)
if [ "X$1" != "X" ]; then
    RUNTIME_ROOT=$1
else
    RUNTIME_ROOT=${CURDIR}
fi

if [ "X$RUNTIME_ROOT" == "X" ]; then
    echo "There is no RUNTIME_ROOT support."
    echo "Usage: ./bootstrap.sh $RUNTIME_ROOT"
    exit -1
fi

RUNTIME_CONF_ROOT=$RUNTIME_ROOT/conf
RUNTIME_LOG_ROOT="YOUR_LOGS_DIR"
CONF_DIR=$RUNTIME_CONF_ROOT/$BinaryName.yaml
LOG_DIR=$RUNTIME_LOG_ROOT

exec "$CURDIR"/bin/${BinaryName} -conf="$CONF_DIR" -log_dir="$LOG_DIR"