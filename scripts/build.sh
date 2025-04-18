#!/bin/bash
ARG_EXEC_NAME=$1
ARG_EXEC_PLATFORMS=$2
ARG_LDFLAGS=$3
ARG_BUILD_FILE=$4
BUILD_DIR=build
CURRENT_PLATFORM=$(go version | awk '{print $4}')
CURRENT_PLATFORM_GOOS="${CURRENT_PLATFORM%%/*}"
CURRENT_PLATFORM_GOARCH="${CURRENT_PLATFORM##*/}"
CURRENT_EXECUTABLE="${ARG_EXEC_NAME}-${CURRENT_PLATFORM_GOOS}-${CURRENT_PLATFORM_GOARCH}"

echo "Building <$ARG_EXEC_NAME> for $ARG_EXEC_PLATFORMS"
echo "Current platform: $CURRENT_PLATFORM_GOOS-$CURRENT_PLATFORM_GOARCH, expected current executable at <${CURRENT_EXECUTABLE}>"

for PLATFORM in $ARG_EXEC_PLATFORMS; do
  GOOS="${PLATFORM%%/*}"
  GOARCH="${PLATFORM##*/}"
  OUTPUT_NAME="$ARG_EXEC_NAME-$GOOS-$GOARCH"
  [ "$GOOS" = "windows" ] && OUTPUT_NAME="$OUTPUT_NAME.exe"

  echo "📦 $GOOS/$GOARCH -> $BUILD_DIR/$OUTPUT_NAME"
  CMD="CGO_ENABLED=0 GOOS=$GOOS GOARCH=$GOARCH go build -o $BUILD_DIR/$OUTPUT_NAME -ldflags=\"$ARG_LDFLAGS\" $ARG_BUILD_FILE"
  echo "$CMD"
  eval $CMD
done

echo ""
echo "==="
echo ""
echo "Signing binaries"
echo ""
echo "==="
echo ""
for file in "$BUILD_DIR"/"$ARG_EXEC_NAME"-*; do
  # Проверка, существует ли файл (если совпадений нет, будет передан шаблон как строка)
  if [ -f "$file" ]; then
    echo "Signing $file to $file.sig"
    cat "$file" | "$BUILD_DIR/$CURRENT_EXECUTABLE" sign -i identity.eye > "$file.sig"
  fi
done
